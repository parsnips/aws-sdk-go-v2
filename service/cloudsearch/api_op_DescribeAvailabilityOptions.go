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

// Gets the availability options configured for a domain. By default, shows the
// configuration with any pending changes. Set the Deployed option to true to show
// the active configuration and exclude pending changes. For more information, see
// Configuring Availability Options
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-availability-options.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) DescribeAvailabilityOptions(ctx context.Context, params *DescribeAvailabilityOptionsInput, optFns ...func(*Options)) (*DescribeAvailabilityOptionsOutput, error) {
	if params == nil {
		params = &DescribeAvailabilityOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAvailabilityOptions", params, optFns, addOperationDescribeAvailabilityOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAvailabilityOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeAvailabilityOptions operation.
// Specifies the name of the domain you want to describe. To show the active
// configuration and exclude any pending changes, set the Deployed option to true.
type DescribeAvailabilityOptionsInput struct {

	// The name of the domain you want to describe.
	//
	// This member is required.
	DomainName *string

	// Whether to display the deployed configuration (true) or include any pending
	// changes (false). Defaults to false.
	Deployed *bool
}

// The result of a DescribeAvailabilityOptions request. Indicates whether or not
// the Multi-AZ option is enabled for the domain specified in the request.
type DescribeAvailabilityOptionsOutput struct {

	// The availability options configured for the domain. Indicates whether Multi-AZ
	// is enabled for the domain.
	AvailabilityOptions *types.AvailabilityOptionsStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAvailabilityOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeAvailabilityOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeAvailabilityOptions{}, middleware.After)
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
	addOpDescribeAvailabilityOptionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAvailabilityOptions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeAvailabilityOptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "DescribeAvailabilityOptions",
	}
}
