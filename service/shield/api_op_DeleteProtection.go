// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an AWS Shield Advanced Protection.
func (c *Client) DeleteProtection(ctx context.Context, params *DeleteProtectionInput, optFns ...func(*Options)) (*DeleteProtectionOutput, error) {
	if params == nil {
		params = &DeleteProtectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteProtection", params, optFns, addOperationDeleteProtectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteProtectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteProtectionInput struct {

	// The unique identifier (ID) for the Protection object to be deleted.
	//
	// This member is required.
	ProtectionId *string
}

type DeleteProtectionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteProtectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteProtection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteProtection{}, middleware.After)
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
	addOpDeleteProtectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteProtection(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteProtection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "DeleteProtection",
	}
}
