// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an SSL/TLS certificate for a Amazon Lightsail content delivery network
// (CDN) distribution. After the certificate is created, use the
// AttachCertificateToDistribution action to attach the certificate to your
// distribution. Only certificates created in the us-east-1 AWS Region can be
// attached to Lightsail distributions. Lightsail distributions are global
// resources that can reference an origin in any AWS Region, and distribute its
// content globally. However, all distributions are located in the us-east-1
// Region.
func (c *Client) CreateCertificate(ctx context.Context, params *CreateCertificateInput, optFns ...func(*Options)) (*CreateCertificateOutput, error) {
	if params == nil {
		params = &CreateCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCertificate", params, optFns, addOperationCreateCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCertificateInput struct {

	// The name for the certificate.
	//
	// This member is required.
	CertificateName *string

	// The domain name (e.g., example.com) for the certificate.
	//
	// This member is required.
	DomainName *string

	// An array of strings that specify the alternate domains (e.g., example2.com) and
	// subdomains (e.g., blog.example.com) for the certificate. You can specify a
	// maximum of nine alternate domains (in addition to the primary domain name).
	// Wildcard domain entries (e.g., *.example.com) are not supported.
	SubjectAlternativeNames []*string

	// The tag keys and optional values to add to the certificate during create. Use
	// the TagResource action to tag a resource after it's created.
	Tags []*types.Tag
}

type CreateCertificateOutput struct {

	// An object that describes the certificate created.
	Certificate *types.CertificateSummary

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCertificate{}, middleware.After)
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
	addOpCreateCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCertificate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateCertificate",
	}
}
