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

// Changes information about a Stage resource.
func (c *Client) UpdateStage(ctx context.Context, params *UpdateStageInput, optFns ...func(*Options)) (*UpdateStageOutput, error) {
	if params == nil {
		params = &UpdateStageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateStage", params, optFns, addOperationUpdateStageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateStageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Requests API Gateway to change information about a Stage resource.
type UpdateStageInput struct {

	// [Required] The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// [Required] The name of the Stage resource to change information about.
	//
	// This member is required.
	StageName *string

	Name *string

	// A list of update operations to be applied to the specified resource and in the
	// order specified in this list.
	PatchOperations []*types.PatchOperation

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// Represents a unique identifier for a version of a deployed RestApi that is
// callable by users. Deploy an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-deploy-api.html)
type UpdateStageOutput struct {

	// Settings for logging access in this stage.
	AccessLogSettings *types.AccessLogSettings

	// Specifies whether a cache cluster is enabled for the stage.
	CacheClusterEnabled *bool

	// The size of the cache cluster for the stage, if enabled.
	CacheClusterSize types.CacheClusterSize

	// The status of the cache cluster for the stage, if enabled.
	CacheClusterStatus types.CacheClusterStatus

	// Settings for the canary deployment in this stage.
	CanarySettings *types.CanarySettings

	// The identifier of a client certificate for an API stage.
	ClientCertificateId *string

	// The timestamp when the stage was created.
	CreatedDate *time.Time

	// The identifier of the Deployment that the stage points to.
	DeploymentId *string

	// The stage's description.
	Description *string

	// The version of the associated API documentation.
	DocumentationVersion *string

	// The timestamp when the stage last updated.
	LastUpdatedDate *time.Time

	// A map that defines the method settings for a Stage resource. Keys (designated as
	// /{method_setting_key below) are method paths defined as
	// {resource_path}/{http_method} for an individual method override, or /\*/\* for
	// overriding all methods in the stage.
	MethodSettings map[string]*types.MethodSetting

	// The name of the stage is the first path segment in the Uniform Resource
	// Identifier (URI) of a call to API Gateway. Stage names can only contain
	// alphanumeric characters, hyphens, and underscores. Maximum length is 128
	// characters.
	StageName *string

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string

	// Specifies whether active tracing with X-ray is enabled for the Stage.
	TracingEnabled *bool

	// A map that defines the stage variables for a Stage resource. Variable names can
	// have alphanumeric and underscore characters, and the values must match
	// [A-Za-z0-9-._~:/?#&=,]+.
	Variables map[string]*string

	// The ARN of the WebAcl associated with the Stage.
	WebAclArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateStageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateStage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateStage{}, middleware.After)
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
	addOpUpdateStageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStage(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateStage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateStage",
	}
}
