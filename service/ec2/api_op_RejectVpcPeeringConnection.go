// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Rejects a VPC peering connection request. The VPC peering connection must be in
// the pending-acceptance state. Use the DescribeVpcPeeringConnections request to
// view your outstanding VPC peering connection requests. To delete an active VPC
// peering connection, or to delete a VPC peering connection request that you
// initiated, use DeleteVpcPeeringConnection.
func (c *Client) RejectVpcPeeringConnection(ctx context.Context, params *RejectVpcPeeringConnectionInput, optFns ...func(*Options)) (*RejectVpcPeeringConnectionOutput, error) {
	if params == nil {
		params = &RejectVpcPeeringConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RejectVpcPeeringConnection", params, optFns, addOperationRejectVpcPeeringConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RejectVpcPeeringConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RejectVpcPeeringConnectionInput struct {

	// The ID of the VPC peering connection.
	//
	// This member is required.
	VpcPeeringConnectionId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type RejectVpcPeeringConnectionOutput struct {

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRejectVpcPeeringConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpRejectVpcPeeringConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpRejectVpcPeeringConnection{}, middleware.After)
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
	addOpRejectVpcPeeringConnectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRejectVpcPeeringConnection(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRejectVpcPeeringConnection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "RejectVpcPeeringConnection",
	}
}
