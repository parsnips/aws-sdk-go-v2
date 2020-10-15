// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attaches tags to an existing Elasticsearch domain. Tags are a set of
// case-sensitive key value pairs. An Elasticsearch domain may have up to 10 tags.
// See  Tagging Amazon Elasticsearch Service Domains for more information.
// (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-managedomains.html#es-managedomains-awsresorcetagging)
func (c *Client) AddTags(ctx context.Context, params *AddTagsInput, optFns ...func(*Options)) (*AddTagsOutput, error) {
	if params == nil {
		params = &AddTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddTags", params, optFns, addOperationAddTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the AddTags operation. Specify the tags that you
// want to attach to the Elasticsearch domain.
type AddTagsInput struct {

	// Specify the ARN for which you want to add the tags.
	//
	// This member is required.
	ARN *string

	// List of Tag that need to be added for the Elasticsearch domain.
	//
	// This member is required.
	TagList []*types.Tag
}

type AddTagsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAddTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAddTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAddTags{}, middleware.After)
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
	addOpAddTagsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddTags(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAddTags(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "AddTags",
	}
}
