// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/emrserverless/types"
)

func ExampleJobDriver_outputUsage() {
	var union types.JobDriver
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.JobDriverMemberHive:
		_ = v.Value // Value is types.Hive

	case *types.JobDriverMemberSparkSubmit:
		_ = v.Value // Value is types.SparkSubmit

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.SparkSubmit
var _ *types.Hive
