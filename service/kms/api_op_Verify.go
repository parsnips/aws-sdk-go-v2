// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Verifies a digital signature that was generated by the Sign operation.
// Verification confirms that an authorized user signed the message with the
// specified CMK and signing algorithm, and the message hasn't changed since it was
// signed. If the signature is verified, the value of the SignatureValid field in
// the response is True. If the signature verification fails, the Verify operation
// fails with an KMSInvalidSignatureException exception. A digital signature is
// generated by using the private key in an asymmetric CMK. The signature is
// verified by using the public key in the same asymmetric CMK. For information
// about symmetric and asymmetric CMKs, see Using Symmetric and Asymmetric CMKs
// (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html)
// in the AWS Key Management Service Developer Guide. To verify a digital
// signature, you can use the Verify operation. Specify the same asymmetric CMK,
// message, and signing algorithm that were used to produce the signature. You can
// also verify the digital signature by using the public key of the CMK outside of
// AWS KMS. Use the GetPublicKey operation to download the public key in the
// asymmetric CMK and then use the public key to verify the signature outside of
// AWS KMS. The advantage of using the Verify operation is that it is performed
// within AWS KMS. As a result, it's easy to call, the operation is performed
// within the FIPS boundary, it is logged in AWS CloudTrail, and you can use key
// policy and IAM policy to determine who is authorized to use the CMK to verify
// signatures. The CMK that you use for this operation must be in a compatible key
// state. For details, see How Key State Affects Use of a Customer Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide.
func (c *Client) Verify(ctx context.Context, params *VerifyInput, optFns ...func(*Options)) (*VerifyOutput, error) {
	if params == nil {
		params = &VerifyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Verify", params, optFns, addOperationVerifyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*VerifyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type VerifyInput struct {

	// Identifies the asymmetric CMK that will be used to verify the signature. This
	// must be the same CMK that was used to generate the signature. If you specify a
	// different CMK, the signature verification fails. To specify a CMK, use its key
	// ID, Amazon Resource Name (ARN), alias name, or alias ARN. When using an alias
	// name, prefix it with "alias/". To specify a CMK in a different AWS account, you
	// must use the key ARN or alias ARN. For example:
	//
	//     * Key ID:
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//     * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//
	// * Alias name: alias/ExampleAlias
	//
	//     * Alias ARN:
	// arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key
	// ARN for a CMK, use ListKeys or DescribeKey. To get the alias name and alias ARN,
	// use ListAliases.
	//
	// This member is required.
	KeyId *string

	// Specifies the message that was signed. You can submit a raw message of up to
	// 4096 bytes, or a hash digest of the message. If you submit a digest, use the
	// MessageType parameter with a value of DIGEST. If the message specified here is
	// different from the message that was signed, the signature verification fails. A
	// message and its hash digest are considered to be the same message.
	//
	// This member is required.
	Message []byte

	// The signature that the Sign operation generated.
	//
	// This member is required.
	Signature []byte

	// The signing algorithm that was used to sign the message. If you submit a
	// different algorithm, the signature verification fails.
	//
	// This member is required.
	SigningAlgorithm types.SigningAlgorithmSpec

	// A list of grant tokens. For more information, see Grant Tokens
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []*string

	// Tells AWS KMS whether the value of the Message parameter is a message or message
	// digest. The default value, RAW, indicates a message. To indicate a message
	// digest, enter DIGEST. Use the DIGEST value only when the value of the Message
	// parameter is a message digest. If you use the DIGEST value with a raw message,
	// the security of the verification operation can be compromised.
	MessageType types.MessageType
}

type VerifyOutput struct {

	// The Amazon Resource Name (key ARN
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN))
	// of the asymmetric CMK that was used to verify the signature.
	KeyId *string

	// A Boolean value that indicates whether the signature was verified. A value of
	// True indicates that the Signature was produced by signing the Message with the
	// specified KeyID and SigningAlgorithm. If the signature is not verified, the
	// Verify operation fails with a KMSInvalidSignatureException exception.
	SignatureValid *bool

	// The signing algorithm that was used to verify the signature.
	SigningAlgorithm types.SigningAlgorithmSpec

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationVerifyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpVerify{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpVerify{}, middleware.After)
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
	addOpVerifyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opVerify(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opVerify(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "Verify",
	}
}
