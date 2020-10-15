// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the JSON text of the resource-based policy document attached to the
// specified secret. The JSON request string input and response output displays
// formatted code with white space and line breaks for better readability. Submit
// your input as a single line JSON string. Minimum permissions To run this
// command, you must have the following permissions:
//
//     *
// secretsmanager:GetResourcePolicy
//
// Related operations
//
//     * To attach a resource
// policy to a secret, use PutResourcePolicy.
//
//     * To delete the resource-based
// policy attached to a secret, use DeleteResourcePolicy.
//
//     * To list all of the
// currently available secrets, use ListSecrets.
func (c *Client) GetResourcePolicy(ctx context.Context, params *GetResourcePolicyInput, optFns ...func(*Options)) (*GetResourcePolicyOutput, error) {
	if params == nil {
		params = &GetResourcePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResourcePolicy", params, optFns, addOperationGetResourcePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourcePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourcePolicyInput struct {

	// Specifies the secret that you want to retrieve the attached resource-based
	// policy for. You can specify either the Amazon Resource Name (ARN) or the
	// friendly name of the secret. If you specify an ARN, we generally recommend that
	// you specify a complete ARN. You can specify a partial ARN too—for example, if
	// you don’t include the final hyphen and six random characters that Secrets
	// Manager adds at the end of the ARN when you created the secret. A partial ARN
	// match can work as long as it uniquely matches only one secret. However, if your
	// secret has a name that ends in a hyphen followed by six characters (before
	// Secrets Manager adds the hyphen and six characters to the ARN) and you try to
	// use that as a partial ARN, then those characters cause Secrets Manager to assume
	// that you’re specifying a complete ARN. This confusion can cause unexpected
	// results. To avoid this situation, we recommend that you don’t create secret
	// names ending with a hyphen followed by six characters. If you specify an
	// incomplete ARN without the random suffix, and instead provide the 'friendly
	// name', you must not include the random suffix. If you do include the random
	// suffix added by Secrets Manager, you receive either a ResourceNotFoundException
	// or an AccessDeniedException error, depending on your permissions.
	//
	// This member is required.
	SecretId *string
}

type GetResourcePolicyOutput struct {

	// The ARN of the secret that the resource-based policy was retrieved for.
	ARN *string

	// The friendly name of the secret that the resource-based policy was retrieved
	// for.
	Name *string

	// A JSON-formatted string that describes the permissions that are associated with
	// the attached secret. These permissions are combined with any permissions that
	// are associated with the user or role that attempts to access this secret. The
	// combined permissions specify who can access the secret and what actions they can
	// perform. For more information, see Authentication and Access Control for AWS
	// Secrets Manager
	// (http://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html)
	// in the AWS Secrets Manager User Guide.
	ResourcePolicy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetResourcePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetResourcePolicy{}, middleware.After)
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
	addOpGetResourcePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetResourcePolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetResourcePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "secretsmanager",
		OperationName: "GetResourcePolicy",
	}
}
