// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified trial component. A trial component must be disassociated
// from all trials before the trial component can be deleted. To disassociate a
// trial component from a trial, call the DisassociateTrialComponent API.
func (c *Client) DeleteTrialComponent(ctx context.Context, params *DeleteTrialComponentInput, optFns ...func(*Options)) (*DeleteTrialComponentOutput, error) {
	if params == nil {
		params = &DeleteTrialComponentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTrialComponent", params, optFns, addOperationDeleteTrialComponentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTrialComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTrialComponentInput struct {

	// The name of the component to delete.
	//
	// This member is required.
	TrialComponentName *string
}

type DeleteTrialComponentOutput struct {

	// The Amazon Resource Name (ARN) of the component is being deleted.
	TrialComponentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteTrialComponentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteTrialComponent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteTrialComponent{}, middleware.After)
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
	addOpDeleteTrialComponentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTrialComponent(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteTrialComponent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DeleteTrialComponent",
	}
}
