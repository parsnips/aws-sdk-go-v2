// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes a dataset group created using the CreateDatasetGroup operation. In
// addition to listing the parameters provided in the CreateDatasetGroup request,
// this operation includes the following properties:
//
//     * DatasetArns - The
// datasets belonging to the group.
//
//     * CreationTime
//
//     *
// LastModificationTime
//
//     * Status
func (c *Client) DescribeDatasetGroup(ctx context.Context, params *DescribeDatasetGroupInput, optFns ...func(*Options)) (*DescribeDatasetGroupOutput, error) {
	if params == nil {
		params = &DescribeDatasetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDatasetGroup", params, optFns, addOperationDescribeDatasetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDatasetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDatasetGroupInput struct {

	// The Amazon Resource Name (ARN) of the dataset group.
	//
	// This member is required.
	DatasetGroupArn *string
}

type DescribeDatasetGroupOutput struct {

	// When the dataset group was created.
	CreationTime *time.Time

	// An array of Amazon Resource Names (ARNs) of the datasets contained in the
	// dataset group.
	DatasetArns []*string

	// The ARN of the dataset group.
	DatasetGroupArn *string

	// The name of the dataset group.
	DatasetGroupName *string

	// The domain associated with the dataset group.
	Domain types.Domain

	// When the dataset group was created or last updated from a call to the
	// UpdateDatasetGroup operation. While the dataset group is being updated,
	// LastModificationTime is the current time of the DescribeDatasetGroup call.
	LastModificationTime *time.Time

	// The status of the dataset group. States include:
	//
	//     * ACTIVE
	//
	//     *
	// CREATE_PENDING, CREATE_IN_PROGRESS, CREATE_FAILED
	//
	//     * DELETE_PENDING,
	// DELETE_IN_PROGRESS, DELETE_FAILED
	//
	//     * UPDATE_PENDING, UPDATE_IN_PROGRESS,
	// UPDATE_FAILED
	//
	// The UPDATE states apply when you call the UpdateDatasetGroup
	// operation. The Status of the dataset group must be ACTIVE before you can use the
	// dataset group to create a predictor.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDatasetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDatasetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDatasetGroup{}, middleware.After)
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
	addOpDescribeDatasetGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDatasetGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeDatasetGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DescribeDatasetGroup",
	}
}
