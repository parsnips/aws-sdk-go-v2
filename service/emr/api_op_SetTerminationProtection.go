// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// SetTerminationProtection locks a cluster (job flow) so the EC2 instances in the
// cluster cannot be terminated by user intervention, an API call, or in the event
// of a job-flow error. The cluster still terminates upon successful completion of
// the job flow. Calling SetTerminationProtection on a cluster is similar to
// calling the Amazon EC2 DisableAPITermination API on all EC2 instances in a
// cluster. SetTerminationProtection is used to prevent accidental termination of a
// cluster and to ensure that in the event of an error, the instances persist so
// that you can recover any data stored in their ephemeral instance storage. To
// terminate a cluster that has been locked by setting SetTerminationProtection to
// true, you must first unlock the job flow by a subsequent call to
// SetTerminationProtection in which you set the value to false.
//
// For more
// information, seeManaging Cluster Termination
// (https://docs.aws.amazon.com/emr/latest/ManagementGuide/UsingEMR_TerminationProtection.html)
// in the Amazon EMR Management Guide.
func (c *Client) SetTerminationProtection(ctx context.Context, params *SetTerminationProtectionInput, optFns ...func(*Options)) (*SetTerminationProtectionOutput, error) {
	if params == nil {
		params = &SetTerminationProtectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetTerminationProtection", params, optFns, addOperationSetTerminationProtectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetTerminationProtectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input argument to the TerminationProtection operation.
type SetTerminationProtectionInput struct {

	// A list of strings that uniquely identify the clusters to protect. This
	// identifier is returned by RunJobFlow and can also be obtained from
	// DescribeJobFlows .
	//
	// This member is required.
	JobFlowIds []*string

	// A Boolean that indicates whether to protect the cluster and prevent the Amazon
	// EC2 instances in the cluster from shutting down due to API calls, user
	// intervention, or job-flow error.
	//
	// This member is required.
	TerminationProtected *bool
}

type SetTerminationProtectionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetTerminationProtectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetTerminationProtection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetTerminationProtection{}, middleware.After)
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
	addOpSetTerminationProtectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetTerminationProtection(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSetTerminationProtection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "SetTerminationProtection",
	}
}
