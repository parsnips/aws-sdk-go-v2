// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// The values of a given attribute, such as Throughput Optimized HDD or Provisioned
// IOPS for the Amazon EC2volumeType attribute.
type AttributeValue struct {

	// The specific value of an attributeName.
	Value *string
}

// The constraints that you want all returned products to match.
type Filter struct {

	// The product metadata field that you want to filter on. You can filter by just
	// the service code to see all products for a specific service, filter by just the
	// attribute name to see a specific attribute for multiple services, or use both a
	// service code and an attribute name to retrieve only products that match both
	// fields. Valid values include: ServiceCode, and all attribute names For example,
	// you can filter by the AmazonEC2 service code and the volumeType attribute name
	// to get the prices for only Amazon EC2 volumes.
	//
	// This member is required.
	Field *string

	// The type of filter that you want to use. Valid values are: TERM_MATCH.
	// TERM_MATCH returns only products that match both the given filter field and the
	// given value.
	//
	// This member is required.
	Type FilterType

	// The service code or attribute value that you want to filter by. If you are
	// filtering by service code this is the actual service code, such as AmazonEC2. If
	// you are filtering by attribute name, this is the attribute value that you want
	// the returned products to match, such as a Provisioned IOPS volume.
	//
	// This member is required.
	Value *string
}

// The metadata for a service, such as the service code and available attribute
// names.
type Service struct {

	// The attributes that are available for this service.
	AttributeNames []*string

	// The code for the AWS service.
	ServiceCode *string
}
