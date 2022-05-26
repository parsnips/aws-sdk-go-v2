// Code generated by smithy-go-codegen DO NOT EDIT.

package emrserverless

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels a job run.
func (c *Client) CancelJobRun(ctx context.Context, params *CancelJobRunInput, optFns ...func(*Options)) (*CancelJobRunOutput, error) {
	if params == nil {
		params = &CancelJobRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelJobRun", params, optFns, c.addOperationCancelJobRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelJobRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelJobRunInput struct {

	// The ID of the application on which the job run will be canceled.
	//
	// This member is required.
	ApplicationId *string

	// The ID of the job run to cancel.
	//
	// This member is required.
	JobRunId *string

	noSmithyDocumentSerde
}

type CancelJobRunOutput struct {

	// The output contains the application ID on which the job run is cancelled.
	//
	// This member is required.
	ApplicationId *string

	// The output contains the ID of the cancelled job run.
	//
	// This member is required.
	JobRunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelJobRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelJobRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelJobRun{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCancelJobRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelJobRun(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCancelJobRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "emr-serverless",
		OperationName: "CancelJobRun",
	}
}
