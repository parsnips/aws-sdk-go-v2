// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all of the organizational units (OUs) or accounts that are contained in
// the specified parent OU or root. This operation, along with ListParents enables
// you to traverse the tree structure that makes up this root. Always check the
// NextToken response parameter for a null value when calling a List* operation.
// These operations can occasionally return an empty set of results even when there
// are more results available. The NextToken response parameter value is null only
// when there are no more results to display. This operation can be called only
// from the organization's master account or by a member account that is a
// delegated administrator for an AWS service.
func (c *Client) ListChildren(ctx context.Context, params *ListChildrenInput, optFns ...func(*Options)) (*ListChildrenOutput, error) {
	if params == nil {
		params = &ListChildrenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListChildren", params, optFns, addOperationListChildrenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListChildrenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListChildrenInput struct {

	// Filters the output to include only the specified child type.
	//
	// This member is required.
	ChildType types.ChildType

	// The unique identifier (ID) for the parent root or OU whose children you want to
	// list. The regex pattern (http://wikipedia.org/wiki/regex) for a parent ID string
	// requires one of the following:
	//
	//     * Root - A string that begins with "r-"
	// followed by from 4 to 32 lowercase letters or digits.
	//
	//     * Organizational unit
	// (OU) - A string that begins with "ou-" followed by from 4 to 32 lowercase
	// letters or digits (the ID of the root that the OU is in). This string is
	// followed by a second "-" dash and from 8 to 32 additional lowercase letters or
	// digits.
	//
	// This member is required.
	ParentId *string

	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Organizations might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string
}

type ListChildrenOutput struct {

	// The list of children of the specified parent container.
	Children []*types.Child

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListChildrenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListChildren{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListChildren{}, middleware.After)
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
	addOpListChildrenValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListChildren(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListChildren(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "ListChildren",
	}
}
