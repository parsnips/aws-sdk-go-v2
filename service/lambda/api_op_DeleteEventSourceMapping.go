// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Deletes an event source mapping
// (https://docs.aws.amazon.com/lambda/latest/dg/intro-invocation-modes.html). You
// can get the identifier of a mapping from the output of ListEventSourceMappings.
// When you delete an event source mapping, it enters a Deleting state and might
// not be completely deleted for several seconds.
func (c *Client) DeleteEventSourceMapping(ctx context.Context, params *DeleteEventSourceMappingInput, optFns ...func(*Options)) (*DeleteEventSourceMappingOutput, error) {
	if params == nil {
		params = &DeleteEventSourceMappingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteEventSourceMapping", params, optFns, addOperationDeleteEventSourceMappingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteEventSourceMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteEventSourceMappingInput struct {

	// The identifier of the event source mapping.
	//
	// This member is required.
	UUID *string
}

// A mapping between an AWS resource and an AWS Lambda function. See
// CreateEventSourceMapping for details.
type DeleteEventSourceMappingOutput struct {

	// The maximum number of items to retrieve in a single batch.
	BatchSize *int32

	// (Streams) If the function returns an error, split the batch in two and retry.
	BisectBatchOnFunctionError *bool

	// (Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded
	// records.
	DestinationConfig *types.DestinationConfig

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

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteEventSourceMappingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteEventSourceMapping{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteEventSourceMapping{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteEventSourceMappingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEventSourceMapping(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteEventSourceMapping(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "DeleteEventSourceMapping",
	}
}
