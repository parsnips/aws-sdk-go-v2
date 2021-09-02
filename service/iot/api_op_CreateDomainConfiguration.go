// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a domain configuration. Requires permission to access the
// CreateDomainConfiguration
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) CreateDomainConfiguration(ctx context.Context, params *CreateDomainConfigurationInput, optFns ...func(*Options)) (*CreateDomainConfigurationOutput, error) {
	if params == nil {
		params = &CreateDomainConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDomainConfiguration", params, optFns, c.addOperationCreateDomainConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDomainConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDomainConfigurationInput struct {

	// The name of the domain configuration. This value must be unique to a region.
	//
	// This member is required.
	DomainConfigurationName *string

	// An object that specifies the authorization service for a domain.
	AuthorizerConfig *types.AuthorizerConfig

	// The name of the domain.
	DomainName *string

	// The ARNs of the certificates that IoT passes to the device during the TLS
	// handshake. Currently you can specify only one certificate ARN. This value is not
	// required for Amazon Web Services-managed domains.
	ServerCertificateArns []string

	// The type of service delivered by the endpoint. Amazon Web Services IoT Core
	// currently supports only the DATA service type.
	ServiceType types.ServiceType

	// Metadata which can be used to manage the domain configuration. For URI Request
	// parameters use format: ...key1=value1&key2=value2... For the CLI command-line
	// parameter use format: &&tags "key1=value1&key2=value2..." For the cli-input-json
	// file use format: "tags": "key1=value1&key2=value2..."
	Tags []types.Tag

	// The certificate used to validate the server certificate and prove domain name
	// ownership. This certificate must be signed by a public certificate authority.
	// This value is not required for Amazon Web Services-managed domains.
	ValidationCertificateArn *string

	noSmithyDocumentSerde
}

type CreateDomainConfigurationOutput struct {

	// The ARN of the domain configuration.
	DomainConfigurationArn *string

	// The name of the domain configuration.
	DomainConfigurationName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDomainConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDomainConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDomainConfiguration{}, middleware.After)
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
	if err = addOpCreateDomainConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDomainConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDomainConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateDomainConfiguration",
	}
}
