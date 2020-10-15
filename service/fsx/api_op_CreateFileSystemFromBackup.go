// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new Amazon FSx file system from an existing Amazon FSx backup. If a
// file system with the specified client request token exists and the parameters
// match, this operation returns the description of the file system. If a client
// request token specified by the file system exists and the parameters don't
// match, this call returns IncompatibleParameterError. If a file system with the
// specified client request token doesn't exist, this operation does the
// following:
//
//     * Creates a new Amazon FSx file system from backup with an
// assigned ID, and an initial lifecycle state of CREATING.
//
//     * Returns the
// description of the file system.
//
// Parameters like Active Directory, default share
// name, automatic backup, and backup settings default to the parameters of the
// file system that was backed up, unless overridden. You can explicitly supply
// other settings. By using the idempotent operation, you can retry a
// CreateFileSystemFromBackup call without the risk of creating an extra file
// system. This approach can be useful when an initial call fails in a way that
// makes it unclear whether a file system was created. Examples are if a transport
// level timeout occurred, or your connection was reset. If you use the same client
// request token and the initial call created a file system, the client receives
// success as long as the parameters are the same. The CreateFileSystemFromBackup
// call returns while the file system's lifecycle state is still CREATING. You can
// check the file-system creation status by calling the DescribeFileSystems
// operation, which returns the file system state along with other information.
func (c *Client) CreateFileSystemFromBackup(ctx context.Context, params *CreateFileSystemFromBackupInput, optFns ...func(*Options)) (*CreateFileSystemFromBackupOutput, error) {
	if params == nil {
		params = &CreateFileSystemFromBackupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFileSystemFromBackup", params, optFns, addOperationCreateFileSystemFromBackupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFileSystemFromBackupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request object for the CreateFileSystemFromBackup operation.
type CreateFileSystemFromBackupInput struct {

	// The ID of the backup. Specifies the backup to use if you're creating a file
	// system from an existing backup.
	//
	// This member is required.
	BackupId *string

	// Specifies the IDs of the subnets that the file system will be accessible from.
	// For Windows MULTI_AZ_1 file system deployment types, provide exactly two subnet
	// IDs, one for the preferred file server and one for the standby file server. You
	// specify one of these subnets as the preferred subnet using the
	// WindowsConfiguration > PreferredSubnetID property. For Windows SINGLE_AZ_1 and
	// SINGLE_AZ_2 deployment types and Lustre file systems, provide exactly one subnet
	// ID. The file server is launched in that subnet's Availability Zone.
	//
	// This member is required.
	SubnetIds []*string

	// A string of up to 64 ASCII characters that Amazon FSx uses to ensure idempotent
	// creation. This string is automatically filled on your behalf when you use the
	// AWS Command Line Interface (AWS CLI) or an AWS SDK.
	ClientRequestToken *string

	// The Lustre configuration for the file system being created.
	LustreConfiguration *types.CreateFileSystemLustreConfiguration

	// A list of IDs for the security groups that apply to the specified network
	// interfaces created for file system access. These security groups apply to all
	// network interfaces. This value isn't returned in later DescribeFileSystem
	// requests.
	SecurityGroupIds []*string

	// Sets the storage type for the Windows file system you're creating from a backup.
	// Valid values are SSD and HDD.
	//
	//     * Set to SSD to use solid state drive
	// storage. Supported on all Windows deployment types.
	//
	//     * Set to HDD to use
	// hard disk drive storage. Supported on SINGLE_AZ_2 and MULTI_AZ_1 Windows file
	// system deployment types.
	//
	// Default value is SSD. HDD and SSD storage types have
	// different minimum storage capacity requirements. A restored file system's
	// storage capacity is tied to the file system that was backed up. You can create a
	// file system that uses HDD storage from a backup of a file system that used SSD
	// storage only if the original SSD file system had a storage capacity of at least
	// 2000 GiB.
	StorageType types.StorageType

	// The tags to be applied to the file system at file system creation. The key value
	// of the Name tag appears in the console as the file system name.
	Tags []*types.Tag

	// The configuration for this Microsoft Windows file system.
	WindowsConfiguration *types.CreateFileSystemWindowsConfiguration
}

// The response object for the CreateFileSystemFromBackup operation.
type CreateFileSystemFromBackupOutput struct {

	// A description of the file system.
	FileSystem *types.FileSystem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateFileSystemFromBackupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFileSystemFromBackup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFileSystemFromBackup{}, middleware.After)
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
	addIdempotencyToken_opCreateFileSystemFromBackupMiddleware(stack, options)
	addOpCreateFileSystemFromBackupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFileSystemFromBackup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpCreateFileSystemFromBackup struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateFileSystemFromBackup) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateFileSystemFromBackup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateFileSystemFromBackupInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateFileSystemFromBackupInput ")
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
func addIdempotencyToken_opCreateFileSystemFromBackupMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateFileSystemFromBackup{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateFileSystemFromBackup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "CreateFileSystemFromBackup",
	}
}
