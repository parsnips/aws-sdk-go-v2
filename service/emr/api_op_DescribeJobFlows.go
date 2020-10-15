// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This API is deprecated and will eventually be removed. We recommend you use
// ListClusters, DescribeCluster, ListSteps, ListInstanceGroups and
// ListBootstrapActions instead. DescribeJobFlows returns a list of job flows that
// match all of the supplied parameters. The parameters can include a list of job
// flow IDs, job flow states, and restrictions on job flow creation date and time.
// Regardless of supplied parameters, only job flows created within the last two
// months are returned. If no parameters are supplied, then job flows matching
// either of the following criteria are returned:
//
//     * Job flows created and
// completed in the last two weeks
//
//     * Job flows created within the last two
// months that are in one of the following states: RUNNING, WAITING, SHUTTING_DOWN,
// STARTING
//
// Amazon EMR can return a maximum of 512 job flow descriptions.
func (c *Client) DescribeJobFlows(ctx context.Context, params *DescribeJobFlowsInput, optFns ...func(*Options)) (*DescribeJobFlowsOutput, error) {
	if params == nil {
		params = &DescribeJobFlowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeJobFlows", params, optFns, addOperationDescribeJobFlowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeJobFlowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DescribeJobFlows operation.
type DescribeJobFlowsInput struct {

	// Return only job flows created after this date and time.
	CreatedAfter *time.Time

	// Return only job flows created before this date and time.
	CreatedBefore *time.Time

	// Return only job flows whose job flow ID is contained in this list.
	JobFlowIds []*string

	// Return only job flows whose state is contained in this list.
	JobFlowStates []types.JobFlowExecutionState
}

// The output for the DescribeJobFlows operation.
type DescribeJobFlowsOutput struct {

	// A list of job flows matching the parameters supplied.
	JobFlows []*types.JobFlowDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeJobFlowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeJobFlows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeJobFlows{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeJobFlows(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeJobFlows(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "DescribeJobFlows",
	}
}
