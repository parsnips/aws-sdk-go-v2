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

// Gets information about the index fields configured for the search domain. Can be
// limited to specific fields by name. By default, shows all fields and includes
// any pending changes to the configuration. Set the Deployed option to true to
// show the active configuration and exclude pending changes. For more information,
// see Getting Domain Information
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-domain-info.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) DescribeIndexFields(ctx context.Context, params *DescribeIndexFieldsInput, optFns ...func(*Options)) (*DescribeIndexFieldsOutput, error) {
	if params == nil {
		params = &DescribeIndexFieldsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeIndexFields", params, optFns, addOperationDescribeIndexFieldsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeIndexFieldsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeIndexFields operation. Specifies the
// name of the domain you want to describe. To restrict the response to particular
// index fields, specify the names of the index fields you want to describe. To
// show the active configuration and exclude any pending changes, set the Deployed
// option to true.
type DescribeIndexFieldsInput struct {

	// The name of the domain you want to describe.
	//
	// This member is required.
	DomainName *string

	// Whether to display the deployed configuration (true) or include any pending
	// changes (false). Defaults to false.
	Deployed *bool

	// A list of the index fields you want to describe. If not specified, information
	// is returned for all configured index fields.
	FieldNames []*string
}

// The result of a DescribeIndexFields request. Contains the index fields
// configured for the domain specified in the request.
type DescribeIndexFieldsOutput struct {

	// The index fields configured for the domain.
	//
	// This member is required.
	IndexFields []*types.IndexFieldStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeIndexFieldsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeIndexFields{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeIndexFields{}, middleware.After)
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
	addOpDescribeIndexFieldsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeIndexFields(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeIndexFields(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "DescribeIndexFields",
	}
}
