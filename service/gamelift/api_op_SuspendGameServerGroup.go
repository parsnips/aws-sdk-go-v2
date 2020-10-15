// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This action is part of Amazon GameLift FleetIQ with game server groups, which is
// in preview release and is subject to change. Temporarily stops activity on a
// game server group without terminating instances or the game server group.
// Activity can be restarted by calling ResumeGameServerGroup. Activities that can
// suspended are:
//
//     * Instance type replacement. This activity evaluates the
// current Spot viability of all instance types that are defined for the game
// server group. It updates the Auto Scaling group to remove nonviable Spot
// instance types (which have a higher chance of game server interruptions) and
// rebalances capacity across the remaining viable Spot instance types. When this
// activity is suspended, the Auto Scaling group continues with its current
// balance, regardless of viability. Instance protection, utilization metrics, and
// capacity autoscaling activities continue to be active.
//
// To suspend activity,
// specify a game server group ARN and the type of activity to be suspended. Learn
// more GameLift FleetIQ Guide
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gsg-intro.html)
// Related operations
//
//     * CreateGameServerGroup
//
//     * ListGameServerGroups
//
//
// * DescribeGameServerGroup
//
//     * UpdateGameServerGroup
//
//     *
// DeleteGameServerGroup
//
//     * ResumeGameServerGroup
//
//     * SuspendGameServerGroup
func (c *Client) SuspendGameServerGroup(ctx context.Context, params *SuspendGameServerGroupInput, optFns ...func(*Options)) (*SuspendGameServerGroupOutput, error) {
	if params == nil {
		params = &SuspendGameServerGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SuspendGameServerGroup", params, optFns, addOperationSuspendGameServerGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SuspendGameServerGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SuspendGameServerGroupInput struct {

	// The unique identifier of the game server group to stop activity on. Use either
	// the GameServerGroup name or ARN value.
	//
	// This member is required.
	GameServerGroupName *string

	// The action to suspend for this game server group.
	//
	// This member is required.
	SuspendActions []types.GameServerGroupAction
}

type SuspendGameServerGroupOutput struct {

	// An object that describes the game server group resource, with the
	// SuspendedActions property updated to reflect the suspended activity.
	GameServerGroup *types.GameServerGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSuspendGameServerGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSuspendGameServerGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSuspendGameServerGroup{}, middleware.After)
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
	addOpSuspendGameServerGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSuspendGameServerGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSuspendGameServerGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "SuspendGameServerGroup",
	}
}
