// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the expressions configured for the search domain. Can be limited to
// specific expressions by name. By default, shows all expressions and includes any
// pending changes to the configuration. Set the Deployed option to true to show
// the active configuration and exclude pending changes. For more information, see
// Configuring Expressions
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-expressions.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) DescribeExpressions(ctx context.Context, params *DescribeExpressionsInput, optFns ...func(*Options)) (*DescribeExpressionsOutput, error) {
	if params == nil {
		params = &DescribeExpressionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeExpressions", params, optFns, addOperationDescribeExpressionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeExpressionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeDomains operation. Specifies the
// name of the domain you want to describe. To restrict the response to particular
// expressions, specify the names of the expressions you want to describe. To show
// the active configuration and exclude any pending changes, set the Deployed
// option to true.
type DescribeExpressionsInput struct {

	// The name of the domain you want to describe.
	//
	// This member is required.
	DomainName *string

	// Whether to display the deployed configuration (true) or include any pending
	// changes (false). Defaults to false.
	Deployed *bool

	// Limits the DescribeExpressions response to the specified expressions. If not
	// specified, all expressions are shown.
	ExpressionNames []*string
}

// The result of a DescribeExpressions request. Contains the expressions configured
// for the domain specified in the request.
type DescribeExpressionsOutput struct {

	// The expressions configured for the domain.
	//
	// This member is required.
	Expressions []*types.ExpressionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeExpressionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeExpressions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeExpressions{}, middleware.After)
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
	addOpDescribeExpressionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeExpressions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeExpressions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "DescribeExpressions",
	}
}
