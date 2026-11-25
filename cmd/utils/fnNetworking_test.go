package utils

import (
	"net"
	"reflect"
	"strings"
	"testing"
)

func mustParseCIDR(t *testing.T, cidr string) *net.IPNet {
	t.Helper()
	_, network, err := net.ParseCIDR(cidr)
	if err != nil {
		t.Fatalf("bad test CIDR %q: %v", cidr, err)
	}
	return network
}

func TestGenSubnetsInNetwork(t *testing.T) {
	tests := []struct {
		name     string
		cidr     string
		maskSize int
		want     []string
	}{
		{"/24 into /26", "10.0.0.0/24", 26, []string{"10.0.0.0/26", "10.0.0.64/26", "10.0.0.128/26", "10.0.0.192/26"}},
		{"/24 into /25", "192.168.1.0/24", 25, []string{"192.168.1.0/25", "192.168.1.128/25"}},
		{"same size returns the network itself", "172.16.0.0/16", 16, []string{"172.16.0.0/16"}},
		{"crosses an octet boundary", "10.1.0.0/22", 24, []string{"10.1.0.0/24", "10.1.1.0/24", "10.1.2.0/24", "10.1.3.0/24"}},
		{"/30 into /32", "10.0.0.4/30", 32, []string{"10.0.0.4/32", "10.0.0.5/32", "10.0.0.6/32", "10.0.0.7/32"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GenSubnetsInNetwork(tc.cidr, tc.maskSize)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGenSubnetsInNetworkCount(t *testing.T) {
	got, err := GenSubnetsInNetwork("10.0.0.0/16", 24)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 256 {
		t.Fatalf("expected 256 /24 subnets in a /16, got %d", len(got))
	}
	if got[0] != "10.0.0.0/24" || got[255] != "10.0.255.0/24" {
		t.Fatalf("unexpected first and last subnets: %q, %q", got[0], got[255])
	}

	// Every generated subnet sits inside the parent network, with no repeats
	parent := mustParseCIDR(t, "10.0.0.0/16")
	seen := map[string]bool{}
	for _, cidr := range got {
		ip, _, err := net.ParseCIDR(cidr)
		if err != nil {
			t.Fatalf("generated an invalid CIDR %q: %v", cidr, err)
		}
		if !parent.Contains(ip) {
			t.Fatalf("%q is outside the parent network", cidr)
		}
		if seen[cidr] {
			t.Fatalf("%q was generated twice", cidr)
		}
		seen[cidr] = true
	}
}

func TestGenSubnetsInNetworkErrors(t *testing.T) {
	tests := []struct {
		name     string
		cidr     string
		maskSize int
		wantErr  string
	}{
		{"not a CIDR", "not-a-cidr", 24, "invalid CIDR"},
		{"missing mask", "10.0.0.0", 24, "invalid CIDR"},
		{"host address rather than network address", "10.0.0.1/24", 26, "not a valid network address"},
		{"subnet mask larger than the network", "10.0.0.0/24", 16, "subnetMaskSize must be greater or equal"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GenSubnetsInNetwork(tc.cidr, tc.maskSize)
			if err == nil {
				t.Fatalf("expected an error, got %v", got)
			}
			if !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("expected error containing %q, got %q", tc.wantErr, err.Error())
			}
		})
	}
}

func TestIp2IntAndInt2Ip(t *testing.T) {
	tests := []struct {
		ip  string
		num uint32
	}{
		{"0.0.0.0", 0},
		{"0.0.0.1", 1},
		{"10.0.0.1", 167772161},
		{"192.168.1.1", 3232235777},
		{"255.255.255.255", 4294967295},
	}
	for _, tc := range tests {
		t.Run(tc.ip, func(t *testing.T) {
			if got := ip2int(net.ParseIP(tc.ip).To4()); got != tc.num {
				t.Fatalf("ip2int: got %d, want %d", got, tc.num)
			}
			if got := int2ip(tc.num).String(); got != tc.ip {
				t.Fatalf("int2ip: got %q, want %q", got, tc.ip)
			}
		})
	}
}

func TestIp2IntPanicsOnIPv6Length(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected a panic for a 16 byte address")
		}
	}()
	ip2int(net.ParseIP("2001:db8::1"))
}

func TestPreviousSubnet(t *testing.T) {
	tests := []struct {
		name      string
		cidr      string
		prefixLen int
		want      string
		wantRoll  bool
	}{
		{"same size", "10.0.1.0/24", 24, "10.0.0.0/24", false},
		{"smaller previous subnet", "10.0.1.0/24", 26, "10.0.0.192/26", false},
		{"larger previous subnet", "10.0.4.0/24", 22, "10.0.0.0/22", false},
		{"crosses second octet", "10.1.0.0/16", 16, "10.0.0.0/16", false},
		{"rolls over below zero", "0.0.0.0/24", 24, "255.255.255.0/24", true},
		{"ipv6", "2001:db8:0:1::/64", 64, "2001:db8::/64", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, rolled := PreviousSubnet(mustParseCIDR(t, tc.cidr), tc.prefixLen)
			if got.String() != tc.want || rolled != tc.wantRoll {
				t.Fatalf("got (%s, %v), want (%s, %v)", got, rolled, tc.want, tc.wantRoll)
			}
		})
	}
}

func TestNextSubnet(t *testing.T) {
	tests := []struct {
		name      string
		cidr      string
		prefixLen int
		want      string
		wantRoll  bool
	}{
		{"same size", "10.0.0.0/24", 24, "10.0.1.0/24", false},
		{"smaller next subnet", "10.0.0.0/24", 26, "10.0.1.0/26", false},
		{"larger next subnet", "10.0.0.0/24", 16, "10.1.0.0/16", false},
		{"crosses second octet", "10.0.255.0/24", 24, "10.1.0.0/24", false},
		{"rolls over past the last address", "255.255.255.0/24", 24, "0.0.0.0/24", true},
		{"ipv6", "2001:db8::/64", 64, "2001:db8:0:1::/64", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, rolled := NextSubnet(mustParseCIDR(t, tc.cidr), tc.prefixLen)
			if got.String() != tc.want || rolled != tc.wantRoll {
				t.Fatalf("got (%s, %v), want (%s, %v)", got, rolled, tc.want, tc.wantRoll)
			}
		})
	}
}

func TestNextAndPreviousSubnetAreInverse(t *testing.T) {
	start := mustParseCIDR(t, "172.16.8.0/22")
	next, rolled := NextSubnet(start, 22)
	if rolled {
		t.Fatal("unexpected rollover")
	}
	back, rolled := PreviousSubnet(next, 22)
	if rolled {
		t.Fatal("unexpected rollover")
	}
	if back.String() != start.String() {
		t.Fatalf("expected to land back on %s, got %s", start, back)
	}
}

func TestAddressRange(t *testing.T) {
	tests := []struct {
		cidr      string
		wantFirst string
		wantLast  string
	}{
		{"192.168.1.0/24", "192.168.1.0", "192.168.1.255"},
		{"10.0.0.0/8", "10.0.0.0", "10.255.255.255"},
		{"10.0.0.64/26", "10.0.0.64", "10.0.0.127"},
		{"10.0.0.4/31", "10.0.0.4", "10.0.0.5"},
		{"10.0.0.7/32", "10.0.0.7", "10.0.0.7"},
		{"0.0.0.0/0", "0.0.0.0", "255.255.255.255"},
		{"2001:db8::/126", "2001:db8::", "2001:db8::3"},
		{"2001:db8::1/128", "2001:db8::1", "2001:db8::1"},
	}
	for _, tc := range tests {
		t.Run(tc.cidr, func(t *testing.T) {
			first, last := AddressRange(mustParseCIDR(t, tc.cidr))
			if first.String() != tc.wantFirst || last.String() != tc.wantLast {
				t.Fatalf("got %s - %s, want %s - %s", first, last, tc.wantFirst, tc.wantLast)
			}
		})
	}
}

func TestAddressRangeSingleHostReturnsDistinctSlices(t *testing.T) {
	first, last := AddressRange(mustParseCIDR(t, "10.0.0.7/32"))
	last[len(last)-1] = 99
	if first.String() != "10.0.0.7" {
		t.Fatalf("changing the last address should not change the first, got %s", first)
	}
}

func TestIncAndDec(t *testing.T) {
	tests := []struct {
		ip      string
		wantInc string
		wantDec string
	}{
		{"10.0.0.1", "10.0.0.2", "10.0.0.0"},
		{"10.0.0.255", "10.0.1.0", "10.0.0.254"},
		{"10.0.1.0", "10.0.1.1", "10.0.0.255"},
		{"10.255.255.255", "11.0.0.0", "10.255.255.254"},
		{"255.255.255.255", "0.0.0.0", "255.255.255.254"},
		{"0.0.0.0", "0.0.0.1", "255.255.255.255"},
		{"2001:db8::ffff", "2001:db8::1:0", "2001:db8::fffe"},
	}
	for _, tc := range tests {
		t.Run(tc.ip, func(t *testing.T) {
			ip := net.ParseIP(tc.ip)
			if got := Inc(ip).String(); got != tc.wantInc {
				t.Fatalf("Inc: got %s, want %s", got, tc.wantInc)
			}
			if got := Dec(ip).String(); got != tc.wantDec {
				t.Fatalf("Dec: got %s, want %s", got, tc.wantDec)
			}
			if ip.String() != tc.ip {
				t.Fatalf("input should not be modified, got %s", ip)
			}
		})
	}
}

func TestGetLargestCIDR(t *testing.T) {
	tests := []struct {
		start, end string
		want       string
	}{
		{"10.0.0.0", "10.0.0.255", "10.0.0.0/24"},
		{"10.0.0.5", "10.0.0.5", "10.0.0.5/32"},
		{"10.0.0.0", "10.0.0.1", "10.0.0.0/31"},
		{"10.0.1.0", "10.0.2.255", "10.0.0.0/22"},
		{"192.168.0.10", "192.168.0.20", "192.168.0.0/27"},
		{"10.0.0.0", "11.255.255.255", "10.0.0.0/7"},
		{"0.0.0.0", "255.255.255.255", "0.0.0.0/0"},
		{"2001:db8::", "2001:db8::ffff", "2001:db8::/112"},
	}
	for _, tc := range tests {
		t.Run(tc.start+"-"+tc.end, func(t *testing.T) {
			got, err := GetLargestCIDR(tc.start, tc.end)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.String() != tc.want {
				t.Fatalf("got %s, want %s", got, tc.want)
			}
		})
	}
}

func TestGetLargestCIDRContainsBothEnds(t *testing.T) {
	pairs := [][2]string{
		{"10.20.30.40", "10.20.99.1"},
		{"172.16.0.1", "172.31.255.254"},
		{"192.168.1.100", "192.168.1.101"},
	}
	for _, pair := range pairs {
		prefix, err := GetLargestCIDR(pair[0], pair[1])
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		network := mustParseCIDR(t, prefix.String())
		for _, ip := range pair {
			if !network.Contains(net.ParseIP(ip)) {
				t.Fatalf("%s should contain %s", prefix, ip)
			}
		}
	}
}

func TestGetLargestCIDRErrors(t *testing.T) {
	tests := []struct {
		name, start, end string
		wantErr          string
	}{
		{"invalid start", "nope", "10.0.0.1", "invalid start IP"},
		{"invalid end", "10.0.0.1", "nope", "invalid end IP"},
		{"mixed families", "10.0.0.1", "2001:db8::1", "mixed IP families"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := GetLargestCIDR(tc.start, tc.end)
			if err == nil {
				t.Fatal("expected an error")
			}
			if !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("expected error containing %q, got %q", tc.wantErr, err.Error())
			}
		})
	}
}
