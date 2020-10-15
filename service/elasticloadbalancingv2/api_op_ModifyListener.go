// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Replaces the specified properties of the specified listener. Any properties that
// you do not specify remain unchanged. Changing the protocol from HTTPS to HTTP,
// or from TLS to TCP, removes the security policy and default certificate
// properties. If you change the protocol from HTTP to HTTPS, or from TCP to TLS,
// you must add the security policy and default certificate properties. To add an
// item to a list, remove an item from a list, or update an item in a list, you
// must provide the entire list. For example, to add an action, specify a list with
// the current actions plus the new action.
func (c *Client) ModifyListener(ctx context.Context, params *ModifyListenerInput, optFns ...func(*Options)) (*ModifyListenerOutput, error) {
	if params == nil {
		params = &ModifyListenerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyListener", params, optFns, addOperationModifyListenerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyListenerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyListenerInput struct {

	// The Amazon Resource Name (ARN) of the listener.
	//
	// This member is required.
	ListenerArn *string

	// [TLS listeners] The name of the Application-Layer Protocol Negotiation (ALPN)
	// policy. You can specify one policy name. The following are the possible
	// values:
	//
	//     * HTTP1Only
	//
	//     * HTTP2Only
	//
	//     * HTTP2Optional
	//
	//     *
	// HTTP2Preferred
	//
	//     * None
	//
	// For more information, see ALPN Policies
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#alpn-policies)
	// in the Network Load Balancers Guide.
	AlpnPolicy []*string

	// [HTTPS and TLS listeners] The default certificate for the listener. You must
	// provide exactly one certificate. Set CertificateArn to the certificate ARN but
	// do not set IsDefault. To create a certificate list, use AddListenerCertificates.
	Certificates []*types.Certificate

	// The actions for the default rule. The rule must include one forward action or
	// one or more fixed-response actions. If the action type is forward, you specify
	// one or more target groups. The protocol of the target group must be HTTP or
	// HTTPS for an Application Load Balancer. The protocol of the target group must be
	// TCP, TLS, UDP, or TCP_UDP for a Network Load Balancer. [HTTPS listeners] If the
	// action type is authenticate-oidc, you authenticate users through an identity
	// provider that is OpenID Connect (OIDC) compliant. [HTTPS listeners] If the
	// action type is authenticate-cognito, you authenticate users through the user
	// pools supported by Amazon Cognito. [Application Load Balancer] If the action
	// type is redirect, you redirect specified client requests from one URL to
	// another. [Application Load Balancer] If the action type is fixed-response, you
	// drop specified client requests and return a custom HTTP response.
	DefaultActions []*types.Action

	// The port for connections from clients to the load balancer.
	Port *int32

	// The protocol for connections from clients to the load balancer. Application Load
	// Balancers support the HTTP and HTTPS protocols. Network Load Balancers support
	// the TCP, TLS, UDP, and TCP_UDP protocols.
	Protocol types.ProtocolEnum

	// [HTTPS and TLS listeners] The security policy that defines which protocols and
	// ciphers are supported. The following are the possible values:
	//
	//     *
	// ELBSecurityPolicy-2016-08
	//
	//     * ELBSecurityPolicy-TLS-1-0-2015-04
	//
	//     *
	// ELBSecurityPolicy-TLS-1-1-2017-01
	//
	//     * ELBSecurityPolicy-TLS-1-2-2017-01
	//
	//
	// * ELBSecurityPolicy-TLS-1-2-Ext-2018-06
	//
	//     * ELBSecurityPolicy-FS-2018-06
	//
	//
	// * ELBSecurityPolicy-FS-1-1-2019-08
	//
	//     * ELBSecurityPolicy-FS-1-2-2019-08
	//
	//
	// * ELBSecurityPolicy-FS-1-2-Res-2019-08
	//
	// For more information, see Security
	// Policies
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies)
	// in the Application Load Balancers Guide and Security Policies
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#describe-ssl-policies)
	// in the Network Load Balancers Guide.
	SslPolicy *string
}

type ModifyListenerOutput struct {

	// Information about the modified listener.
	Listeners []*types.Listener

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyListenerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyListener{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyListener{}, middleware.After)
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
	addOpModifyListenerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyListener(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opModifyListener(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "ModifyListener",
	}
}
