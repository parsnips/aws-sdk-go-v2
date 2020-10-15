// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a user's Secure Shell (SSH) public key. No response is returned from
// this operation.
func (c *Client) DeleteSshPublicKey(ctx context.Context, params *DeleteSshPublicKeyInput, optFns ...func(*Options)) (*DeleteSshPublicKeyOutput, error) {
	if params == nil {
		params = &DeleteSshPublicKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSshPublicKey", params, optFns, addOperationDeleteSshPublicKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSshPublicKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSshPublicKeyInput struct {

	// A system-assigned unique identifier for a file transfer protocol-enabled server
	// instance that has the user assigned to it.
	//
	// This member is required.
	ServerId *string

	// A unique identifier used to reference your user's specific SSH key.
	//
	// This member is required.
	SshPublicKeyId *string

	// A unique string that identifies a user whose public key is being deleted.
	//
	// This member is required.
	UserName *string
}

type DeleteSshPublicKeyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteSshPublicKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteSshPublicKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteSshPublicKey{}, middleware.After)
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
	addOpDeleteSshPublicKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSshPublicKey(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteSshPublicKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "DeleteSshPublicKey",
	}
}
