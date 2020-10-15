// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a delivery stream and its data. To check the state of a delivery stream,
// use DescribeDeliveryStream. You can delete a delivery stream only if it is in
// one of the following states: ACTIVE, DELETING, CREATING_FAILED, or
// DELETING_FAILED. You can't delete a delivery stream that is in the CREATING
// state. While the deletion request is in process, the delivery stream is in the
// DELETING state. While the delivery stream is in the DELETING state, the service
// might continue to accept records, but it doesn't make any guarantees with
// respect to delivering the data. Therefore, as a best practice, first stop any
// applications that are sending records before you delete a delivery stream.
func (c *Client) DeleteDeliveryStream(ctx context.Context, params *DeleteDeliveryStreamInput, optFns ...func(*Options)) (*DeleteDeliveryStreamOutput, error) {
	if params == nil {
		params = &DeleteDeliveryStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDeliveryStream", params, optFns, addOperationDeleteDeliveryStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDeliveryStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDeliveryStreamInput struct {

	// The name of the delivery stream.
	//
	// This member is required.
	DeliveryStreamName *string

	// Set this to true if you want to delete the delivery stream even if Kinesis Data
	// Firehose is unable to retire the grant for the CMK. Kinesis Data Firehose might
	// be unable to retire the grant due to a customer error, such as when the CMK or
	// the grant are in an invalid state. If you force deletion, you can then use the
	// RevokeGrant
	// (https://docs.aws.amazon.com/kms/latest/APIReference/API_RevokeGrant.html)
	// operation to revoke the grant you gave to Kinesis Data Firehose. If a failure to
	// retire the grant happens due to an AWS KMS issue, Kinesis Data Firehose keeps
	// retrying the delete operation. The default value is false.
	AllowForceDelete *bool
}

type DeleteDeliveryStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDeliveryStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDeliveryStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDeliveryStream{}, middleware.After)
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
	addOpDeleteDeliveryStreamValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDeliveryStream(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteDeliveryStream(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "firehose",
		OperationName: "DeleteDeliveryStream",
	}
}
