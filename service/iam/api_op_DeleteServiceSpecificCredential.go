// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified service-specific credential.
func (c *Client) DeleteServiceSpecificCredential(ctx context.Context, params *DeleteServiceSpecificCredentialInput, optFns ...func(*Options)) (*DeleteServiceSpecificCredentialOutput, error) {
	if params == nil {
		params = &DeleteServiceSpecificCredentialInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteServiceSpecificCredential", params, optFns, addOperationDeleteServiceSpecificCredentialMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteServiceSpecificCredentialOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteServiceSpecificCredentialInput struct {

	// The unique identifier of the service-specific credential. You can get this value
	// by calling ListServiceSpecificCredentials. This parameter allows (through its
	// regex pattern (http://wikipedia.org/wiki/regex)) a string of characters that can
	// consist of any upper or lowercased letter or digit.
	//
	// This member is required.
	ServiceSpecificCredentialId *string

	// The name of the IAM user associated with the service-specific credential. If
	// this value is not specified, then the operation assumes the user whose
	// credentials are used to call the operation. This parameter allows (through its
	// regex pattern (http://wikipedia.org/wiki/regex)) a string of characters
	// consisting of upper and lowercase alphanumeric characters with no spaces. You
	// can also include any of the following characters: _+=,.@-
	UserName *string
}

type DeleteServiceSpecificCredentialOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteServiceSpecificCredentialMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteServiceSpecificCredential{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteServiceSpecificCredential{}, middleware.After)
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
	addOpDeleteServiceSpecificCredentialValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteServiceSpecificCredential(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteServiceSpecificCredential(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "DeleteServiceSpecificCredential",
	}
}
