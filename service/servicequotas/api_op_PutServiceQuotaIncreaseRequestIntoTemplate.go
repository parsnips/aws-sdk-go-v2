// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Defines and adds a quota to the service quota template. To add a quota to the
// template, you must provide the ServiceCode, QuotaCode, AwsRegion, and
// DesiredValue. Once you add a quota to the template, use
// ListServiceQuotaIncreaseRequestsInTemplate to see the list of quotas in the
// template.
func (c *Client) PutServiceQuotaIncreaseRequestIntoTemplate(ctx context.Context, params *PutServiceQuotaIncreaseRequestIntoTemplateInput, optFns ...func(*Options)) (*PutServiceQuotaIncreaseRequestIntoTemplateOutput, error) {
	if params == nil {
		params = &PutServiceQuotaIncreaseRequestIntoTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutServiceQuotaIncreaseRequestIntoTemplate", params, optFns, addOperationPutServiceQuotaIncreaseRequestIntoTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutServiceQuotaIncreaseRequestIntoTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutServiceQuotaIncreaseRequestIntoTemplateInput struct {

	// Specifies the AWS Region for the quota.
	//
	// This member is required.
	AwsRegion *string

	// Specifies the new, increased value for the quota.
	//
	// This member is required.
	DesiredValue *float64

	// Specifies the service quota that you want to use.
	//
	// This member is required.
	QuotaCode *string

	// Specifies the service that you want to use.
	//
	// This member is required.
	ServiceCode *string
}

type PutServiceQuotaIncreaseRequestIntoTemplateOutput struct {

	// A structure that contains information about one service quota increase request.
	ServiceQuotaIncreaseRequestInTemplate *types.ServiceQuotaIncreaseRequestInTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutServiceQuotaIncreaseRequestIntoTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutServiceQuotaIncreaseRequestIntoTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutServiceQuotaIncreaseRequestIntoTemplate{}, middleware.After)
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
	addOpPutServiceQuotaIncreaseRequestIntoTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutServiceQuotaIncreaseRequestIntoTemplate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutServiceQuotaIncreaseRequestIntoTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "PutServiceQuotaIncreaseRequestIntoTemplate",
	}
}
