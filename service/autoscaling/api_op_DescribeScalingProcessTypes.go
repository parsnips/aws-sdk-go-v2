// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the scaling process types for use with the ResumeProcesses and
// SuspendProcesses APIs.
func (c *Client) DescribeScalingProcessTypes(ctx context.Context, params *DescribeScalingProcessTypesInput, optFns ...func(*Options)) (*DescribeScalingProcessTypesOutput, error) {
	if params == nil {
		params = &DescribeScalingProcessTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeScalingProcessTypes", params, optFns, addOperationDescribeScalingProcessTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeScalingProcessTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeScalingProcessTypesInput struct {
}

type DescribeScalingProcessTypesOutput struct {

	// The names of the process types.
	Processes []*types.ProcessType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeScalingProcessTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeScalingProcessTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeScalingProcessTypes{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeScalingProcessTypes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeScalingProcessTypes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeScalingProcessTypes",
	}
}
