// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the queue specified by the QueueUrl, regardless of the queue's contents.
// Be careful with the DeleteQueue action: When you delete a queue, any messages in
// the queue are no longer available. When you delete a queue, the deletion process
// takes up to 60 seconds. Requests you send involving that queue during the 60
// seconds might succeed. For example, a SendMessage request might succeed, but
// after 60 seconds the queue and the message you sent no longer exist. When you
// delete a queue, you must wait at least 60 seconds before creating a queue with
// the same name. Cross-account permissions don't apply to this action. For more
// information, see Grant Cross-Account Permissions to a Role and a User Name
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon Simple Queue Service Developer Guide.
func (c *Client) DeleteQueue(ctx context.Context, params *DeleteQueueInput, optFns ...func(*Options)) (*DeleteQueueOutput, error) {
	if params == nil {
		params = &DeleteQueueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteQueue", params, optFns, addOperationDeleteQueueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteQueueInput struct {

	// The URL of the Amazon SQS queue to delete. Queue URLs and names are
	// case-sensitive.
	//
	// This member is required.
	QueueUrl *string
}

type DeleteQueueOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteQueueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteQueue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteQueue{}, middleware.After)
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
	addOpDeleteQueueValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteQueue(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteQueue(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "DeleteQueue",
	}
}
