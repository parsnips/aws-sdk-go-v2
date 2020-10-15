// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
	"time"
)

// Adds an object to a bucket. You must have WRITE permissions on a bucket to add
// an object to it. Amazon S3 never adds partial objects; if you receive a success
// response, Amazon S3 added the entire object to the bucket. Amazon S3 is a
// distributed system. If it receives multiple write requests for the same object
// simultaneously, it overwrites all but the last object written. Amazon S3 does
// not provide object locking; if you need this, make sure to build it into your
// application layer or use versioning instead. To ensure that data is not
// corrupted traversing the network, use the Content-MD5 header. When you use this
// header, Amazon S3 checks the object against the provided MD5 value and, if they
// do not match, returns an error. Additionally, you can calculate the MD5 while
// putting an object to Amazon S3 and compare the returned ETag to the calculated
// MD5 value. The Content-MD5 header is required for any request to upload an
// object with a retention period configured using Amazon S3 Object Lock. For more
// information about Amazon S3 Object Lock, see Amazon S3 Object Lock Overview
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html) in
// the Amazon Simple Storage Service Developer Guide. Server-side Encryption You
// can optionally request server-side encryption. With server-side encryption,
// Amazon S3 encrypts your data as it writes it to disks in its data centers and
// decrypts the data when you access it. You have the option to provide your own
// encryption key or use AWS managed encryption keys. For more information, see
// Using Server-Side Encryption
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html).
// Access Control List (ACL)-Specific Request Headers
//
// You can use headers to grant
// ACL- based permissions. By default, all objects are private. Only the owner has
// full access control. When adding a new object, you can grant permissions to
// individual AWS accounts or to predefined groups defined by Amazon S3. These
// permissions are then added to the ACL on the object. For more information, see
// Access Control List (ACL) Overview
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html) and Managing
// ACLs Using the REST API
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-using-rest-api.html).
// Storage Class Options By default, Amazon S3 uses the STANDARD storage class to
// store newly created objects. The STANDARD storage class provides high durability
// and high availability. Depending on performance needs, you can specify a
// different storage class. For more information, see Storage Classes
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html) in
// the Amazon S3 Service Developer Guide. Versioning If you enable versioning for a
// bucket, Amazon S3 automatically generates a unique version ID for the object
// being stored. Amazon S3 returns this ID in the response. When you enable
// versioning for a bucket, if Amazon S3 receives multiple write requests for the
// same object simultaneously, it stores all of the objects. For more information
// about versioning, see Adding Objects to Versioning Enabled Buckets
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/AddingObjectstoVersioningEnabledBuckets.html).
// For information about returning the versioning state of a bucket, see
// GetBucketVersioning. Related Resources
//
//     * CopyObject
//
//     * DeleteObject
func (c *Client) PutObject(ctx context.Context, params *PutObjectInput, optFns ...func(*Options)) (*PutObjectOutput, error) {
	if params == nil {
		params = &PutObjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutObject", params, optFns, addOperationPutObjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutObjectInput struct {

	// Bucket name to which the PUT operation was initiated. When using this API with
	// an access point, you must direct requests to the access point hostname. The
	// access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// operation using an access point through the AWS SDKs, you provide the access
	// point ARN in place of the bucket name. For more information about access point
	// ARNs, see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) in
	// the Amazon Simple Storage Service Developer Guide.
	//
	// This member is required.
	Bucket *string

	// Object key for which the PUT operation was initiated.
	//
	// This member is required.
	Key *string

	// The canned ACL to apply to the object. For more information, see Canned ACL
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL).
	ACL types.ObjectCannedACL

	// Object data.
	Body io.Reader

	// Can be used to specify caching behavior along the request/reply chain. For more
	// information, see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9
	// (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9).
	CacheControl *string

	// Specifies presentational information for the object. For more information, see
	// http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1
	// (http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1).
	ContentDisposition *string

	// Specifies what content encodings have been applied to the object and thus what
	// decoding mechanisms must be applied to obtain the media-type referenced by the
	// Content-Type header field. For more information, see
	// http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11
	// (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11).
	ContentEncoding *string

	// The language the content is in.
	ContentLanguage *string

	// Size of the body in bytes. This parameter is useful when the size of the body
	// cannot be determined automatically. For more information, see
	// http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.13
	// (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.13).
	ContentLength *int64

	// The base64-encoded 128-bit MD5 digest of the message (without the headers)
	// according to RFC 1864. This header can be used as a message integrity check to
	// verify that the data is the same data that was originally sent. Although it is
	// optional, we recommend using the Content-MD5 mechanism as an end-to-end
	// integrity check. For more information about REST request authentication, see
	// REST Authentication
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html).
	ContentMD5 *string

	// A standard MIME type describing the format of the contents. For more
	// information, see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17
	// (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
	ContentType *string

	// The date and time at which the object is no longer cacheable. For more
	// information, see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.21
	// (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.21).
	Expires *time.Time

	// Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.
	GrantFullControl *string

	// Allows grantee to read the object data and its metadata.
	GrantRead *string

	// Allows grantee to read the object ACL.
	GrantReadACP *string

	// Allows grantee to write the ACL for the applicable object.
	GrantWriteACP *string

	// A map of metadata to store with the object in S3.
	Metadata map[string]*string

	// Specifies whether a legal hold will be applied to this object. For more
	// information about S3 Object Lock, see Object Lock
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html).
	ObjectLockLegalHoldStatus types.ObjectLockLegalHoldStatus

	// The Object Lock mode that you want to apply to this object.
	ObjectLockMode types.ObjectLockMode

	// The date and time when you want this object's Object Lock to expire.
	ObjectLockRetainUntilDate *time.Time

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer

	// Specifies the algorithm to use to when encrypting the object (for example,
	// AES256).
	SSECustomerAlgorithm *string

	// Specifies the customer-provided encryption key for Amazon S3 to use in
	// encrypting data. This value is used to store the object and then it is
	// discarded; Amazon S3 does not store the encryption key. The key must be
	// appropriate for use with the algorithm specified in the
	// x-amz-server-side-encryption-customer-algorithm header.
	SSECustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	SSECustomerKeyMD5 *string

	// Specifies the AWS KMS Encryption Context to use for object encryption. The value
	// of this header is a base64-encoded UTF-8 string holding JSON with the encryption
	// context key-value pairs.
	SSEKMSEncryptionContext *string

	// If x-amz-server-side-encryption is present and has the value of aws:kms, this
	// header specifies the ID of the AWS Key Management Service (AWS KMS) symmetrical
	// customer managed customer master key (CMK) that was used for the object. If the
	// value of x-amz-server-side-encryption is aws:kms, this header specifies the ID
	// of the symmetric customer managed AWS KMS CMK that will be used for the object.
	// If you specify x-amz-server-side-encryption:aws:kms, but do not provide
	// x-amz-server-side-encryption-aws-kms-key-id, Amazon S3 uses the AWS managed CMK
	// in AWS to protect the data.
	SSEKMSKeyId *string

	// The server-side encryption algorithm used when storing this object in Amazon S3
	// (for example, AES256, aws:kms).
	ServerSideEncryption types.ServerSideEncryption

	// If you don't specify, S3 Standard is the default storage class. Amazon S3
	// supports other storage classes.
	StorageClass types.StorageClass

	// The tag-set for the object. The tag-set must be encoded as URL Query parameters.
	// (For example, "Key1=Value1")
	Tagging *string

	// If the bucket is configured as a website, redirects requests for this object to
	// another object in the same bucket or to an external URL. Amazon S3 stores the
	// value of this header in the object metadata. For information about object
	// metadata, see Object Key and Metadata
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html). In the
	// following example, the request header sets the redirect to an object
	// (anotherPage.html) in the same bucket: x-amz-website-redirect-location:
	// /anotherPage.html In the following example, the request header sets the object
	// redirect to another website: x-amz-website-redirect-location:
	// http://www.example.com/ For more information about website hosting in Amazon S3,
	// see Hosting Websites on Amazon S3
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html) and How to
	// Configure Website Page Redirects
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
	WebsiteRedirectLocation *string
}

type PutObjectOutput struct {

	// Entity tag for the uploaded object.
	ETag *string

	// If the expiration is configured for the object (see
	// PutBucketLifecycleConfiguration), the response includes this header. It includes
	// the expiry-date and rule-id key-value pairs that provide information about
	// object expiration. The value of the rule-id is URL encoded.
	Expiration *string

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm used.
	SSECustomerAlgorithm *string

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round-trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string

	// If present, specifies the AWS KMS Encryption Context to use for object
	// encryption. The value of this header is a base64-encoded UTF-8 string holding
	// JSON with the encryption context key-value pairs.
	SSEKMSEncryptionContext *string

	// If x-amz-server-side-encryption is present and has the value of aws:kms, this
	// header specifies the ID of the AWS Key Management Service (AWS KMS) symmetric
	// customer managed customer master key (CMK) that was used for the object.
	SSEKMSKeyId *string

	// If you specified server-side encryption either with an AWS KMS customer master
	// key (CMK) or Amazon S3-managed encryption key in your PUT request, the response
	// includes this header. It confirms the encryption algorithm that Amazon S3 used
	// to encrypt the object.
	ServerSideEncryption types.ServerSideEncryption

	// Version of the object.
	VersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutObjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutObject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutObject{}, middleware.After)
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
	addOpPutObjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutObject(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutObject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutObject",
	}
}
