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

// Sets up a method's integration.
func (c *Client) PutIntegration(ctx context.Context, params *PutIntegrationInput, optFns ...func(*Options)) (*PutIntegrationOutput, error) {
	if params == nil {
		params = &PutIntegrationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutIntegration", params, optFns, addOperationPutIntegrationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutIntegrationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Sets up a method's integration.
type PutIntegrationInput struct {

	// [Required] Specifies a put integration request's HTTP method.
	//
	// This member is required.
	HttpMethod *string

	// [Required] Specifies a put integration request's resource ID.
	//
	// This member is required.
	ResourceId *string

	// [Required] The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// [Required] Specifies a put integration input's type.
	//
	// This member is required.
	Type types.IntegrationType

	// A list of request parameters whose values API Gateway caches. To be valid values
	// for cacheKeyParameters, these parameters must also be specified for
	// MethodrequestParameters.
	CacheKeyParameters []*string

	// Specifies a group of related cached parameters. By default, API Gateway uses the
	// resource ID as the cacheNamespace. You can specify the same cacheNamespace
	// across resources to return the same cached data for requests to different
	// resources.
	CacheNamespace *string

	// The (id
	// (https://docs.aws.amazon.com/apigateway/api-reference/resource/vpc-link/#id)) of
	// the VpcLink used for the integration when connectionType=VPC_LINK and undefined,
	// otherwise.
	ConnectionId *string

	// The type of the network connection to the integration endpoint. The valid value
	// is INTERNET for connections through the public routable internet or VPC_LINK for
	// private connections between API Gateway and a network load balancer in a VPC.
	// The default value is INTERNET.
	ConnectionType types.ConnectionType

	// Specifies how to handle request payload content type conversions. Supported
	// values are CONVERT_TO_BINARY and CONVERT_TO_TEXT, with the following
	// behaviors:
	//
	//     * CONVERT_TO_BINARY: Converts a request payload from a
	// Base64-encoded string to the corresponding binary blob.
	//
	//     * CONVERT_TO_TEXT:
	// Converts a request payload from a binary blob to a Base64-encoded string.
	//
	// If
	// this property is not defined, the request payload will be passed through from
	// the method request to integration request without modification, provided that
	// the passthroughBehavior is configured to support payload pass-through.
	ContentHandling types.ContentHandlingStrategy

	// Specifies whether credentials are required for a put integration.
	Credentials *string

	// Specifies a put integration HTTP method. When the integration type is HTTP or
	// AWS, this field is required.
	IntegrationHttpMethod *string

	Name *string

	// Specifies the pass-through behavior for incoming requests based on the
	// Content-Type header in the request, and the available mapping templates
	// specified as the requestTemplates property on the Integration resource. There
	// are three valid values: WHEN_NO_MATCH, WHEN_NO_TEMPLATES, and NEVER.
	//
	//     *
	// WHEN_NO_MATCH passes the request body for unmapped content types through to the
	// integration back end without transformation.
	//
	//     * NEVER rejects unmapped
	// content types with an HTTP 415 'Unsupported Media Type' response.
	//
	//     *
	// WHEN_NO_TEMPLATES allows pass-through when the integration has NO content types
	// mapped to templates. However if there is at least one content type defined,
	// unmapped content types will be rejected with the same 415 response.
	PassthroughBehavior *string

	// A key-value map specifying request parameters that are passed from the method
	// request to the back end. The key is an integration request parameter name and
	// the associated value is a method request parameter value or static value that
	// must be enclosed within single quotes and pre-encoded as required by the back
	// end. The method request parameter value must match the pattern of
	// method.request.{location}.{name}, where location is querystring, path, or header
	// and name must be a valid and unique method request parameter name.
	RequestParameters map[string]*string

	// Represents a map of Velocity templates that are applied on the request payload
	// based on the value of the Content-Type header sent by the client. The content
	// type value is the key in this map, and the template (as a String) is the value.
	RequestTemplates map[string]*string

	Template *bool

	TemplateSkipList []*string

	// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000
	// milliseconds or 29 seconds.
	TimeoutInMillis *int32

	Title *string

	TlsConfig *types.TlsConfig

	// Specifies Uniform Resource Identifier (URI) of the integration endpoint.
	//
	//     *
	// For HTTP or HTTP_PROXY integrations, the URI must be a fully formed, encoded
	// HTTP(S) URL according to the RFC-3986 specification
	// (https://en.wikipedia.org/wiki/Uniform_Resource_Identifier), for either standard
	// integration, where connectionType is not VPC_LINK, or private integration, where
	// connectionType is VPC_LINK. For a private HTTP integration, the URI is not used
	// for routing.
	//
	//     * For AWS or AWS_PROXY integrations, the URI is of the form
	// arn:aws:apigateway:{region}:{subdomain.service|service}:path|action/{service_api}.
	// Here, {Region} is the API Gateway region (e.g., us-east-1); {service} is the
	// name of the integrated AWS service (e.g., s3); and {subdomain} is a designated
	// subdomain supported by certain AWS service for fast host-name lookup. action can
	// be used for an AWS service action-based API, using an
	// Action={name}&{p1}={v1}&p2={v2}... query string. The ensuing {service_api}
	// refers to a supported action {name} plus any required input parameters.
	// Alternatively, path can be used for an AWS service path-based API. The ensuing
	// service_api refers to the path to an AWS service resource, including the region
	// of the integrated AWS service, if applicable. For example, for integration with
	// the S3 API of GetObject
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectGET.html), the uri
	// can be either
	// arn:aws:apigateway:us-west-2:s3:action/GetObject&Bucket={bucket}&Key={key} or
	// arn:aws:apigateway:us-west-2:s3:path/{bucket}/{key}
	Uri *string
}

// Represents an HTTP, HTTP_PROXY, AWS, AWS_PROXY, or Mock integration. In the API
// Gateway console, the built-in Lambda integration is an AWS integration. Creating
// an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
type PutIntegrationOutput struct {

	// A list of request parameters whose values API Gateway caches. To be valid values
	// for cacheKeyParameters, these parameters must also be specified for
	// MethodrequestParameters.
	CacheKeyParameters []*string

	// Specifies a group of related cached parameters. By default, API Gateway uses the
	// resource ID as the cacheNamespace. You can specify the same cacheNamespace
	// across resources to return the same cached data for requests to different
	// resources.
	CacheNamespace *string

	// The (id
	// (https://docs.aws.amazon.com/apigateway/api-reference/resource/vpc-link/#id)) of
	// the VpcLink used for the integration when connectionType=VPC_LINK and undefined,
	// otherwise.
	ConnectionId *string

	// The type of the network connection to the integration endpoint. The valid value
	// is INTERNET for connections through the public routable internet or VPC_LINK for
	// private connections between API Gateway and a network load balancer in a VPC.
	// The default value is INTERNET.
	ConnectionType types.ConnectionType

	// Specifies how to handle request payload content type conversions. Supported
	// values are CONVERT_TO_BINARY and CONVERT_TO_TEXT, with the following
	// behaviors:
	//
	//     * CONVERT_TO_BINARY: Converts a request payload from a
	// Base64-encoded string to the corresponding binary blob.
	//
	//     * CONVERT_TO_TEXT:
	// Converts a request payload from a binary blob to a Base64-encoded string.
	//
	// If
	// this property is not defined, the request payload will be passed through from
	// the method request to integration request without modification, provided that
	// the passthroughBehavior is configured to support payload pass-through.
	ContentHandling types.ContentHandlingStrategy

	// Specifies the credentials required for the integration, if any. For AWS
	// integrations, three options are available. To specify an IAM Role for API
	// Gateway to assume, use the role's Amazon Resource Name (ARN). To require that
	// the caller's identity be passed through from the request, specify the string
	// arn:aws:iam::\*:user/\*. To use resource-based permissions on supported AWS
	// services, specify null.
	Credentials *string

	// Specifies the integration's HTTP method type.
	HttpMethod *string

	// Specifies the integration's responses.
	// Example: Get integration responses of a
	// method
	//
	// Request
	//
	//     GET
	// /restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200
	// HTTP/1.1 Content-Type: application/json Host: apigateway.us-east-1.amazonaws.com
	// X-Amz-Date: 20160607T191449Z Authorization: AWS4-HMAC-SHA256
	// Credential={access_key_ID}/20160607/us-east-1/apigateway/aws4_request,
	// SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
	//
	// Response
	//
	// The
	// successful response returns 200 OK status and a payload as follows: { "_links":
	// { "curies": { "href":
	// "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-integration-response-{rel}.html",
	// "name": "integrationresponse", "templated": true }, "self": { "href":
	// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200",
	// "title": "200" }, "integrationresponse:delete": { "href":
	// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200"
	// }, "integrationresponse:update": { "href":
	// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200"
	// } }, "responseParameters": { "method.response.header.Content-Type":
	// "'application/xml'" }, "responseTemplates": { "application/json":
	// "$util.urlDecode(\"%3CkinesisStreams%3E#foreach($stream in
	// $input.path('$.StreamNames'))%3Cstream%3E%3Cname%3E$stream%3C/name%3E%3C/stream%3E#end%3C/kinesisStreams%3E\")\n"
	// }, "statusCode": "200" }Creating an API
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
	IntegrationResponses map[string]*types.IntegrationResponse

	// Specifies how the method request body of an unmapped content type will be passed
	// through the integration request to the back end without transformation. A
	// content type is unmapped if no mapping template is defined in the integration or
	// the content type does not match any of the mapped content types, as specified in
	// requestTemplates. The valid value is one of the following:
	//
	//     * WHEN_NO_MATCH:
	// passes the method request body through the integration request to the back end
	// without transformation when the method request content type does not match any
	// content type associated with the mapping templates defined in the integration
	// request.
	//
	//     * WHEN_NO_TEMPLATES: passes the method request body through the
	// integration request to the back end without transformation when no mapping
	// template is defined in the integration request. If a template is defined when
	// this option is selected, the method request of an unmapped content-type will be
	// rejected with an HTTP 415 Unsupported Media Type response.
	//
	//     * NEVER: rejects
	// the method request with an HTTP 415 Unsupported Media Type response when either
	// the method request content type does not match any content type associated with
	// the mapping templates defined in the integration request or no mapping template
	// is defined in the integration request.
	PassthroughBehavior *string

	// A key-value map specifying request parameters that are passed from the method
	// request to the back end. The key is an integration request parameter name and
	// the associated value is a method request parameter value or static value that
	// must be enclosed within single quotes and pre-encoded as required by the back
	// end. The method request parameter value must match the pattern of
	// method.request.{location}.{name}, where location is querystring, path, or header
	// and name must be a valid and unique method request parameter name.
	RequestParameters map[string]*string

	// Represents a map of Velocity templates that are applied on the request payload
	// based on the value of the Content-Type header sent by the client. The content
	// type value is the key in this map, and the template (as a String) is the value.
	RequestTemplates map[string]*string

	// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000
	// milliseconds or 29 seconds.
	TimeoutInMillis *int32

	// Specifies the TLS configuration for an integration.
	TlsConfig *types.TlsConfig

	// Specifies an API method integration type. The valid value is one of the
	// following:
	//
	//     * AWS: for integrating the API method request with an AWS
	// service action, including the Lambda function-invoking action. With the Lambda
	// function-invoking action, this is referred to as the Lambda custom integration.
	// With any other AWS service action, this is known as AWS integration.
	//
	//     *
	// AWS_PROXY: for integrating the API method request with the Lambda
	// function-invoking action with the client request passed through as-is. This
	// integration is also referred to as the Lambda proxy integration.
	//
	//     * HTTP:
	// for integrating the API method request with an HTTP endpoint, including a
	// private HTTP endpoint within a VPC. This integration is also referred to as the
	// HTTP custom integration.
	//
	//     * HTTP_PROXY: for integrating the API method
	// request with an HTTP endpoint, including a private HTTP endpoint within a VPC,
	// with the client request passed through as-is. This is also referred to as the
	// HTTP proxy integration.
	//
	//     * MOCK: for integrating the API method request with
	// API Gateway as a "loop-back" endpoint without invoking any backend.
	//
	// For the
	// HTTP and HTTP proxy integrations, each integration can specify a protocol
	// (http/https), port and path. Standard 80 and 443 ports are supported as well as
	// custom ports above 1024. An HTTP or HTTP proxy integration with a connectionType
	// of VPC_LINK is referred to as a private integration and uses a VpcLink to
	// connect API Gateway to a network load balancer of a VPC.
	Type types.IntegrationType

	// Specifies Uniform Resource Identifier (URI) of the integration endpoint.
	//
	//     *
	// For HTTP or HTTP_PROXY integrations, the URI must be a fully formed, encoded
	// HTTP(S) URL according to the RFC-3986 specification
	// (https://en.wikipedia.org/wiki/Uniform_Resource_Identifier), for either standard
	// integration, where connectionType is not VPC_LINK, or private integration, where
	// connectionType is VPC_LINK. For a private HTTP integration, the URI is not used
	// for routing.
	//
	//     * For AWS or AWS_PROXY integrations, the URI is of the form
	// arn:aws:apigateway:{region}:{subdomain.service|service}:path|action/{service_api}.
	// Here, {Region} is the API Gateway region (e.g., us-east-1); {service} is the
	// name of the integrated AWS service (e.g., s3); and {subdomain} is a designated
	// subdomain supported by certain AWS service for fast host-name lookup. action can
	// be used for an AWS service action-based API, using an
	// Action={name}&{p1}={v1}&p2={v2}... query string. The ensuing {service_api}
	// refers to a supported action {name} plus any required input parameters.
	// Alternatively, path can be used for an AWS service path-based API. The ensuing
	// service_api refers to the path to an AWS service resource, including the region
	// of the integrated AWS service, if applicable. For example, for integration with
	// the S3 API of GetObject
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectGET.html), the uri
	// can be either
	// arn:aws:apigateway:us-west-2:s3:action/GetObject&Bucket={bucket}&Key={key} or
	// arn:aws:apigateway:us-west-2:s3:path/{bucket}/{key}
	Uri *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutIntegrationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutIntegration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutIntegration{}, middleware.After)
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
	addOpPutIntegrationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutIntegration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutIntegration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "PutIntegration",
	}
}
