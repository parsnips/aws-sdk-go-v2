// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
)

// The request was rejected because multiple requests to change this object were
// submitted simultaneously. Wait a few minutes and submit your request again.
type ConcurrentModificationException struct {
	Message *string
}

func (e *ConcurrentModificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModificationException) ErrorCode() string {
	return "ConcurrentModificationException"
}
func (e *ConcurrentModificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the most recent credential report has expired.
// To generate a new credential report, use GenerateCredentialReport. For more
// information about credential report expiration, see Getting Credential Reports
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/credential-reports.html) in
// the IAM User Guide.
type CredentialReportExpiredException struct {
	Message *string
}

func (e *CredentialReportExpiredException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CredentialReportExpiredException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CredentialReportExpiredException) ErrorCode() string {
	return "CredentialReportExpiredException"
}
func (e *CredentialReportExpiredException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the credential report does not exist. To
// generate a credential report, use GenerateCredentialReport.
type CredentialReportNotPresentException struct {
	Message *string
}

func (e *CredentialReportNotPresentException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CredentialReportNotPresentException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CredentialReportNotPresentException) ErrorCode() string {
	return "CredentialReportNotPresentException"
}
func (e *CredentialReportNotPresentException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The request was rejected because the credential report is still being generated.
type CredentialReportNotReadyException struct {
	Message *string
}

func (e *CredentialReportNotReadyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CredentialReportNotReadyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CredentialReportNotReadyException) ErrorCode() string {
	return "CredentialReportNotReadyException"
}
func (e *CredentialReportNotReadyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because it attempted to delete a resource that has
// attached subordinate entities. The error message describes these entities.
type DeleteConflictException struct {
	Message *string
}

func (e *DeleteConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DeleteConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DeleteConflictException) ErrorCode() string             { return "DeleteConflictException" }
func (e *DeleteConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the same certificate is associated with an IAM
// user in the account.
type DuplicateCertificateException struct {
	Message *string
}

func (e *DuplicateCertificateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateCertificateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateCertificateException) ErrorCode() string             { return "DuplicateCertificateException" }
func (e *DuplicateCertificateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the SSH public key is already associated with
// the specified IAM user.
type DuplicateSSHPublicKeyException struct {
	Message *string
}

func (e *DuplicateSSHPublicKeyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateSSHPublicKeyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateSSHPublicKeyException) ErrorCode() string             { return "DuplicateSSHPublicKeyException" }
func (e *DuplicateSSHPublicKeyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because it attempted to create a resource that already
// exists.
type EntityAlreadyExistsException struct {
	Message *string
}

func (e *EntityAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EntityAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EntityAlreadyExistsException) ErrorCode() string             { return "EntityAlreadyExistsException" }
func (e *EntityAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because it referenced an entity that is temporarily
// unmodifiable, such as a user name that was deleted and then recreated. The error
// indicates that the request is likely to succeed if you try again after waiting
// several minutes. The error message describes the entity.
type EntityTemporarilyUnmodifiableException struct {
	Message *string
}

func (e *EntityTemporarilyUnmodifiableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EntityTemporarilyUnmodifiableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EntityTemporarilyUnmodifiableException) ErrorCode() string {
	return "EntityTemporarilyUnmodifiableException"
}
func (e *EntityTemporarilyUnmodifiableException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The request was rejected because the authentication code was not recognized. The
// error message describes the specific error.
type InvalidAuthenticationCodeException struct {
	Message *string
}

func (e *InvalidAuthenticationCodeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidAuthenticationCodeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidAuthenticationCodeException) ErrorCode() string {
	return "InvalidAuthenticationCodeException"
}
func (e *InvalidAuthenticationCodeException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The request was rejected because the certificate is invalid.
type InvalidCertificateException struct {
	Message *string
}

func (e *InvalidCertificateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidCertificateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidCertificateException) ErrorCode() string             { return "InvalidCertificateException" }
func (e *InvalidCertificateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because an invalid or out-of-range value was supplied
// for an input parameter.
type InvalidInputException struct {
	Message *string
}

func (e *InvalidInputException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidInputException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidInputException) ErrorCode() string             { return "InvalidInputException" }
func (e *InvalidInputException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the public key is malformed or otherwise
// invalid.
type InvalidPublicKeyException struct {
	Message *string
}

func (e *InvalidPublicKeyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidPublicKeyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidPublicKeyException) ErrorCode() string             { return "InvalidPublicKeyException" }
func (e *InvalidPublicKeyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the type of user for the transaction was
// incorrect.
type InvalidUserTypeException struct {
	Message *string
}

func (e *InvalidUserTypeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidUserTypeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidUserTypeException) ErrorCode() string             { return "InvalidUserTypeException" }
func (e *InvalidUserTypeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the public key certificate and the private key
// do not match.
type KeyPairMismatchException struct {
	Message *string
}

func (e *KeyPairMismatchException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KeyPairMismatchException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KeyPairMismatchException) ErrorCode() string             { return "KeyPairMismatchException" }
func (e *KeyPairMismatchException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because it attempted to create resources beyond the
// current AWS account limitations. The error message describes the limit exceeded.
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

// The request was rejected because the certificate was malformed or expired. The
// error message describes the specific error.
type MalformedCertificateException struct {
	Message *string
}

func (e *MalformedCertificateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MalformedCertificateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MalformedCertificateException) ErrorCode() string             { return "MalformedCertificateException" }
func (e *MalformedCertificateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the policy document was malformed. The error
// message describes the specific error.
type MalformedPolicyDocumentException struct {
	Message *string
}

func (e *MalformedPolicyDocumentException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MalformedPolicyDocumentException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MalformedPolicyDocumentException) ErrorCode() string {
	return "MalformedPolicyDocumentException"
}
func (e *MalformedPolicyDocumentException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because it referenced a resource entity that does not
// exist. The error message describes the resource.
type NoSuchEntityException struct {
	Message *string
}

func (e *NoSuchEntityException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoSuchEntityException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoSuchEntityException) ErrorCode() string             { return "NoSuchEntityException" }
func (e *NoSuchEntityException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the provided password did not meet the
// requirements imposed by the account password policy.
type PasswordPolicyViolationException struct {
	Message *string
}

func (e *PasswordPolicyViolationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PasswordPolicyViolationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PasswordPolicyViolationException) ErrorCode() string {
	return "PasswordPolicyViolationException"
}
func (e *PasswordPolicyViolationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request failed because a provided policy could not be successfully
// evaluated. An additional detailed message indicates the source of the failure.
type PolicyEvaluationException struct {
	Message *string
}

func (e *PolicyEvaluationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PolicyEvaluationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PolicyEvaluationException) ErrorCode() string             { return "PolicyEvaluationException" }
func (e *PolicyEvaluationException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The request failed because AWS service role policies can only be attached to the
// service-linked role for that service.
type PolicyNotAttachableException struct {
	Message *string
}

func (e *PolicyNotAttachableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PolicyNotAttachableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PolicyNotAttachableException) ErrorCode() string             { return "PolicyNotAttachableException" }
func (e *PolicyNotAttachableException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request failed because the maximum number of concurrent requests for this
// account are already running.
type ReportGenerationLimitExceededException struct {
	Message *string
}

func (e *ReportGenerationLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReportGenerationLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReportGenerationLimitExceededException) ErrorCode() string {
	return "ReportGenerationLimitExceededException"
}
func (e *ReportGenerationLimitExceededException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The request processing has failed because of an unknown error, exception or
// failure.
type ServiceFailureException struct {
	Message *string
}

func (e *ServiceFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceFailureException) ErrorCode() string             { return "ServiceFailureException" }
func (e *ServiceFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The specified service does not support service-specific credentials.
type ServiceNotSupportedException struct {
	Message *string
}

func (e *ServiceNotSupportedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceNotSupportedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceNotSupportedException) ErrorCode() string             { return "ServiceNotSupportedException" }
func (e *ServiceNotSupportedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because only the service that depends on the
// service-linked role can modify or delete the role on your behalf. The error
// message includes the name of the service that depends on this service-linked
// role. You must request the change through that service.
type UnmodifiableEntityException struct {
	Message *string
}

func (e *UnmodifiableEntityException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnmodifiableEntityException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnmodifiableEntityException) ErrorCode() string             { return "UnmodifiableEntityException" }
func (e *UnmodifiableEntityException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the public key encoding format is unsupported
// or unrecognized.
type UnrecognizedPublicKeyEncodingException struct {
	Message *string
}

func (e *UnrecognizedPublicKeyEncodingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnrecognizedPublicKeyEncodingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnrecognizedPublicKeyEncodingException) ErrorCode() string {
	return "UnrecognizedPublicKeyEncodingException"
}
func (e *UnrecognizedPublicKeyEncodingException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
