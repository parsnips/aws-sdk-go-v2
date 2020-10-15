// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Accepts the transfer of a domain from another AWS account to the current AWS
// account. You initiate a transfer between AWS accounts using
// TransferDomainToAnotherAwsAccount
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html).
// Use either ListOperations
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ListOperations.html)
// or GetOperationDetail
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html)
// to determine whether the operation succeeded. GetOperationDetail
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html)
// provides additional information, for example, Domain Transfer from Aws Account
// 111122223333 has been cancelled.
func (c *Client) AcceptDomainTransferFromAnotherAwsAccount(ctx context.Context, params *AcceptDomainTransferFromAnotherAwsAccountInput, optFns ...func(*Options)) (*AcceptDomainTransferFromAnotherAwsAccountOutput, error) {
	if params == nil {
		params = &AcceptDomainTransferFromAnotherAwsAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AcceptDomainTransferFromAnotherAwsAccount", params, optFns, addOperationAcceptDomainTransferFromAnotherAwsAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AcceptDomainTransferFromAnotherAwsAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The AcceptDomainTransferFromAnotherAwsAccount request includes the following
// elements.
type AcceptDomainTransferFromAnotherAwsAccountInput struct {

	// The name of the domain that was specified when another AWS account submitted a
	// TransferDomainToAnotherAwsAccount
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html)
	// request.
	//
	// This member is required.
	DomainName *string

	// The password that was returned by the TransferDomainToAnotherAwsAccount
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html)
	// request.
	//
	// This member is required.
	Password *string
}

// The AcceptDomainTransferFromAnotherAwsAccount response includes the following
// element.
type AcceptDomainTransferFromAnotherAwsAccountOutput struct {

	// Identifier for tracking the progress of the request. To query the operation
	// status, use GetOperationDetail
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html).
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAcceptDomainTransferFromAnotherAwsAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAcceptDomainTransferFromAnotherAwsAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAcceptDomainTransferFromAnotherAwsAccount{}, middleware.After)
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
	addOpAcceptDomainTransferFromAnotherAwsAccountValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptDomainTransferFromAnotherAwsAccount(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAcceptDomainTransferFromAnotherAwsAccount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "AcceptDomainTransferFromAnotherAwsAccount",
	}
}
