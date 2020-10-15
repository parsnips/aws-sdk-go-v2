// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Domain used by SageMaker Studio. A domain consists of an associated
// directory, a list of authorized users, and a variety of security, application,
// policy, and Amazon Virtual Private Cloud (VPC) configurations. An AWS account is
// limited to one domain per region. Users within a domain can share notebook files
// and other artifacts with each other. When a domain is created, an Amazon Elastic
// File System (EFS) volume is also created for use by all of the users within the
// domain. Each user receives a private home directory within the EFS for
// notebooks, Git repositories, and data files. All traffic between the domain and
// the EFS volume is communicated through the specified subnet IDs. All other
// traffic goes over the Internet through an Amazon SageMaker system VPC. The EFS
// traffic uses the NFS/TCP protocol over port 2049. NFS traffic over TCP on port
// 2049 needs to be allowed in both inbound and outbound rules in order to launch a
// SageMaker Studio app successfully.
func (c *Client) CreateDomain(ctx context.Context, params *CreateDomainInput, optFns ...func(*Options)) (*CreateDomainOutput, error) {
	if params == nil {
		params = &CreateDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDomain", params, optFns, addOperationCreateDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDomainInput struct {

	// The mode of authentication that members use to access the domain.
	//
	// This member is required.
	AuthMode types.AuthMode

	// The default user settings.
	//
	// This member is required.
	DefaultUserSettings *types.UserSettings

	// A name for the domain.
	//
	// This member is required.
	DomainName *string

	// The VPC subnets to use for communication with the EFS volume.
	//
	// This member is required.
	SubnetIds []*string

	// The ID of the Amazon Virtual Private Cloud (VPC) to use for communication with
	// the EFS volume.
	//
	// This member is required.
	VpcId *string

	// The AWS Key Management Service (KMS) encryption key ID. Encryption with a
	// customer master key (CMK) is not supported.
	HomeEfsFileSystemKmsKeyId *string

	// Tags to associated with the Domain. Each tag consists of a key and an optional
	// value. Tag keys must be unique per resource. Tags are searchable using the
	// Search API.
	Tags []*types.Tag
}

type CreateDomainOutput struct {

	// The Amazon Resource Name (ARN) of the created domain.
	DomainArn *string

	// The URL to the created domain.
	Url *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDomain{}, middleware.After)
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
	addOpCreateDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDomain(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateDomain",
	}
}
