// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of virtual tape recovery points that are available for the
// specified tape gateway. A recovery point is a point-in-time view of a virtual
// tape at which all the data on the virtual tape is consistent. If your gateway
// crashes, virtual tapes that have recovery points can be recovered to a new
// gateway. This operation is only supported in the tape gateway type.
func (c *Client) DescribeTapeRecoveryPoints(ctx context.Context, params *DescribeTapeRecoveryPointsInput, optFns ...func(*Options)) (*DescribeTapeRecoveryPointsOutput, error) {
	if params == nil {
		params = &DescribeTapeRecoveryPointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTapeRecoveryPoints", params, optFns, addOperationDescribeTapeRecoveryPointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTapeRecoveryPointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeTapeRecoveryPointsInput
type DescribeTapeRecoveryPointsInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string

	// Specifies that the number of virtual tape recovery points that are described be
	// limited to the specified number.
	Limit *int32

	// An opaque string that indicates the position at which to begin describing the
	// virtual tape recovery points.
	Marker *string
}

// DescribeTapeRecoveryPointsOutput
type DescribeTapeRecoveryPointsOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// An opaque string that indicates the position at which the virtual tape recovery
	// points that were listed for description ended. Use this marker in your next
	// request to list the next set of virtual tape recovery points in the list. If
	// there are no more recovery points to describe, this field does not appear in the
	// response.
	Marker *string

	// An array of TapeRecoveryPointInfos that are available for the specified gateway.
	TapeRecoveryPointInfos []*types.TapeRecoveryPointInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTapeRecoveryPointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTapeRecoveryPoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTapeRecoveryPoints{}, middleware.After)
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
	addOpDescribeTapeRecoveryPointsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTapeRecoveryPoints(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeTapeRecoveryPoints(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeTapeRecoveryPoints",
	}
}
