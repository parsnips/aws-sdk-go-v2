// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Schedules a query of a log group using CloudWatch Logs Insights. You specify the
// log group and time range to query, and the query string to use. For more
// information, see CloudWatch Logs Insights Query Syntax
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
// Queries time out after 15 minutes of execution. If your queries are timing out,
// reduce the time range being searched, or partition your query into a number of
// queries.
func (c *Client) StartQuery(ctx context.Context, params *StartQueryInput, optFns ...func(*Options)) (*StartQueryOutput, error) {
	if params == nil {
		params = &StartQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartQuery", params, optFns, addOperationStartQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartQueryInput struct {

	// The end of the time range to query. The range is inclusive, so the specified end
	// time is included in the query. Specified as epoch time, the number of seconds
	// since January 1, 1970, 00:00:00 UTC.
	//
	// This member is required.
	EndTime *int64

	// The query string to use. For more information, see CloudWatch Logs Insights
	// Query Syntax
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
	//
	// This member is required.
	QueryString *string

	// The beginning of the time range to query. The range is inclusive, so the
	// specified start time is included in the query. Specified as epoch time, the
	// number of seconds since January 1, 1970, 00:00:00 UTC.
	//
	// This member is required.
	StartTime *int64

	// The maximum number of log events to return in the query. If the query string
	// uses the fields command, only the specified fields and their values are
	// returned. The default is 1000.
	Limit *int32

	// The log group on which to perform the query. A StartQuery operation must include
	// a logGroupNames or a logGroupName parameter, but not both.
	LogGroupName *string

	// The list of log groups to be queried. You can include up to 20 log groups. A
	// StartQuery operation must include a logGroupNames or a logGroupName parameter,
	// but not both.
	LogGroupNames []*string
}

type StartQueryOutput struct {

	// The unique ID of the query.
	QueryId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartQuery{}, middleware.After)
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
	addOpStartQueryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartQuery(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartQuery(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "StartQuery",
	}
}
