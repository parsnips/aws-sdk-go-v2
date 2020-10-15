// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the parameters of a cluster parameter group. To modify more than one
// parameter, submit a list of the following: ParameterName, ParameterValue, and
// ApplyMethod. A maximum of 20 parameters can be modified in a single request.
// Changes to dynamic parameters are applied immediately. Changes to static
// parameters require a reboot or maintenance window before the change can take
// effect. After you create a cluster parameter group, you should wait at least 5
// minutes before creating your first cluster that uses that cluster parameter
// group as the default parameter group. This allows Amazon DocumentDB to fully
// complete the create action before the parameter group is used as the default for
// a new cluster. This step is especially important for parameters that are
// critical when creating the default database for a cluster, such as the character
// set for the default database defined by the character_set_database parameter.
func (c *Client) ModifyDBClusterParameterGroup(ctx context.Context, params *ModifyDBClusterParameterGroupInput, optFns ...func(*Options)) (*ModifyDBClusterParameterGroupOutput, error) {
	if params == nil {
		params = &ModifyDBClusterParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBClusterParameterGroup", params, optFns, addOperationModifyDBClusterParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBClusterParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to ModifyDBClusterParameterGroup.
type ModifyDBClusterParameterGroupInput struct {

	// The name of the cluster parameter group to modify.
	//
	// This member is required.
	DBClusterParameterGroupName *string

	// A list of parameters in the cluster parameter group to modify.
	//
	// This member is required.
	Parameters []*types.Parameter
}

// Contains the name of a cluster parameter group.
type ModifyDBClusterParameterGroupOutput struct {

	// The name of a cluster parameter group. Constraints:
	//
	//     * Must be from 1 to 255
	// letters or numbers.
	//
	//     * The first character must be a letter.
	//
	//     * Cannot
	// end with a hyphen or contain two consecutive hyphens.
	//
	// This value is stored as a
	// lowercase string.
	DBClusterParameterGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyDBClusterParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBClusterParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBClusterParameterGroup{}, middleware.After)
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
	addOpModifyDBClusterParameterGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBClusterParameterGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opModifyDBClusterParameterGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBClusterParameterGroup",
	}
}
