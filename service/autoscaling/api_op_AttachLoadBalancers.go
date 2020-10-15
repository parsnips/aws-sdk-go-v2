// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// To attach an Application Load Balancer or a Network Load Balancer, use the
// AttachLoadBalancerTargetGroups API operation instead. Attaches one or more
// Classic Load Balancers to the specified Auto Scaling group. Amazon EC2 Auto
// Scaling registers the running instances with these Classic Load Balancers. To
// describe the load balancers for an Auto Scaling group, call the
// DescribeLoadBalancers API. To detach the load balancer from the Auto Scaling
// group, call the DetachLoadBalancers API. For more information, see Attaching a
// Load Balancer to Your Auto Scaling Group
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/attach-load-balancer-asg.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) AttachLoadBalancers(ctx context.Context, params *AttachLoadBalancersInput, optFns ...func(*Options)) (*AttachLoadBalancersOutput, error) {
	if params == nil {
		params = &AttachLoadBalancersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachLoadBalancers", params, optFns, addOperationAttachLoadBalancersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachLoadBalancersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachLoadBalancersInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The names of the load balancers. You can specify up to 10 load balancers.
	//
	// This member is required.
	LoadBalancerNames []*string
}

type AttachLoadBalancersOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAttachLoadBalancersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAttachLoadBalancers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAttachLoadBalancers{}, middleware.After)
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
	addOpAttachLoadBalancersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAttachLoadBalancers(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAttachLoadBalancers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "AttachLoadBalancers",
	}
}
