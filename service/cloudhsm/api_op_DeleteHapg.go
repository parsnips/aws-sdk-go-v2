// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is documentation for AWS CloudHSM Classic. For more information, see AWS
// CloudHSM Classic FAQs (http://aws.amazon.com/cloudhsm/faqs-classic/), the AWS
// CloudHSM Classic User Guide
// (https://docs.aws.amazon.com/cloudhsm/classic/userguide/), and the AWS CloudHSM
// Classic API Reference
// (https://docs.aws.amazon.com/cloudhsm/classic/APIReference/). For information
// about the current version of AWS CloudHSM, see AWS CloudHSM
// (http://aws.amazon.com/cloudhsm/), the AWS CloudHSM User Guide
// (https://docs.aws.amazon.com/cloudhsm/latest/userguide/), and the AWS CloudHSM
// API Reference (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/).
// Deletes a high-availability partition group.
func (c *Client) DeleteHapg(ctx context.Context, params *DeleteHapgInput, optFns ...func(*Options)) (*DeleteHapgOutput, error) {
	if params == nil {
		params = &DeleteHapgInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteHapg", params, optFns, addOperationDeleteHapgMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteHapgOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the DeleteHapg action.
type DeleteHapgInput struct {

	// The ARN of the high-availability partition group to delete.
	//
	// This member is required.
	HapgArn *string
}

// Contains the output of the DeleteHapg action.
type DeleteHapgOutput struct {

	// The status of the action.
	//
	// This member is required.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteHapgMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteHapg{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteHapg{}, middleware.After)
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
	addOpDeleteHapgValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteHapg(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteHapg(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudhsm",
		OperationName: "DeleteHapg",
	}
}
