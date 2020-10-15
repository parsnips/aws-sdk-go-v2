// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Decreases the Kinesis data stream's retention period, which is the length of
// time data records are accessible after they are added to the stream. The minimum
// value of a stream's retention period is 24 hours. This operation may result in
// lost data. For example, if the stream's retention period is 48 hours and is
// decreased to 24 hours, any data already in the stream that is older than 24
// hours is inaccessible.
func (c *Client) DecreaseStreamRetentionPeriod(ctx context.Context, params *DecreaseStreamRetentionPeriodInput, optFns ...func(*Options)) (*DecreaseStreamRetentionPeriodOutput, error) {
	if params == nil {
		params = &DecreaseStreamRetentionPeriodInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DecreaseStreamRetentionPeriod", params, optFns, addOperationDecreaseStreamRetentionPeriodMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DecreaseStreamRetentionPeriodOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for DecreaseStreamRetentionPeriod.
type DecreaseStreamRetentionPeriodInput struct {

	// The new retention period of the stream, in hours. Must be less than the current
	// retention period.
	//
	// This member is required.
	RetentionPeriodHours *int32

	// The name of the stream to modify.
	//
	// This member is required.
	StreamName *string
}

type DecreaseStreamRetentionPeriodOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDecreaseStreamRetentionPeriodMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDecreaseStreamRetentionPeriod{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDecreaseStreamRetentionPeriod{}, middleware.After)
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
	addOpDecreaseStreamRetentionPeriodValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDecreaseStreamRetentionPeriod(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDecreaseStreamRetentionPeriod(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "DecreaseStreamRetentionPeriod",
	}
}
