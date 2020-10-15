// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Limits that are related to concurrency and storage. All file and storage sizes
// are in bytes.
type AccountLimit struct {

	// The maximum size of a function's deployment package and layers when they're
	// extracted.
	CodeSizeUnzipped *int64

	// The maximum size of a deployment package when it's uploaded directly to AWS
	// Lambda. Use Amazon S3 for larger files.
	CodeSizeZipped *int64

	// The maximum number of simultaneous function executions.
	ConcurrentExecutions *int32

	// The amount of storage space that you can use for all deployment packages and
	// layer archives.
	TotalCodeSize *int64

	// The maximum number of simultaneous function executions, minus the capacity
	// that's reserved for individual functions with PutFunctionConcurrency.
	UnreservedConcurrentExecutions *int32
}

// The number of functions and amount of storage in use.
type AccountUsage struct {

	// The number of Lambda functions.
	FunctionCount *int64

	// The amount of storage space, in bytes, that's being used by deployment packages
	// and layer archives.
	TotalCodeSize *int64
}

// Provides configuration information about a Lambda function alias
// (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html).
type AliasConfiguration struct {

	// The Amazon Resource Name (ARN) of the alias.
	AliasArn *string

	// A description of the alias.
	Description *string

	// The function version that the alias invokes.
	FunctionVersion *string

	// The name of the alias.
	Name *string

	// A unique identifier that changes when you update the alias.
	RevisionId *string

	// The routing configuration
	// (https://docs.aws.amazon.com/lambda/latest/dg/lambda-traffic-shifting-using-aliases.html)
	// of the alias.
	RoutingConfig *AliasRoutingConfiguration
}

// The traffic-shifting
// (https://docs.aws.amazon.com/lambda/latest/dg/lambda-traffic-shifting-using-aliases.html)
// configuration of a Lambda function alias.
type AliasRoutingConfiguration struct {

	// The second version, and the percentage of traffic that's routed to it.
	AdditionalVersionWeights map[string]*float64
}

type Concurrency struct {

	// The number of concurrent executions that are reserved for this function. For
	// more information, see Managing Concurrency
	// (https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html).
	ReservedConcurrentExecutions *int32
}

// The dead-letter queue
// (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#dlq) for
// failed asynchronous invocations.
type DeadLetterConfig struct {

	// The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.
	TargetArn *string
}

// A configuration object that specifies the destination of an event after Lambda
// processes it.
type DestinationConfig struct {

	// The destination configuration for failed invocations.
	OnFailure *OnFailure

	// The destination configuration for successful invocations.
	OnSuccess *OnSuccess
}

// A function's environment variable settings.
type Environment struct {

	// Environment variable key-value pairs.
	Variables map[string]*string
}

// Error messages for environment variables that couldn't be applied.
type EnvironmentError struct {

	// The error code.
	ErrorCode *string

	// The error message.
	Message *string
}

// The results of an operation to update or read environment variables. If the
// operation is successful, the response contains the environment variables. If it
// failed, the response contains details about the error.
type EnvironmentResponse struct {

	// Error messages for environment variables that couldn't be applied.
	Error *EnvironmentError

	// Environment variable key-value pairs.
	Variables map[string]*string
}

// A mapping between an AWS resource and an AWS Lambda function. See
// CreateEventSourceMapping for details.
type EventSourceMappingConfiguration struct {

	// The maximum number of items to retrieve in a single batch.
	BatchSize *int32

	// (Streams) If the function returns an error, split the batch in two and retry.
	BisectBatchOnFunctionError *bool

	// (Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded
	// records.
	DestinationConfig *DestinationConfig

	// The Amazon Resource Name (ARN) of the event source.
	EventSourceArn *string

	// The ARN of the Lambda function.
	FunctionArn *string

	// The date that the event source mapping was last updated, or its state changed.
	LastModified *time.Time

	// The result of the last AWS Lambda invocation of your Lambda function.
	LastProcessingResult *string

	// (Streams) The maximum amount of time to gather records before invoking the
	// function, in seconds.
	MaximumBatchingWindowInSeconds *int32

	// (Streams) The maximum age of a record that Lambda sends to a function for
	// processing.
	MaximumRecordAgeInSeconds *int32

	// (Streams) The maximum number of times to retry when the function returns an
	// error.
	MaximumRetryAttempts *int32

	// (Streams) The number of batches to process from each shard concurrently.
	ParallelizationFactor *int32

	// The state of the event source mapping. It can be one of the following: Creating,
	// Enabling, Enabled, Disabling, Disabled, Updating, or Deleting.
	State *string

	// Indicates whether the last change to the event source mapping was made by a
	// user, or by the Lambda service.
	StateTransitionReason *string

	// The identifier of the event source mapping.
	UUID *string
}

// Details about the connection between a Lambda function and an Amazon EFS file
// system.
type FileSystemConfig struct {

	// The Amazon Resource Name (ARN) of the Amazon EFS access point that provides
	// access to the file system.
	//
	// This member is required.
	Arn *string

	// The path where the function can access the file system, starting with /mnt/.
	//
	// This member is required.
	LocalMountPath *string
}

// The code for the Lambda function. You can specify either an object in Amazon S3,
// or upload a deployment package directly.
type FunctionCode struct {

	// An Amazon S3 bucket in the same AWS Region as your function. The bucket can be
	// in a different AWS account.
	S3Bucket *string

	// The Amazon S3 key of the deployment package.
	S3Key *string

	// For versioned objects, the version of the deployment package object to use.
	S3ObjectVersion *string

	// The base64-encoded contents of the deployment package. AWS SDK and AWS CLI
	// clients handle the encoding for you.
	ZipFile []byte
}

// Details about a function's deployment package.
type FunctionCodeLocation struct {

	// A presigned URL that you can use to download the deployment package.
	Location *string

	// The service that's hosting the file.
	RepositoryType *string
}

// Details about a function's configuration.
type FunctionConfiguration struct {

	// The SHA256 hash of the function's deployment package.
	CodeSha256 *string

	// The size of the function's deployment package, in bytes.
	CodeSize *int64

	// The function's dead letter queue.
	DeadLetterConfig *DeadLetterConfig

	// The function's description.
	Description *string

	// The function's environment variables.
	Environment *EnvironmentResponse

	// Connection settings for an Amazon EFS file system.
	FileSystemConfigs []*FileSystemConfig

	// The function's Amazon Resource Name (ARN).
	FunctionArn *string

	// The name of the function.
	FunctionName *string

	// The function that Lambda calls to begin executing your function.
	Handler *string

	// The KMS key that's used to encrypt the function's environment variables. This
	// key is only returned if you've configured a customer managed CMK.
	KMSKeyArn *string

	// The date and time that the function was last updated, in ISO-8601 format
	// (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
	LastModified *string

	// The status of the last update that was performed on the function. This is first
	// set to Successful after function creation completes.
	LastUpdateStatus LastUpdateStatus

	// The reason for the last update that was performed on the function.
	LastUpdateStatusReason *string

	// The reason code for the last update that was performed on the function.
	LastUpdateStatusReasonCode LastUpdateStatusReasonCode

	// The function's  layers
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
	Layers []*Layer

	// For Lambda@Edge functions, the ARN of the master function.
	MasterArn *string

	// The memory that's allocated to the function.
	MemorySize *int32

	// The latest updated revision of the function or alias.
	RevisionId *string

	// The function's execution role.
	Role *string

	// The runtime environment for the Lambda function.
	Runtime Runtime

	// The current state of the function. When the state is Inactive, you can
	// reactivate the function by invoking it.
	State State

	// The reason for the function's current state.
	StateReason *string

	// The reason code for the function's current state. When the code is Creating, you
	// can't invoke or modify the function.
	StateReasonCode StateReasonCode

	// The amount of time in seconds that Lambda allows a function to run before
	// stopping it.
	Timeout *int32

	// The function's AWS X-Ray tracing configuration.
	TracingConfig *TracingConfigResponse

	// The version of the Lambda function.
	Version *string

	// The function's networking configuration.
	VpcConfig *VpcConfigResponse
}

type FunctionEventInvokeConfig struct {

	// A destination for events after they have been sent to a function for processing.
	// Destinations
	//
	//     * Function - The Amazon Resource Name (ARN) of a Lambda
	// function.
	//
	//     * Queue - The ARN of an SQS queue.
	//
	//     * Topic - The ARN of an
	// SNS topic.
	//
	//     * Event Bus - The ARN of an Amazon EventBridge event bus.
	DestinationConfig *DestinationConfig

	// The Amazon Resource Name (ARN) of the function.
	FunctionArn *string

	// The date and time that the configuration was last updated.
	LastModified *time.Time

	// The maximum age of a request that Lambda sends to a function for processing.
	MaximumEventAgeInSeconds *int32

	// The maximum number of times to retry when the function returns an error.
	MaximumRetryAttempts *int32
}

// An AWS Lambda layer
// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
type Layer struct {

	// The Amazon Resource Name (ARN) of the function layer.
	Arn *string

	// The size of the layer archive in bytes.
	CodeSize *int64
}

// Details about an AWS Lambda layer
// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
type LayersListItem struct {

	// The newest version of the layer.
	LatestMatchingVersion *LayerVersionsListItem

	// The Amazon Resource Name (ARN) of the function layer.
	LayerArn *string

	// The name of the layer.
	LayerName *string
}

// A ZIP archive that contains the contents of an AWS Lambda layer
// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html). You
// can specify either an Amazon S3 location, or upload a layer archive directly.
type LayerVersionContentInput struct {

	// The Amazon S3 bucket of the layer archive.
	S3Bucket *string

	// The Amazon S3 key of the layer archive.
	S3Key *string

	// For versioned objects, the version of the layer archive object to use.
	S3ObjectVersion *string

	// The base64-encoded contents of the layer archive. AWS SDK and AWS CLI clients
	// handle the encoding for you.
	ZipFile []byte
}

// Details about a version of an AWS Lambda layer
// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
type LayerVersionContentOutput struct {

	// The SHA-256 hash of the layer archive.
	CodeSha256 *string

	// The size of the layer archive in bytes.
	CodeSize *int64

	// A link to the layer archive in Amazon S3 that is valid for 10 minutes.
	Location *string
}

// Details about a version of an AWS Lambda layer
// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
type LayerVersionsListItem struct {

	// The layer's compatible runtimes.
	CompatibleRuntimes []Runtime

	// The date that the version was created, in ISO 8601 format. For example,
	// 2018-11-27T15:10:45.123+0000.
	CreatedDate *string

	// The description of the version.
	Description *string

	// The ARN of the layer version.
	LayerVersionArn *string

	// The layer's open-source license.
	LicenseInfo *string

	// The version number.
	Version *int64
}

// A destination for events that failed processing.
type OnFailure struct {

	// The Amazon Resource Name (ARN) of the destination resource.
	Destination *string
}

// A destination for events that were processed successfully.
type OnSuccess struct {

	// The Amazon Resource Name (ARN) of the destination resource.
	Destination *string
}

// Details about the provisioned concurrency configuration for a function alias or
// version.
type ProvisionedConcurrencyConfigListItem struct {

	// The amount of provisioned concurrency allocated.
	AllocatedProvisionedConcurrentExecutions *int32

	// The amount of provisioned concurrency available.
	AvailableProvisionedConcurrentExecutions *int32

	// The Amazon Resource Name (ARN) of the alias or version.
	FunctionArn *string

	// The date and time that a user last updated the configuration, in ISO 8601 format
	// (https://www.iso.org/iso-8601-date-and-time-format.html).
	LastModified *string

	// The amount of provisioned concurrency requested.
	RequestedProvisionedConcurrentExecutions *int32

	// The status of the allocation process.
	Status ProvisionedConcurrencyStatusEnum

	// For failed allocations, the reason that provisioned concurrency could not be
	// allocated.
	StatusReason *string
}

// The function's AWS X-Ray tracing configuration. To sample and record incoming
// requests, set Mode to Active.
type TracingConfig struct {

	// The tracing mode.
	Mode TracingMode
}

// The function's AWS X-Ray tracing configuration.
type TracingConfigResponse struct {

	// The tracing mode.
	Mode TracingMode
}

// The VPC security groups and subnets that are attached to a Lambda function. For
// more information, see VPC Settings
// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html).
type VpcConfig struct {

	// A list of VPC security groups IDs.
	SecurityGroupIds []*string

	// A list of VPC subnet IDs.
	SubnetIds []*string
}

// The VPC security groups and subnets that are attached to a Lambda function.
type VpcConfigResponse struct {

	// A list of VPC security groups IDs.
	SecurityGroupIds []*string

	// A list of VPC subnet IDs.
	SubnetIds []*string

	// The ID of the VPC.
	VpcId *string
}
