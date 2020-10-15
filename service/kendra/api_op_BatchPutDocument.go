// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds one or more documents to an index. The BatchPutDocument operation enables
// you to ingest inline documents or a set of documents stored in an Amazon S3
// bucket. Use this operation to ingest your text and unstructured text into an
// index, add custom attributes to the documents, and to attach an access control
// list to the documents added to the index. The documents are indexed
// asynchronously. You can see the progress of the batch using AWS CloudWatch. Any
// error messages related to processing the batch are sent to your AWS CloudWatch
// log.
func (c *Client) BatchPutDocument(ctx context.Context, params *BatchPutDocumentInput, optFns ...func(*Options)) (*BatchPutDocumentOutput, error) {
	if params == nil {
		params = &BatchPutDocumentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchPutDocument", params, optFns, addOperationBatchPutDocumentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchPutDocumentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchPutDocumentInput struct {

	// One or more documents to add to the index. Documents have the following file
	// size limits.
	//
	//     * 5 MB total size for inline documents
	//
	//     * 50 MB total size
	// for files from an S3 bucket
	//
	//     * 5 MB extracted text for any file
	//
	// For more
	// information about file size and transaction per second quotas, see Quotas
	// (https://docs.aws.amazon.com/kendra/latest/dg/quotas.html).
	//
	// This member is required.
	Documents []*types.Document

	// The identifier of the index to add the documents to. You need to create the
	// index first using the CreateIndex operation.
	//
	// This member is required.
	IndexId *string

	// The Amazon Resource Name (ARN) of a role that is allowed to run the
	// BatchPutDocument operation. For more information, see IAM Roles for Amazon
	// Kendra (https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html).
	RoleArn *string
}

type BatchPutDocumentOutput struct {

	// A list of documents that were not added to the index because the document failed
	// a validation check. Each document contains an error message that indicates why
	// the document couldn't be added to the index. If there was an error adding a
	// document to an index the error is reported in your AWS CloudWatch log. For more
	// information, see Monitoring Amazon Kendra with Amazon CloudWatch Logs
	// (https://docs.aws.amazon.com/kendra/latest/dg/cloudwatch-logs.html)
	FailedDocuments []*types.BatchPutDocumentResponseFailedDocument

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchPutDocumentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchPutDocument{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchPutDocument{}, middleware.After)
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
	addOpBatchPutDocumentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchPutDocument(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchPutDocument(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "BatchPutDocument",
	}
}
