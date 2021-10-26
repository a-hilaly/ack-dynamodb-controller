// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package table

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws-controllers-k8s/dynamodb-controller/apis/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	svcsdk "github.com/aws/aws-sdk-go/service/dynamodb"
	corev1 "k8s.io/api/core/v1"
)

var (
	ErrTableDeleting = errors.New("Table in 'DELETING' state, cannot be modified or deleted")
	ErrTableCreating = errors.New("Table in 'CREATING' state, cannot be modified or deleted")
	ErrTableUpdating = errors.New("Table in 'UPDATING' state, cannot be modified or deleted")
)

var (
	// TerminalStatuses are the status strings that are terminal states for a
	// DynamoDB table
	TerminalStatuses = []v1alpha1.TableStatus_SDK{
		v1alpha1.TableStatus_SDK_ARCHIVING,
		v1alpha1.TableStatus_SDK_DELETING,
	}
)

var (
	requeueWaitWhileDeleting = ackrequeue.NeededAfter(
		ErrTableDeleting,
		5*time.Second,
	)
	requeueWaitWhileCreating = ackrequeue.NeededAfter(
		ErrTableCreating,
		5*time.Second,
	)
	requeueWaitWhileUpdating = ackrequeue.NeededAfter(
		ErrTableUpdating,
		5*time.Second,
	)
)

// tableHasTerminalStatus returns whether the supplied Dynamodb table is in a
// terminal state
func tableHasTerminalStatus(r *resource) bool {
	if r.ko.Status.TableStatus == nil {
		return false
	}
	ts := *r.ko.Status.TableStatus
	for _, s := range TerminalStatuses {
		if ts == string(s) {
			return true
		}
	}
	return false
}

// isTableCreating returns true if the supplied DynamodbDB table is in the process
// of being created
func isTableCreating(r *resource) bool {
	if r.ko.Status.TableStatus == nil {
		return false
	}
	dbis := *r.ko.Status.TableStatus
	return dbis == string(v1alpha1.TableStatus_SDK_CREATING)
}

// isTableDeleting returns true if the supplied DynamodbDB table is in the process
// of being deleted
func isTableDeleting(r *resource) bool {
	if r.ko.Status.TableStatus == nil {
		return false
	}
	dbis := *r.ko.Status.TableStatus
	return dbis == string(v1alpha1.TableStatus_SDK_DELETING)
}

// isTableUpdating returns true if the supplied DynamodbDB table is in the process
// of being deleted
func isTableUpdating(r *resource) bool {
	if r.ko.Status.TableStatus == nil {
		return false
	}
	dbis := *r.ko.Status.TableStatus
	return dbis == string(v1alpha1.TableStatus_SDK_UPDATING)
}

func (rm *resourceManager) customUpdateTable(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	if isTableDeleting(latest) {
		msg := "table is currently being deleted"
		setSyncedCondition(desired, corev1.ConditionFalse, &msg, nil)
		return desired, requeueWaitWhileDeleting
	}
	if isTableCreating(latest) {
		msg := "table is currently being created"
		setSyncedCondition(desired, corev1.ConditionFalse, &msg, nil)
		return desired, requeueWaitWhileCreating
	}
	if isTableUpdating(latest) {
		msg := "table is currently being updated"
		setSyncedCondition(desired, corev1.ConditionFalse, &msg, nil)
		return desired, requeueWaitWhileUpdating
	}
	if tableHasTerminalStatus(latest) {
		msg := "table is in '" + *latest.ko.Status.TableStatus + "' status"
		setTerminalCondition(desired, corev1.ConditionTrue, &msg, nil)
		setSyncedCondition(desired, corev1.ConditionTrue, nil, nil)
		return desired, nil
	}

	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.customUpdateTable")
	defer exit(err)

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// You can only perform one of the following operations at once:
	// - Modify the provisioned throughput settings of the table.
	// - Enable or disable DynamoDB Streams on the table.
	// - Remove a global secondary index from the table.
	//
	// Create a new global secondary index on the table. After the index begins
	// backfilling, you can use UpdateTable to perform other operations. UpdateTable
	// is an asynchronous operation; while it is executing, the table status changes
	// from ACTIVE to UPDATING. While it is UPDATING, you cannot issue another
	// UpdateTable request. When the table returns to the ACTIVE state, the
	// UpdateTable operation is complete.

	if delta.DifferentAt("Spec.AttributeDefinitions") ||
		delta.DifferentAt("Spec.BillingMode") ||
		delta.DifferentAt("Spec.StreamSpecification") ||
		delta.DifferentAt("Spec.SEESpecification") {
		if err := rm.syncTable(ctx, desired); err != nil {
			return nil, err
		}
	}

	/* 	if delta.DifferentAt("Spec.Tags") {
		if err := rm.syncTableTags(ctx, desired); err != nil {
			return nil, err
		}
	} */

	// We only want to call one those updates at once. Priority to the fastest
	// operations.
	switch {
	case delta.DifferentAt("Spec.ProvisionedThroughput"):
		if err := rm.syncTableProvisionedThroughput(ctx, desired); err != nil {
			return nil, err
		}
	case delta.DifferentAt("Spec.GlobalSecondaryIndexes"):
		if err := rm.syncTableGlobalSecondaryIndexes(ctx, desired); err != nil {
			return nil, err
		}
	}

	return &resource{ko}, nil
}

func (rm *resourceManager) syncTableGlobalSecondaryIndexes(
	ctx context.Context,
	r *resource,
) (err error) {
	fmt.Println("UPDATE CALL")
	return
}

func (rm *resourceManager) syncTableProvisionedThroughput(
	ctx context.Context,
	r *resource,
) (err error) {
	return
}

func (rm *resourceManager) syncTable(
	ctx context.Context,
	r *resource,
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.syncTable")
	defer exit(err)

	input, err := rm.newUpdateTablePayload(ctx, r)
	if err != nil {
		return err
	}

	//rm.sdkapi.UpdateTable
	_, err = rm.sdkapi.UpdateTable(input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateTable", err)
	if err != nil {
		return err
	}

	return nil
}

func (rm *resourceManager) newUpdateTablePayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.UpdateTableInput, error) {
	res := &svcsdk.UpdateTableInput{}
	res.SetTableName(*r.ko.Spec.TableName)

	if r.ko.Spec.AttributeDefinitions != nil {

	} else {

	}
	fmt.Println("UPDATE PAYLOAD")

	return nil, nil
}

/* func (rm *resourceManager) syncTableTags(
	ctx context.Context,
	r *resource,
) (err error) {
	return
} */
