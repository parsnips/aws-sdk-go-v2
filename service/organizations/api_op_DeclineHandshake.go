// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Declines a handshake request. This sets the handshake state to DECLINED and
// effectively deactivates the request. This operation can be called only from the
// account that received the handshake. The originator of the handshake can use
// CancelHandshake instead. The originator can't reactivate a declined request, but
// can reinitiate the process with a new handshake request. After you decline a
// handshake, it continues to appear in the results of relevant APIs for only 30
// days. After that, it's deleted.
func (c *Client) DeclineHandshake(ctx context.Context, params *DeclineHandshakeInput, optFns ...func(*Options)) (*DeclineHandshakeOutput, error) {
	if params == nil {
		params = &DeclineHandshakeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeclineHandshake", params, optFns, addOperationDeclineHandshakeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeclineHandshakeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeclineHandshakeInput struct {

	// The unique identifier (ID) of the handshake that you want to decline. You can
	// get the ID from the ListHandshakesForAccount operation. The regex pattern
	// (http://wikipedia.org/wiki/regex) for handshake ID string requires "h-" followed
	// by from 8 to 32 lowercase letters or digits.
	//
	// This member is required.
	HandshakeId *string
}

type DeclineHandshakeOutput struct {

	// A structure that contains details about the declined handshake. The state is
	// updated to show the value DECLINED.
	Handshake *types.Handshake

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeclineHandshakeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeclineHandshake{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeclineHandshake{}, middleware.After)
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
	addOpDeclineHandshakeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeclineHandshake(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeclineHandshake(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DeclineHandshake",
	}
}
