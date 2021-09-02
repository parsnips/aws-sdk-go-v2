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

// Lists the targets (thing groups) associated with a given Device Defender
// security profile. Requires permission to access the
// ListTargetsForSecurityProfile
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) ListTargetsForSecurityProfile(ctx context.Context, params *ListTargetsForSecurityProfileInput, optFns ...func(*Options)) (*ListTargetsForSecurityProfileOutput, error) {
	if params == nil {
		params = &ListTargetsForSecurityProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTargetsForSecurityProfile", params, optFns, c.addOperationListTargetsForSecurityProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTargetsForSecurityProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTargetsForSecurityProfileInput struct {

	// The security profile.
	//
	// This member is required.
	SecurityProfileName *string

	// The maximum number of results to return at one time.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTargetsForSecurityProfileOutput struct {

	// A token that can be used to retrieve the next set of results, or null if there
	// are no additional results.
	NextToken *string

	// The thing groups to which the security profile is attached.
	SecurityProfileTargets []types.SecurityProfileTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTargetsForSecurityProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTargetsForSecurityProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTargetsForSecurityProfile{}, middleware.After)
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
	if err = addOpListTargetsForSecurityProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTargetsForSecurityProfile(options.Region), middleware.Before); err != nil {
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

// ListTargetsForSecurityProfileAPIClient is a client that implements the
// ListTargetsForSecurityProfile operation.
type ListTargetsForSecurityProfileAPIClient interface {
	ListTargetsForSecurityProfile(context.Context, *ListTargetsForSecurityProfileInput, ...func(*Options)) (*ListTargetsForSecurityProfileOutput, error)
}

var _ ListTargetsForSecurityProfileAPIClient = (*Client)(nil)

// ListTargetsForSecurityProfilePaginatorOptions is the paginator options for
// ListTargetsForSecurityProfile
type ListTargetsForSecurityProfilePaginatorOptions struct {
	// The maximum number of results to return at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTargetsForSecurityProfilePaginator is a paginator for
// ListTargetsForSecurityProfile
type ListTargetsForSecurityProfilePaginator struct {
	options   ListTargetsForSecurityProfilePaginatorOptions
	client    ListTargetsForSecurityProfileAPIClient
	params    *ListTargetsForSecurityProfileInput
	nextToken *string
	firstPage bool
}

// NewListTargetsForSecurityProfilePaginator returns a new
// ListTargetsForSecurityProfilePaginator
func NewListTargetsForSecurityProfilePaginator(client ListTargetsForSecurityProfileAPIClient, params *ListTargetsForSecurityProfileInput, optFns ...func(*ListTargetsForSecurityProfilePaginatorOptions)) *ListTargetsForSecurityProfilePaginator {
	if params == nil {
		params = &ListTargetsForSecurityProfileInput{}
	}

	options := ListTargetsForSecurityProfilePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTargetsForSecurityProfilePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTargetsForSecurityProfilePaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListTargetsForSecurityProfile page.
func (p *ListTargetsForSecurityProfilePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTargetsForSecurityProfileOutput, error) {
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

	result, err := p.client.ListTargetsForSecurityProfile(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTargetsForSecurityProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListTargetsForSecurityProfile",
	}
}
