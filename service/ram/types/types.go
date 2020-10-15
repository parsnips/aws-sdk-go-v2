// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Describes a principal for use with AWS Resource Access Manager.
type Principal struct {

	// The time when the principal was associated with the resource share.
	CreationTime *time.Time

	// Indicates whether the principal belongs to the same AWS organization as the AWS
	// account that owns the resource share.
	External *bool

	// The ID of the principal.
	Id *string

	// The time when the association was last updated.
	LastUpdatedTime *time.Time

	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn *string
}

// Describes a resource associated with a resource share.
type Resource struct {

	// The Amazon Resource Name (ARN) of the resource.
	Arn *string

	// The time when the resource was associated with the resource share.
	CreationTime *time.Time

	// The time when the association was last updated.
	LastUpdatedTime *time.Time

	// The ARN of the resource group. This value is returned only if the resource is a
	// resource group.
	ResourceGroupArn *string

	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn *string

	// The status of the resource.
	Status ResourceStatus

	// A message about the status of the resource.
	StatusMessage *string

	// The resource type.
	Type *string
}

// Describes a resource share.
type ResourceShare struct {

	// Indicates whether principals outside your AWS organization can be associated
	// with a resource share.
	AllowExternalPrincipals *bool

	// The time when the resource share was created.
	CreationTime *time.Time

	// Indicates how the resource share was created. Possible values include:
	//
	//     *
	// CREATED_FROM_POLICY - Indicates that the resource share was created from an AWS
	// Identity and Access Management (AWS IAM) policy attached to a resource. These
	// resource shares are visible only to the AWS account that created it. They cannot
	// be modified in AWS RAM.
	//
	//     * PROMOTING_TO_STANDARD - The resource share is in
	// the process of being promoted. For more information, see
	// PromoteResourceShareCreatedFromPolicy.
	//
	//     * STANDARD - Indicates that the
	// resource share was created in AWS RAM using the console or APIs. These resource
	// shares are visible to all principals. They can be modified in AWS RAM.
	FeatureSet ResourceShareFeatureSet

	// The time when the resource share was last updated.
	LastUpdatedTime *time.Time

	// The name of the resource share.
	Name *string

	// The ID of the AWS account that owns the resource share.
	OwningAccountId *string

	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn *string

	// The status of the resource share.
	Status ResourceShareStatus

	// A message about the status of the resource share.
	StatusMessage *string

	// The tags for the resource share.
	Tags []*Tag
}

// Describes an association with a resource share.
type ResourceShareAssociation struct {

	// The associated entity. For resource associations, this is the ARN of the
	// resource. For principal associations, this is the ID of an AWS account or the
	// ARN of an OU or organization from AWS Organizations.
	AssociatedEntity *string

	// The association type.
	AssociationType ResourceShareAssociationType

	// The time when the association was created.
	CreationTime *time.Time

	// Indicates whether the principal belongs to the same AWS organization as the AWS
	// account that owns the resource share.
	External *bool

	// The time when the association was last updated.
	LastUpdatedTime *time.Time

	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn *string

	// The name of the resource share.
	ResourceShareName *string

	// The status of the association.
	Status ResourceShareAssociationStatus

	// A message about the status of the association.
	StatusMessage *string
}

// Describes an invitation to join a resource share.
type ResourceShareInvitation struct {

	// The date and time when the invitation was sent.
	InvitationTimestamp *time.Time

	// The ID of the AWS account that received the invitation.
	ReceiverAccountId *string

	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn *string

	// To view the resources associated with a pending resource share invitation, use
	// ListPendingInvitationResources
	// (https://docs.aws.amazon.com/ram/latest/APIReference/API_ListPendingInvitationResources.html).
	ResourceShareAssociations []*ResourceShareAssociation

	// The Amazon Resource Name (ARN) of the invitation.
	ResourceShareInvitationArn *string

	// The name of the resource share.
	ResourceShareName *string

	// The ID of the AWS account that sent the invitation.
	SenderAccountId *string

	// The status of the invitation.
	Status ResourceShareInvitationStatus
}

// Information about an AWS RAM permission.
type ResourceSharePermissionDetail struct {

	// The ARN of the permission.
	Arn *string

	// The date and time when the permission was created.
	CreationTime *time.Time

	// The identifier for the version of the permission that is set as the default
	// version.
	DefaultVersion *bool

	// The date and time when the permission was last updated.
	LastUpdatedTime *time.Time

	// The name of the permission.
	Name *string

	// The permission's effect and actions in JSON format. The effect indicates whether
	// the actions are allowed or denied. The actions list the API actions to which the
	// principal is granted or denied access.
	Permission *string

	// The resource type to which the permission applies.
	ResourceType *string

	// The identifier for the version of the permission.
	Version *string
}

// Information about a permission that is associated with a resource share.
type ResourceSharePermissionSummary struct {

	// The ARN of the permission.
	Arn *string

	// The date and time when the permission was created.
	CreationTime *time.Time

	// The identifier for the version of the permission that is set as the default
	// version.
	DefaultVersion *bool

	// The date and time when the permission was last updated.
	LastUpdatedTime *time.Time

	// The name of the permission.
	Name *string

	// The type of resource to which the permission applies.
	ResourceType *string

	// The current status of the permission.
	Status *string

	// The identifier for the version of the permission.
	Version *string
}

// Information about the shareable resource types and the AWS services to which
// they belong.
type ServiceNameAndResourceType struct {

	// The shareable resource types.
	ResourceType *string

	// The name of the AWS services to which the resources belong.
	ServiceName *string
}

// Information about a tag.
type Tag struct {

	// The key of the tag.
	Key *string

	// The value of the tag.
	Value *string
}

// Used to filter information based on tags.
type TagFilter struct {

	// The tag key.
	TagKey *string

	// The tag values.
	TagValues []*string
}
