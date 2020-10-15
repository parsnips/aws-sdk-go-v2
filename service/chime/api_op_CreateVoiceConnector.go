// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an Amazon Chime Voice Connector under the administrator's AWS account.
// You can choose to create an Amazon Chime Voice Connector in a specific AWS
// Region. Enabling CreateVoiceConnectorRequest$RequireEncryption configures your
// Amazon Chime Voice Connector to use TLS transport for SIP signaling and Secure
// RTP (SRTP) for media. Inbound calls use TLS transport, and unencrypted outbound
// calls are blocked.
func (c *Client) CreateVoiceConnector(ctx context.Context, params *CreateVoiceConnectorInput, optFns ...func(*Options)) (*CreateVoiceConnectorOutput, error) {
	if params == nil {
		params = &CreateVoiceConnectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVoiceConnector", params, optFns, addOperationCreateVoiceConnectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVoiceConnectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVoiceConnectorInput struct {

	// The name of the Amazon Chime Voice Connector.
	//
	// This member is required.
	Name *string

	// When enabled, requires encryption for the Amazon Chime Voice Connector.
	//
	// This member is required.
	RequireEncryption *bool

	// The AWS Region in which the Amazon Chime Voice Connector is created. Default
	// value: us-east-1.
	AwsRegion types.VoiceConnectorAwsRegion
}

type CreateVoiceConnectorOutput struct {

	// The Amazon Chime Voice Connector details.
	VoiceConnector *types.VoiceConnector

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateVoiceConnectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateVoiceConnector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateVoiceConnector{}, middleware.After)
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
	addOpCreateVoiceConnectorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVoiceConnector(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateVoiceConnector(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreateVoiceConnector",
	}
}
