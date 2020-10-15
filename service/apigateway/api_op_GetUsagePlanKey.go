// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a usage plan key of a given key identifier.
func (c *Client) GetUsagePlanKey(ctx context.Context, params *GetUsagePlanKeyInput, optFns ...func(*Options)) (*GetUsagePlanKeyOutput, error) {
	if params == nil {
		params = &GetUsagePlanKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUsagePlanKey", params, optFns, addOperationGetUsagePlanKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUsagePlanKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GET request to get a usage plan key of a given key identifier.
type GetUsagePlanKeyInput struct {

	// [Required] The key Id of the to-be-retrieved UsagePlanKey resource representing
	// a plan customer.
	//
	// This member is required.
	KeyId *string

	// [Required] The Id of the UsagePlan resource representing the usage plan
	// containing the to-be-retrieved UsagePlanKey resource representing a plan
	// customer.
	//
	// This member is required.
	UsagePlanId *string

	Name *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// Represents a usage plan key to identify a plan customer. To associate an API
// stage with a selected API key in a usage plan, you must create a UsagePlanKey
// resource to represent the selected ApiKey. " Create and Use Usage Plans
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html)
type GetUsagePlanKeyOutput struct {

	// The Id of a usage plan key.
	Id *string

	// The name of a usage plan key.
	Name *string

	// The type of a usage plan key. Currently, the valid key type is API_KEY.
	Type *string

	// The value of a usage plan key.
	Value *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetUsagePlanKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUsagePlanKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUsagePlanKey{}, middleware.After)
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
	addOpGetUsagePlanKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetUsagePlanKey(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetUsagePlanKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetUsagePlanKey",
	}
}
