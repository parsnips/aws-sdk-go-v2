// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Records the configuration state for the resource provided in the request. The
// configuration state of a resource is represented in AWS Config as Configuration
// Items. Once this API records the configuration item, you can retrieve the list
// of configuration items for the custom resource type using existing AWS Config
// APIs. The custom resource type must be registered with AWS CloudFormation. This
// API accepts the configuration item registered with AWS CloudFormation. When you
// call this API, AWS Config only stores configuration state of the resource
// provided in the request. This API does not change or remediate the configuration
// of the resource. Write-only schema properites are not recorded as part of the
// published configuration item.
func (c *Client) PutResourceConfig(ctx context.Context, params *PutResourceConfigInput, optFns ...func(*Options)) (*PutResourceConfigOutput, error) {
	if params == nil {
		params = &PutResourceConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutResourceConfig", params, optFns, addOperationPutResourceConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutResourceConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutResourceConfigInput struct {

	// The configuration object of the resource in valid JSON format. It must match the
	// schema registered with AWS CloudFormation. The configuration JSON must not
	// exceed 64 KB.
	//
	// This member is required.
	Configuration *string

	// Unique identifier of the resource.
	//
	// This member is required.
	ResourceId *string

	// The type of the resource. The custom resource type must be registered with AWS
	// CloudFormation. You cannot use the organization names “aws”, “amzn”, “amazon”,
	// “alexa”, “custom” with custom resource types. It is the first part of the
	// ResourceType up to the first ::.
	//
	// This member is required.
	ResourceType *string

	// Version of the schema registered for the ResourceType in AWS CloudFormation.
	//
	// This member is required.
	SchemaVersionId *string

	// Name of the resource.
	ResourceName *string

	// Tags associated with the resource.
	Tags map[string]*string
}

type PutResourceConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutResourceConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutResourceConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutResourceConfig{}, middleware.After)
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
	addOpPutResourceConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutResourceConfig(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutResourceConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutResourceConfig",
	}
}
