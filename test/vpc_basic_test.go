package test

import (
	"testing"

	"github.com/cloudposse/terratest-helpers/pkg/atmos"
	"github.com/stretchr/testify/require"
)

func TestVPCBasic(t *testing.T) {
	testOpts := atmos.NewAwsComponentTestOptions("us-east-2", "vpc", "test-use2-basic")
	atmos.AwsComponentTestHelper(t, testOpts, func(t *testing.T, opts *atmos.Options, output string) {
		vpcCIDR := atmos.Output(t, opts, "vpc_cidr")

		require.Equal(t, "10.14.0.0/16", vpcCIDR)
	})
}
