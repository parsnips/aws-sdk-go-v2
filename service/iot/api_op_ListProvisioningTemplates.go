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

// Lists the fleet provisioning templates in your Amazon Web Services account.
// Requires permission to access the ListProvisioningTemplates
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) ListProvisioningTemplates(ctx context.Context, params *ListProvisioningTemplatesInput, optFns ...func(*Options)) (*ListProvisioningTemplatesOutput, error) {
	if params == nil {
		params = &ListProvisioningTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProvisioningTemplates", params, optFns, c.addOperationListProvisioningTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProvisioningTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProvisioningTemplatesInput struct {

	// The maximum number of results to return at one time.
	MaxResults *int32

	// A token to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListProvisioningTemplatesOutput struct {

	// A token to retrieve the next set of results.
	NextToken *string

	// A list of fleet provisioning templates
	Templates []types.ProvisioningTemplateSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProvisioningTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProvisioningTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProvisioningTemplates{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProvisioningTemplates(options.Region), middleware.Before); err != nil {
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

// ListProvisioningTemplatesAPIClient is a client that implements the
// ListProvisioningTemplates operation.
type ListProvisioningTemplatesAPIClient interface {
	ListProvisioningTemplates(context.Context, *ListProvisioningTemplatesInput, ...func(*Options)) (*ListProvisioningTemplatesOutput, error)
}

var _ ListProvisioningTemplatesAPIClient = (*Client)(nil)

// ListProvisioningTemplatesPaginatorOptions is the paginator options for
// ListProvisioningTemplates
type ListProvisioningTemplatesPaginatorOptions struct {
	// The maximum number of results to return at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProvisioningTemplatesPaginator is a paginator for ListProvisioningTemplates
type ListProvisioningTemplatesPaginator struct {
	options   ListProvisioningTemplatesPaginatorOptions
	client    ListProvisioningTemplatesAPIClient
	params    *ListProvisioningTemplatesInput
	nextToken *string
	firstPage bool
}

// NewListProvisioningTemplatesPaginator returns a new
// ListProvisioningTemplatesPaginator
func NewListProvisioningTemplatesPaginator(client ListProvisioningTemplatesAPIClient, params *ListProvisioningTemplatesInput, optFns ...func(*ListProvisioningTemplatesPaginatorOptions)) *ListProvisioningTemplatesPaginator {
	if params == nil {
		params = &ListProvisioningTemplatesInput{}
	}

	options := ListProvisioningTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProvisioningTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProvisioningTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListProvisioningTemplates page.
func (p *ListProvisioningTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProvisioningTemplatesOutput, error) {
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

	result, err := p.client.ListProvisioningTemplates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProvisioningTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListProvisioningTemplates",
	}
}
