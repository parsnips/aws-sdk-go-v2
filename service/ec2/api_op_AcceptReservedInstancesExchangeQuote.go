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

// Accepts the Convertible Reserved Instance exchange quote described in the
// GetReservedInstancesExchangeQuote call.
func (c *Client) AcceptReservedInstancesExchangeQuote(ctx context.Context, params *AcceptReservedInstancesExchangeQuoteInput, optFns ...func(*Options)) (*AcceptReservedInstancesExchangeQuoteOutput, error) {
	if params == nil {
		params = &AcceptReservedInstancesExchangeQuoteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AcceptReservedInstancesExchangeQuote", params, optFns, addOperationAcceptReservedInstancesExchangeQuoteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AcceptReservedInstancesExchangeQuoteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for accepting the quote.
type AcceptReservedInstancesExchangeQuoteInput struct {

	// The IDs of the Convertible Reserved Instances to exchange for another
	// Convertible Reserved Instance of the same or higher value.
	//
	// This member is required.
	ReservedInstanceIds []*string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The configuration of the target Convertible Reserved Instance to exchange for
	// your current Convertible Reserved Instances.
	TargetConfigurations []*types.TargetConfigurationRequest
}

// The result of the exchange and whether it was successful.
type AcceptReservedInstancesExchangeQuoteOutput struct {

	// The ID of the successful exchange.
	ExchangeId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAcceptReservedInstancesExchangeQuoteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpAcceptReservedInstancesExchangeQuote{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAcceptReservedInstancesExchangeQuote{}, middleware.After)
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
	addOpAcceptReservedInstancesExchangeQuoteValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptReservedInstancesExchangeQuote(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAcceptReservedInstancesExchangeQuote(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AcceptReservedInstancesExchangeQuote",
	}
}
