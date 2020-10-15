// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Updates the environment description, deploys a new application version, updates
// the configuration settings to an entirely new configuration template, or updates
// select configuration option values in the running environment. Attempting to
// update both the release and configuration is not allowed and AWS Elastic
// Beanstalk returns an InvalidParameterCombination error. When updating the
// configuration settings to a new template or individual settings, a draft
// configuration is created and DescribeConfigurationSettings for this environment
// returns two setting descriptions with different DeploymentStatus values.
func (c *Client) UpdateEnvironment(ctx context.Context, params *UpdateEnvironmentInput, optFns ...func(*Options)) (*UpdateEnvironmentOutput, error) {
	if params == nil {
		params = &UpdateEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEnvironment", params, optFns, addOperationUpdateEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to update an environment.
type UpdateEnvironmentInput struct {

	// The name of the application with which the environment is associated.
	ApplicationName *string

	// If this parameter is specified, AWS Elastic Beanstalk updates the description of
	// this environment.
	Description *string

	// The ID of the environment to update. If no environment with this ID exists, AWS
	// Elastic Beanstalk returns an InvalidParameterValue error. Condition: You must
	// specify either this or an EnvironmentName, or both. If you do not specify
	// either, AWS Elastic Beanstalk returns MissingRequiredParameter error.
	EnvironmentId *string

	// The name of the environment to update. If no environment with this name exists,
	// AWS Elastic Beanstalk returns an InvalidParameterValue error. Condition: You
	// must specify either this or an EnvironmentId, or both. If you do not specify
	// either, AWS Elastic Beanstalk returns MissingRequiredParameter error.
	EnvironmentName *string

	// The name of the group to which the target environment belongs. Specify a group
	// name only if the environment's name is specified in an environment manifest and
	// not with the environment name or environment ID parameters. See Environment
	// Manifest (env.yaml)
	// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/environment-cfg-manifest.html)
	// for details.
	GroupName *string

	// If specified, AWS Elastic Beanstalk updates the configuration set associated
	// with the running environment and sets the specified configuration options to the
	// requested value.
	OptionSettings []*types.ConfigurationOptionSetting

	// A list of custom user-defined configuration options to remove from the
	// configuration set for this environment.
	OptionsToRemove []*types.OptionSpecification

	// The ARN of the platform, if used.
	PlatformArn *string

	// This specifies the platform version that the environment will run after the
	// environment is updated.
	SolutionStackName *string

	// If this parameter is specified, AWS Elastic Beanstalk deploys this configuration
	// template to the environment. If no such configuration template is found, AWS
	// Elastic Beanstalk returns an InvalidParameterValue error.
	TemplateName *string

	// This specifies the tier to use to update the environment. Condition: At this
	// time, if you change the tier version, name, or type, AWS Elastic Beanstalk
	// returns InvalidParameterValue error.
	Tier *types.EnvironmentTier

	// If this parameter is specified, AWS Elastic Beanstalk deploys the named
	// application version to the environment. If no such application version is found,
	// returns an InvalidParameterValue error.
	VersionLabel *string
}

// Describes the properties of an environment.
type UpdateEnvironmentOutput struct {

	// Indicates if there is an in-progress environment configuration update or
	// application version deployment that you can cancel. true: There is an update in
	// progress. false: There are no updates currently in progress.
	AbortableOperationInProgress *bool

	// The name of the application associated with this environment.
	ApplicationName *string

	// The URL to the CNAME for this environment.
	CNAME *string

	// The creation date for this environment.
	DateCreated *time.Time

	// The last modified date for this environment.
	DateUpdated *time.Time

	// Describes this environment.
	Description *string

	// For load-balanced, autoscaling environments, the URL to the LoadBalancer. For
	// single-instance environments, the IP address of the instance.
	EndpointURL *string

	// The environment's Amazon Resource Name (ARN), which can be used in other API
	// requests that require an ARN.
	EnvironmentArn *string

	// The ID of this environment.
	EnvironmentId *string

	// A list of links to other environments in the same group.
	EnvironmentLinks []*types.EnvironmentLink

	// The name of this environment.
	EnvironmentName *string

	// Describes the health status of the environment. AWS Elastic Beanstalk indicates
	// the failure levels for a running environment:
	//
	//     * Red: Indicates the
	// environment is not responsive. Occurs when three or more consecutive failures
	// occur for an environment.
	//
	//     * Yellow: Indicates that something is wrong.
	// Occurs when two consecutive failures occur for an environment.
	//
	//     * Green:
	// Indicates the environment is healthy and fully functional.
	//
	//     * Grey: Default
	// health for a new environment. The environment is not fully launched and health
	// checks have not started or health checks are suspended during an
	// UpdateEnvironment or RestartEnvironment request.
	//
	// Default: Grey
	Health types.EnvironmentHealth

	// Returns the health status of the application running in your environment. For
	// more information, see Health Colors and Statuses
	// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/health-enhanced-status.html).
	HealthStatus types.EnvironmentHealthStatus

	// The Amazon Resource Name (ARN) of the environment's operations role. For more
	// information, see Operations roles
	// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/iam-operationsrole.html)
	// in the AWS Elastic Beanstalk Developer Guide.
	OperationsRole *string

	// The ARN of the platform version.
	PlatformArn *string

	// The description of the AWS resources used by this environment.
	Resources *types.EnvironmentResourcesDescription

	// The name of the SolutionStack deployed with this environment.
	SolutionStackName *string

	// The current operational status of the environment:
	//
	//     * Launching: Environment
	// is in the process of initial deployment.
	//
	//     * Updating: Environment is in the
	// process of updating its configuration settings or application version.
	//
	//     *
	// Ready: Environment is available to have an action performed on it, such as
	// update or terminate.
	//
	//     * Terminating: Environment is in the shut-down
	// process.
	//
	//     * Terminated: Environment is not running.
	Status types.EnvironmentStatus

	// The name of the configuration template used to originally launch this
	// environment.
	TemplateName *string

	// Describes the current tier of this environment.
	Tier *types.EnvironmentTier

	// The application version deployed in this environment.
	VersionLabel *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateEnvironment{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEnvironment(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateEnvironment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "UpdateEnvironment",
	}
}
