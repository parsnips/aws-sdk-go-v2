// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disconnects a volume from an iSCSI connection and then detaches the volume from
// the specified gateway. Detaching and attaching a volume enables you to recover
// your data from one gateway to a different gateway without creating a snapshot.
// It also makes it easier to move your volumes from an on-premises gateway to a
// gateway hosted on an Amazon EC2 instance. This operation is only supported in
// the volume gateway type.
func (c *Client) DetachVolume(ctx context.Context, params *DetachVolumeInput, optFns ...func(*Options)) (*DetachVolumeOutput, error) {
	if params == nil {
		params = &DetachVolumeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachVolume", params, optFns, addOperationDetachVolumeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// AttachVolumeInput
type DetachVolumeInput struct {

	// The Amazon Resource Name (ARN) of the volume to detach from the gateway.
	//
	// This member is required.
	VolumeARN *string

	// Set to true to forcibly remove the iSCSI connection of the target volume and
	// detach the volume. The default is false. If this value is set to false, you must
	// manually disconnect the iSCSI connection from the target volume. Valid Values:
	// true | false
	ForceDetach *bool
}

// AttachVolumeOutput
type DetachVolumeOutput struct {

	// The Amazon Resource Name (ARN) of the volume that was detached.
	VolumeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDetachVolumeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetachVolume{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetachVolume{}, middleware.After)
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
	addOpDetachVolumeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetachVolume(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDetachVolume(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DetachVolume",
	}
}
