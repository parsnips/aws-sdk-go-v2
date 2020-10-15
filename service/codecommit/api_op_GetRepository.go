// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about a repository. The description field for a repository
// accepts all HTML characters and all valid Unicode characters. Applications that
// do not HTML-encode the description and display it in a webpage can expose users
// to potentially malicious code. Make sure that you HTML-encode the description
// field in any application that uses this API to display the repository
// description on a webpage.
func (c *Client) GetRepository(ctx context.Context, params *GetRepositoryInput, optFns ...func(*Options)) (*GetRepositoryOutput, error) {
	if params == nil {
		params = &GetRepositoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRepository", params, optFns, addOperationGetRepositoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a get repository operation.
type GetRepositoryInput struct {

	// The name of the repository to get information about.
	//
	// This member is required.
	RepositoryName *string
}

// Represents the output of a get repository operation.
type GetRepositoryOutput struct {

	// Information about the repository.
	RepositoryMetadata *types.RepositoryMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRepositoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRepository{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRepository{}, middleware.After)
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
	addOpGetRepositoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRepository(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetRepository(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "GetRepository",
	}
}
