package net_test

import (
	"testing"

	tfnet "github.com/terraform-providers/terraform-provider-aws/aws/internal/net"
)

func Test_CIDRBlocksEqual(t *testing.T) {
	for _, ts := range []struct {
		cidr1 string
		cidr2 string
		equal bool
	}{
		{"10.2.2.0/24", "10.2.2.0/24", true},
		{"10.2.2.0/1234", "10.2.2.0/24", false},
		{"10.2.2.0/24", "10.2.2.0/1234", false},
		{"2001::/15", "2001::/15", true},
		{"::/0", "2001::/15", false},
		{"::/0", "::0/0", true},
		{"", "", false},
	} {
		equal := tfnet.CIDRBlocksEqual(ts.cidr1, ts.cidr2)
		if ts.equal != equal {
			t.Fatalf("CIDRBlocksEqual(%q, %q) should be: %t", ts.cidr1, ts.cidr2, ts.equal)
		}
	}
}
