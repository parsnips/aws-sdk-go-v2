// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates or updates a lifecycle hook for the specified Auto Scaling group. A
// lifecycle hook tells Amazon EC2 Auto Scaling to perform an action on an instance
// when the instance launches (before it is put into service) or as the instance
// terminates (before it is fully terminated). This step is a part of the procedure
// for adding a lifecycle hook to an Auto Scaling group:
//
//     * (Optional) Create a
// Lambda function and a rule that allows CloudWatch Events to invoke your Lambda
// function when Amazon EC2 Auto Scaling launches or terminates instances.
//
//     *
// (Optional) Create a notification target and an IAM role. The target can be
// either an Amazon SQS queue or an Amazon SNS topic. The role allows Amazon EC2
// Auto Scaling to publish lifecycle notifications to the target.
//
//     * Create the
// lifecycle hook. Specify whether the hook is used when the instances launch or
// terminate.
//
//     * If you need more time, record the lifecycle action heartbeat
// to keep the instance in a pending state using the RecordLifecycleActionHeartbeat
// API call.
//
//     * If you finish before the timeout period ends, complete the
// lifecycle action using the CompleteLifecycleAction API call.
//
// For more
// information, see Amazon EC2 Auto Scaling Lifecycle Hooks
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/lifecycle-hooks.html) in
// the Amazon EC2 Auto Scaling User Guide. If you exceed your maximum limit of
// lifecycle hooks, which by default is 50 per Auto Scaling group, the call fails.
// You can view the lifecycle hooks for an Auto Scaling group using the
// DescribeLifecycleHooks API call. If you are no longer using a lifecycle hook,
// you can delete it by calling the DeleteLifecycleHook API.
func (c *Client) PutLifecycleHook(ctx context.Context, params *PutLifecycleHookInput, optFns ...func(*Options)) (*PutLifecycleHookOutput, error) {
	if params == nil {
		params = &PutLifecycleHookInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutLifecycleHook", params, optFns, addOperationPutLifecycleHookMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutLifecycleHookOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLifecycleHookInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The name of the lifecycle hook.
	//
	// This member is required.
	LifecycleHookName *string

	// Defines the action the Auto Scaling group should take when the lifecycle hook
	// timeout elapses or if an unexpected failure occurs. This parameter can be either
	// CONTINUE or ABANDON. The default value is ABANDON.
	DefaultResult *string

	// The maximum time, in seconds, that can elapse before the lifecycle hook times
	// out. The range is from 30 to 7200 seconds. The default value is 3600 seconds (1
	// hour). If the lifecycle hook times out, Amazon EC2 Auto Scaling performs the
	// action that you specified in the DefaultResult parameter. You can prevent the
	// lifecycle hook from timing out by calling the RecordLifecycleActionHeartbeat
	// API.
	HeartbeatTimeout *int32

	// The instance state to which you want to attach the lifecycle hook. The valid
	// values are:
	//
	//     * autoscaling:EC2_INSTANCE_LAUNCHING
	//
	//     *
	// autoscaling:EC2_INSTANCE_TERMINATING
	//
	// Required for new lifecycle hooks, but
	// optional when updating existing hooks.
	LifecycleTransition *string

	// Additional information that you want to include any time Amazon EC2 Auto Scaling
	// sends a message to the notification target.
	NotificationMetadata *string

	// The ARN of the notification target that Amazon EC2 Auto Scaling uses to notify
	// you when an instance is in the transition state for the lifecycle hook. This
	// target can be either an SQS queue or an SNS topic. If you specify an empty
	// string, this overrides the current ARN. This operation uses the JSON format when
	// sending notifications to an Amazon SQS queue, and an email key-value pair format
	// when sending notifications to an Amazon SNS topic. When you specify a
	// notification target, Amazon EC2 Auto Scaling sends it a test message. Test
	// messages contain the following additional key-value pair: "Event":
	// "autoscaling:TEST_NOTIFICATION".
	NotificationTargetARN *string

	// The ARN of the IAM role that allows the Auto Scaling group to publish to the
	// specified notification target, for example, an Amazon SNS topic or an Amazon SQS
	// queue. Required for new lifecycle hooks, but optional when updating existing
	// hooks.
	RoleARN *string
}

type PutLifecycleHookOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutLifecycleHookMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutLifecycleHook{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutLifecycleHook{}, middleware.After)
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
	addOpPutLifecycleHookValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutLifecycleHook(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutLifecycleHook(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "PutLifecycleHook",
	}
}
