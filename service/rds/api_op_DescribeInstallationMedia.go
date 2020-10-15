// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the available installation media for a DB engine that requires an
// on-premises customer provided license, such as Microsoft SQL Server.
func (c *Client) DescribeInstallationMedia(ctx context.Context, params *DescribeInstallationMediaInput, optFns ...func(*Options)) (*DescribeInstallationMediaOutput, error) {
	if params == nil {
		params = &DescribeInstallationMediaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstallationMedia", params, optFns, addOperationDescribeInstallationMediaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstallationMediaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstallationMediaInput struct {

	// A filter that specifies one or more installation media to describe. Supported
	// filters include the following:
	//
	//     * custom-availability-zone-id - Accepts
	// custom Availability Zone (AZ) identifiers. The results list includes information
	// about only the custom AZs identified by these identifiers.
	//
	//     * engine -
	// Accepts database engines. The results list includes information about only the
	// database engines identified by these identifiers. For more information about the
	// valid engines for installation media, see ImportInstallationMedia.
	Filters []*types.Filter

	// The installation medium ID.
	InstallationMediaId *string

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// An optional pagination token provided by a previous DescribeInstallationMedia
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	MaxRecords *int32
}

type DescribeInstallationMediaOutput struct {

	// The list of InstallationMedia objects for the AWS account.
	InstallationMedia []*types.InstallationMedia

	// An optional pagination token provided by a previous DescribeInstallationMedia
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeInstallationMediaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeInstallationMedia{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeInstallationMedia{}, middleware.After)
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
	addOpDescribeInstallationMediaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstallationMedia(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeInstallationMedia(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeInstallationMedia",
	}
}
