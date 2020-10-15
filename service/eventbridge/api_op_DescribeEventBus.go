// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Displays details about an event bus in your account. This can include the
// external AWS accounts that are permitted to write events to your default event
// bus, and the associated policy. For custom event buses and partner event buses,
// it displays the name, ARN, policy, state, and creation time. To enable your
// account to receive events from other accounts on its default event bus, use
// PutPermission. For more information about partner event buses, see
// CreateEventBus.
func (c *Client) DescribeEventBus(ctx context.Context, params *DescribeEventBusInput, optFns ...func(*Options)) (*DescribeEventBusOutput, error) {
	if params == nil {
		params = &DescribeEventBusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventBus", params, optFns, addOperationDescribeEventBusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventBusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventBusInput struct {

	// The name of the event bus to show details for. If you omit this, the default
	// event bus is displayed.
	Name *string
}

type DescribeEventBusOutput struct {

	// The Amazon Resource Name (ARN) of the account permitted to write events to the
	// current account.
	Arn *string

	// The name of the event bus. Currently, this is always default.
	Name *string

	// The policy that enables the external account to send events to your account.
	Policy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEventBusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEventBus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEventBus{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventBus(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeEventBus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "DescribeEventBus",
	}
}
