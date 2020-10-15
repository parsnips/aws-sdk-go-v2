// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Use this operation to delete a workforce. If you want to create a new workforce
// in an AWS Region where the a workforce already exists, use this operation to
// delete the existing workforce and then use to create a new workforce.
func (c *Client) DeleteWorkforce(ctx context.Context, params *DeleteWorkforceInput, optFns ...func(*Options)) (*DeleteWorkforceOutput, error) {
	if params == nil {
		params = &DeleteWorkforceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteWorkforce", params, optFns, addOperationDeleteWorkforceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteWorkforceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteWorkforceInput struct {

	// The name of the workforce.
	//
	// This member is required.
	WorkforceName *string
}

type DeleteWorkforceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteWorkforceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteWorkforce{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteWorkforce{}, middleware.After)
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
	addOpDeleteWorkforceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteWorkforce(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteWorkforce(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DeleteWorkforce",
	}
}
