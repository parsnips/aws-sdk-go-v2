// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/snowball/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// While a cluster's ClusterState value is in the AwaitingQuorum state, you can
// update some of the information associated with a cluster. Once the cluster
// changes to a different job state, usually 60 minutes after the cluster being
// created, this action is no longer available.
func (c *Client) UpdateCluster(ctx context.Context, params *UpdateClusterInput, optFns ...func(*Options)) (*UpdateClusterOutput, error) {
	if params == nil {
		params = &UpdateClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCluster", params, optFns, addOperationUpdateClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateClusterInput struct {

	// The cluster ID of the cluster that you want to update, for example
	// CID123e4567-e89b-12d3-a456-426655440000.
	//
	// This member is required.
	ClusterId *string

	// The ID of the updated Address object.
	AddressId *string

	// The updated description of this cluster.
	Description *string

	// The updated ID for the forwarding address for a cluster. This field is not
	// supported in most regions.
	ForwardingAddressId *string

	// The new or updated Notification object.
	Notification *types.Notification

	// The updated arrays of JobResource objects that can include updated S3Resource
	// objects or LambdaResource objects.
	Resources *types.JobResource

	// The new role Amazon Resource Name (ARN) that you want to associate with this
	// cluster. To create a role ARN, use the CreateRole
	// (https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html) API
	// action in AWS Identity and Access Management (IAM).
	RoleARN *string

	// The updated shipping option value of this cluster's ShippingDetails object.
	ShippingOption types.ShippingOption
}

type UpdateClusterOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCluster{}, middleware.After)
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
	addOpUpdateClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCluster(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "UpdateCluster",
	}
}
