// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about an association between a specified resolver rule and a
// VPC. You associate a resolver rule and a VPC using AssociateResolverRule.
func (c *Client) GetResolverRuleAssociation(ctx context.Context, params *GetResolverRuleAssociationInput, optFns ...func(*Options)) (*GetResolverRuleAssociationOutput, error) {
	if params == nil {
		params = &GetResolverRuleAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResolverRuleAssociation", params, optFns, addOperationGetResolverRuleAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResolverRuleAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResolverRuleAssociationInput struct {

	// The ID of the resolver rule association that you want to get information about.
	//
	// This member is required.
	ResolverRuleAssociationId *string
}

type GetResolverRuleAssociationOutput struct {

	// Information about the resolver rule association that you specified in a
	// GetResolverRuleAssociation request.
	ResolverRuleAssociation *types.ResolverRuleAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetResolverRuleAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetResolverRuleAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetResolverRuleAssociation{}, middleware.After)
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
	addOpGetResolverRuleAssociationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetResolverRuleAssociation(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetResolverRuleAssociation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "GetResolverRuleAssociation",
	}
}
