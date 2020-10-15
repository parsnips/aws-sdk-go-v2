// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Restores a certificate authority (CA) that is in the DELETED state. You can
// restore a CA during the period that you defined in the
// PermanentDeletionTimeInDays parameter of the DeleteCertificateAuthority action.
// Currently, you can specify 7 to 30 days. If you did not specify a
// PermanentDeletionTimeInDays value, by default you can restore the CA at any time
// in a 30 day period. You can check the time remaining in the restoration period
// of a private CA in the DELETED state by calling the DescribeCertificateAuthority
// or ListCertificateAuthorities actions. The status of a restored CA is set to its
// pre-deletion status when the RestoreCertificateAuthority action returns. To
// change its status to ACTIVE, call the UpdateCertificateAuthority action. If the
// private CA was in the PENDING_CERTIFICATE state at deletion, you must use the
// ImportCertificateAuthorityCertificate action to import a certificate authority
// into the private CA before it can be activated. You cannot restore a CA after
// the restoration period has ended.
func (c *Client) RestoreCertificateAuthority(ctx context.Context, params *RestoreCertificateAuthorityInput, optFns ...func(*Options)) (*RestoreCertificateAuthorityOutput, error) {
	if params == nil {
		params = &RestoreCertificateAuthorityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreCertificateAuthority", params, optFns, addOperationRestoreCertificateAuthorityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreCertificateAuthorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreCertificateAuthorityInput struct {

	// The Amazon Resource Name (ARN) that was returned when you called the
	// CreateCertificateAuthority action. This must be of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	//
	// This member is required.
	CertificateAuthorityArn *string
}

type RestoreCertificateAuthorityOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRestoreCertificateAuthorityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRestoreCertificateAuthority{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRestoreCertificateAuthority{}, middleware.After)
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
	addOpRestoreCertificateAuthorityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreCertificateAuthority(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRestoreCertificateAuthority(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "RestoreCertificateAuthority",
	}
}
