// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists model compilation jobs that satisfy various filters. To create a model
// compilation job, use CreateCompilationJob. To get information about a particular
// model compilation job you have created, use DescribeCompilationJob.
func (c *Client) ListCompilationJobs(ctx context.Context, params *ListCompilationJobsInput, optFns ...func(*Options)) (*ListCompilationJobsOutput, error) {
	if params == nil {
		params = &ListCompilationJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCompilationJobs", params, optFns, addOperationListCompilationJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCompilationJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCompilationJobsInput struct {

	// A filter that returns the model compilation jobs that were created after a
	// specified time.
	CreationTimeAfter *time.Time

	// A filter that returns the model compilation jobs that were created before a
	// specified time.
	CreationTimeBefore *time.Time

	// A filter that returns the model compilation jobs that were modified after a
	// specified time.
	LastModifiedTimeAfter *time.Time

	// A filter that returns the model compilation jobs that were modified before a
	// specified time.
	LastModifiedTimeBefore *time.Time

	// The maximum number of model compilation jobs to return in the response.
	MaxResults *int32

	// A filter that returns the model compilation jobs whose name contains a specified
	// string.
	NameContains *string

	// If the result of the previous ListCompilationJobs request was truncated, the
	// response includes a NextToken. To retrieve the next set of model compilation
	// jobs, use the token in the next request.
	NextToken *string

	// The field by which to sort results. The default is CreationTime.
	SortBy types.ListCompilationJobsSortBy

	// The sort order for results. The default is Ascending.
	SortOrder types.SortOrder

	// A filter that retrieves model compilation jobs with a specific
	// DescribeCompilationJobResponse$CompilationJobStatus status.
	StatusEquals types.CompilationJobStatus
}

type ListCompilationJobsOutput struct {

	// An array of CompilationJobSummary objects, each describing a model compilation
	// job.
	//
	// This member is required.
	CompilationJobSummaries []*types.CompilationJobSummary

	// If the response is truncated, Amazon SageMaker returns this NextToken. To
	// retrieve the next set of model compilation jobs, use this token in the next
	// request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListCompilationJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCompilationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCompilationJobs{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListCompilationJobs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListCompilationJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListCompilationJobs",
	}
}
