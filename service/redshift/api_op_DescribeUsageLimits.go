// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Shows usage limits on a cluster. Results are filtered based on the combination
// of input usage limit identifier, cluster identifier, and feature type
// parameters:
//
//     * If usage limit identifier, cluster identifier, and feature
// type are not provided, then all usage limit objects for the current account in
// the current region are returned.
//
//     * If usage limit identifier is provided,
// then the corresponding usage limit object is returned.
//
//     * If cluster
// identifier is provided, then all usage limit objects for the specified cluster
// are returned.
//
//     * If cluster identifier and feature type are provided, then
// all usage limit objects for the combination of cluster and feature are returned.
func (c *Client) DescribeUsageLimits(ctx context.Context, params *DescribeUsageLimitsInput, optFns ...func(*Options)) (*DescribeUsageLimitsOutput, error) {
	if params == nil {
		params = &DescribeUsageLimitsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeUsageLimits", params, optFns, addOperationDescribeUsageLimitsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeUsageLimitsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUsageLimitsInput struct {

	// The identifier of the cluster for which you want to describe usage limits.
	ClusterIdentifier *string

	// The feature type for which you want to describe usage limits.
	FeatureType types.UsageLimitFeatureType

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeUsageLimits request exceed the
	// value specified in MaxRecords, AWS returns a value in the Marker field of the
	// response. You can retrieve the next set of response records by providing the
	// returned marker value in the Marker parameter and retrying the request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32

	// A tag key or keys for which you want to return all matching usage limit objects
	// that are associated with the specified key or keys. For example, suppose that
	// you have parameter groups that are tagged with keys called owner and
	// environment. If you specify both of these tag keys in the request, Amazon
	// Redshift returns a response with the usage limit objects have either or both of
	// these tag keys associated with them.
	TagKeys []*string

	// A tag value or values for which you want to return all matching usage limit
	// objects that are associated with the specified tag value or values. For example,
	// suppose that you have parameter groups that are tagged with values called admin
	// and test. If you specify both of these tag values in the request, Amazon
	// Redshift returns a response with the usage limit objects that have either or
	// both of these tag values associated with them.
	TagValues []*string

	// The identifier of the usage limit to describe.
	UsageLimitId *string
}

type DescribeUsageLimitsOutput struct {

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// Contains the output from the DescribeUsageLimits action.
	UsageLimits []*types.UsageLimit

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeUsageLimitsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeUsageLimits{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeUsageLimits{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUsageLimits(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeUsageLimits(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeUsageLimits",
	}
}
