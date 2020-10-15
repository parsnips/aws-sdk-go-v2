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

// Returns some or all (up to 1,000) of the objects in a bucket. You can use the
// request parameters as selection criteria to return a subset of the objects in a
// bucket. A 200 OK response can contain valid or invalid XML. Make sure to design
// your application to parse the contents of the response and handle it
// appropriately. To use this operation, you must have READ access to the bucket.
// To use this operation in an AWS Identity and Access Management (IAM) policy, you
// must have permissions to perform the s3:ListBucket action. The bucket owner has
// this permission by default and can grant this permission to others. For more
// information about permissions, see Permissions Related to Bucket Subresource
// Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html). This
// section describes the latest revision of the API. We recommend that you use this
// revised API for application development. For backward compatibility, Amazon S3
// continues to support the prior version of this API, ListObjects. To get a list
// of your buckets, see ListBuckets. The following operations are related to
// ListObjectsV2:
//
//     * GetObject
//
//     * PutObject
//
//     * CreateBucket
func (c *Client) ListObjectsV2(ctx context.Context, params *ListObjectsV2Input, optFns ...func(*Options)) (*ListObjectsV2Output, error) {
	if params == nil {
		params = &ListObjectsV2Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListObjectsV2", params, optFns, addOperationListObjectsV2Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListObjectsV2Output)
	out.ResultMetadata = metadata
	return out, nil
}

type ListObjectsV2Input struct {

	// Bucket name to list. When using this API with an access point, you must direct
	// requests to the access point hostname. The access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// operation using an access point through the AWS SDKs, you provide the access
	// point ARN in place of the bucket name. For more information about access point
	// ARNs, see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) in
	// the Amazon Simple Storage Service Developer Guide.
	//
	// This member is required.
	Bucket *string

	// ContinuationToken indicates Amazon S3 that the list is being continued on this
	// bucket with a token. ContinuationToken is obfuscated and is not a real key.
	ContinuationToken *string

	// A delimiter is a character you use to group keys.
	Delimiter *string

	// Encoding type used by Amazon S3 to encode object keys in the response.
	EncodingType types.EncodingType

	// The owner field is not present in listV2 by default, if you want to return owner
	// field with each key in the result then set the fetch owner field to true.
	FetchOwner *bool

	// Sets the maximum number of keys returned in the response. By default the API
	// returns up to 1,000 key names. The response might contain fewer keys but will
	// never contain more.
	MaxKeys *int32

	// Limits the response to keys that begin with the specified prefix.
	Prefix *string

	// Confirms that the requester knows that she or he will be charged for the list
	// objects request in V2 style. Bucket owners need not specify this parameter in
	// their requests.
	RequestPayer types.RequestPayer

	// StartAfter is where you want Amazon S3 to start listing from. Amazon S3 starts
	// listing after this specified key. StartAfter can be any key in the bucket.
	StartAfter *string
}

type ListObjectsV2Output struct {

	// All of the keys rolled up into a common prefix count as a single return when
	// calculating the number of returns. A response can contain CommonPrefixes only if
	// you specify a delimiter. CommonPrefixes contains all (if there are any) keys
	// between Prefix and the next occurrence of the string specified by a delimiter.
	// CommonPrefixes lists keys that act like subdirectories in the directory
	// specified by Prefix. For example, if the prefix is notes/ and the delimiter is a
	// slash (/) as in notes/summer/july, the common prefix is notes/summer/. All of
	// the keys that roll up into a common prefix count as a single return when
	// calculating the number of returns.
	CommonPrefixes []*types.CommonPrefix

	// Metadata about each object returned.
	Contents []*types.Object

	// If ContinuationToken was sent with the request, it is included in the response.
	ContinuationToken *string

	// Causes keys that contain the same string between the prefix and the first
	// occurrence of the delimiter to be rolled up into a single result element in the
	// CommonPrefixes collection. These rolled-up keys are not returned elsewhere in
	// the response. Each rolled-up result counts as only one return against the
	// MaxKeys value.
	Delimiter *string

	// Encoding type used by Amazon S3 to encode object key names in the XML response.
	// If you specify the encoding-type request parameter, Amazon S3 includes this
	// element in the response, and returns encoded key name values in the following
	// response elements: Delimiter, Prefix, Key, and StartAfter.
	EncodingType types.EncodingType

	// Set to false if all of the results were returned. Set to true if more keys are
	// available to return. If the number of results exceeds that specified by MaxKeys,
	// all of the results might not be returned.
	IsTruncated *bool

	// KeyCount is the number of keys returned with this request. KeyCount will always
	// be less than equals to MaxKeys field. Say you ask for 50 keys, your result will
	// include less than equals 50 keys
	KeyCount *int32

	// Sets the maximum number of keys returned in the response. By default the API
	// returns up to 1,000 key names. The response might contain fewer keys but will
	// never contain more.
	MaxKeys *int32

	// Bucket name. When using this API with an access point, you must direct requests
	// to the access point hostname. The access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// operation using an access point through the AWS SDKs, you provide the access
	// point ARN in place of the bucket name. For more information about access point
	// ARNs, see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) in
	// the Amazon Simple Storage Service Developer Guide.
	Name *string

	// NextContinuationToken is sent when isTruncated is true, which means there are
	// more keys in the bucket that can be listed. The next list requests to Amazon S3
	// can be continued with this NextContinuationToken. NextContinuationToken is
	// obfuscated and is not a real key
	NextContinuationToken *string

	// Keys that begin with the indicated prefix.
	Prefix *string

	// If StartAfter was sent with the request, it is included in the response.
	StartAfter *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListObjectsV2Middlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListObjectsV2{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListObjectsV2{}, middleware.After)
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
	addOpListObjectsV2ValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListObjectsV2(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opListObjectsV2(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListObjectsV2",
	}
}
