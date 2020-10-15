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

// Updates a usage plan of a given plan Id.
func (c *Client) UpdateUsagePlan(ctx context.Context, params *UpdateUsagePlanInput, optFns ...func(*Options)) (*UpdateUsagePlanOutput, error) {
	if params == nil {
		params = &UpdateUsagePlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateUsagePlan", params, optFns, addOperationUpdateUsagePlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateUsagePlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The PATCH request to update a usage plan of a given plan Id.
type UpdateUsagePlanInput struct {

	// [Required] The Id of the to-be-updated usage plan.
	//
	// This member is required.
	UsagePlanId *string

	Name *string

	// A list of update operations to be applied to the specified resource and in the
	// order specified in this list.
	PatchOperations []*types.PatchOperation

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// Represents a usage plan than can specify who can assess associated API stages
// with specified request limits and quotas. In a usage plan, you associate an API
// by specifying the API's Id and a stage name of the specified API. You add plan
// customers by adding API keys to the plan. Create and Use Usage Plans
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html)
type UpdateUsagePlanOutput struct {

	// The associated API stages of a usage plan.
	ApiStages []*types.ApiStage

	// The description of a usage plan.
	Description *string

	// The identifier of a UsagePlan resource.
	Id *string

	// The name of a usage plan.
	Name *string

	// The AWS Markeplace product identifier to associate with the usage plan as a SaaS
	// product on AWS Marketplace.
	ProductCode *string

	// The maximum number of permitted requests per a given unit time interval.
	Quota *types.QuotaSettings

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string

	// The request throttle limits of a usage plan.
	Throttle *types.ThrottleSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateUsagePlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateUsagePlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateUsagePlan{}, middleware.After)
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
	addOpUpdateUsagePlanValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUsagePlan(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateUsagePlan(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateUsagePlan",
	}
}
