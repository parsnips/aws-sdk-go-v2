// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Copies a snapshot of a DB cluster. To copy a DB cluster snapshot from a shared
// manual DB cluster snapshot, SourceDBClusterSnapshotIdentifier must be the Amazon
// Resource Name (ARN) of the shared DB cluster snapshot.
func (c *Client) CopyDBClusterSnapshot(ctx context.Context, params *CopyDBClusterSnapshotInput, optFns ...func(*Options)) (*CopyDBClusterSnapshotOutput, error) {
	if params == nil {
		params = &CopyDBClusterSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyDBClusterSnapshot", params, optFns, addOperationCopyDBClusterSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyDBClusterSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyDBClusterSnapshotInput struct {

	// The identifier of the DB cluster snapshot to copy. This parameter is not
	// case-sensitive. You can't copy from one AWS Region to another. Constraints:
	//
	//
	// * Must specify a valid system snapshot in the "available" state.
	//
	//     * Specify
	// a valid DB snapshot identifier.
	//
	// Example: my-cluster-snapshot1
	//
	// This member is required.
	SourceDBClusterSnapshotIdentifier *string

	// The identifier of the new DB cluster snapshot to create from the source DB
	// cluster snapshot. This parameter is not case-sensitive. Constraints:
	//
	//     * Must
	// contain from 1 to 63 letters, numbers, or hyphens.
	//
	//     * First character must
	// be a letter.
	//
	//     * Cannot end with a hyphen or contain two consecutive
	// hyphens.
	//
	// Example: my-cluster-snapshot2
	//
	// This member is required.
	TargetDBClusterSnapshotIdentifier *string

	// True to copy all tags from the source DB cluster snapshot to the target DB
	// cluster snapshot, and otherwise false. The default is false.
	CopyTags *bool

	// The AWS AWS KMS key ID for an encrypted DB cluster snapshot. The KMS key ID is
	// the Amazon Resource Name (ARN), KMS key identifier, or the KMS key alias for the
	// KMS encryption key. If you copy an encrypted DB cluster snapshot from your AWS
	// account, you can specify a value for KmsKeyId to encrypt the copy with a new KMS
	// encryption key. If you don't specify a value for KmsKeyId, then the copy of the
	// DB cluster snapshot is encrypted with the same KMS key as the source DB cluster
	// snapshot. If you copy an encrypted DB cluster snapshot that is shared from
	// another AWS account, then you must specify a value for KmsKeyId. KMS encryption
	// keys are specific to the AWS Region that they are created in, and you can't use
	// encryption keys from one AWS Region in another AWS Region. You cannot encrypt an
	// unencrypted DB cluster snapshot when you copy it. If you try to copy an
	// unencrypted DB cluster snapshot and specify a value for the KmsKeyId parameter,
	// an error is returned.
	KmsKeyId *string

	// Not currently supported.
	PreSignedUrl *string

	// The tags to assign to the new DB cluster snapshot copy.
	Tags []*types.Tag
}

type CopyDBClusterSnapshotOutput struct {

	// Contains the details for an Amazon Neptune DB cluster snapshot This data type is
	// used as a response element in the DescribeDBClusterSnapshots action.
	DBClusterSnapshot *types.DBClusterSnapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCopyDBClusterSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCopyDBClusterSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCopyDBClusterSnapshot{}, middleware.After)
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
	addOpCopyDBClusterSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCopyDBClusterSnapshot(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCopyDBClusterSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CopyDBClusterSnapshot",
	}
}
