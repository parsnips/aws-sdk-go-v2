// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the status of audit mitigation action tasks that were executed. Requires
// permission to access the ListAuditMitigationActionsExecutions
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) ListAuditMitigationActionsExecutions(ctx context.Context, params *ListAuditMitigationActionsExecutionsInput, optFns ...func(*Options)) (*ListAuditMitigationActionsExecutionsOutput, error) {
	if params == nil {
		params = &ListAuditMitigationActionsExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAuditMitigationActionsExecutions", params, optFns, c.addOperationListAuditMitigationActionsExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAuditMitigationActionsExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAuditMitigationActionsExecutionsInput struct {

	// Specify this filter to limit results to those that were applied to a specific
	// audit finding.
	//
	// This member is required.
	FindingId *string

	// Specify this filter to limit results to actions for a specific audit mitigation
	// actions task.
	//
	// This member is required.
	TaskId *string

	// Specify this filter to limit results to those with a specific status.
	ActionStatus types.AuditMitigationActionsExecutionStatus

	// The maximum number of results to return at one time. The default is 25.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAuditMitigationActionsExecutionsOutput struct {

	// A set of task execution results based on the input parameters. Details include
	// the mitigation action applied, start time, and task status.
	ActionsExecutions []types.AuditMitigationActionExecutionMetadata

	// The token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAuditMitigationActionsExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAuditMitigationActionsExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAuditMitigationActionsExecutions{}, middleware.After)
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
	if err = addOpListAuditMitigationActionsExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAuditMitigationActionsExecutions(options.Region), middleware.Before); err != nil {
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

// ListAuditMitigationActionsExecutionsAPIClient is a client that implements the
// ListAuditMitigationActionsExecutions operation.
type ListAuditMitigationActionsExecutionsAPIClient interface {
	ListAuditMitigationActionsExecutions(context.Context, *ListAuditMitigationActionsExecutionsInput, ...func(*Options)) (*ListAuditMitigationActionsExecutionsOutput, error)
}

var _ ListAuditMitigationActionsExecutionsAPIClient = (*Client)(nil)

// ListAuditMitigationActionsExecutionsPaginatorOptions is the paginator options
// for ListAuditMitigationActionsExecutions
type ListAuditMitigationActionsExecutionsPaginatorOptions struct {
	// The maximum number of results to return at one time. The default is 25.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAuditMitigationActionsExecutionsPaginator is a paginator for
// ListAuditMitigationActionsExecutions
type ListAuditMitigationActionsExecutionsPaginator struct {
	options   ListAuditMitigationActionsExecutionsPaginatorOptions
	client    ListAuditMitigationActionsExecutionsAPIClient
	params    *ListAuditMitigationActionsExecutionsInput
	nextToken *string
	firstPage bool
}

// NewListAuditMitigationActionsExecutionsPaginator returns a new
// ListAuditMitigationActionsExecutionsPaginator
func NewListAuditMitigationActionsExecutionsPaginator(client ListAuditMitigationActionsExecutionsAPIClient, params *ListAuditMitigationActionsExecutionsInput, optFns ...func(*ListAuditMitigationActionsExecutionsPaginatorOptions)) *ListAuditMitigationActionsExecutionsPaginator {
	if params == nil {
		params = &ListAuditMitigationActionsExecutionsInput{}
	}

	options := ListAuditMitigationActionsExecutionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAuditMitigationActionsExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAuditMitigationActionsExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListAuditMitigationActionsExecutions page.
func (p *ListAuditMitigationActionsExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAuditMitigationActionsExecutionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListAuditMitigationActionsExecutions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListAuditMitigationActionsExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListAuditMitigationActionsExecutions",
	}
}
