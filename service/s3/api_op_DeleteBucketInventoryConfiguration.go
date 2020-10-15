// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an inventory configuration (identified by the inventory ID) from the
// bucket. To use this operation, you must have permissions to perform the
// s3:PutInventoryConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see Permissions Related to Bucket Subresource
// Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html). For
// information about the Amazon S3 inventory feature, see Amazon S3 Inventory
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html).
// Operations related to DeleteBucketInventoryConfiguration include:
//
//     *
// GetBucketInventoryConfiguration
//
//     * PutBucketInventoryConfiguration
//
//     *
// ListBucketInventoryConfigurations
func (c *Client) DeleteBucketInventoryConfiguration(ctx context.Context, params *DeleteBucketInventoryConfigurationInput, optFns ...func(*Options)) (*DeleteBucketInventoryConfigurationOutput, error) {
	if params == nil {
		params = &DeleteBucketInventoryConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBucketInventoryConfiguration", params, optFns, addOperationDeleteBucketInventoryConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBucketInventoryConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBucketInventoryConfigurationInput struct {

	// The name of the bucket containing the inventory configuration to delete.
	//
	// This member is required.
	Bucket *string

	// The ID used to identify the inventory configuration.
	//
	// This member is required.
	Id *string
}

type DeleteBucketInventoryConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteBucketInventoryConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteBucketInventoryConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteBucketInventoryConfiguration{}, middleware.After)
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
	addOpDeleteBucketInventoryConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBucketInventoryConfiguration(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteBucketInventoryConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteBucketInventoryConfiguration",
	}
}
