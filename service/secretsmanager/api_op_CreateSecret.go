// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new secret. A secret in Secrets Manager consists of both the protected
// secret data and the important information needed to manage the secret. Secrets
// Manager stores the encrypted secret data in one of a collection of "versions"
// associated with the secret. Each version contains a copy of the encrypted secret
// data. Each version is associated with one or more "staging labels" that identify
// where the version is in the rotation cycle. The SecretVersionsToStages field of
// the secret contains the mapping of staging labels to the active versions of the
// secret. Versions without a staging label are considered deprecated and not
// included in the list. You provide the secret data to be encrypted by putting
// text in either the SecretString parameter or binary data in the SecretBinary
// parameter, but not both. If you include SecretString or SecretBinary then
// Secrets Manager also creates an initial secret version and automatically
// attaches the staging label AWSCURRENT to the new version.
//
//     * If you call an
// operation to encrypt or decrypt the SecretString or SecretBinary for a secret in
// the same account as the calling user and that secret doesn't specify a AWS KMS
// encryption key, Secrets Manager uses the account's default AWS managed customer
// master key (CMK) with the alias aws/secretsmanager. If this key doesn't already
// exist in your account then Secrets Manager creates it for you automatically. All
// users and roles in the same AWS account automatically have access to use the
// default CMK. Note that if an Secrets Manager API call results in AWS creating
// the account's AWS-managed CMK, it can result in a one-time significant delay in
// returning the result.
//
//     * If the secret resides in a different AWS account
// from the credentials calling an API that requires encryption or decryption of
// the secret value then you must create and use a custom AWS KMS CMK because you
// can't access the default CMK for the account using credentials from a different
// AWS account. Store the ARN of the CMK in the secret when you create the secret
// or when you update it by including it in the KMSKeyId. If you call an API that
// must encrypt or decrypt SecretString or SecretBinary using credentials from a
// different account then the AWS KMS key policy must grant cross-account access to
// that other account's user or role for both the kms:GenerateDataKey and
// kms:Decrypt operations.
//
// Minimum permissions To run this command, you must have
// the following permissions:
//
//     * secretsmanager:CreateSecret
//
//     *
// kms:GenerateDataKey - needed only if you use a customer-managed AWS KMS key to
// encrypt the secret. You do not need this permission to use the account default
// AWS managed CMK for Secrets Manager.
//
//     * kms:Decrypt - needed only if you use
// a customer-managed AWS KMS key to encrypt the secret. You do not need this
// permission to use the account default AWS managed CMK for Secrets Manager.
//
//
// * secretsmanager:TagResource - needed only if you include the Tags
// parameter.
//
// Related operations
//
//     * To delete a secret, use DeleteSecret.
//
//
// * To modify an existing secret, use UpdateSecret.
//
//     * To create a new version
// of a secret, use PutSecretValue.
//
//     * To retrieve the encrypted secure string
// and secure binary values, use GetSecretValue.
//
//     * To retrieve all other
// details for a secret, use DescribeSecret. This does not include the encrypted
// secure string and secure binary values.
//
//     * To retrieve the list of secret
// versions associated with the current secret, use DescribeSecret and examine the
// SecretVersionsToStages response value.
func (c *Client) CreateSecret(ctx context.Context, params *CreateSecretInput, optFns ...func(*Options)) (*CreateSecretOutput, error) {
	if params == nil {
		params = &CreateSecretInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSecret", params, optFns, addOperationCreateSecretMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSecretOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSecretInput struct {

	// Specifies the friendly name of the new secret. The secret name must be ASCII
	// letters, digits, or the following characters : /_+=.@- Do not end your secret
	// name with a hyphen followed by six characters. If you do so, you risk confusion
	// and unexpected results when searching for a secret by partial ARN. Secrets
	// Manager automatically adds a hyphen and six random characters at the end of the
	// ARN.
	//
	// This member is required.
	Name *string

	// (Optional) If you include SecretString or SecretBinary, then an initial version
	// is created as part of the secret, and this parameter specifies a unique
	// identifier for the new version. If you use the AWS CLI or one of the AWS SDK to
	// call this operation, then you can leave this parameter empty. The CLI or SDK
	// generates a random UUID for you and includes it as the value for this parameter
	// in the request. If you don't use the SDK and instead generate a raw HTTP request
	// to the Secrets Manager service endpoint, then you must generate a
	// ClientRequestToken yourself for the new version and include the value in the
	// request. This value helps ensure idempotency. Secrets Manager uses this value to
	// prevent the accidental creation of duplicate versions if there are failures and
	// retries during a rotation. We recommend that you generate a UUID-type
	// (https://wikipedia.org/wiki/Universally_unique_identifier) value to ensure
	// uniqueness of your versions within the specified secret.
	//
	//     * If the
	// ClientRequestToken value isn't already associated with a version of the secret
	// then a new version of the secret is created.
	//
	//     * If a version with this value
	// already exists and the version SecretString and SecretBinary values are the same
	// as those in the request, then the request is ignored.
	//
	//     * If a version with
	// this value already exists and that version's SecretString and SecretBinary
	// values are different from those in the request then the request fails because
	// you cannot modify an existing version. Instead, use PutSecretValue to create a
	// new version.
	//
	// This value becomes the VersionId of the new version.
	ClientRequestToken *string

	// (Optional) Specifies a user-provided description of the secret.
	Description *string

	// (Optional) Specifies the ARN, Key ID, or alias of the AWS KMS customer master
	// key (CMK) to be used to encrypt the SecretString or SecretBinary values in the
	// versions stored in this secret. You can specify any of the supported ways to
	// identify a AWS KMS key ID. If you need to reference a CMK in a different
	// account, you can use only the key ARN or the alias ARN. If you don't specify
	// this value, then Secrets Manager defaults to using the AWS account's default CMK
	// (the one named aws/secretsmanager). If a AWS KMS CMK with that name doesn't yet
	// exist, then Secrets Manager creates it for you automatically the first time it
	// needs to encrypt a version's SecretString or SecretBinary fields. You can use
	// the account default CMK to encrypt and decrypt only if you call this operation
	// using credentials from the same account that owns the secret. If the secret
	// resides in a different account, then you must create a custom CMK and specify
	// the ARN in this field.
	KmsKeyId *string

	// (Optional) Specifies binary data that you want to encrypt and store in the new
	// version of the secret. To use this parameter in the command-line tools, we
	// recommend that you store your binary data in a file and then use the appropriate
	// technique for your tool to pass the contents of the file as a parameter. Either
	// SecretString or SecretBinary must have a value, but not both. They cannot both
	// be empty. This parameter is not available using the Secrets Manager console. It
	// can be accessed only by using the AWS CLI or one of the AWS SDKs.
	SecretBinary []byte

	// (Optional) Specifies text data that you want to encrypt and store in this new
	// version of the secret. Either SecretString or SecretBinary must have a value,
	// but not both. They cannot both be empty. If you create a secret by using the
	// Secrets Manager console then Secrets Manager puts the protected secret text in
	// only the SecretString parameter. The Secrets Manager console stores the
	// information as a JSON structure of key/value pairs that the Lambda rotation
	// function knows how to parse. For storing multiple values, we recommend that you
	// use a JSON text string argument and specify key/value pairs. For information on
	// how to format a JSON parameter for the various command line tool environments,
	// see Using JSON for Parameters
	// (https://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#cli-using-param-json)
	// in the AWS CLI User Guide. For example:
	// {"username":"bob","password":"abc123xyz456"} If your command-line tool or SDK
	// requires quotation marks around the parameter, you should use single quotes to
	// avoid confusion with the double quotes required in the JSON text.
	SecretString *string

	// (Optional) Specifies a list of user-defined tags that are attached to the
	// secret. Each tag is a "Key" and "Value" pair of strings. This operation only
	// appends tags to the existing list of tags. To remove tags, you must use
	// UntagResource.
	//
	//     * Secrets Manager tag key names are case sensitive. A tag
	// with the key "ABC" is a different tag from one with key "abc".
	//
	//     * If you
	// check tags in IAM policy Condition elements as part of your security strategy,
	// then adding or removing a tag can change permissions. If the successful
	// completion of this operation would result in you losing your permissions for
	// this secret, then this operation is blocked and returns an Access Denied
	// error.
	//
	// This parameter requires a JSON text string argument. For information on
	// how to format a JSON parameter for the various command line tool environments,
	// see Using JSON for Parameters
	// (https://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#cli-using-param-json)
	// in the AWS CLI User Guide. For example:
	// [{"Key":"CostCenter","Value":"12345"},{"Key":"environment","Value":"production"}]
	// If your command-line tool or SDK requires quotation marks around the parameter,
	// you should use single quotes to avoid confusion with the double quotes required
	// in the JSON text. The following basic restrictions apply to tags:
	//
	//     * Maximum
	// number of tags per secret—50
	//
	//     * Maximum key length—127 Unicode characters in
	// UTF-8
	//
	//     * Maximum value length—255 Unicode characters in UTF-8
	//
	//     * Tag
	// keys and values are case sensitive.
	//
	//     * Do not use the aws: prefix in your
	// tag names or values because AWS reserves it for AWS use. You can't edit or
	// delete tag names or values with this prefix. Tags with this prefix do not count
	// against your tags per secret limit.
	//
	//     * If you use your tagging schema across
	// multiple services and resources, remember other services might have restrictions
	// on allowed characters. Generally allowed characters: letters, spaces, and
	// numbers representable in UTF-8, plus the following special characters: + - = . _
	// : / @.
	Tags []*types.Tag
}

type CreateSecretOutput struct {

	// The Amazon Resource Name (ARN) of the secret that you just created. Secrets
	// Manager automatically adds several random characters to the name at the end of
	// the ARN when you initially create a secret. This affects only the ARN and not
	// the actual friendly name. This ensures that if you create a new secret with the
	// same name as an old secret that you previously deleted, then users with access
	// to the old secret don't automatically get access to the new secret because the
	// ARNs are different.
	ARN *string

	// The friendly name of the secret that you just created.
	Name *string

	// The unique identifier associated with the version of the secret you just
	// created.
	VersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateSecretMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSecret{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSecret{}, middleware.After)
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
	addIdempotencyToken_opCreateSecretMiddleware(stack, options)
	addOpCreateSecretValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSecret(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpCreateSecret struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateSecret) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateSecret) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateSecretInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateSecretInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateSecretMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateSecret{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateSecret(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "secretsmanager",
		OperationName: "CreateSecret",
	}
}
