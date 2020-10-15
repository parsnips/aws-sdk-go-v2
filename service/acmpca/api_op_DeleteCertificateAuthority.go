// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a private certificate authority (CA). You must provide the Amazon
// Resource Name (ARN) of the private CA that you want to delete. You can find the
// ARN by calling the ListCertificateAuthorities action. Deleting a CA will
// invalidate other CAs and certificates below it in your CA hierarchy. Before you
// can delete a CA that you have created and activated, you must disable it. To do
// this, call the UpdateCertificateAuthority action and set the
// CertificateAuthorityStatus parameter to DISABLED. Additionally, you can delete a
// CA if you are waiting for it to be created (that is, the status of the CA is
// CREATING). You can also delete it if the CA has been created but you haven't yet
// imported the signed certificate into ACM Private CA (that is, the status of the
// CA is PENDING_CERTIFICATE). When you successfully call
// DeleteCertificateAuthority, the CA's status changes to DELETED. However, the CA
// won't be permanently deleted until the restoration period has passed. By
// default, if you do not set the PermanentDeletionTimeInDays parameter, the CA
// remains restorable for 30 days. You can set the parameter from 7 to 30 days. The
// DescribeCertificateAuthority action returns the time remaining in the
// restoration window of a private CA in the DELETED state. To restore an eligible
// CA, call the RestoreCertificateAuthority action.
func (c *Client) DeleteCertificateAuthority(ctx context.Context, params *DeleteCertificateAuthorityInput, optFns ...func(*Options)) (*DeleteCertificateAuthorityOutput, error) {
	if params == nil {
		params = &DeleteCertificateAuthorityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCertificateAuthority", params, optFns, addOperationDeleteCertificateAuthorityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCertificateAuthorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCertificateAuthorityInput struct {

	// The Amazon Resource Name (ARN) that was returned when you called
	// CreateCertificateAuthority. This must have the following form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	// .
	//
	// This member is required.
	CertificateAuthorityArn *string

	// The number of days to make a CA restorable after it has been deleted. This can
	// be anywhere from 7 to 30 days, with 30 being the default.
	PermanentDeletionTimeInDays *int32
}

type DeleteCertificateAuthorityOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteCertificateAuthorityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteCertificateAuthority{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteCertificateAuthority{}, middleware.After)
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
	addOpDeleteCertificateAuthorityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCertificateAuthority(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteCertificateAuthority(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "DeleteCertificateAuthority",
	}
}
