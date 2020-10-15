// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of inventory configurations for the bucket. You can have up to
// 1,000 analytics configurations per bucket. This operation supports list
// pagination and does not return more than 100 configurations at a time. Always
// check the IsTruncated element in the response. If there are no more
// configurations to list, IsTruncated is set to false. If there are more
// configurations to list, IsTruncated is set to true, and there is a value in
// NextContinuationToken. You use the NextContinuationToken value to continue the
// pagination of the list by passing the value in continuation-token in the request
// to GET the next page. To use this operation, you must have permissions to
// perform the s3:GetInventoryConfiguration action. The bucket owner has this
// permission by default. The bucket owner can grant this permission to others. For
// more information about permissions, see Permissions Related to Bucket
// Subresource Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html). For
// information about the Amazon S3 inventory feature, see Amazon S3 Inventory
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html) The
// following operations are related to ListBucketInventoryConfigurations:
//
//     *
// GetBucketInventoryConfiguration
//
//     * DeleteBucketInventoryConfiguration
//
//     *
// PutBucketInventoryConfiguration
func (c *Client) ListBucketInventoryConfigurations(ctx context.Context, params *ListBucketInventoryConfigurationsInput, optFns ...func(*Options)) (*ListBucketInventoryConfigurationsOutput, error) {
	if params == nil {
		params = &ListBucketInventoryConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBucketInventoryConfigurations", params, optFns, addOperationListBucketInventoryConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBucketInventoryConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBucketInventoryConfigurationsInput struct {

	// The name of the bucket containing the inventory configurations to retrieve.
	//
	// This member is required.
	Bucket *string

	// The marker used to continue an inventory configuration listing that has been
	// truncated. Use the NextContinuationToken from a previously truncated list
	// response to continue the listing. The continuation token is an opaque value that
	// Amazon S3 understands.
	ContinuationToken *string
}

type ListBucketInventoryConfigurationsOutput struct {

	// If sent in the request, the marker that is used as a starting point for this
	// inventory configuration list response.
	ContinuationToken *string

	// The list of inventory configurations for a bucket.
	InventoryConfigurationList []*types.InventoryConfiguration

	// Tells whether the returned list of inventory configurations is complete. A value
	// of true indicates that the list is not complete and the NextContinuationToken is
	// provided for a subsequent request.
	IsTruncated *bool

	// The marker used to continue this inventory configuration listing. Use the
	// NextContinuationToken from this response to continue the listing in a subsequent
	// request. The continuation token is an opaque value that Amazon S3 understands.
	NextContinuationToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListBucketInventoryConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListBucketInventoryConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListBucketInventoryConfigurations{}, middleware.After)
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
	addOpListBucketInventoryConfigurationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListBucketInventoryConfigurations(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opListBucketInventoryConfigurations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListBucketInventoryConfigurations",
	}
}
