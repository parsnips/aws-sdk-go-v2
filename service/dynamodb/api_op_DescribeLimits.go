// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the current provisioned-capacity limits for your AWS account in a
// Region, both for the Region as a whole and for any one DynamoDB table that you
// create there. When you establish an AWS account, the account has initial limits
// on the maximum read capacity units and write capacity units that you can
// provision across all of your DynamoDB tables in a given Region. Also, there are
// per-table limits that apply when you create a table there. For more information,
// see Limits
// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html)
// page in the Amazon DynamoDB Developer Guide. Although you can increase these
// limits by filing a case at AWS Support Center
// (https://console.aws.amazon.com/support/home#/), obtaining the increase is not
// instantaneous. The DescribeLimits action lets you write code to compare the
// capacity you are currently using to those limits imposed by your account so that
// you have enough time to apply for an increase before you hit a limit. For
// example, you could use one of the AWS SDKs to do the following:
//
//     * Call
// DescribeLimits for a particular Region to obtain your current account limits on
// provisioned capacity there.
//
//     * Create a variable to hold the aggregate read
// capacity units provisioned for all your tables in that Region, and one to hold
// the aggregate write capacity units. Zero them both.
//
//     * Call ListTables to
// obtain a list of all your DynamoDB tables.
//
//     * For each table name listed by
// ListTables, do the following:
//
//         * Call DescribeTable with the table
// name.
//
//         * Use the data returned by DescribeTable to add the read capacity
// units and write capacity units provisioned for the table itself to your
// variables.
//
//         * If the table has one or more global secondary indexes
// (GSIs), loop over these GSIs and add their provisioned capacity values to your
// variables as well.
//
//     * Report the account limits for that Region returned by
// DescribeLimits, along with the total current provisioned capacity levels you
// have calculated.
//
// This will let you see whether you are getting close to your
// account-level limits. The per-table limits apply only when you are creating a
// new table. They restrict the sum of the provisioned capacity of the new table
// itself and all its global secondary indexes. For existing tables and their GSIs,
// DynamoDB doesn't let you increase provisioned capacity extremely rapidly. But
// the only upper limit that applies is that the aggregate provisioned capacity
// over all your tables and GSIs cannot exceed either of the per-account limits.
// DescribeLimits should only be called periodically. You can expect throttling
// errors if you call it more than once in a minute. The DescribeLimits Request
// element has no content.
func (c *Client) DescribeLimits(ctx context.Context, params *DescribeLimitsInput, optFns ...func(*Options)) (*DescribeLimitsOutput, error) {
	if params == nil {
		params = &DescribeLimitsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLimits", params, optFns, addOperationDescribeLimitsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLimitsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DescribeLimits operation. Has no content.
type DescribeLimitsInput struct {
}

// Represents the output of a DescribeLimits operation.
type DescribeLimitsOutput struct {

	// The maximum total read capacity units that your account allows you to provision
	// across all of your tables in this Region.
	AccountMaxReadCapacityUnits *int64

	// The maximum total write capacity units that your account allows you to provision
	// across all of your tables in this Region.
	AccountMaxWriteCapacityUnits *int64

	// The maximum read capacity units that your account allows you to provision for a
	// new table that you are creating in this Region, including the read capacity
	// units provisioned for its global secondary indexes (GSIs).
	TableMaxReadCapacityUnits *int64

	// The maximum write capacity units that your account allows you to provision for a
	// new table that you are creating in this Region, including the write capacity
	// units provisioned for its global secondary indexes (GSIs).
	TableMaxWriteCapacityUnits *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeLimitsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeLimits{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeLimits{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLimits(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addValidateResponseChecksum(stack, options)
	addAcceptEncodingGzip(stack, options)
	return nil
}

func newServiceMetadataMiddleware_opDescribeLimits(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "DescribeLimits",
	}
}
