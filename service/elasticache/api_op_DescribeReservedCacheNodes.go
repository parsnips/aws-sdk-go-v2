// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about reserved cache nodes for this account, or about a
// specified reserved cache node.
func (c *Client) DescribeReservedCacheNodes(ctx context.Context, params *DescribeReservedCacheNodesInput, optFns ...func(*Options)) (*DescribeReservedCacheNodesOutput, error) {
	if params == nil {
		params = &DescribeReservedCacheNodesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReservedCacheNodes", params, optFns, addOperationDescribeReservedCacheNodesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReservedCacheNodesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DescribeReservedCacheNodes operation.
type DescribeReservedCacheNodesInput struct {

	// The cache node type filter value. Use this parameter to show only those
	// reservations matching the specified cache node type. The following node types
	// are supported by ElastiCache. Generally speaking, the current generation types
	// provide more memory and computational power at lower cost when compared to their
	// equivalent previous generation counterparts.
	//
	//     * General purpose:
	//
	//         *
	// Current generation: M5 node types: cache.m5.large, cache.m5.xlarge,
	// cache.m5.2xlarge, cache.m5.4xlarge, cache.m5.12xlarge, cache.m5.24xlarge M4 node
	// types: cache.m4.large, cache.m4.xlarge, cache.m4.2xlarge, cache.m4.4xlarge,
	// cache.m4.10xlarge T3 node types: cache.t3.micro, cache.t3.small, cache.t3.medium
	// T2 node types: cache.t2.micro, cache.t2.small, cache.t2.medium
	//
	//         *
	// Previous generation: (not recommended) T1 node types: cache.t1.micro M1 node
	// types: cache.m1.small, cache.m1.medium, cache.m1.large, cache.m1.xlarge M3 node
	// types: cache.m3.medium, cache.m3.large, cache.m3.xlarge, cache.m3.2xlarge
	//
	//     *
	// Compute optimized:
	//
	//         * Previous generation: (not recommended) C1 node
	// types: cache.c1.xlarge
	//
	//     * Memory optimized:
	//
	//         * Current generation:
	// R5 node types: cache.r5.large, cache.r5.xlarge, cache.r5.2xlarge,
	// cache.r5.4xlarge, cache.r5.12xlarge, cache.r5.24xlarge R4 node types:
	// cache.r4.large, cache.r4.xlarge, cache.r4.2xlarge, cache.r4.4xlarge,
	// cache.r4.8xlarge, cache.r4.16xlarge
	//
	//         * Previous generation: (not
	// recommended) M2 node types: cache.m2.xlarge, cache.m2.2xlarge, cache.m2.4xlarge
	// R3 node types: cache.r3.large, cache.r3.xlarge, cache.r3.2xlarge,
	//
	//
	// cache.r3.4xlarge, cache.r3.8xlarge
	//
	// Additional node type info
	//
	//     * All current
	// generation instance types are created in Amazon VPC by default.
	//
	//     * Redis
	// append-only files (AOF) are not supported for T1 or T2 instances.
	//
	//     * Redis
	// Multi-AZ with automatic failover is not supported on T1 instances.
	//
	//     * Redis
	// configuration variables appendonly and appendfsync are not supported on Redis
	// version 2.8.22 and later.
	CacheNodeType *string

	// The duration filter value, specified in years or seconds. Use this parameter to
	// show only reservations for this duration. Valid Values: 1 | 3 | 31536000 |
	// 94608000
	Duration *string

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved. Default: 100 Constraints: minimum
	// 20; maximum 100.
	MaxRecords *int32

	// The offering type filter value. Use this parameter to show only the available
	// offerings matching the specified offering type. Valid values: "Light
	// Utilization"|"Medium Utilization"|"Heavy Utilization"
	OfferingType *string

	// The product description filter value. Use this parameter to show only those
	// reservations matching the specified product description.
	ProductDescription *string

	// The reserved cache node identifier filter value. Use this parameter to show only
	// the reservation that matches the specified reservation ID.
	ReservedCacheNodeId *string

	// The offering identifier filter value. Use this parameter to show only purchased
	// reservations matching the specified offering identifier.
	ReservedCacheNodesOfferingId *string
}

// Represents the output of a DescribeReservedCacheNodes operation.
type DescribeReservedCacheNodesOutput struct {

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string

	// A list of reserved cache nodes. Each element in the list contains detailed
	// information about one node.
	ReservedCacheNodes []*types.ReservedCacheNode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeReservedCacheNodesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeReservedCacheNodes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeReservedCacheNodes{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedCacheNodes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeReservedCacheNodes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DescribeReservedCacheNodes",
	}
}
