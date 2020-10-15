// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list of artifacts created during the session.
func (c *Client) ListTestGridSessionArtifacts(ctx context.Context, params *ListTestGridSessionArtifactsInput, optFns ...func(*Options)) (*ListTestGridSessionArtifactsOutput, error) {
	if params == nil {
		params = &ListTestGridSessionArtifactsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTestGridSessionArtifacts", params, optFns, addOperationListTestGridSessionArtifactsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTestGridSessionArtifactsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTestGridSessionArtifactsInput struct {

	// The ARN of a TestGridSession.
	//
	// This member is required.
	SessionArn *string

	// The maximum number of results to be returned by a request.
	MaxResult *int32

	// Pagination token.
	NextToken *string

	// Limit results to a specified type of artifact.
	Type types.TestGridSessionArtifactCategory
}

type ListTestGridSessionArtifactsOutput struct {

	// A list of test grid session artifacts for a TestGridSession.
	Artifacts []*types.TestGridSessionArtifact

	// Pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTestGridSessionArtifactsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTestGridSessionArtifacts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTestGridSessionArtifacts{}, middleware.After)
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
	addOpListTestGridSessionArtifactsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTestGridSessionArtifacts(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListTestGridSessionArtifacts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListTestGridSessionArtifacts",
	}
}
