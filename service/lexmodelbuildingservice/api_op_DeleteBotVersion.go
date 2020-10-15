// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a specific version of a bot. To delete all versions of a bot, use the
// DeleteBot operation. This operation requires permissions for the
// lex:DeleteBotVersion action.
func (c *Client) DeleteBotVersion(ctx context.Context, params *DeleteBotVersionInput, optFns ...func(*Options)) (*DeleteBotVersionOutput, error) {
	if params == nil {
		params = &DeleteBotVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBotVersion", params, optFns, addOperationDeleteBotVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBotVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBotVersionInput struct {

	// The name of the bot.
	//
	// This member is required.
	Name *string

	// The version of the bot to delete. You cannot delete the $LATEST version of the
	// bot. To delete the $LATEST version, use the DeleteBot operation.
	//
	// This member is required.
	Version *string
}

type DeleteBotVersionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteBotVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteBotVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteBotVersion{}, middleware.After)
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
	addOpDeleteBotVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBotVersion(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteBotVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DeleteBotVersion",
	}
}
