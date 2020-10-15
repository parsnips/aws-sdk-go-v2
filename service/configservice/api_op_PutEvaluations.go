// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Used by an AWS Lambda function to deliver evaluation results to AWS Config. This
// action is required in every AWS Lambda function that is invoked by an AWS Config
// rule.
func (c *Client) PutEvaluations(ctx context.Context, params *PutEvaluationsInput, optFns ...func(*Options)) (*PutEvaluationsOutput, error) {
	if params == nil {
		params = &PutEvaluationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutEvaluations", params, optFns, addOperationPutEvaluationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutEvaluationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type PutEvaluationsInput struct {

	// An encrypted token that associates an evaluation with an AWS Config rule.
	// Identifies the rule and the event that triggered the evaluation.
	//
	// This member is required.
	ResultToken *string

	// The assessments that the AWS Lambda function performs. Each evaluation
	// identifies an AWS resource and indicates whether it complies with the AWS Config
	// rule that invokes the AWS Lambda function.
	Evaluations []*types.Evaluation

	// Use this parameter to specify a test run for PutEvaluations. You can verify
	// whether your AWS Lambda function will deliver evaluation results to AWS Config.
	// No updates occur to your existing evaluations, and evaluation results are not
	// sent to AWS Config. When TestMode is true, PutEvaluations doesn't require a
	// valid value for the ResultToken parameter, but the value cannot be null.
	TestMode *bool
}

//
type PutEvaluationsOutput struct {

	// Requests that failed because of a client or server error.
	FailedEvaluations []*types.Evaluation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutEvaluationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutEvaluations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutEvaluations{}, middleware.After)
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
	addOpPutEvaluationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutEvaluations(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutEvaluations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutEvaluations",
	}
}
