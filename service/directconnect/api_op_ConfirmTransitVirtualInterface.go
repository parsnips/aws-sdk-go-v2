// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Accepts ownership of a transit virtual interface created by another AWS account.
// After the owner of the transit virtual interface makes this call, the specified
// transit virtual interface is created and made available to handle traffic.
func (c *Client) ConfirmTransitVirtualInterface(ctx context.Context, params *ConfirmTransitVirtualInterfaceInput, optFns ...func(*Options)) (*ConfirmTransitVirtualInterfaceOutput, error) {
	if params == nil {
		params = &ConfirmTransitVirtualInterfaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ConfirmTransitVirtualInterface", params, optFns, addOperationConfirmTransitVirtualInterfaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ConfirmTransitVirtualInterfaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ConfirmTransitVirtualInterfaceInput struct {

	// The ID of the Direct Connect gateway.
	//
	// This member is required.
	DirectConnectGatewayId *string

	// The ID of the virtual interface.
	//
	// This member is required.
	VirtualInterfaceId *string
}

type ConfirmTransitVirtualInterfaceOutput struct {

	// The state of the virtual interface. The following are the possible values:
	//
	//
	// * confirming: The creation of the virtual interface is pending confirmation from
	// the virtual interface owner. If the owner of the virtual interface is different
	// from the owner of the connection on which it is provisioned, then the virtual
	// interface will remain in this state until it is confirmed by the virtual
	// interface owner.
	//
	//     * verifying: This state only applies to public virtual
	// interfaces. Each public virtual interface needs validation before the virtual
	// interface can be created.
	//
	//     * pending: A virtual interface is in this state
	// from the time that it is created until the virtual interface is ready to forward
	// traffic.
	//
	//     * available: A virtual interface that is able to forward
	// traffic.
	//
	//     * down: A virtual interface that is BGP down.
	//
	//     * deleting: A
	// virtual interface is in this state immediately after calling
	// DeleteVirtualInterface until it can no longer forward traffic.
	//
	//     * deleted: A
	// virtual interface that cannot forward traffic.
	//
	//     * rejected: The virtual
	// interface owner has declined creation of the virtual interface. If a virtual
	// interface in the Confirming state is deleted by the virtual interface owner, the
	// virtual interface enters the Rejected state.
	//
	//     * unknown: The state of the
	// virtual interface is not available.
	VirtualInterfaceState types.VirtualInterfaceState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationConfirmTransitVirtualInterfaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpConfirmTransitVirtualInterface{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpConfirmTransitVirtualInterface{}, middleware.After)
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
	addOpConfirmTransitVirtualInterfaceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opConfirmTransitVirtualInterface(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opConfirmTransitVirtualInterface(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "ConfirmTransitVirtualInterface",
	}
}
