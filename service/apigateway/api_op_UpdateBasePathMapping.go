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

// Changes information about the BasePathMapping resource.
func (c *Client) UpdateBasePathMapping(ctx context.Context, params *UpdateBasePathMappingInput, optFns ...func(*Options)) (*UpdateBasePathMappingOutput, error) {
	if params == nil {
		params = &UpdateBasePathMappingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateBasePathMapping", params, optFns, addOperationUpdateBasePathMappingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateBasePathMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to change information about the BasePathMapping resource.
type UpdateBasePathMappingInput struct {

	// [Required] The base path of the BasePathMapping resource to change. To specify
	// an empty base path, set this parameter to '(none)'.
	//
	// This member is required.
	BasePath *string

	// [Required] The domain name of the BasePathMapping resource to change.
	//
	// This member is required.
	DomainName *string

	Name *string

	// A list of update operations to be applied to the specified resource and in the
	// order specified in this list.
	PatchOperations []*types.PatchOperation

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// Represents the base path that callers of the API must provide as part of the URL
// after the domain name. A custom domain name plus a BasePathMapping specification
// identifies a deployed RestApi in a given stage of the owner Account. Use Custom
// Domain Names
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html)
type UpdateBasePathMappingOutput struct {

	// The base path name that callers of the API must provide as part of the URL after
	// the domain name.
	BasePath *string

	// The string identifier of the associated RestApi.
	RestApiId *string

	// The name of the associated stage.
	Stage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateBasePathMappingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateBasePathMapping{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateBasePathMapping{}, middleware.After)
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
	addOpUpdateBasePathMappingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBasePathMapping(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateBasePathMapping(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateBasePathMapping",
	}
}
