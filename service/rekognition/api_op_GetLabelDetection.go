// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the label detection results of a Amazon Rekognition Video analysis started
// by StartLabelDetection. The label detection operation is started by a call to
// StartLabelDetection which returns a job identifier (JobId). When the label
// detection operation finishes, Amazon Rekognition publishes a completion status
// to the Amazon Simple Notification Service topic registered in the initial call
// to StartlabelDetection. To get the results of the label detection operation,
// first check that the status value published to the Amazon SNS topic is
// SUCCEEDED. If so, call GetLabelDetection and pass the job identifier (JobId)
// from the initial call to StartLabelDetection. GetLabelDetection returns an array
// of detected labels (Labels) sorted by the time the labels were detected. You can
// also sort by the label name by specifying NAME for the SortBy input parameter.
// The labels returned include the label name, the percentage confidence in the
// accuracy of the detected label, and the time the label was detected in the
// video. The returned labels also include bounding box information for common
// objects, a hierarchical taxonomy of detected labels, and the version of the
// label model used for detection. Use MaxResults parameter to limit the number of
// labels returned. If there are more results than specified in MaxResults, the
// value of NextToken in the operation response contains a pagination token for
// getting the next set of results. To get the next page of results, call
// GetlabelDetection and populate the NextToken request parameter with the token
// value returned from the previous call to GetLabelDetection.
func (c *Client) GetLabelDetection(ctx context.Context, params *GetLabelDetectionInput, optFns ...func(*Options)) (*GetLabelDetectionOutput, error) {
	if params == nil {
		params = &GetLabelDetectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLabelDetection", params, optFns, addOperationGetLabelDetectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLabelDetectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLabelDetectionInput struct {

	// Job identifier for the label detection operation for which you want results
	// returned. You get the job identifer from an initial call to StartlabelDetection.
	//
	// This member is required.
	JobId *string

	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000. If you specify a value greater than 1000, a maximum of 1000
	// results is returned. The default value is 1000.
	MaxResults *int32

	// If the previous response was incomplete (because there are more labels to
	// retrieve), Amazon Rekognition Video returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of labels.
	NextToken *string

	// Sort to use for elements in the Labels array. Use TIMESTAMP to sort array
	// elements by the time labels are detected. Use NAME to alphabetically group
	// elements for a label together. Within each label group, the array element are
	// sorted by detection confidence. The default sort is by TIMESTAMP.
	SortBy types.LabelDetectionSortBy
}

type GetLabelDetectionOutput struct {

	// The current status of the label detection job.
	JobStatus types.VideoJobStatus

	// Version number of the label detection model that was used to detect labels.
	LabelModelVersion *string

	// An array of labels detected in the video. Each element contains the detected
	// label and the time, in milliseconds from the start of the video, that the label
	// was detected.
	Labels []*types.LabelDetection

	// If the response is truncated, Amazon Rekognition Video returns this token that
	// you can use in the subsequent request to retrieve the next set of labels.
	NextToken *string

	// If the job fails, StatusMessage provides a descriptive error message.
	StatusMessage *string

	// Information about a video that Amazon Rekognition Video analyzed. Videometadata
	// is returned in every page of paginated responses from a Amazon Rekognition video
	// operation.
	VideoMetadata *types.VideoMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetLabelDetectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetLabelDetection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLabelDetection{}, middleware.After)
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
	addOpGetLabelDetectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetLabelDetection(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetLabelDetection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "GetLabelDetection",
	}
}
