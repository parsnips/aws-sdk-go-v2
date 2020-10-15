// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Writes multiple data records into a delivery stream in a single call, which can
// achieve higher throughput per producer than when writing single records. To
// write single data records into a delivery stream, use PutRecord. Applications
// using these operations are referred to as producers. For information about
// service quota, see Amazon Kinesis Data Firehose Quota
// (https://docs.aws.amazon.com/firehose/latest/dev/limits.html). Each
// PutRecordBatch request supports up to 500 records. Each record in the request
// can be as large as 1,000 KB (before 64-bit encoding), up to a limit of 4 MB for
// the entire request. These limits cannot be changed. You must specify the name of
// the delivery stream and the data record when using PutRecord. The data record
// consists of a data blob that can be up to 1,000 KB in size, and any kind of
// data. For example, it could be a segment from a log file, geographic location
// data, website clickstream data, and so on. Kinesis Data Firehose buffers records
// before delivering them to the destination. To disambiguate the data blobs at the
// destination, a common solution is to use delimiters in the data, such as a
// newline (\n) or some other character unique within the data. This allows the
// consumer application to parse individual data items when reading the data from
// the destination. The PutRecordBatch response includes a count of failed records,
// FailedPutCount, and an array of responses, RequestResponses. Even if the
// PutRecordBatch call succeeds, the value of FailedPutCount may be greater than 0,
// indicating that there are records for which the operation didn't succeed. Each
// entry in the RequestResponses array provides additional information about the
// processed record. It directly correlates with a record in the request array
// using the same ordering, from the top to the bottom. The response array always
// includes the same number of records as the request array. RequestResponses
// includes both successfully and unsuccessfully processed records. Kinesis Data
// Firehose tries to process all records in each PutRecordBatch request. A single
// record failure does not stop the processing of subsequent records. A
// successfully processed record includes a RecordId value, which is unique for the
// record. An unsuccessfully processed record includes ErrorCode and ErrorMessage
// values. ErrorCode reflects the type of error, and is one of the following
// values: ServiceUnavailableException or InternalFailure. ErrorMessage provides
// more detailed information about the error. If there is an internal server error
// or a timeout, the write might have completed or it might have failed. If
// FailedPutCount is greater than 0, retry the request, resending only those
// records that might have failed processing. This minimizes the possible duplicate
// records and also reduces the total bytes sent (and corresponding charges). We
// recommend that you handle any duplicates at the destination. If PutRecordBatch
// throws ServiceUnavailableException, back off and retry. If the exception
// persists, it is possible that the throughput limits have been exceeded for the
// delivery stream. Data records sent to Kinesis Data Firehose are stored for 24
// hours from the time they are added to a delivery stream as it attempts to send
// the records to the destination. If the destination is unreachable for more than
// 24 hours, the data is no longer available. Don't concatenate two or more base64
// strings to form the data fields of your records. Instead, concatenate the raw
// data, then perform base64 encoding.
func (c *Client) PutRecordBatch(ctx context.Context, params *PutRecordBatchInput, optFns ...func(*Options)) (*PutRecordBatchOutput, error) {
	if params == nil {
		params = &PutRecordBatchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRecordBatch", params, optFns, addOperationPutRecordBatchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRecordBatchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRecordBatchInput struct {

	// The name of the delivery stream.
	//
	// This member is required.
	DeliveryStreamName *string

	// One or more records.
	//
	// This member is required.
	Records []*types.Record
}

type PutRecordBatchOutput struct {

	// The number of records that might have failed processing. This number might be
	// greater than 0 even if the PutRecordBatch call succeeds. Check FailedPutCount to
	// determine whether there are records that you need to resend.
	//
	// This member is required.
	FailedPutCount *int32

	// The results array. For each record, the index of the response element is the
	// same as the index used in the request array.
	//
	// This member is required.
	RequestResponses []*types.PutRecordBatchResponseEntry

	// Indicates whether server-side encryption (SSE) was enabled during this
	// operation.
	Encrypted *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutRecordBatchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutRecordBatch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutRecordBatch{}, middleware.After)
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
	addOpPutRecordBatchValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutRecordBatch(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutRecordBatch(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "firehose",
		OperationName: "PutRecordBatch",
	}
}
