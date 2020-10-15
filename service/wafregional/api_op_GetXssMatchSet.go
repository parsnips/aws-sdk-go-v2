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
// global use. Returns the XssMatchSet that is specified by XssMatchSetId.
func (c *Client) GetXssMatchSet(ctx context.Context, params *GetXssMatchSetInput, optFns ...func(*Options)) (*GetXssMatchSetOutput, error) {
	if params == nil {
		params = &GetXssMatchSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetXssMatchSet", params, optFns, addOperationGetXssMatchSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetXssMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get an XssMatchSet.
type GetXssMatchSetInput struct {

	// The XssMatchSetId of the XssMatchSet that you want to get. XssMatchSetId is
	// returned by CreateXssMatchSet and by ListXssMatchSets.
	//
	// This member is required.
	XssMatchSetId *string
}

// The response to a GetXssMatchSet request.
type GetXssMatchSetOutput struct {

	// Information about the XssMatchSet that you specified in the GetXssMatchSet
	// request. For more information, see the following topics:
	//
	//     * XssMatchSet:
	// Contains Name, XssMatchSetId, and an array of XssMatchTuple objects
	//
	//     *
	// XssMatchTuple: Each XssMatchTuple object contains FieldToMatch and
	// TextTransformation
	//
	//     * FieldToMatch: Contains Data and Type
	XssMatchSet *types.XssMatchSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetXssMatchSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetXssMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetXssMatchSet{}, middleware.After)
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
	addOpGetXssMatchSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetXssMatchSet(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetXssMatchSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "GetXssMatchSet",
	}
}
