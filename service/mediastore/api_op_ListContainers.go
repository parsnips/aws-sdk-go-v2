// Code generated by smithy-go-codegen DO NOT EDIT.

package mediastore

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediastore/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the properties of all containers in AWS Elemental MediaStore. You can
// query to receive all the containers in one response. Or you can include the
// MaxResults parameter to receive a limited number of containers in each response.
// In this case, the response includes a token. To get the next set of containers,
// send the command again, this time with the NextToken parameter (with the
// returned token as its value). The next set of responses appears, with a token if
// there are still more containers to receive. See also DescribeContainer, which
// gets the properties of one container.
func (c *Client) ListContainers(ctx context.Context, params *ListContainersInput, optFns ...func(*Options)) (*ListContainersOutput, error) {
	if params == nil {
		params = &ListContainersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListContainers", params, optFns, addOperationListContainersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListContainersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListContainersInput struct {

	// Enter the maximum number of containers in the response. Use from 1 to 255
	// characters.
	MaxResults *int32

	// Only if you used MaxResults in the first command, enter the token (which was
	// included in the previous response) to obtain the next set of containers. This
	// token is included in a response only if there actually are more containers to
	// list.
	NextToken *string
}

type ListContainersOutput struct {

	// The names of the containers.
	//
	// This member is required.
	Containers []*types.Container

	// NextToken is the token to use in the next call to ListContainers. This token is
	// returned only if you included the MaxResults tag in the original command, and
	// only if there are still containers to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListContainersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListContainers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListContainers{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListContainers(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListContainers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediastore",
		OperationName: "ListContainers",
	}
}
