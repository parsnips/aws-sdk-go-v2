// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Move a dedicated IP address to an existing dedicated IP pool. The dedicated IP
// address that you specify must already exist, and must be associated with your
// AWS account. The dedicated IP pool you specify must already exist. You can
// create a new pool by using the CreateDedicatedIpPool operation.
func (c *Client) PutDedicatedIpInPool(ctx context.Context, params *PutDedicatedIpInPoolInput, optFns ...func(*Options)) (*PutDedicatedIpInPoolOutput, error) {
	if params == nil {
		params = &PutDedicatedIpInPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutDedicatedIpInPool", params, optFns, addOperationPutDedicatedIpInPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutDedicatedIpInPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to move a dedicated IP address to a dedicated IP pool.
type PutDedicatedIpInPoolInput struct {

	// The name of the IP pool that you want to add the dedicated IP address to. You
	// have to specify an IP pool that already exists.
	//
	// This member is required.
	DestinationPoolName *string

	// The IP address that you want to move to the dedicated IP pool. The value you
	// specify has to be a dedicated IP address that's associated with your AWS
	// account.
	//
	// This member is required.
	Ip *string
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type PutDedicatedIpInPoolOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutDedicatedIpInPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutDedicatedIpInPool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutDedicatedIpInPool{}, middleware.After)
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
	addOpPutDedicatedIpInPoolValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutDedicatedIpInPool(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutDedicatedIpInPool(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutDedicatedIpInPool",
	}
}
