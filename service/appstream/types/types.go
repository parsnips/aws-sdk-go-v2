// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Describes an interface VPC endpoint (interface endpoint) that lets you create a
// private connection between the virtual private cloud (VPC) that you specify and
// AppStream 2.0. When you specify an interface endpoint for a stack, users of the
// stack can connect to AppStream 2.0 only through that endpoint. When you specify
// an interface endpoint for an image builder, administrators can connect to the
// image builder only through that endpoint.
type AccessEndpoint struct {

	// The type of interface endpoint.
	//
	// This member is required.
	EndpointType AccessEndpointType

	// The identifier (ID) of the VPC in which the interface endpoint is used.
	VpceId *string
}

// Describes an application in the application catalog.
type Application struct {

	// The application name to display.
	DisplayName *string

	// If there is a problem, the application can be disabled after image creation.
	Enabled *bool

	// The URL for the application icon. This URL might be time-limited.
	IconURL *string

	// The arguments that are passed to the application at launch.
	LaunchParameters *string

	// The path to the application executable in the instance.
	LaunchPath *string

	// Additional attributes that describe the application.
	Metadata map[string]*string

	// The name of the application.
	Name *string
}

// The persistent application settings for users of a stack.
type ApplicationSettings struct {

	// Enables or disables persistent application settings for users during their
	// streaming sessions.
	//
	// This member is required.
	Enabled *bool

	// The path prefix for the S3 bucket where users’ persistent application settings
	// are stored. You can allow the same persistent application settings to be used
	// across multiple stacks by specifying the same settings group for each stack.
	SettingsGroup *string
}

// Describes the persistent application settings for users of a stack.
type ApplicationSettingsResponse struct {

	// Specifies whether persistent application settings are enabled for users during
	// their streaming sessions.
	Enabled *bool

	// The S3 bucket where users’ persistent application settings are stored. When
	// persistent application settings are enabled for the first time for an account in
	// an AWS Region, an S3 bucket is created. The bucket is unique to the AWS account
	// and the Region.
	S3BucketName *string

	// The path prefix for the S3 bucket where users’ persistent application settings
	// are stored.
	SettingsGroup *string
}

// Describes the capacity for a fleet.
type ComputeCapacity struct {

	// The desired number of streaming instances.
	//
	// This member is required.
	DesiredInstances *int32
}

// Describes the capacity status for a fleet.
type ComputeCapacityStatus struct {

	// The desired number of streaming instances.
	//
	// This member is required.
	Desired *int32

	// The number of currently available instances that can be used to stream sessions.
	Available *int32

	// The number of instances in use for streaming.
	InUse *int32

	// The total number of simultaneous streaming instances that are running.
	Running *int32
}

// Describes the configuration information required to join fleets and image
// builders to Microsoft Active Directory domains.
type DirectoryConfig struct {

	// The fully qualified name of the directory (for example, corp.example.com).
	//
	// This member is required.
	DirectoryName *string

	// The time the directory configuration was created.
	CreatedTime *time.Time

	// The distinguished names of the organizational units for computer accounts.
	OrganizationalUnitDistinguishedNames []*string

	// The credentials for the service account used by the fleet or image builder to
	// connect to the directory.
	ServiceAccountCredentials *ServiceAccountCredentials
}

// Describes the configuration information required to join fleets and image
// builders to Microsoft Active Directory domains.
type DomainJoinInfo struct {

	// The fully qualified name of the directory (for example, corp.example.com).
	DirectoryName *string

	// The distinguished name of the organizational unit for computer accounts.
	OrganizationalUnitDistinguishedName *string
}

// Describes a fleet.
type Fleet struct {

	// The Amazon Resource Name (ARN) for the fleet.
	//
	// This member is required.
	Arn *string

	// The capacity status for the fleet.
	//
	// This member is required.
	ComputeCapacityStatus *ComputeCapacityStatus

	// The instance type to use when launching fleet instances. The following instance
	// types are available:
	//
	//     * stream.standard.medium
	//
	//     *
	// stream.standard.large
	//
	//     * stream.compute.large
	//
	//     * stream.compute.xlarge
	//
	//
	// * stream.compute.2xlarge
	//
	//     * stream.compute.4xlarge
	//
	//     *
	// stream.compute.8xlarge
	//
	//     * stream.memory.large
	//
	//     * stream.memory.xlarge
	//
	//
	// * stream.memory.2xlarge
	//
	//     * stream.memory.4xlarge
	//
	//     *
	// stream.memory.8xlarge
	//
	//     * stream.graphics-design.large
	//
	//     *
	// stream.graphics-design.xlarge
	//
	//     * stream.graphics-design.2xlarge
	//
	//     *
	// stream.graphics-design.4xlarge
	//
	//     * stream.graphics-desktop.2xlarge
	//
	//     *
	// stream.graphics-pro.4xlarge
	//
	//     * stream.graphics-pro.8xlarge
	//
	//     *
	// stream.graphics-pro.16xlarge
	//
	// This member is required.
	InstanceType *string

	// The name of the fleet.
	//
	// This member is required.
	Name *string

	// The current state for the fleet.
	//
	// This member is required.
	State FleetState

	// The time the fleet was created.
	CreatedTime *time.Time

	// The description to display.
	Description *string

	// The amount of time that a streaming session remains active after users
	// disconnect. If they try to reconnect to the streaming session after a
	// disconnection or network interruption within this time interval, they are
	// connected to their previous session. Otherwise, they are connected to a new
	// session with a new streaming instance. Specify a value between 60 and 360000.
	DisconnectTimeoutInSeconds *int32

	// The fleet name to display.
	DisplayName *string

	// The name of the directory and organizational unit (OU) to use to join the fleet
	// to a Microsoft Active Directory domain.
	DomainJoinInfo *DomainJoinInfo

	// Indicates whether default internet access is enabled for the fleet.
	EnableDefaultInternetAccess *bool

	// The fleet errors.
	FleetErrors []*FleetError

	// The fleet type. ALWAYS_ON Provides users with instant-on access to their apps.
	// You are charged for all running instances in your fleet, even if no users are
	// streaming apps. ON_DEMAND Provide users with access to applications after they
	// connect, which takes one to two minutes. You are charged for instance streaming
	// when users are connected and a small hourly fee for instances that are not
	// streaming apps.
	FleetType FleetType

	// The ARN of the IAM role that is applied to the fleet. To assume a role, the
	// fleet instance calls the AWS Security Token Service (STS) AssumeRole API
	// operation and passes the ARN of the role to use. The operation creates a new
	// session with temporary credentials. AppStream 2.0 retrieves the temporary
	// credentials and creates the AppStream_Machine_Role credential profile on the
	// instance. For more information, see Using an IAM Role to Grant Permissions to
	// Applications and Scripts Running on AppStream 2.0 Streaming Instances
	// (https://docs.aws.amazon.com/appstream2/latest/developerguide/using-iam-roles-to-grant-permissions-to-applications-scripts-streaming-instances.html)
	// in the Amazon AppStream 2.0 Administration Guide.
	IamRoleArn *string

	// The amount of time that users can be idle (inactive) before they are
	// disconnected from their streaming session and the DisconnectTimeoutInSeconds
	// time interval begins. Users are notified before they are disconnected due to
	// inactivity. If users try to reconnect to the streaming session before the time
	// interval specified in DisconnectTimeoutInSeconds elapses, they are connected to
	// their previous session. Users are considered idle when they stop providing
	// keyboard or mouse input during their streaming session. File uploads and
	// downloads, audio in, audio out, and pixels changing do not qualify as user
	// activity. If users continue to be idle after the time interval in
	// IdleDisconnectTimeoutInSeconds elapses, they are disconnected. To prevent users
	// from being disconnected due to inactivity, specify a value of 0. Otherwise,
	// specify a value between 60 and 3600. The default value is 0. If you enable this
	// feature, we recommend that you specify a value that corresponds exactly to a
	// whole number of minutes (for example, 60, 120, and 180). If you don't do this,
	// the value is rounded to the nearest minute. For example, if you specify a value
	// of 70, users are disconnected after 1 minute of inactivity. If you specify a
	// value that is at the midpoint between two different minutes, the value is
	// rounded up. For example, if you specify a value of 90, users are disconnected
	// after 2 minutes of inactivity.
	IdleDisconnectTimeoutInSeconds *int32

	// The ARN for the public, private, or shared image.
	ImageArn *string

	// The name of the image used to create the fleet.
	ImageName *string

	// The maximum amount of time that a streaming session can remain active, in
	// seconds. If users are still connected to a streaming instance five minutes
	// before this limit is reached, they are prompted to save any open documents
	// before being disconnected. After this time elapses, the instance is terminated
	// and replaced by a new instance. Specify a value between 600 and 360000.
	MaxUserDurationInSeconds *int32

	// The VPC configuration for the fleet.
	VpcConfig *VpcConfig
}

// Describes a fleet error.
type FleetError struct {

	// The error code.
	ErrorCode FleetErrorCode

	// The error message.
	ErrorMessage *string
}

// Describes an image.
type Image struct {

	// The name of the image.
	//
	// This member is required.
	Name *string

	// The applications associated with the image.
	Applications []*Application

	// The version of the AppStream 2.0 agent to use for instances that are launched
	// from this image.
	AppstreamAgentVersion *string

	// The ARN of the image.
	Arn *string

	// The ARN of the image from which this image was created.
	BaseImageArn *string

	// The time the image was created.
	CreatedTime *time.Time

	// The description to display.
	Description *string

	// The image name to display.
	DisplayName *string

	// The name of the image builder that was used to create the private image. If the
	// image is shared, this value is null.
	ImageBuilderName *string

	// Indicates whether an image builder can be launched from this image.
	ImageBuilderSupported *bool

	// The permissions to provide to the destination AWS account for the specified
	// image.
	ImagePermissions *ImagePermissions

	// The operating system platform of the image.
	Platform PlatformType

	// The release date of the public base image. For private images, this date is the
	// release date of the base image from which the image was created.
	PublicBaseImageReleasedDate *time.Time

	// The image starts in the PENDING state. If image creation succeeds, the state is
	// AVAILABLE. If image creation fails, the state is FAILED.
	State ImageState

	// The reason why the last state change occurred.
	StateChangeReason *ImageStateChangeReason

	// Indicates whether the image is public or private.
	Visibility VisibilityType
}

// Describes a virtual machine that is used to create an image.
type ImageBuilder struct {

	// The name of the image builder.
	//
	// This member is required.
	Name *string

	// The list of virtual private cloud (VPC) interface endpoint objects.
	// Administrators can connect to the image builder only through the specified
	// endpoints.
	AccessEndpoints []*AccessEndpoint

	// The version of the AppStream 2.0 agent that is currently being used by the image
	// builder.
	AppstreamAgentVersion *string

	// The ARN for the image builder.
	Arn *string

	// The time stamp when the image builder was created.
	CreatedTime *time.Time

	// The description to display.
	Description *string

	// The image builder name to display.
	DisplayName *string

	// The name of the directory and organizational unit (OU) to use to join the image
	// builder to a Microsoft Active Directory domain.
	DomainJoinInfo *DomainJoinInfo

	// Enables or disables default internet access for the image builder.
	EnableDefaultInternetAccess *bool

	// The ARN of the IAM role that is applied to the image builder. To assume a role,
	// the image builder calls the AWS Security Token Service (STS) AssumeRole API
	// operation and passes the ARN of the role to use. The operation creates a new
	// session with temporary credentials. AppStream 2.0 retrieves the temporary
	// credentials and creates the AppStream_Machine_Role credential profile on the
	// instance. For more information, see Using an IAM Role to Grant Permissions to
	// Applications and Scripts Running on AppStream 2.0 Streaming Instances
	// (https://docs.aws.amazon.com/appstream2/latest/developerguide/using-iam-roles-to-grant-permissions-to-applications-scripts-streaming-instances.html)
	// in the Amazon AppStream 2.0 Administration Guide.
	IamRoleArn *string

	// The ARN of the image from which this builder was created.
	ImageArn *string

	// The image builder errors.
	ImageBuilderErrors []*ResourceError

	// The instance type for the image builder. The following instance types are
	// available:
	//
	//     * stream.standard.medium
	//
	//     * stream.standard.large
	//
	//     *
	// stream.compute.large
	//
	//     * stream.compute.xlarge
	//
	//     *
	// stream.compute.2xlarge
	//
	//     * stream.compute.4xlarge
	//
	//     *
	// stream.compute.8xlarge
	//
	//     * stream.memory.large
	//
	//     * stream.memory.xlarge
	//
	//
	// * stream.memory.2xlarge
	//
	//     * stream.memory.4xlarge
	//
	//     *
	// stream.memory.8xlarge
	//
	//     * stream.graphics-design.large
	//
	//     *
	// stream.graphics-design.xlarge
	//
	//     * stream.graphics-design.2xlarge
	//
	//     *
	// stream.graphics-design.4xlarge
	//
	//     * stream.graphics-desktop.2xlarge
	//
	//     *
	// stream.graphics-pro.4xlarge
	//
	//     * stream.graphics-pro.8xlarge
	//
	//     *
	// stream.graphics-pro.16xlarge
	InstanceType *string

	// Describes the network details of the fleet or image builder instance.
	NetworkAccessConfiguration *NetworkAccessConfiguration

	// The operating system platform of the image builder.
	Platform PlatformType

	// The state of the image builder.
	State ImageBuilderState

	// The reason why the last state change occurred.
	StateChangeReason *ImageBuilderStateChangeReason

	// The VPC configuration of the image builder.
	VpcConfig *VpcConfig
}

// Describes the reason why the last image builder state change occurred.
type ImageBuilderStateChangeReason struct {

	// The state change reason code.
	Code ImageBuilderStateChangeReasonCode

	// The state change reason message.
	Message *string
}

// Describes the permissions for an image.
type ImagePermissions struct {

	// Indicates whether the image can be used for a fleet.
	AllowFleet *bool

	// Indicates whether the image can be used for an image builder.
	AllowImageBuilder *bool
}

// Describes the reason why the last image state change occurred.
type ImageStateChangeReason struct {

	// The state change reason code.
	Code ImageStateChangeReasonCode

	// The state change reason message.
	Message *string
}

// Describes the error that is returned when a usage report can't be generated.
type LastReportGenerationExecutionError struct {

	// The error code for the error that is returned when a usage report can't be
	// generated.
	ErrorCode UsageReportExecutionErrorCode

	// The error message for the error that is returned when a usage report can't be
	// generated.
	ErrorMessage *string
}

// Describes the network details of the fleet or image builder instance.
type NetworkAccessConfiguration struct {

	// The resource identifier of the elastic network interface that is attached to
	// instances in your VPC. All network interfaces have the eni-xxxxxxxx resource
	// identifier.
	EniId *string

	// The private IP address of the elastic network interface that is attached to
	// instances in your VPC.
	EniPrivateIpAddress *string
}

// Describes a resource error.
type ResourceError struct {

	// The error code.
	ErrorCode FleetErrorCode

	// The error message.
	ErrorMessage *string

	// The time the error occurred.
	ErrorTimestamp *time.Time
}

// Describes the credentials for the service account used by the fleet or image
// builder to connect to the directory.
type ServiceAccountCredentials struct {

	// The user name of the account. This account must have the following privileges:
	// create computer objects, join computers to the domain, and change/reset the
	// password on descendant computer objects for the organizational units specified.
	//
	// This member is required.
	AccountName *string

	// The password for the account.
	//
	// This member is required.
	AccountPassword *string
}

// Describes a streaming session.
type Session struct {

	// The name of the fleet for the streaming session.
	//
	// This member is required.
	FleetName *string

	// The identifier of the streaming session.
	//
	// This member is required.
	Id *string

	// The name of the stack for the streaming session.
	//
	// This member is required.
	StackName *string

	// The current state of the streaming session.
	//
	// This member is required.
	State SessionState

	// The identifier of the user for whom the session was created.
	//
	// This member is required.
	UserId *string

	// The authentication method. The user is authenticated using a streaming URL (API)
	// or SAML 2.0 federation (SAML).
	AuthenticationType AuthenticationType

	// Specifies whether a user is connected to the streaming session.
	ConnectionState SessionConnectionState

	// The time when the streaming session is set to expire. This time is based on the
	// MaxUserDurationinSeconds value, which determines the maximum length of time that
	// a streaming session can run. A streaming session might end earlier than the time
	// specified in SessionMaxExpirationTime, when the DisconnectTimeOutInSeconds
	// elapses or the user chooses to end his or her session. If the
	// DisconnectTimeOutInSeconds elapses, or the user chooses to end his or her
	// session, the streaming instance is terminated and the streaming session ends.
	MaxExpirationTime *time.Time

	// The network details for the streaming session.
	NetworkAccessConfiguration *NetworkAccessConfiguration

	// The time when a streaming instance is dedicated for the user.
	StartTime *time.Time
}

// Describes the permissions that are available to the specified AWS account for a
// shared image.
type SharedImagePermissions struct {

	// Describes the permissions for a shared image.
	//
	// This member is required.
	ImagePermissions *ImagePermissions

	// The 12-digit identifier of the AWS account with which the image is shared.
	//
	// This member is required.
	SharedAccountId *string
}

// Describes a stack.
type Stack struct {

	// The name of the stack.
	//
	// This member is required.
	Name *string

	// The list of virtual private cloud (VPC) interface endpoint objects. Users of the
	// stack can connect to AppStream 2.0 only through the specified endpoints.
	AccessEndpoints []*AccessEndpoint

	// The persistent application settings for users of the stack.
	ApplicationSettings *ApplicationSettingsResponse

	// The ARN of the stack.
	Arn *string

	// The time the stack was created.
	CreatedTime *time.Time

	// The description to display.
	Description *string

	// The stack name to display.
	DisplayName *string

	// The domains where AppStream 2.0 streaming sessions can be embedded in an iframe.
	// You must approve the domains that you want to host embedded AppStream 2.0
	// streaming sessions.
	EmbedHostDomains []*string

	// The URL that users are redirected to after they click the Send Feedback link. If
	// no URL is specified, no Send Feedback link is displayed.
	FeedbackURL *string

	// The URL that users are redirected to after their streaming session ends.
	RedirectURL *string

	// The errors for the stack.
	StackErrors []*StackError

	// The storage connectors to enable.
	StorageConnectors []*StorageConnector

	// The actions that are enabled or disabled for users during their streaming
	// sessions. By default these actions are enabled.
	UserSettings []*UserSetting
}

// Describes a stack error.
type StackError struct {

	// The error code.
	ErrorCode StackErrorCode

	// The error message.
	ErrorMessage *string
}

// Describes a connector that enables persistent storage for users.
type StorageConnector struct {

	// The type of storage connector.
	//
	// This member is required.
	ConnectorType StorageConnectorType

	// The names of the domains for the account.
	Domains []*string

	// The ARN of the storage connector.
	ResourceIdentifier *string
}

// Describes information about the usage report subscription.
type UsageReportSubscription struct {

	// The time when the last usage report was generated.
	LastGeneratedReportDate *time.Time

	// The Amazon S3 bucket where generated reports are stored. If you enabled
	// on-instance session scripts and Amazon S3 logging for your session script
	// configuration, AppStream 2.0 created an S3 bucket to store the script output.
	// The bucket is unique to your account and Region. When you enable usage reporting
	// in this case, AppStream 2.0 uses the same bucket to store your usage reports. If
	// you haven't already enabled on-instance session scripts, when you enable usage
	// reports, AppStream 2.0 creates a new S3 bucket.
	S3BucketName *string

	// The schedule for generating usage reports.
	Schedule UsageReportSchedule

	// The errors that were returned if usage reports couldn't be generated.
	SubscriptionErrors []*LastReportGenerationExecutionError
}

// Describes a user in the user pool.
type User struct {

	// The authentication type for the user.
	//
	// This member is required.
	AuthenticationType AuthenticationType

	// The ARN of the user.
	Arn *string

	// The date and time the user was created in the user pool.
	CreatedTime *time.Time

	// Specifies whether the user in the user pool is enabled.
	Enabled *bool

	// The first name, or given name, of the user.
	FirstName *string

	// The last name, or surname, of the user.
	LastName *string

	// The status of the user in the user pool. The status can be one of the
	// following:
	//
	//     * UNCONFIRMED – The user is created but not confirmed.
	//
	//     *
	// CONFIRMED – The user is confirmed.
	//
	//     * ARCHIVED – The user is no longer
	// active.
	//
	//     * COMPROMISED – The user is disabled because of a potential
	// security threat.
	//
	//     * UNKNOWN – The user status is not known.
	Status *string

	// The email address of the user. Users' email addresses are case-sensitive.
	UserName *string
}

// Describes an action and whether the action is enabled or disabled for users
// during their streaming sessions.
type UserSetting struct {

	// The action that is enabled or disabled.
	//
	// This member is required.
	Action Action

	// Indicates whether the action is enabled or disabled.
	//
	// This member is required.
	Permission Permission
}

// Describes a user in the user pool and the associated stack.
type UserStackAssociation struct {

	// The authentication type for the user.
	//
	// This member is required.
	AuthenticationType AuthenticationType

	// The name of the stack that is associated with the user.
	//
	// This member is required.
	StackName *string

	// The email address of the user who is associated with the stack. Users' email
	// addresses are case-sensitive.
	//
	// This member is required.
	UserName *string

	// Specifies whether a welcome email is sent to a user after the user is created in
	// the user pool.
	SendEmailNotification *bool
}

// Describes the error that is returned when a user can’t be associated with or
// disassociated from a stack.
type UserStackAssociationError struct {

	// The error code for the error that is returned when a user can’t be associated
	// with or disassociated from a stack.
	ErrorCode UserStackAssociationErrorCode

	// The error message for the error that is returned when a user can’t be associated
	// with or disassociated from a stack.
	ErrorMessage *string

	// Information about the user and associated stack.
	UserStackAssociation *UserStackAssociation
}

// Describes VPC configuration information for fleets and image builders.
type VpcConfig struct {

	// The identifiers of the security groups for the fleet or image builder.
	SecurityGroupIds []*string

	// The identifiers of the subnets to which a network interface is attached from the
	// fleet instance or image builder instance. Fleet instances use one or more
	// subnets. Image builder instances use one subnet.
	SubnetIds []*string
}
