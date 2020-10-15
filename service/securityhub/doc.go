// Code generated by smithy-go-codegen DO NOT EDIT.

// Package securityhub provides the API client, operations, and parameter types for
// AWS SecurityHub.
//
// Security Hub provides you with a comprehensive view of the security state of
// your AWS environment and resources. It also provides you with the readiness
// status of your environment based on controls from supported security standards.
// Security Hub collects security data from AWS accounts, services, and integrated
// third-party products and helps you analyze security trends in your environment
// to identify the highest priority security issues. For more information about
// Security Hub, see the AWS Security Hub User Guide
// (https://docs.aws.amazon.com/securityhub/latest/userguide/what-is-securityhub.html)
// . When you use operations in the Security Hub API, the requests are executed
// only in the AWS Region that is currently active or in the specific AWS Region
// that you specify in your request. Any configuration or settings change that
// results from the operation is applied only to that Region. To make the same
// change in other Regions, execute the same command for each Region to apply the
// change to. For example, if your Region is set to us-west-2, when you use
// CreateMembers to add a member account to Security Hub, the association of the
// member account with the master account is created only in the us-west-2 Region.
// Security Hub must be enabled for the member account in the same Region that the
// invitation was sent from. The following throttling limits apply to using
// Security Hub API operations.
//
//     * GetFindings - RateLimit of 3 requests per
// second. BurstLimit of 6 requests per second.
//
//     * UpdateFindings - RateLimit
// of 1 request per second. BurstLimit of 5 requests per second.
//
//     * All other
// operations - RateLimit of 10 requests per second. BurstLimit of 30 requests per
// second.
package securityhub
