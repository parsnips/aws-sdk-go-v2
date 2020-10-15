// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// A feature of the API Gateway control service for creating a new API from an
// external API definition file.
func (c *Client) ImportRestApi(ctx context.Context, params *ImportRestApiInput, optFns ...func(*Options)) (*ImportRestApiOutput, error) {
	if params == nil {
		params = &ImportRestApiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportRestApi", params, optFns, addOperationImportRestApiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportRestApiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A POST request to import an API to API Gateway using an input of an API
// definition file.
type ImportRestApiInput struct {

	// A query parameter to indicate whether to rollback the API creation (true) or not
	// (false) when a warning is encountered. The default value is false.
	FailOnWarnings *bool

	Name *string

	// A key-value map of context-specific query string parameters specifying the
	// behavior of different API importing operations. The following shows
	// operation-specific parameters and their supported values. To exclude
	// DocumentationParts from the import, set parameters as ignore=documentation. To
	// configure the endpoint type, set parameters as endpointConfigurationTypes=EDGE,
	// endpointConfigurationTypes=REGIONAL, or endpointConfigurationTypes=PRIVATE. The
	// default endpoint type is EDGE. To handle imported basepath, set parameters as
	// basepath=ignore, basepath=prepend or basepath=split. For example, the AWS CLI
	// command to exclude documentation from the imported API is: aws apigateway
	// import-rest-api --parameters ignore=documentation --body
	// 'file:///path/to/imported-api-body.json' The AWS CLI command to set the regional
	// endpoint on the imported API is: aws apigateway import-rest-api --parameters
	// endpointConfigurationTypes=REGIONAL --body
	// 'file:///path/to/imported-api-body.json'
	Parameters map[string]*string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// Represents a REST API. Create an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
type ImportRestApiOutput struct {

	// The source of the API key for metering requests according to a usage plan. Valid
	// values are:
	//
	//     * HEADER to read the API key from the X-API-Key header of a
	// request.
	//
	//     * AUTHORIZER to read the API key from the UsageIdentifierKey from
	// a custom authorizer.
	ApiKeySource types.ApiKeySourceType

	// The list of binary media types supported by the RestApi. By default, the RestApi
	// supports only UTF-8-encoded text payloads.
	BinaryMediaTypes []*string

	// The timestamp when the API was created.
	CreatedDate *time.Time

	// The API's description.
	Description *string

	// The endpoint configuration of this RestApi showing the endpoint types of the
	// API.
	EndpointConfiguration *types.EndpointConfiguration

	// The API's identifier. This identifier is unique across all of your APIs in API
	// Gateway.
	Id *string

	// A nullable integer that is used to enable compression (with non-negative between
	// 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null
	// value) on an API. When compression is enabled, compression or decompression is
	// not applied on the payload if the payload size is smaller than this value.
	// Setting it to zero allows compression for any payload size.
	MinimumCompressionSize *int32

	// The API's name.
	Name *string

	// A stringified JSON policy document that applies to this RestApi regardless of
	// the caller and Method configuration.
	Policy *string

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string

	// A version identifier for the API.
	Version *string

	// The warning messages reported when failonwarnings is turned on during API
	// import.
	Warnings []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationImportRestApiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpImportRestApi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpImportRestApi{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opImportRestApi(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opImportRestApi(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "ImportRestApi",
	}
}
