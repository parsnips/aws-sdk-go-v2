// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Represents a collection of DomainName resources.
func (c *Client) GetDomainNames(ctx context.Context, params *GetDomainNamesInput, optFns ...func(*Options)) (*GetDomainNamesOutput, error) {
	if params == nil {
		params = &GetDomainNamesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDomainNames", params, optFns, addOperationGetDomainNamesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDomainNamesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to describe a collection of DomainName resources.
type GetDomainNamesInput struct {

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32

	Name *string

	// The current pagination position in the paged result set.
	Position *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// Represents a collection of DomainName resources. Use Client-Side Certificate
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html)
type GetDomainNamesOutput struct {

	// The current page of elements from this collection.
	Items []*types.DomainName

	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDomainNamesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDomainNames{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDomainNames{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDomainNames(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetDomainNames(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetDomainNames",
	}
}
