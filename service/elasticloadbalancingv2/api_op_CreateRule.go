// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a rule for the specified listener. The listener must be associated with
// an Application Load Balancer. Rules are evaluated in priority order, from the
// lowest value to the highest value. When the conditions for a rule are met, its
// actions are performed. If the conditions for no rules are met, the actions for
// the default rule are performed. For more information, see Listener Rules
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#listener-rules)
// in the Application Load Balancers Guide. To view your current rules, use
// DescribeRules. To update a rule, use ModifyRule. To set the priorities of your
// rules, use SetRulePriorities. To delete a rule, use DeleteRule.
func (c *Client) CreateRule(ctx context.Context, params *CreateRuleInput, optFns ...func(*Options)) (*CreateRuleOutput, error) {
	if params == nil {
		params = &CreateRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRule", params, optFns, addOperationCreateRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRuleInput struct {

	// The actions. Each rule must include exactly one of the following types of
	// actions: forward, fixed-response, or redirect, and it must be the last action to
	// be performed. If the action type is forward, you specify one or more target
	// groups. The protocol of the target group must be HTTP or HTTPS for an
	// Application Load Balancer. The protocol of the target group must be TCP, TLS,
	// UDP, or TCP_UDP for a Network Load Balancer. [HTTPS listeners] If the action
	// type is authenticate-oidc, you authenticate users through an identity provider
	// that is OpenID Connect (OIDC) compliant. [HTTPS listeners] If the action type is
	// authenticate-cognito, you authenticate users through the user pools supported by
	// Amazon Cognito. [Application Load Balancer] If the action type is redirect, you
	// redirect specified client requests from one URL to another. [Application Load
	// Balancer] If the action type is fixed-response, you drop specified client
	// requests and return a custom HTTP response.
	//
	// This member is required.
	Actions []*types.Action

	// The conditions. Each rule can include zero or one of the following conditions:
	// http-request-method, host-header, path-pattern, and source-ip, and zero or more
	// of the following conditions: http-header and query-string.
	//
	// This member is required.
	Conditions []*types.RuleCondition

	// The Amazon Resource Name (ARN) of the listener.
	//
	// This member is required.
	ListenerArn *string

	// The rule priority. A listener can't have multiple rules with the same priority.
	//
	// This member is required.
	Priority *int32
}

type CreateRuleOutput struct {

	// Information about the rule.
	Rules []*types.Rule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateRule{}, middleware.After)
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
	addOpCreateRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRule(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "CreateRule",
	}
}
