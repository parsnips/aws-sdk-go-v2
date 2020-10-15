// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns policies attached to an object in pagination fashion.
func (c *Client) ListObjectPolicies(ctx context.Context, params *ListObjectPoliciesInput, optFns ...func(*Options)) (*ListObjectPoliciesOutput, error) {
	if params == nil {
		params = &ListObjectPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListObjectPolicies", params, optFns, addOperationListObjectPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListObjectPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListObjectPoliciesInput struct {

	// The Amazon Resource Name (ARN) that is associated with the Directory where
	// objects reside. For more information, see arns.
	//
	// This member is required.
	DirectoryArn *string

	// Reference that identifies the object for which policies will be listed.
	//
	// This member is required.
	ObjectReference *types.ObjectReference

	// Represents the manner and timing in which the successful write or update of an
	// object is reflected in a subsequent read operation of that same object.
	ConsistencyLevel types.ConsistencyLevel

	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	MaxResults *int32

	// The pagination token.
	NextToken *string
}

type ListObjectPoliciesOutput struct {

	// A list of policy ObjectIdentifiers, that are attached to the object.
	AttachedPolicyIds []*string

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListObjectPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListObjectPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListObjectPolicies{}, middleware.After)
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
	addOpListObjectPoliciesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListObjectPolicies(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListObjectPolicies(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListObjectPolicies",
	}
}
