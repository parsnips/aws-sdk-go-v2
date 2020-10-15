// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Returns the SqlInjectionMatchSet that is specified by
// SqlInjectionMatchSetId.
func (c *Client) GetSqlInjectionMatchSet(ctx context.Context, params *GetSqlInjectionMatchSetInput, optFns ...func(*Options)) (*GetSqlInjectionMatchSetOutput, error) {
	if params == nil {
		params = &GetSqlInjectionMatchSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSqlInjectionMatchSet", params, optFns, addOperationGetSqlInjectionMatchSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSqlInjectionMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get a SqlInjectionMatchSet.
type GetSqlInjectionMatchSetInput struct {

	// The SqlInjectionMatchSetId of the SqlInjectionMatchSet that you want to get.
	// SqlInjectionMatchSetId is returned by CreateSqlInjectionMatchSet and by
	// ListSqlInjectionMatchSets.
	//
	// This member is required.
	SqlInjectionMatchSetId *string
}

// The response to a GetSqlInjectionMatchSet request.
type GetSqlInjectionMatchSetOutput struct {

	// Information about the SqlInjectionMatchSet that you specified in the
	// GetSqlInjectionMatchSet request. For more information, see the following
	// topics:
	//
	//     * SqlInjectionMatchSet: Contains Name, SqlInjectionMatchSetId, and
	// an array of SqlInjectionMatchTuple objects
	//
	//     * SqlInjectionMatchTuple: Each
	// SqlInjectionMatchTuple object contains FieldToMatch and TextTransformation
	//
	//
	// * FieldToMatch: Contains Data and Type
	SqlInjectionMatchSet *types.SqlInjectionMatchSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSqlInjectionMatchSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSqlInjectionMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSqlInjectionMatchSet{}, middleware.After)
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
	addOpGetSqlInjectionMatchSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSqlInjectionMatchSet(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetSqlInjectionMatchSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "GetSqlInjectionMatchSet",
	}
}
