// Code generated by smithy-go-codegen DO NOT EDIT.

package ebs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ebs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a new Amazon EBS snapshot. The new snapshot enters the pending state
// after the request completes. After creating the snapshot, use  PutSnapshotBlock
// (https://docs.aws.amazon.com/ebs/latest/APIReference/API_PutSnapshotBlock.html)
// to write blocks of data to the snapshot.
func (c *Client) StartSnapshot(ctx context.Context, params *StartSnapshotInput, optFns ...func(*Options)) (*StartSnapshotOutput, error) {
	if params == nil {
		params = &StartSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSnapshot", params, optFns, addOperationStartSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSnapshotInput struct {

	// The size of the volume, in GiB. The maximum size is 16384 GiB (16 TiB).
	//
	// This member is required.
	VolumeSize *int64

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. Idempotency ensures that an API request completes only once.
	// With an idempotent request, if the original request completes successfully. The
	// subsequent retries with the same client token return the result from the
	// original successful request and they have no additional effect. If you do not
	// specify a client token, one is automatically generated by the AWS SDK. For more
	// information, see  Idempotency for StartSnapshot API
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-direct-api-idempotency.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	ClientToken *string

	// A description for the snapshot.
	Description *string

	// Indicates whether to encrypt the snapshot. To create an encrypted snapshot,
	// specify true. To create an unencrypted snapshot, omit this parameter. If you
	// specify a value for ParentSnapshotId, omit this parameter. If you specify true,
	// the snapshot is encrypted using the CMK specified using the KmsKeyArn parameter.
	// If no value is specified for KmsKeyArn, the default CMK for your account is
	// used. If no default CMK has been specified for your account, the AWS managed CMK
	// is used. To set a default CMK for your account, use  ModifyEbsDefaultKmsKeyId
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyEbsDefaultKmsKeyId.html).
	// If your account is enabled for encryption by default, you cannot set this
	// parameter to false. In this case, you can omit this parameter. For more
	// information, see  Using encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-accessing-snapshot.html#ebsapis-using-encryption)
	// in the Amazon Elastic Compute Cloud User Guide.
	Encrypted *bool

	// The Amazon Resource Name (ARN) of the AWS Key Management Service (AWS KMS)
	// customer master key (CMK) to be used to encrypt the snapshot. If you do not
	// specify a CMK, the default AWS managed CMK is used. If you specify a
	// ParentSnapshotId, omit this parameter; the snapshot will be encrypted using the
	// same CMK that was used to encrypt the parent snapshot. If Encrypted is set to
	// true, you must specify a CMK ARN.
	KmsKeyArn *string

	// The ID of the parent snapshot. If there is no parent snapshot, or if you are
	// creating the first snapshot for an on-premises volume, omit this parameter. If
	// your account is enabled for encryption by default, you cannot use an unencrypted
	// snapshot as a parent snapshot. You must first create an encrypted copy of the
	// parent snapshot using CopySnapshot
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CopySnapshot.html).
	ParentSnapshotId *string

	// The tags to apply to the snapshot.
	Tags []*types.Tag

	// The amount of time (in minutes) after which the snapshot is automatically
	// cancelled if:
	//
	//     * No blocks are written to the snapshot.
	//
	//     * The snapshot
	// is not completed after writing the last block of data.
	//
	// If no value is
	// specified, the timeout defaults to 60 minutes.
	Timeout *int32
}

type StartSnapshotOutput struct {

	// The size of the blocks in the snapshot, in bytes.
	BlockSize *int32

	// The description of the snapshot.
	Description *string

	// The Amazon Resource Name (ARN) of the AWS Key Management Service (AWS KMS)
	// customer master key (CMK) used to encrypt the snapshot.
	KmsKeyArn *string

	// The AWS account ID of the snapshot owner.
	OwnerId *string

	// The ID of the parent snapshot.
	ParentSnapshotId *string

	// The ID of the snapshot.
	SnapshotId *string

	// The timestamp when the snapshot was created.
	StartTime *time.Time

	// The status of the snapshot.
	Status types.Status

	// The tags applied to the snapshot. You can specify up to 50 tags per snapshot.
	// For more information, see  Tagging your Amazon EC2 resources
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	Tags []*types.Tag

	// The size of the volume, in GiB.
	VolumeSize *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartSnapshot{}, middleware.After)
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
	addIdempotencyToken_opStartSnapshotMiddleware(stack, options)
	addOpStartSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartSnapshot(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpStartSnapshot struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartSnapshot) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartSnapshot) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartSnapshotInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartSnapshotInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartSnapshotMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpStartSnapshot{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ebs",
		OperationName: "StartSnapshot",
	}
}
