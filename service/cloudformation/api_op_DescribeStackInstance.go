// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the stack instance that's associated with the specified stack set, AWS
// account, and Region. For a list of stack instances that are associated with a
// specific stack set, use ListStackInstances.
func (c *Client) DescribeStackInstance(ctx context.Context, params *DescribeStackInstanceInput, optFns ...func(*Options)) (*DescribeStackInstanceOutput, error) {
	if params == nil {
		params = &DescribeStackInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStackInstance", params, optFns, addOperationDescribeStackInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStackInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStackInstanceInput struct {

	// The ID of an AWS account that's associated with this stack instance.
	//
	// This member is required.
	StackInstanceAccount *string

	// The name of a Region that's associated with this stack instance.
	//
	// This member is required.
	StackInstanceRegion *string

	// The name or the unique stack ID of the stack set that you want to get stack
	// instance information for.
	//
	// This member is required.
	StackSetName *string
}

type DescribeStackInstanceOutput struct {

	// The stack instance that matches the specified request parameters.
	StackInstance *types.StackInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeStackInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeStackInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeStackInstance{}, middleware.After)
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
	addOpDescribeStackInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStackInstance(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeStackInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DescribeStackInstance",
	}
}
