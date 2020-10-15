// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the list of severity levels that you can assign to an AWS Support case.
// The severity level for a case is also a field in the CaseDetails data type that
// you include for a CreateCase request.
//
//     * You must have a Business or
// Enterprise support plan to use the AWS Support API.
//
//     * If you call the AWS
// Support API from an account that does not have a Business or Enterprise support
// plan, the SubscriptionRequiredException error message appears. For information
// about changing your support plan, see AWS Support
// (http://aws.amazon.com/premiumsupport/).
func (c *Client) DescribeSeverityLevels(ctx context.Context, params *DescribeSeverityLevelsInput, optFns ...func(*Options)) (*DescribeSeverityLevelsOutput, error) {
	if params == nil {
		params = &DescribeSeverityLevelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSeverityLevels", params, optFns, addOperationDescribeSeverityLevelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSeverityLevelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSeverityLevelsInput struct {

	// The ISO 639-1 code for the language in which AWS provides support. AWS Support
	// currently supports English ("en") and Japanese ("ja"). Language parameters must
	// be passed explicitly for operations that take them.
	Language *string
}

// The list of severity levels returned by the DescribeSeverityLevels operation.
type DescribeSeverityLevelsOutput struct {

	// The available severity levels for the support case. Available severity levels
	// are defined by your service level agreement with AWS.
	SeverityLevels []*types.SeverityLevel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSeverityLevelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSeverityLevels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSeverityLevels{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSeverityLevels(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeSeverityLevels(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "DescribeSeverityLevels",
	}
}
