// Code generated by smithy-go-codegen DO NOT EDIT.

package sts

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a set of temporary security credentials for users who have been
// authenticated in a mobile or web application with a web identity provider.
// Example providers include Amazon Cognito, Login with Amazon, Facebook, Google,
// or any OpenID Connect-compatible identity provider. For mobile applications, we
// recommend that you use Amazon Cognito. You can use Amazon Cognito with the AWS
// SDK for iOS Developer Guide (http://aws.amazon.com/sdkforios/) and the AWS SDK
// for Android Developer Guide (http://aws.amazon.com/sdkforandroid/) to uniquely
// identify a user. You can also supply the user with a consistent identity
// throughout the lifetime of an application. To learn more about Amazon Cognito,
// see Amazon Cognito Overview
// (https://docs.aws.amazon.com/mobile/sdkforandroid/developerguide/cognito-auth.html#d0e840)
// in AWS SDK for Android Developer Guide and Amazon Cognito Overview
// (https://docs.aws.amazon.com/mobile/sdkforios/developerguide/cognito-auth.html#d0e664)
// in the AWS SDK for iOS Developer Guide. Calling AssumeRoleWithWebIdentity does
// not require the use of AWS security credentials. Therefore, you can distribute
// an application (for example, on mobile devices) that requests temporary security
// credentials without including long-term AWS credentials in the application. You
// also don't need to deploy server-based proxy services that use long-term AWS
// credentials. Instead, the identity of the caller is validated by using a token
// from the web identity provider. For a comparison of AssumeRoleWithWebIdentity
// with the other API operations that produce temporary credentials, see Requesting
// Temporary Security Credentials
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html)
// and Comparing the AWS STS API operations
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#stsapi_comparison)
// in the IAM User Guide. The temporary security credentials returned by this API
// consist of an access key ID, a secret access key, and a security token.
// Applications can use these temporary security credentials to sign calls to AWS
// service API operations. Session Duration By default, the temporary security
// credentials created by AssumeRoleWithWebIdentity last for one hour. However, you
// can use the optional DurationSeconds parameter to specify the duration of your
// session. You can provide a value from 900 seconds (15 minutes) up to the maximum
// session duration setting for the role. This setting can have a value from 1 hour
// to 12 hours. To learn how to view the maximum value for your role, see View the
// Maximum Session Duration Setting for a Role
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html#id_roles_use_view-role-max-session)
// in the IAM User Guide. The maximum session duration limit applies when you use
// the AssumeRole* API operations or the assume-role* CLI commands. However the
// limit does not apply when you use those operations to create a console URL. For
// more information, see Using IAM Roles
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html) in the IAM
// User Guide. Permissions The temporary security credentials created by
// AssumeRoleWithWebIdentity can be used to make API calls to any AWS service with
// the following exception: you cannot call the STS GetFederationToken or
// GetSessionToken API operations. (Optional) You can pass inline or managed
// session policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
// to this operation. You can pass a single JSON policy document to use as an
// inline session policy. You can also specify up to 10 managed policies to use as
// managed session policies. The plain text that you use for both inline and
// managed session policies can't exceed 2,048 characters. Passing policies to this
// operation returns new temporary credentials. The resulting session's permissions
// are the intersection of the role's identity-based policy and the session
// policies. You can use the role's temporary credentials in subsequent AWS API
// calls to access resources in the account that owns the role. You cannot use
// session policies to grant more permissions than those allowed by the
// identity-based policy of the role that is being assumed. For more information,
// see Session Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
// in the IAM User Guide. Tags (Optional) You can configure your IdP to pass
// attributes into your web identity token as session tags. Each session tag
// consists of a key name and an associated value. For more information about
// session tags, see Passing Session Tags in STS
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html) in the
// IAM User Guide. You can pass up to 50 session tags. The plain text session tag
// keys can’t exceed 128 characters and the values can’t exceed 256 characters. For
// these and additional limits, see IAM and STS Character Limits
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-limits.html#reference_iam-limits-entity-length)
// in the IAM User Guide. An AWS conversion compresses the passed session policies
// and session tags into a packed binary format that has a separate limit. Your
// request can fail for this limit even if your plain text meets the other
// requirements. The PackedPolicySize response element indicates by percentage how
// close the policies and tags for your request are to the upper size limit. You
// can pass a session tag with the same key as a tag that is attached to the role.
// When you do, the session tag overrides the role tag with the same key. An
// administrator must grant you the permissions necessary to pass session tags. The
// administrator can also create granular permissions to allow you to pass only
// specific session tags. For more information, see Tutorial: Using Tags for
// Attribute-Based Access Control
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_attribute-based-access-control.html)
// in the IAM User Guide. You can set the session tags as transitive. Transitive
// tags persist during role chaining. For more information, see Chaining Roles with
// Session Tags
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html#id_session-tags_role-chaining)
// in the IAM User Guide. Identities Before your application can call
// AssumeRoleWithWebIdentity, you must have an identity token from a supported
// identity provider and create a role that the application can assume. The role
// that your application assumes must trust the identity provider that is
// associated with the identity token. In other words, the identity provider must
// be specified in the role's trust policy. Calling AssumeRoleWithWebIdentity can
// result in an entry in your AWS CloudTrail logs. The entry includes the Subject
// (http://openid.net/specs/openid-connect-core-1_0.html#Claims) of the provided
// Web Identity Token. We recommend that you avoid using any personally
// identifiable information (PII) in this field. For example, you could instead use
// a GUID or a pairwise identifier, as suggested in the OIDC specification
// (http://openid.net/specs/openid-connect-core-1_0.html#SubjectIDTypes). For more
// information about how to use web identity federation and the
// AssumeRoleWithWebIdentity API, see the following resources:
//
//     * Using Web
// Identity Federation API Operations for Mobile Apps
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_oidc_manual.html)
// and Federation Through a Web-based Identity Provider
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#api_assumerolewithwebidentity).
//
//
// * Web Identity Federation Playground
// (https://web-identity-federation-playground.s3.amazonaws.com/index.html). Walk
// through the process of authenticating through Login with Amazon, Facebook, or
// Google, getting temporary security credentials, and then using those credentials
// to make a request to AWS.
//
//     * AWS SDK for iOS Developer Guide
// (http://aws.amazon.com/sdkforios/) and AWS SDK for Android Developer Guide
// (http://aws.amazon.com/sdkforandroid/). These toolkits contain sample apps that
// show how to invoke the identity providers. The toolkits then show how to use the
// information from these providers to get and use temporary security
// credentials.
//
//     * Web Identity Federation with Mobile Applications
// (http://aws.amazon.com/articles/web-identity-federation-with-mobile-applications).
// This article discusses web identity federation and shows an example of how to
// use web identity federation to get access to content in Amazon S3.
func (c *Client) AssumeRoleWithWebIdentity(ctx context.Context, params *AssumeRoleWithWebIdentityInput, optFns ...func(*Options)) (*AssumeRoleWithWebIdentityOutput, error) {
	if params == nil {
		params = &AssumeRoleWithWebIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssumeRoleWithWebIdentity", params, optFns, addOperationAssumeRoleWithWebIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssumeRoleWithWebIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssumeRoleWithWebIdentityInput struct {

	// The Amazon Resource Name (ARN) of the role that the caller is assuming.
	//
	// This member is required.
	RoleArn *string

	// An identifier for the assumed role session. Typically, you pass the name or
	// identifier that is associated with the user who is using your application. That
	// way, the temporary security credentials that your application will use are
	// associated with that user. This session name is included as part of the ARN and
	// assumed role ID in the AssumedRoleUser response element. The regex used to
	// validate this parameter is a string of characters consisting of upper- and
	// lower-case alphanumeric characters with no spaces. You can also include
	// underscores or any of the following characters: =,.@-
	//
	// This member is required.
	RoleSessionName *string

	// The OAuth 2.0 access token or OpenID Connect ID token that is provided by the
	// identity provider. Your application must get this token by authenticating the
	// user who is using your application with a web identity provider before the
	// application makes an AssumeRoleWithWebIdentity call.
	//
	// This member is required.
	WebIdentityToken *string

	// The duration, in seconds, of the role session. The value can range from 900
	// seconds (15 minutes) up to the maximum session duration setting for the role.
	// This setting can have a value from 1 hour to 12 hours. If you specify a value
	// higher than this setting, the operation fails. For example, if you specify a
	// session duration of 12 hours, but your administrator set the maximum session
	// duration to 6 hours, your operation fails. To learn how to view the maximum
	// value for your role, see View the Maximum Session Duration Setting for a Role
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html#id_roles_use_view-role-max-session)
	// in the IAM User Guide. By default, the value is set to 3600 seconds. The
	// DurationSeconds parameter is separate from the duration of a console session
	// that you might request using the returned credentials. The request to the
	// federation endpoint for a console sign-in token takes a SessionDuration
	// parameter that specifies the maximum length of the console session. For more
	// information, see Creating a URL that Enables Federated Users to Access the AWS
	// Management Console
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_enable-console-custom-url.html)
	// in the IAM User Guide.
	DurationSeconds *int32

	// An IAM policy in JSON format that you want to use as an inline session policy.
	// This parameter is optional. Passing policies to this operation returns new
	// temporary credentials. The resulting session's permissions are the intersection
	// of the role's identity-based policy and the session policies. You can use the
	// role's temporary credentials in subsequent AWS API calls to access resources in
	// the account that owns the role. You cannot use session policies to grant more
	// permissions than those allowed by the identity-based policy of the role that is
	// being assumed. For more information, see Session Policies
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
	// in the IAM User Guide. The plain text that you use for both inline and managed
	// session policies can't exceed 2,048 characters. The JSON policy characters can
	// be any ASCII character from the space character to the end of the valid
	// character list (\u0020 through \u00FF). It can also include the tab (\u0009),
	// linefeed (\u000A), and carriage return (\u000D) characters. An AWS conversion
	// compresses the passed session policies and session tags into a packed binary
	// format that has a separate limit. Your request can fail for this limit even if
	// your plain text meets the other requirements. The PackedPolicySize response
	// element indicates by percentage how close the policies and tags for your request
	// are to the upper size limit.
	Policy *string

	// The Amazon Resource Names (ARNs) of the IAM managed policies that you want to
	// use as managed session policies. The policies must exist in the same account as
	// the role. This parameter is optional. You can provide up to 10 managed policy
	// ARNs. However, the plain text that you use for both inline and managed session
	// policies can't exceed 2,048 characters. For more information about ARNs, see
	// Amazon Resource Names (ARNs) and AWS Service Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference. An AWS conversion compresses the passed session
	// policies and session tags into a packed binary format that has a separate limit.
	// Your request can fail for this limit even if your plain text meets the other
	// requirements. The PackedPolicySize response element indicates by percentage how
	// close the policies and tags for your request are to the upper size limit.
	// Passing policies to this operation returns new temporary credentials. The
	// resulting session's permissions are the intersection of the role's
	// identity-based policy and the session policies. You can use the role's temporary
	// credentials in subsequent AWS API calls to access resources in the account that
	// owns the role. You cannot use session policies to grant more permissions than
	// those allowed by the identity-based policy of the role that is being assumed.
	// For more information, see Session Policies
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
	// in the IAM User Guide.
	PolicyArns []*types.PolicyDescriptorType

	// The fully qualified host component of the domain name of the identity provider.
	// Specify this value only for OAuth 2.0 access tokens. Currently www.amazon.com
	// and graph.facebook.com are the only supported identity providers for OAuth 2.0
	// access tokens. Do not include URL schemes and port numbers. Do not specify this
	// value for OpenID Connect ID tokens.
	ProviderId *string
}

// Contains the response to a successful AssumeRoleWithWebIdentity request,
// including temporary AWS credentials that can be used to make AWS requests.
type AssumeRoleWithWebIdentityOutput struct {

	// The Amazon Resource Name (ARN) and the assumed role ID, which are identifiers
	// that you can use to refer to the resulting temporary security credentials. For
	// example, you can reference these credentials as a principal in a resource-based
	// policy by using the ARN or assumed role ID. The ARN and ID include the
	// RoleSessionName that you specified when you called AssumeRole.
	AssumedRoleUser *types.AssumedRoleUser

	// The intended audience (also known as client ID) of the web identity token. This
	// is traditionally the client identifier issued to the application that requested
	// the web identity token.
	Audience *string

	// The temporary security credentials, which include an access key ID, a secret
	// access key, and a security token. The size of the security token that STS API
	// operations return is not fixed. We strongly recommend that you make no
	// assumptions about the maximum size.
	Credentials *types.Credentials

	// A percentage value that indicates the packed size of the session policies and
	// session tags combined passed in the request. The request fails if the packed
	// size is greater than 100 percent, which means the policies and tags exceeded the
	// allowed space.
	PackedPolicySize *int32

	// The issuing authority of the web identity token presented. For OpenID Connect ID
	// tokens, this contains the value of the iss field. For OAuth 2.0 access tokens,
	// this contains the value of the ProviderId parameter that was passed in the
	// AssumeRoleWithWebIdentity request.
	Provider *string

	// The unique user identifier that is returned by the identity provider. This
	// identifier is associated with the WebIdentityToken that was submitted with the
	// AssumeRoleWithWebIdentity call. The identifier is typically unique to the user
	// and the application that acquired the WebIdentityToken (pairwise identifier).
	// For OpenID Connect ID tokens, this field contains the value returned by the
	// identity provider as the token's sub (Subject) claim.
	SubjectFromWebIdentityToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssumeRoleWithWebIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAssumeRoleWithWebIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAssumeRoleWithWebIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpAssumeRoleWithWebIdentityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssumeRoleWithWebIdentity(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAssumeRoleWithWebIdentity(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sts",
		OperationName: "AssumeRoleWithWebIdentity",
	}
}
