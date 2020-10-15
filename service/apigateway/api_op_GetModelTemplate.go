// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Generates a sample mapping template that can be used to transform a payload into
// the structure of a model.
func (c *Client) GetModelTemplate(ctx context.Context, params *GetModelTemplateInput, optFns ...func(*Options)) (*GetModelTemplateOutput, error) {
	if params == nil {
		params = &GetModelTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetModelTemplate", params, optFns, addOperationGetModelTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetModelTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to generate a sample mapping template used to transform the payload.
type GetModelTemplateInput struct {

	// [Required] The name of the model for which to generate a template.
	//
	// This member is required.
	ModelName *string

	// [Required] The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	Name *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// Represents a mapping template used to transform a payload. Mapping Templates
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/models-mappings.html#models-mappings-mappings)
type GetModelTemplateOutput struct {

	// The Apache Velocity Template Language (VTL)
	// (https://velocity.apache.org/engine/devel/vtl-reference.html) template content
	// used for the template resource.
	Value *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetModelTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetModelTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetModelTemplate{}, middleware.After)
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
	addOpGetModelTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetModelTemplate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetModelTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetModelTemplate",
	}
}
