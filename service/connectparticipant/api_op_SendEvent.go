// Code generated by smithy-go-codegen DO NOT EDIT.

package connectparticipant

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sends an event. Note that ConnectionToken is used for invoking this API instead
// of ParticipantToken.
func (c *Client) SendEvent(ctx context.Context, params *SendEventInput, optFns ...func(*Options)) (*SendEventOutput, error) {
	if params == nil {
		params = &SendEventInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendEvent", params, optFns, addOperationSendEventMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendEventOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendEventInput struct {

	// The authentication token associated with the participant's connection.
	//
	// This member is required.
	ConnectionToken *string

	// The content type of the request. Supported types are:
	//
	//     *
	// application/vnd.amazonaws.connect.event.typing
	//
	//     *
	// application/vnd.amazonaws.connect.event.connection.acknowledged
	//
	// This member is required.
	ContentType *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// The content of the event to be sent (for example, message text). This is not yet
	// supported.
	Content *string
}

type SendEventOutput struct {

	// The time when the event was sent. It's specified in ISO 8601 format:
	// yyyy-MM-ddThh:mm:ss.SSSZ. For example, 2019-11-08T02:41:28.172Z.
	AbsoluteTime *string

	// The ID of the response.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSendEventMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSendEvent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSendEvent{}, middleware.After)
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
	addIdempotencyToken_opSendEventMiddleware(stack, options)
	addOpSendEventValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSendEvent(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpSendEvent struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpSendEvent) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpSendEvent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*SendEventInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *SendEventInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opSendEventMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpSendEvent{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opSendEvent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "SendEvent",
	}
}
