// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates a node from an AWS OpsWorks CM server, and removes the node from
// the server's managed nodes. After a node is disassociated, the node key pair is
// no longer valid for accessing the configuration manager's API. For more
// information about how to associate a node, see AssociateNode. A node can can
// only be disassociated from a server that is in a HEALTHY state. Otherwise, an
// InvalidStateException is thrown. A ResourceNotFoundException is thrown when the
// server does not exist. A ValidationException is raised when parameters of the
// request are not valid.
func (c *Client) DisassociateNode(ctx context.Context, params *DisassociateNodeInput, optFns ...func(*Options)) (*DisassociateNodeOutput, error) {
	if params == nil {
		params = &DisassociateNodeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateNode", params, optFns, addOperationDisassociateNodeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateNodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateNodeInput struct {

	// The name of the client node.
	//
	// This member is required.
	NodeName *string

	// The name of the server from which to disassociate the node.
	//
	// This member is required.
	ServerName *string

	// Engine attributes that are used for disassociating the node. No attributes are
	// required for Puppet. Attributes required in a DisassociateNode request for
	// Chef
	//
	//     * CHEF_ORGANIZATION: The Chef organization with which the node was
	// associated. By default only one organization named default can exist.
	EngineAttributes []*types.EngineAttribute
}

type DisassociateNodeOutput struct {

	// Contains a token which can be passed to the DescribeNodeAssociationStatus API
	// call to get the status of the disassociation request.
	NodeAssociationStatusToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateNodeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateNode{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateNode{}, middleware.After)
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
	addOpDisassociateNodeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateNode(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisassociateNode(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "DisassociateNode",
	}
}
