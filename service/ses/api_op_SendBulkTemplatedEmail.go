// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Composes an email message to multiple destinations. The message body is created
// using an email template. In order to send email using the SendBulkTemplatedEmail
// operation, your call to the API must meet the following requirements:
//
//     * The
// call must refer to an existing email template. You can create email templates
// using the CreateTemplate operation.
//
//     * The message must be sent from a
// verified email address or domain.
//
//     * If your account is still in the Amazon
// SES sandbox, you may only send to verified addresses or domains, or to email
// addresses associated with the Amazon SES Mailbox Simulator. For more
// information, see Verifying Email Addresses and Domains
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-addresses-and-domains.html)
// in the Amazon SES Developer Guide.
//
//     * The maximum message size is 10 MB.
//
//
// * Each Destination parameter must include at least one recipient email address.
// The recipient address can be a To: address, a CC: address, or a BCC: address. If
// a recipient email address is invalid (that is, it is not in the format
// UserName@[SubDomain.]Domain.TopLevelDomain), the entire message will be
// rejected, even if the message contains other recipients that are valid.
//
//     *
// The message may not include more than 50 recipients, across the To:, CC: and
// BCC: fields. If you need to send an email message to a larger audience, you can
// divide your recipient list into groups of 50 or fewer, and then call the
// SendBulkTemplatedEmail operation several times to send the message to each
// group.
//
//     * The number of destinations you can contact in a single call to the
// API may be limited by your account's maximum sending rate.
func (c *Client) SendBulkTemplatedEmail(ctx context.Context, params *SendBulkTemplatedEmailInput, optFns ...func(*Options)) (*SendBulkTemplatedEmailOutput, error) {
	if params == nil {
		params = &SendBulkTemplatedEmailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendBulkTemplatedEmail", params, optFns, addOperationSendBulkTemplatedEmailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendBulkTemplatedEmailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to send a templated email to multiple destinations using
// Amazon SES. For more information, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-personalized-email-api.html).
type SendBulkTemplatedEmailInput struct {

	// One or more Destination objects. All of the recipients in a Destination will
	// receive the same version of the email. You can specify up to 50 Destination
	// objects within a Destinations array.
	//
	// This member is required.
	Destinations []*types.BulkEmailDestination

	// The email address that is sending the email. This email address must be either
	// individually verified with Amazon SES, or from a domain that has been verified
	// with Amazon SES. For information about verifying identities, see the Amazon SES
	// Developer Guide
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-addresses-and-domains.html).
	// If you are sending on behalf of another user and have been permitted to do so by
	// a sending authorization policy, then you must also specify the SourceArn
	// parameter. For more information about sending authorization, see the Amazon SES
	// Developer Guide
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
	// Amazon SES does not support the SMTPUTF8 extension, as described in RFC6531
	// (https://tools.ietf.org/html/rfc6531). For this reason, the local part of a
	// source email address (the part of the email address that precedes the @ sign)
	// may only contain 7-bit ASCII characters
	// (https://en.wikipedia.org/wiki/Email_address#Local-part). If the domain part of
	// an address (the part after the @ sign) contains non-ASCII characters, they must
	// be encoded using Punycode, as described in RFC3492
	// (https://tools.ietf.org/html/rfc3492.html). The sender name (also known as the
	// friendly name) may contain non-ASCII characters. These characters must be
	// encoded using MIME encoded-word syntax, as described in RFC 2047
	// (https://tools.ietf.org/html/rfc2047). MIME encoded-word syntax uses the
	// following form: =?charset?encoding?encoded-text?=.
	//
	// This member is required.
	Source *string

	// The template to use when sending this email.
	//
	// This member is required.
	Template *string

	// The name of the configuration set to use when you send an email using
	// SendBulkTemplatedEmail.
	ConfigurationSetName *string

	// A list of tags, in the form of name/value pairs, to apply to an email that you
	// send to a destination using SendBulkTemplatedEmail.
	DefaultTags []*types.MessageTag

	// A list of replacement values to apply to the template when replacement data is
	// not specified in a Destination object. These values act as a default or fallback
	// option when no other data is available. The template data is a JSON object,
	// typically consisting of key-value pairs in which the keys correspond to
	// replacement tags in the email template.
	DefaultTemplateData *string

	// The reply-to email address(es) for the message. If the recipient replies to the
	// message, each reply-to address will receive the reply.
	ReplyToAddresses []*string

	// The email address that bounces and complaints will be forwarded to when feedback
	// forwarding is enabled. If the message cannot be delivered to the recipient, then
	// an error message will be returned from the recipient's ISP; this message will
	// then be forwarded to the email address specified by the ReturnPath parameter.
	// The ReturnPath parameter is never overwritten. This email address must be either
	// individually verified with Amazon SES, or from a domain that has been verified
	// with Amazon SES.
	ReturnPath *string

	// This parameter is used only for sending authorization. It is the ARN of the
	// identity that is associated with the sending authorization policy that permits
	// you to use the email address specified in the ReturnPath parameter. For example,
	// if the owner of example.com (which has ARN
	// arn:aws:ses:us-east-1:123456789012:identity/example.com) attaches a policy to it
	// that authorizes you to use feedback@example.com, then you would specify the
	// ReturnPathArn to be arn:aws:ses:us-east-1:123456789012:identity/example.com, and
	// the ReturnPath to be feedback@example.com. For more information about sending
	// authorization, see the Amazon SES Developer Guide
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
	ReturnPathArn *string

	// This parameter is used only for sending authorization. It is the ARN of the
	// identity that is associated with the sending authorization policy that permits
	// you to send for the email address specified in the Source parameter. For
	// example, if the owner of example.com (which has ARN
	// arn:aws:ses:us-east-1:123456789012:identity/example.com) attaches a policy to it
	// that authorizes you to send from user@example.com, then you would specify the
	// SourceArn to be arn:aws:ses:us-east-1:123456789012:identity/example.com, and the
	// Source to be user@example.com. For more information about sending authorization,
	// see the Amazon SES Developer Guide
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
	SourceArn *string

	// The ARN of the template to use when sending this email.
	TemplateArn *string
}

type SendBulkTemplatedEmailOutput struct {

	// The unique message identifier returned from the SendBulkTemplatedEmail action.
	//
	// This member is required.
	Status []*types.BulkEmailDestinationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSendBulkTemplatedEmailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSendBulkTemplatedEmail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSendBulkTemplatedEmail{}, middleware.After)
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
	addOpSendBulkTemplatedEmailValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSendBulkTemplatedEmail(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSendBulkTemplatedEmail(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "SendBulkTemplatedEmail",
	}
}
