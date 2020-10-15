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

// Schedules a run.
func (c *Client) ScheduleRun(ctx context.Context, params *ScheduleRunInput, optFns ...func(*Options)) (*ScheduleRunOutput, error) {
	if params == nil {
		params = &ScheduleRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ScheduleRun", params, optFns, addOperationScheduleRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ScheduleRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to the schedule run operation.
type ScheduleRunInput struct {

	// The ARN of the project for the run to be scheduled.
	//
	// This member is required.
	ProjectArn *string

	// Information about the test for the run to be scheduled.
	//
	// This member is required.
	Test *types.ScheduleRunTest

	// The ARN of an application package to run tests against, created with
	// CreateUpload. See ListUploads.
	AppArn *string

	// Information about the settings for the run to be scheduled.
	Configuration *types.ScheduleRunConfiguration

	// The ARN of the device pool for the run to be scheduled.
	DevicePoolArn *string

	// The filter criteria used to dynamically select a set of devices for a test run
	// and the maximum number of devices to be included in the run. Either
	// devicePoolArn or deviceSelectionConfiguration is required in a request.
	DeviceSelectionConfiguration *types.DeviceSelectionConfiguration

	// Specifies configuration information about a test run, such as the execution
	// timeout (in minutes).
	ExecutionConfiguration *types.ExecutionConfiguration

	// The name for the run to be scheduled.
	Name *string
}

// Represents the result of a schedule run request.
type ScheduleRunOutput struct {

	// Information about the scheduled run.
	Run *types.Run

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationScheduleRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpScheduleRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpScheduleRun{}, middleware.After)
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
	addOpScheduleRunValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opScheduleRun(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opScheduleRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ScheduleRun",
	}
}
