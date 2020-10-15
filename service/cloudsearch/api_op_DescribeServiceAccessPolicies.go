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

// Gets information about the access policies that control access to the domain's
// document and search endpoints. By default, shows the configuration with any
// pending changes. Set the Deployed option to true to show the active
// configuration and exclude pending changes. For more information, see Configuring
// Access for a Search Domain
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-access.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) DescribeServiceAccessPolicies(ctx context.Context, params *DescribeServiceAccessPoliciesInput, optFns ...func(*Options)) (*DescribeServiceAccessPoliciesOutput, error) {
	if params == nil {
		params = &DescribeServiceAccessPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeServiceAccessPolicies", params, optFns, addOperationDescribeServiceAccessPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeServiceAccessPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeServiceAccessPolicies operation.
// Specifies the name of the domain you want to describe. To show the active
// configuration and exclude any pending changes, set the Deployed option to true.
type DescribeServiceAccessPoliciesInput struct {

	// The name of the domain you want to describe.
	//
	// This member is required.
	DomainName *string

	// Whether to display the deployed configuration (true) or include any pending
	// changes (false). Defaults to false.
	Deployed *bool
}

// The result of a DescribeServiceAccessPolicies request.
type DescribeServiceAccessPoliciesOutput struct {

	// The access rules configured for the domain specified in the request.
	//
	// This member is required.
	AccessPolicies *types.AccessPoliciesStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeServiceAccessPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeServiceAccessPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeServiceAccessPolicies{}, middleware.After)
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
	addOpDescribeServiceAccessPoliciesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeServiceAccessPolicies(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeServiceAccessPolicies(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "DescribeServiceAccessPolicies",
	}
}
