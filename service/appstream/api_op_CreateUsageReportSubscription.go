// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a usage report subscription. Usage reports are generated daily.
func (c *Client) CreateUsageReportSubscription(ctx context.Context, params *CreateUsageReportSubscriptionInput, optFns ...func(*Options)) (*CreateUsageReportSubscriptionOutput, error) {
	if params == nil {
		params = &CreateUsageReportSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUsageReportSubscription", params, optFns, addOperationCreateUsageReportSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUsageReportSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUsageReportSubscriptionInput struct {
}

type CreateUsageReportSubscriptionOutput struct {

	// The Amazon S3 bucket where generated reports are stored. If you enabled
	// on-instance session scripts and Amazon S3 logging for your session script
	// configuration, AppStream 2.0 created an S3 bucket to store the script output.
	// The bucket is unique to your account and Region. When you enable usage reporting
	// in this case, AppStream 2.0 uses the same bucket to store your usage reports. If
	// you haven't already enabled on-instance session scripts, when you enable usage
	// reports, AppStream 2.0 creates a new S3 bucket.
	S3BucketName *string

	// The schedule for generating usage reports.
	Schedule types.UsageReportSchedule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateUsageReportSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUsageReportSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUsageReportSubscription{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUsageReportSubscription(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateUsageReportSubscription(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "CreateUsageReportSubscription",
	}
}
