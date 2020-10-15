// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Resets an attribute of an instance to its default value. To reset the kernel or
// ramdisk, the instance must be in a stopped state. To reset the sourceDestCheck,
// the instance can be either running or stopped. The sourceDestCheck attribute
// controls whether source/destination checking is enabled. The default value is
// true, which means checking is enabled. This value must be false for a NAT
// instance to perform NAT. For more information, see NAT Instances
// (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_NAT_Instance.html)
// in the Amazon Virtual Private Cloud User Guide.
func (c *Client) ResetInstanceAttribute(ctx context.Context, params *ResetInstanceAttributeInput, optFns ...func(*Options)) (*ResetInstanceAttributeOutput, error) {
	if params == nil {
		params = &ResetInstanceAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResetInstanceAttribute", params, optFns, addOperationResetInstanceAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResetInstanceAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetInstanceAttributeInput struct {

	// The attribute to reset. You can only reset the following attributes: kernel |
	// ramdisk | sourceDestCheck. To change an instance attribute, use
	// ModifyInstanceAttribute.
	//
	// This member is required.
	Attribute types.InstanceAttributeName

	// The ID of the instance.
	//
	// This member is required.
	InstanceId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type ResetInstanceAttributeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationResetInstanceAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpResetInstanceAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpResetInstanceAttribute{}, middleware.After)
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
	addOpResetInstanceAttributeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResetInstanceAttribute(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opResetInstanceAttribute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ResetInstanceAttribute",
	}
}
