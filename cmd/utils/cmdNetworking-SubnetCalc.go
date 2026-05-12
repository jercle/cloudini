package utils

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/jercle/cloudini/lib"
	"github.com/rmasci/ipsubnet"
	"github.com/spf13/cobra"
)

var (
	cidrRange         string
	getPossibleRanges string
	rangeSize         int
	nextSubnet        string
	prevSubnet        string
)

// utilsCmd represents the util command
var cmdSubnetCalc = &cobra.Command{
	Use:     "subnetCalc",
	Short:   "Subnet related functions",
	Aliases: []string{"snet"},
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	//
	Run: func(cmd *cobra.Command, args []string) {

		if getPossibleRanges != "" {

			ranges, _ := GenSubnetsInNetwork(getPossibleRanges, 28)
			for _, snet := range ranges {
				fmt.Println("*", snet)
			}
		}

		if cidrRange != "" {
			ipAddr, _, err := net.ParseCIDR(cidrRange)
			lib.CheckFatalError(err)
			cidrNotation, err := strconv.Atoi(strings.Split(cidrRange, "/")[1])
			lib.CheckFatalError(err)
			cidrIp := strings.Split(cidrRange, "/")[0]
			cidrIpParsed := net.ParseIP(cidrIp)
			if cidrIpParsed == nil {
				fmt.Println("Invalid IP Address")
				os.Exit(1)
			}

			snet := ipsubnet.SubnetCalculator(cidrIp, int64(cidrNotation))

			startAddress := snet.GetIPAddressRange()[0]
			endAddress := snet.GetIPAddressRange()[1]
			if cidrNotation < 32 {
				fmt.Println("First IP address:", startAddress)
				fmt.Println("Last IP address:", endAddress)
				fmt.Println("Broadcast address:", snet.GetBroadcastAddress())
				fmt.Println("Number of addresses:", snet.GetNumberIPAddresses())
				fmt.Println("Number of addressable hosts:", snet.GetNumberAddressableHosts())
				fmt.Println("Subnet Mask:", snet.GetSubnetMask())

				fmt.Println("")
				fmt.Println("GetHostPortion:", snet.GetHostPortion())
				fmt.Println("GetNetworkPortion:", snet.GetNetworkPortion())
				fmt.Println("GetNetworkSize:", snet.GetNetworkSize())
				fmt.Println("GetIPAddressRange:", snet.GetIPAddressRange())
				fmt.Println("GetIPAddress:", snet.GetIPAddress())

			} else {
				fmt.Println("Number of addresses:", 1)
				fmt.Println("Address:", ipAddr.String())
			}
		}

	},
}

func init() {
	cmdNetworking.AddCommand(cmdSubnetCalc)

	cmdSubnetCalc.Flags().StringVarP(&cidrRange, "cidrRange", "c", "", "CIDR range to get IP Address range and details from")
	cmdSubnetCalc.Flags().StringVarP(&getPossibleRanges, "getPossibleRanges", "r", "", "CIDR range. Get possible ranges within given CIDR range. Must be used with 'networkSize' flag.")
	cmdSubnetCalc.Flags().IntVarP(&rangeSize, "rangeSize", "s", 0, "Network size (CIDR notation) to use with getPossibleRanges.")

	cmdSubnetCalc.Flags().StringVarP(&nextSubnet, "nextSubnet", "n", "", "CIDR range. Gets the next available subnet of the desired mask size")
	cmdSubnetCalc.Flags().StringVarP(&prevSubnet, "prevSubnet", "p", "", "CIDR range. Gets the subnet of the desired mask in the IP space just lower than the start of IPNet provided.")

	cmdSubnetCalc.MarkFlagsMutuallyExclusive("cidrRange", "getPossibleRanges", "nextSubnet", "prevSubnet")
	// cmdSubnetCalc.MarkFlagsMutuallyExclusive("cidrRange", "rangeSize")
	cmdSubnetCalc.MarkFlagsRequiredTogether("getPossibleRanges", "rangeSize")

	// utilsCmd.Flags().IntVarP("pwdgen", )
	// StringVarP("pwdgen", "p", 0)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// utilsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// utilsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
