// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified Kinesis data stream. The information returned includes
// the stream name, Amazon Resource Name (ARN), creation time, enhanced metric
// configuration, and shard map. The shard map is an array of shard objects. For
// each shard object, there is the hash key and sequence number ranges that the
// shard spans, and the IDs of any earlier shards that played in a role in creating
// the shard. Every record ingested in the stream is identified by a sequence
// number, which is assigned when the record is put into the stream. You can limit
// the number of shards returned by each call. For more information, see Retrieving
// Shards from a Stream
// (https://docs.aws.amazon.com/kinesis/latest/dev/kinesis-using-sdk-java-retrieve-shards.html)
// in the Amazon Kinesis Data Streams Developer Guide. There are no guarantees
// about the chronological order shards returned. To process shards in
// chronological order, use the ID of the parent shard to track the lineage to the
// oldest shard. This operation has a limit of 10 transactions per second per
// account.
func (c *Client) DescribeStream(ctx context.Context, params *DescribeStreamInput, optFns ...func(*Options)) (*DescribeStreamOutput, error) {
	if params == nil {
		params = &DescribeStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStream", params, optFns, addOperationDescribeStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for DescribeStream.
type DescribeStreamInput struct {

	// The name of the stream to describe.
	//
	// This member is required.
	StreamName *string

	// The shard ID of the shard to start with.
	ExclusiveStartShardId *string

	// The maximum number of shards to return in a single call. The default value is
	// 100. If you specify a value greater than 100, at most 100 shards are returned.
	Limit *int32
}

// Represents the output for DescribeStream.
type DescribeStreamOutput struct {

	// The current status of the stream, the stream Amazon Resource Name (ARN), an
	// array of shard objects that comprise the stream, and whether there are more
	// shards available.
	//
	// This member is required.
	StreamDescription *types.StreamDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStream{}, middleware.After)
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
	addOpDescribeStreamValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStream(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeStream(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "DescribeStream",
	}
}
