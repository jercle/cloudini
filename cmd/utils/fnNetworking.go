package utils

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"math/big"
	"net"
	"strconv"
)

func GenSubnetsInNetwork(netCIDR string, subnetMaskSize int) ([]string, error) {
	ip, ipNet, err := net.ParseCIDR(netCIDR)
	if err != nil {
		return nil, err
	}
	if !ip.Equal(ipNet.IP) {
		return nil, errors.New("netCIDR is not a valid network address")
	}
	netMaskSize, _ := ipNet.Mask.Size()
	if netMaskSize > int(subnetMaskSize) {
		return nil, errors.New("subnetMaskSize must be greater or equal than netMaskSize")
	}

	totalSubnetsInNetwork := math.Pow(2, float64(subnetMaskSize)-float64(netMaskSize))
	totalHostsInSubnet := math.Pow(2, 32-float64(subnetMaskSize))
	subnetIntAddresses := make([]uint32, int(totalSubnetsInNetwork))
	// first subnet address is same as the network address
	subnetIntAddresses[0] = ip2int(ip.To4())
	for i := 1; i < int(totalSubnetsInNetwork); i++ {
		subnetIntAddresses[i] = subnetIntAddresses[i-1] + uint32(totalHostsInSubnet)
	}

	subnetCIDRs := make([]string, 0)
	for _, sia := range subnetIntAddresses {
		subnetCIDRs = append(
			subnetCIDRs,
			int2ip(sia).String()+"/"+strconv.Itoa(int(subnetMaskSize)),
		)
	}
	return subnetCIDRs, nil
}

func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		panic("cannot convert IPv6 into uint32")
	}
	return binary.BigEndian.Uint32(ip)
}
func int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

// https://github.com/apparentlymart/go-cidr/tree/master
// PreviousSubnet returns the subnet of the desired mask in the IP space
// just lower than the start of IPNet provided. If the IP space rolls over
// then the second return value is true
func PreviousSubnet(network *net.IPNet, prefixLen int) (*net.IPNet, bool) {
	startIP := checkIPv4(network.IP)
	previousIP := make(net.IP, len(startIP))
	copy(previousIP, startIP)
	cMask := net.CIDRMask(prefixLen, 8*len(previousIP))
	previousIP = Dec(previousIP)
	previous := &net.IPNet{IP: previousIP.Mask(cMask), Mask: cMask}
	if startIP.Equal(net.IPv4zero) || startIP.Equal(net.IPv6zero) {
		return previous, true
	}
	return previous, false
}

// NextSubnet returns the next available subnet of the desired mask size
// starting for the maximum IP of the offset subnet
// If the IP exceeds the maxium IP then the second return value is true
func NextSubnet(network *net.IPNet, prefixLen int) (*net.IPNet, bool) {
	_, currentLast := AddressRange(network)
	mask := net.CIDRMask(prefixLen, 8*len(currentLast))
	currentSubnet := &net.IPNet{IP: currentLast.Mask(mask), Mask: mask}
	_, last := AddressRange(currentSubnet)
	last = Inc(last)
	next := &net.IPNet{IP: last.Mask(mask), Mask: mask}
	if last.Equal(net.IPv4zero) || last.Equal(net.IPv6zero) {
		return next, true
	}
	return next, false
}

// AddressRange returns the first and last addresses in the given CIDR range.
func AddressRange(network *net.IPNet) (net.IP, net.IP) {
	// the first IP is easy
	firstIP := network.IP

	// the last IP is the network address OR NOT the mask address
	prefixLen, bits := network.Mask.Size()
	if prefixLen == bits {
		// Easy!
		// But make sure that our two slices are distinct, since they
		// would be in all other cases.
		lastIP := make([]byte, len(firstIP))
		copy(lastIP, firstIP)
		return firstIP, lastIP
	}

	firstIPInt, bits := ipToInt(firstIP)
	hostLen := uint(bits) - uint(prefixLen)
	lastIPInt := big.NewInt(1)
	lastIPInt.Lsh(lastIPInt, hostLen)
	lastIPInt.Sub(lastIPInt, big.NewInt(1))
	lastIPInt.Or(lastIPInt, firstIPInt)

	return firstIP, intToIP(lastIPInt, bits)
}

// Inc increases the IP by one this returns a new []byte for the IP
func Inc(IP net.IP) net.IP {
	IP = checkIPv4(IP)
	incIP := make([]byte, len(IP))
	copy(incIP, IP)
	for j := len(incIP) - 1; j >= 0; j-- {
		incIP[j]++
		if incIP[j] > 0 {
			break
		}
	}
	return incIP
}

// Dec decreases the IP by one this returns a new []byte for the IP
func Dec(IP net.IP) net.IP {
	IP = checkIPv4(IP)
	decIP := make([]byte, len(IP))
	copy(decIP, IP)
	decIP = checkIPv4(decIP)
	for j := len(decIP) - 1; j >= 0; j-- {
		decIP[j]--
		if decIP[j] < 255 {
			break
		}
	}
	return decIP
}

func checkIPv4(ip net.IP) net.IP {
	// Go for some reason allocs IPv6len for IPv4 so we have to correct it
	if v4 := ip.To4(); v4 != nil {
		return v4
	}
	return ip
}

//
//

func intToIP(ipInt *big.Int, bits int) net.IP {
	ipBytes := ipInt.Bytes()
	ret := make([]byte, bits/8)
	// Pack our IP bytes into the end of the return array,
	// since big.Int.Bytes() removes front zero padding.
	for i := 1; i <= len(ipBytes); i++ {
		ret[len(ret)-i] = ipBytes[len(ipBytes)-i]
	}
	return net.IP(ret)
}

func ipToInt(ip net.IP) (*big.Int, int) {
	val := &big.Int{}
	val.SetBytes([]byte(ip))
	if len(ip) == net.IPv4len {
		return val, 32
	} else if len(ip) == net.IPv6len {
		return val, 128
	} else {
		panic(fmt.Errorf("Unsupported address length %d", len(ip)))
	}
}
