// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of your queues that have the RedrivePolicy queue attribute
// configured with a dead-letter queue. For more information about using
// dead-letter queues, see Using Amazon SQS Dead-Letter Queues
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html)
// in the Amazon Simple Queue Service Developer Guide.
func (c *Client) ListDeadLetterSourceQueues(ctx context.Context, params *ListDeadLetterSourceQueuesInput, optFns ...func(*Options)) (*ListDeadLetterSourceQueuesOutput, error) {
	if params == nil {
		params = &ListDeadLetterSourceQueuesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeadLetterSourceQueues", params, optFns, addOperationListDeadLetterSourceQueuesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeadLetterSourceQueuesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ListDeadLetterSourceQueuesInput struct {

	// The URL of a dead-letter queue. Queue URLs and names are case-sensitive.
	//
	// This member is required.
	QueueUrl *string

	// Maximum number of results to include in the response.
	MaxResults *int32

	// Pagination token to request the next set of results.
	NextToken *string
}

// A list of your dead letter source queues.
type ListDeadLetterSourceQueuesOutput struct {

	// A list of source queue URLs that have the RedrivePolicy queue attribute
	// configured with a dead-letter queue.
	//
	// This member is required.
	QueueUrls []*string

	// Pagination token to include in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDeadLetterSourceQueuesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListDeadLetterSourceQueues{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListDeadLetterSourceQueues{}, middleware.After)
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
	addOpListDeadLetterSourceQueuesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDeadLetterSourceQueues(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListDeadLetterSourceQueues(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "ListDeadLetterSourceQueues",
	}
}
