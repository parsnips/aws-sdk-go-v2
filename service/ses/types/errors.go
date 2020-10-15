// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
)

// Indicates that email sending is disabled for your entire Amazon SES account. You
// can enable or disable email sending for your Amazon SES account using
// UpdateAccountSendingEnabled.
type AccountSendingPausedException struct {
	Message *string
}

func (e *AccountSendingPausedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccountSendingPausedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccountSendingPausedException) ErrorCode() string             { return "AccountSendingPausedException" }
func (e *AccountSendingPausedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that a resource could not be created because of a naming conflict.
type AlreadyExistsException struct {
	Message *string

	Name *string
}

func (e *AlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AlreadyExistsException) ErrorCode() string             { return "AlreadyExistsException" }
func (e *AlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the delete operation could not be completed.
type CannotDeleteException struct {
	Message *string

	Name *string
}

func (e *CannotDeleteException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CannotDeleteException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CannotDeleteException) ErrorCode() string             { return "CannotDeleteException" }
func (e *CannotDeleteException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the configuration set could not be created because of a naming
// conflict.
type ConfigurationSetAlreadyExistsException struct {
	Message *string

	ConfigurationSetName *string
}

func (e *ConfigurationSetAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConfigurationSetAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConfigurationSetAlreadyExistsException) ErrorCode() string {
	return "ConfigurationSetAlreadyExistsException"
}
func (e *ConfigurationSetAlreadyExistsException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Indicates that the configuration set does not exist.
type ConfigurationSetDoesNotExistException struct {
	Message *string

	ConfigurationSetName *string
}

func (e *ConfigurationSetDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConfigurationSetDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConfigurationSetDoesNotExistException) ErrorCode() string {
	return "ConfigurationSetDoesNotExistException"
}
func (e *ConfigurationSetDoesNotExistException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Indicates that email sending is disabled for the configuration set. You can
// enable or disable email sending for a configuration set using
// UpdateConfigurationSetSendingEnabled.
type ConfigurationSetSendingPausedException struct {
	Message *string

	ConfigurationSetName *string
}

func (e *ConfigurationSetSendingPausedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConfigurationSetSendingPausedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConfigurationSetSendingPausedException) ErrorCode() string {
	return "ConfigurationSetSendingPausedException"
}
func (e *ConfigurationSetSendingPausedException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Indicates that custom verification email template provided content is invalid.
type CustomVerificationEmailInvalidContentException struct {
	Message *string
}

func (e *CustomVerificationEmailInvalidContentException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CustomVerificationEmailInvalidContentException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CustomVerificationEmailInvalidContentException) ErrorCode() string {
	return "CustomVerificationEmailInvalidContentException"
}
func (e *CustomVerificationEmailInvalidContentException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Indicates that a custom verification email template with the name you specified
// already exists.
type CustomVerificationEmailTemplateAlreadyExistsException struct {
	Message *string

	CustomVerificationEmailTemplateName *string
}

func (e *CustomVerificationEmailTemplateAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CustomVerificationEmailTemplateAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CustomVerificationEmailTemplateAlreadyExistsException) ErrorCode() string {
	return "CustomVerificationEmailTemplateAlreadyExistsException"
}
func (e *CustomVerificationEmailTemplateAlreadyExistsException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Indicates that a custom verification email template with the name you specified
// does not exist.
type CustomVerificationEmailTemplateDoesNotExistException struct {
	Message *string

	CustomVerificationEmailTemplateName *string
}

func (e *CustomVerificationEmailTemplateDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CustomVerificationEmailTemplateDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CustomVerificationEmailTemplateDoesNotExistException) ErrorCode() string {
	return "CustomVerificationEmailTemplateDoesNotExistException"
}
func (e *CustomVerificationEmailTemplateDoesNotExistException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Indicates that the event destination could not be created because of a naming
// conflict.
type EventDestinationAlreadyExistsException struct {
	Message *string

	ConfigurationSetName *string
	EventDestinationName *string
}

func (e *EventDestinationAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EventDestinationAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EventDestinationAlreadyExistsException) ErrorCode() string {
	return "EventDestinationAlreadyExistsException"
}
func (e *EventDestinationAlreadyExistsException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Indicates that the event destination does not exist.
type EventDestinationDoesNotExistException struct {
	Message *string

	EventDestinationName *string
	ConfigurationSetName *string
}

func (e *EventDestinationDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EventDestinationDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EventDestinationDoesNotExistException) ErrorCode() string {
	return "EventDestinationDoesNotExistException"
}
func (e *EventDestinationDoesNotExistException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Indicates that the sender address specified for a custom verification email is
// not verified, and is therefore not eligible to send the custom verification
// email.
type FromEmailAddressNotVerifiedException struct {
	Message *string

	FromEmailAddress *string
}

func (e *FromEmailAddressNotVerifiedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FromEmailAddressNotVerifiedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FromEmailAddressNotVerifiedException) ErrorCode() string {
	return "FromEmailAddressNotVerifiedException"
}
func (e *FromEmailAddressNotVerifiedException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Indicates that the Amazon CloudWatch destination is invalid. See the error
// message for details.
type InvalidCloudWatchDestinationException struct {
	Message *string

	EventDestinationName *string
	ConfigurationSetName *string
}

func (e *InvalidCloudWatchDestinationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidCloudWatchDestinationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidCloudWatchDestinationException) ErrorCode() string {
	return "InvalidCloudWatchDestinationException"
}
func (e *InvalidCloudWatchDestinationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Indicates that the configuration set is invalid. See the error message for
// details.
type InvalidConfigurationSetException struct {
	Message *string
}

func (e *InvalidConfigurationSetException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidConfigurationSetException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidConfigurationSetException) ErrorCode() string {
	return "InvalidConfigurationSetException"
}
func (e *InvalidConfigurationSetException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that provided delivery option is invalid.
type InvalidDeliveryOptionsException struct {
	Message *string
}

func (e *InvalidDeliveryOptionsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDeliveryOptionsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDeliveryOptionsException) ErrorCode() string {
	return "InvalidDeliveryOptionsException"
}
func (e *InvalidDeliveryOptionsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the Amazon Kinesis Firehose destination is invalid. See the error
// message for details.
type InvalidFirehoseDestinationException struct {
	Message *string

	EventDestinationName *string
	ConfigurationSetName *string
}

func (e *InvalidFirehoseDestinationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidFirehoseDestinationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidFirehoseDestinationException) ErrorCode() string {
	return "InvalidFirehoseDestinationException"
}
func (e *InvalidFirehoseDestinationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Indicates that the provided AWS Lambda function is invalid, or that Amazon SES
// could not execute the provided function, possibly due to permissions issues. For
// information about giving permissions, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-permissions.html).
type InvalidLambdaFunctionException struct {
	Message *string

	FunctionArn *string
}

func (e *InvalidLambdaFunctionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidLambdaFunctionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidLambdaFunctionException) ErrorCode() string             { return "InvalidLambdaFunctionException" }
func (e *InvalidLambdaFunctionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the provided policy is invalid. Check the error stack for more
// information about what caused the error.
type InvalidPolicyException struct {
	Message *string
}

func (e *InvalidPolicyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidPolicyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidPolicyException) ErrorCode() string             { return "InvalidPolicyException" }
func (e *InvalidPolicyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that one or more of the replacement values you provided is invalid.
// This error may occur when the TemplateData object contains invalid JSON.
type InvalidRenderingParameterException struct {
	Message *string

	TemplateName *string
}

func (e *InvalidRenderingParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRenderingParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRenderingParameterException) ErrorCode() string {
	return "InvalidRenderingParameterException"
}
func (e *InvalidRenderingParameterException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Indicates that the provided Amazon S3 bucket or AWS KMS encryption key is
// invalid, or that Amazon SES could not publish to the bucket, possibly due to
// permissions issues. For information about giving permissions, see the Amazon SES
// Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-permissions.html).
type InvalidS3ConfigurationException struct {
	Message *string

	Bucket *string
}

func (e *InvalidS3ConfigurationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidS3ConfigurationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidS3ConfigurationException) ErrorCode() string {
	return "InvalidS3ConfigurationException"
}
func (e *InvalidS3ConfigurationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the Amazon Simple Notification Service (Amazon SNS) destination
// is invalid. See the error message for details.
type InvalidSNSDestinationException struct {
	Message *string

	ConfigurationSetName *string
	EventDestinationName *string
}

func (e *InvalidSNSDestinationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSNSDestinationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSNSDestinationException) ErrorCode() string             { return "InvalidSNSDestinationException" }
func (e *InvalidSNSDestinationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the provided Amazon SNS topic is invalid, or that Amazon SES
// could not publish to the topic, possibly due to permissions issues. For
// information about giving permissions, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-permissions.html).
type InvalidSnsTopicException struct {
	Message *string

	Topic *string
}

func (e *InvalidSnsTopicException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSnsTopicException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSnsTopicException) ErrorCode() string             { return "InvalidSnsTopicException" }
func (e *InvalidSnsTopicException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the template that you specified could not be rendered. This issue
// may occur when a template refers to a partial that does not exist.
type InvalidTemplateException struct {
	Message *string

	TemplateName *string
}

func (e *InvalidTemplateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTemplateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTemplateException) ErrorCode() string             { return "InvalidTemplateException" }
func (e *InvalidTemplateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the custom domain to be used for open and click tracking
// redirects is invalid. This error appears most often in the following
// situations:
//
//     * When the tracking domain you specified is not verified in
// Amazon SES.
//
//     * When the tracking domain you specified is not a valid domain
// or subdomain.
type InvalidTrackingOptionsException struct {
	Message *string
}

func (e *InvalidTrackingOptionsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTrackingOptionsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTrackingOptionsException) ErrorCode() string {
	return "InvalidTrackingOptionsException"
}
func (e *InvalidTrackingOptionsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that a resource could not be created because of service limits. For a
// list of Amazon SES limits, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/limits.html).
type LimitExceededException struct {
	Message *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the message could not be sent because Amazon SES could not read
// the MX record required to use the specified MAIL FROM domain. For information
// about editing the custom MAIL FROM domain settings for an identity, see the
// Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/mail-from-edit.html).
type MailFromDomainNotVerifiedException struct {
	Message *string
}

func (e *MailFromDomainNotVerifiedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MailFromDomainNotVerifiedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MailFromDomainNotVerifiedException) ErrorCode() string {
	return "MailFromDomainNotVerifiedException"
}
func (e *MailFromDomainNotVerifiedException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Indicates that the action failed, and the message could not be sent. Check the
// error stack for more information about what caused the error.
type MessageRejected struct {
	Message *string
}

func (e *MessageRejected) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MessageRejected) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MessageRejected) ErrorCode() string             { return "MessageRejected" }
func (e *MessageRejected) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that one or more of the replacement values for the specified template
// was not specified. Ensure that the TemplateData object contains references to
// all of the replacement tags in the specified template.
type MissingRenderingAttributeException struct {
	Message *string

	TemplateName *string
}

func (e *MissingRenderingAttributeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MissingRenderingAttributeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MissingRenderingAttributeException) ErrorCode() string {
	return "MissingRenderingAttributeException"
}
func (e *MissingRenderingAttributeException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Indicates that the account has not been granted production access.
type ProductionAccessNotGrantedException struct {
	Message *string
}

func (e *ProductionAccessNotGrantedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ProductionAccessNotGrantedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ProductionAccessNotGrantedException) ErrorCode() string {
	return "ProductionAccessNotGrantedException"
}
func (e *ProductionAccessNotGrantedException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Indicates that the provided receipt rule does not exist.
type RuleDoesNotExistException struct {
	Message *string

	Name *string
}

func (e *RuleDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RuleDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RuleDoesNotExistException) ErrorCode() string             { return "RuleDoesNotExistException" }
func (e *RuleDoesNotExistException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the provided receipt rule set does not exist.
type RuleSetDoesNotExistException struct {
	Message *string

	Name *string
}

func (e *RuleSetDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RuleSetDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RuleSetDoesNotExistException) ErrorCode() string             { return "RuleSetDoesNotExistException" }
func (e *RuleSetDoesNotExistException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the Template object you specified does not exist in your Amazon
// SES account.
type TemplateDoesNotExistException struct {
	Message *string

	TemplateName *string
}

func (e *TemplateDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TemplateDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TemplateDoesNotExistException) ErrorCode() string             { return "TemplateDoesNotExistException" }
func (e *TemplateDoesNotExistException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the configuration set you specified already contains a
// TrackingOptions object.
type TrackingOptionsAlreadyExistsException struct {
	Message *string

	ConfigurationSetName *string
}

func (e *TrackingOptionsAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TrackingOptionsAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TrackingOptionsAlreadyExistsException) ErrorCode() string {
	return "TrackingOptionsAlreadyExistsException"
}
func (e *TrackingOptionsAlreadyExistsException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Indicates that the TrackingOptions object you specified does not exist.
type TrackingOptionsDoesNotExistException struct {
	Message *string

	ConfigurationSetName *string
}

func (e *TrackingOptionsDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TrackingOptionsDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TrackingOptionsDoesNotExistException) ErrorCode() string {
	return "TrackingOptionsDoesNotExistException"
}
func (e *TrackingOptionsDoesNotExistException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
