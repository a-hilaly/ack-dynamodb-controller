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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type AttributeAction string

const (
	AttributeAction_ADD    AttributeAction = "ADD"
	AttributeAction_PUT    AttributeAction = "PUT"
	AttributeAction_DELETE AttributeAction = "DELETE"
)

type BackupStatus_SDK string

const (
	BackupStatus_SDK_CREATING  BackupStatus_SDK = "CREATING"
	BackupStatus_SDK_DELETED   BackupStatus_SDK = "DELETED"
	BackupStatus_SDK_AVAILABLE BackupStatus_SDK = "AVAILABLE"
)

type BackupType string

const (
	BackupType_USER       BackupType = "USER"
	BackupType_SYSTEM     BackupType = "SYSTEM"
	BackupType_AWS_BACKUP BackupType = "AWS_BACKUP"
)

type BackupTypeFilter string

const (
	BackupTypeFilter_USER       BackupTypeFilter = "USER"
	BackupTypeFilter_SYSTEM     BackupTypeFilter = "SYSTEM"
	BackupTypeFilter_AWS_BACKUP BackupTypeFilter = "AWS_BACKUP"
	BackupTypeFilter_ALL        BackupTypeFilter = "ALL"
)

type BatchStatementErrorCodeEnum string

const (
	BatchStatementErrorCodeEnum_ConditionalCheckFailed          BatchStatementErrorCodeEnum = "ConditionalCheckFailed"
	BatchStatementErrorCodeEnum_ItemCollectionSizeLimitExceeded BatchStatementErrorCodeEnum = "ItemCollectionSizeLimitExceeded"
	BatchStatementErrorCodeEnum_RequestLimitExceeded            BatchStatementErrorCodeEnum = "RequestLimitExceeded"
	BatchStatementErrorCodeEnum_ValidationError                 BatchStatementErrorCodeEnum = "ValidationError"
	BatchStatementErrorCodeEnum_ProvisionedThroughputExceeded   BatchStatementErrorCodeEnum = "ProvisionedThroughputExceeded"
	BatchStatementErrorCodeEnum_TransactionConflict             BatchStatementErrorCodeEnum = "TransactionConflict"
	BatchStatementErrorCodeEnum_ThrottlingError                 BatchStatementErrorCodeEnum = "ThrottlingError"
	BatchStatementErrorCodeEnum_InternalServerError             BatchStatementErrorCodeEnum = "InternalServerError"
	BatchStatementErrorCodeEnum_ResourceNotFound                BatchStatementErrorCodeEnum = "ResourceNotFound"
	BatchStatementErrorCodeEnum_AccessDenied                    BatchStatementErrorCodeEnum = "AccessDenied"
	BatchStatementErrorCodeEnum_DuplicateItem                   BatchStatementErrorCodeEnum = "DuplicateItem"
)

type BillingMode string

const (
	BillingMode_PROVISIONED     BillingMode = "PROVISIONED"
	BillingMode_PAY_PER_REQUEST BillingMode = "PAY_PER_REQUEST"
)

type ComparisonOperator string

const (
	ComparisonOperator_EQ           ComparisonOperator = "EQ"
	ComparisonOperator_NE           ComparisonOperator = "NE"
	ComparisonOperator_IN           ComparisonOperator = "IN"
	ComparisonOperator_LE           ComparisonOperator = "LE"
	ComparisonOperator_LT           ComparisonOperator = "LT"
	ComparisonOperator_GE           ComparisonOperator = "GE"
	ComparisonOperator_GT           ComparisonOperator = "GT"
	ComparisonOperator_BETWEEN      ComparisonOperator = "BETWEEN"
	ComparisonOperator_NOT_NULL     ComparisonOperator = "NOT_NULL"
	ComparisonOperator_NULL         ComparisonOperator = "NULL"
	ComparisonOperator_CONTAINS     ComparisonOperator = "CONTAINS"
	ComparisonOperator_NOT_CONTAINS ComparisonOperator = "NOT_CONTAINS"
	ComparisonOperator_BEGINS_WITH  ComparisonOperator = "BEGINS_WITH"
)

type ConditionalOperator string

const (
	ConditionalOperator_AND ConditionalOperator = "AND"
	ConditionalOperator_OR  ConditionalOperator = "OR"
)

type ContinuousBackupsStatus string

const (
	ContinuousBackupsStatus_ENABLED  ContinuousBackupsStatus = "ENABLED"
	ContinuousBackupsStatus_DISABLED ContinuousBackupsStatus = "DISABLED"
)

type ContributorInsightsAction string

const (
	ContributorInsightsAction_ENABLE  ContributorInsightsAction = "ENABLE"
	ContributorInsightsAction_DISABLE ContributorInsightsAction = "DISABLE"
)

type ContributorInsightsStatus string

const (
	ContributorInsightsStatus_ENABLING  ContributorInsightsStatus = "ENABLING"
	ContributorInsightsStatus_ENABLED   ContributorInsightsStatus = "ENABLED"
	ContributorInsightsStatus_DISABLING ContributorInsightsStatus = "DISABLING"
	ContributorInsightsStatus_DISABLED  ContributorInsightsStatus = "DISABLED"
	ContributorInsightsStatus_FAILED    ContributorInsightsStatus = "FAILED"
)

type DestinationStatus string

const (
	DestinationStatus_ENABLING      DestinationStatus = "ENABLING"
	DestinationStatus_ACTIVE        DestinationStatus = "ACTIVE"
	DestinationStatus_DISABLING     DestinationStatus = "DISABLING"
	DestinationStatus_DISABLED      DestinationStatus = "DISABLED"
	DestinationStatus_ENABLE_FAILED DestinationStatus = "ENABLE_FAILED"
)

type ExportFormat string

const (
	ExportFormat_DYNAMODB_JSON ExportFormat = "DYNAMODB_JSON"
	ExportFormat_ION           ExportFormat = "ION"
)

type ExportStatus string

const (
	ExportStatus_IN_PROGRESS ExportStatus = "IN_PROGRESS"
	ExportStatus_COMPLETED   ExportStatus = "COMPLETED"
	ExportStatus_FAILED      ExportStatus = "FAILED"
)

type GlobalTableStatus_SDK string

const (
	GlobalTableStatus_SDK_CREATING GlobalTableStatus_SDK = "CREATING"
	GlobalTableStatus_SDK_ACTIVE   GlobalTableStatus_SDK = "ACTIVE"
	GlobalTableStatus_SDK_DELETING GlobalTableStatus_SDK = "DELETING"
	GlobalTableStatus_SDK_UPDATING GlobalTableStatus_SDK = "UPDATING"
)

type IndexStatus string

const (
	IndexStatus_CREATING IndexStatus = "CREATING"
	IndexStatus_UPDATING IndexStatus = "UPDATING"
	IndexStatus_DELETING IndexStatus = "DELETING"
	IndexStatus_ACTIVE   IndexStatus = "ACTIVE"
)

type KeyType string

const (
	KeyType_HASH  KeyType = "HASH"
	KeyType_RANGE KeyType = "RANGE"
)

type PointInTimeRecoveryStatus string

const (
	PointInTimeRecoveryStatus_ENABLED  PointInTimeRecoveryStatus = "ENABLED"
	PointInTimeRecoveryStatus_DISABLED PointInTimeRecoveryStatus = "DISABLED"
)

type ProjectionType string

const (
	ProjectionType_ALL       ProjectionType = "ALL"
	ProjectionType_KEYS_ONLY ProjectionType = "KEYS_ONLY"
	ProjectionType_INCLUDE   ProjectionType = "INCLUDE"
)

type ReplicaStatus string

const (
	ReplicaStatus_CREATING                            ReplicaStatus = "CREATING"
	ReplicaStatus_CREATION_FAILED                     ReplicaStatus = "CREATION_FAILED"
	ReplicaStatus_UPDATING                            ReplicaStatus = "UPDATING"
	ReplicaStatus_DELETING                            ReplicaStatus = "DELETING"
	ReplicaStatus_ACTIVE                              ReplicaStatus = "ACTIVE"
	ReplicaStatus_REGION_DISABLED                     ReplicaStatus = "REGION_DISABLED"
	ReplicaStatus_INACCESSIBLE_ENCRYPTION_CREDENTIALS ReplicaStatus = "INACCESSIBLE_ENCRYPTION_CREDENTIALS"
)

type ReturnConsumedCapacity string

const (
	ReturnConsumedCapacity_INDEXES ReturnConsumedCapacity = "INDEXES"
	ReturnConsumedCapacity_TOTAL   ReturnConsumedCapacity = "TOTAL"
	ReturnConsumedCapacity_NONE    ReturnConsumedCapacity = "NONE"
)

type ReturnItemCollectionMetrics string

const (
	ReturnItemCollectionMetrics_SIZE ReturnItemCollectionMetrics = "SIZE"
	ReturnItemCollectionMetrics_NONE ReturnItemCollectionMetrics = "NONE"
)

type ReturnValue string

const (
	ReturnValue_NONE        ReturnValue = "NONE"
	ReturnValue_ALL_OLD     ReturnValue = "ALL_OLD"
	ReturnValue_UPDATED_OLD ReturnValue = "UPDATED_OLD"
	ReturnValue_ALL_NEW     ReturnValue = "ALL_NEW"
	ReturnValue_UPDATED_NEW ReturnValue = "UPDATED_NEW"
)

type ReturnValuesOnConditionCheckFailure string

const (
	ReturnValuesOnConditionCheckFailure_ALL_OLD ReturnValuesOnConditionCheckFailure = "ALL_OLD"
	ReturnValuesOnConditionCheckFailure_NONE    ReturnValuesOnConditionCheckFailure = "NONE"
)

type S3SSEAlgorithm string

const (
	S3SSEAlgorithm_AES256 S3SSEAlgorithm = "AES256"
	S3SSEAlgorithm_KMS    S3SSEAlgorithm = "KMS"
)

type SSEStatus string

const (
	SSEStatus_ENABLING  SSEStatus = "ENABLING"
	SSEStatus_ENABLED   SSEStatus = "ENABLED"
	SSEStatus_DISABLING SSEStatus = "DISABLING"
	SSEStatus_DISABLED  SSEStatus = "DISABLED"
	SSEStatus_UPDATING  SSEStatus = "UPDATING"
)

type SSEType string

const (
	SSEType_AES256 SSEType = "AES256"
	SSEType_KMS    SSEType = "KMS"
)

type ScalarAttributeType string

const (
	ScalarAttributeType_S ScalarAttributeType = "S"
	ScalarAttributeType_N ScalarAttributeType = "N"
	ScalarAttributeType_B ScalarAttributeType = "B"
)

type Select string

const (
	Select_ALL_ATTRIBUTES           Select = "ALL_ATTRIBUTES"
	Select_ALL_PROJECTED_ATTRIBUTES Select = "ALL_PROJECTED_ATTRIBUTES"
	Select_SPECIFIC_ATTRIBUTES      Select = "SPECIFIC_ATTRIBUTES"
	Select_COUNT                    Select = "COUNT"
)

type StreamViewType string

const (
	StreamViewType_NEW_IMAGE          StreamViewType = "NEW_IMAGE"
	StreamViewType_OLD_IMAGE          StreamViewType = "OLD_IMAGE"
	StreamViewType_NEW_AND_OLD_IMAGES StreamViewType = "NEW_AND_OLD_IMAGES"
	StreamViewType_KEYS_ONLY          StreamViewType = "KEYS_ONLY"
)

type TableStatus_SDK string

const (
	TableStatus_SDK_CREATING                            TableStatus_SDK = "CREATING"
	TableStatus_SDK_UPDATING                            TableStatus_SDK = "UPDATING"
	TableStatus_SDK_DELETING                            TableStatus_SDK = "DELETING"
	TableStatus_SDK_ACTIVE                              TableStatus_SDK = "ACTIVE"
	TableStatus_SDK_INACCESSIBLE_ENCRYPTION_CREDENTIALS TableStatus_SDK = "INACCESSIBLE_ENCRYPTION_CREDENTIALS"
	TableStatus_SDK_ARCHIVING                           TableStatus_SDK = "ARCHIVING"
	TableStatus_SDK_ARCHIVED                            TableStatus_SDK = "ARCHIVED"
)

type TimeToLiveStatus string

const (
	TimeToLiveStatus_ENABLING  TimeToLiveStatus = "ENABLING"
	TimeToLiveStatus_DISABLING TimeToLiveStatus = "DISABLING"
	TimeToLiveStatus_ENABLED   TimeToLiveStatus = "ENABLED"
	TimeToLiveStatus_DISABLED  TimeToLiveStatus = "DISABLED"
)
