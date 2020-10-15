// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provisions an IPv4 or IPv6 address range for use with your AWS resources through
// bring your own IP addresses (BYOIP) and creates a corresponding address pool.
// After the address range is provisioned, it is ready to be advertised using
// AdvertiseByoipCidr. AWS verifies that you own the address range and are
// authorized to advertise it. You must ensure that the address range is registered
// to you and that you created an RPKI ROA to authorize Amazon ASNs 16509 and 14618
// to advertise the address range. For more information, see Bring Your Own IP
// Addresses (BYOIP)
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html) in the
// Amazon Elastic Compute Cloud User Guide. Provisioning an address range is an
// asynchronous operation, so the call returns immediately, but the address range
// is not ready to use until its status changes from pending-provision to
// provisioned. To monitor the status of an address range, use DescribeByoipCidrs.
// To allocate an Elastic IP address from your IPv4 address pool, use
// AllocateAddress with either the specific address from the address pool or the ID
// of the address pool.
func (c *Client) ProvisionByoipCidr(ctx context.Context, params *ProvisionByoipCidrInput, optFns ...func(*Options)) (*ProvisionByoipCidrOutput, error) {
	if params == nil {
		params = &ProvisionByoipCidrInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ProvisionByoipCidr", params, optFns, addOperationProvisionByoipCidrMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ProvisionByoipCidrOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ProvisionByoipCidrInput struct {

	// The public IPv4 or IPv6 address range, in CIDR notation. The most specific IPv4
	// prefix that you can specify is /24. The most specific IPv6 prefix you can
	// specify is /56. The address range cannot overlap with another address range that
	// you've brought to this or another Region.
	//
	// This member is required.
	Cidr *string

	// A signed document that proves that you are authorized to bring the specified IP
	// address range to Amazon using BYOIP.
	CidrAuthorizationContext *types.CidrAuthorizationContext

	// A description for the address range and the address pool.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The tags to apply to the address pool.
	PoolTagSpecifications []*types.TagSpecification

	// (IPv6 only) Indicate whether the address range will be publicly advertised to
	// the internet. Default: true
	PubliclyAdvertisable *bool
}

type ProvisionByoipCidrOutput struct {

	// Information about the address range.
	ByoipCidr *types.ByoipCidr

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationProvisionByoipCidrMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpProvisionByoipCidr{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpProvisionByoipCidr{}, middleware.After)
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
	addOpProvisionByoipCidrValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opProvisionByoipCidr(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opProvisionByoipCidr(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ProvisionByoipCidr",
	}
}
