package azure

import (
	"encoding/json"
	"fmt"

	"github.com/jercle/cloudini/lib"
	"github.com/spf13/cobra"
)

var (
	configuredVMs              []string
	startVMs                   bool
	stopVMs                    bool
	deallocateVMs              bool
	statusVMs                  bool
	listConfiguredVMs          bool
	runAgainstAllConfiguredVMs bool
)

// configCmd represents the subs command
var cmdVirtualMachines = &cobra.Command{
	Use:   "vm",
	Short: "Virtual Machine commands",
	Long: `Some useful VM commands using VMs configured in cloudini config file.

Location in config file json: azure.virtualMachines
Format for config: Key: Value

Note: You can only run this command for vms in one tenant at a time
`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			token  *lib.AzureMultiAuthToken
			vmList []string
		)

		config := lib.GetCldConfig(nil)

		tokOpts := lib.AzureMultiAuthTokenRequestOptions{
			TenantName: tenantName,
		}

		if startVMs || stopVMs || deallocateVMs {
			tokOpts.GetWriteToken = true
		}

		// lib.JsonMarshalAndPrint(tokOpts)
		// os.Exit(0)

		if !listConfiguredVMs {
			tokenReq, err := GetTenantSPToken(tokOpts, nil)
			lib.CheckFatalError(err)
			token = tokenReq

			if runAgainstAllConfiguredVMs {
				for _, vm := range config.Azure.VirtualMachines {
					vmList = append(vmList, vm)
				}
			} else {
				for _, vm := range configuredVMs {
					if config.Azure.VirtualMachines[vm] != "" {
						vmList = append(vmList, config.Azure.VirtualMachines[vm])
					}
				}
			}
		}

		// jsonStr, _ := json.MarshalIndent(vmList, "", "  ")
		// fmt.Println(string(jsonStr))
		// os.Exit(0)

		if listConfiguredVMs {
			jsonStr, _ := json.MarshalIndent(config.Azure.VirtualMachines, "", "  ")
			fmt.Println(string(jsonStr))
		}
		if startVMs {
			StartMultipleAzureVms(vmList, token)
		}
		if stopVMs {
			StopMultipleAzureVms(vmList, token)
		}
		if deallocateVMs {
			DeallocateMultipleAzureVms(vmList, token)
		}
		if statusVMs {
			GetMultipleAzureVMStatuses(vmList, token)
		}
	},
}

func init() {
	azCmd.AddCommand(cmdVirtualMachines)
	cmdVirtualMachines.Flags().StringSliceVarP(&configuredVMs, "configured-vms", "v", nil, "List of configured vms")
	cmdVirtualMachines.Flags().BoolVarP(&startVMs, "start", "s", false, "Allocates and/or starts the provided VMs")
	cmdVirtualMachines.Flags().BoolVarP(&stopVMs, "stop", "x", false, "Stops the provided VMs")
	cmdVirtualMachines.Flags().BoolVarP(&deallocateVMs, "deallocate", "d", false, "Deallocates the provided VMs")
	cmdVirtualMachines.Flags().BoolVarP(&statusVMs, "status-info", "i", false, "Gets the current status of the provided VMs")
	cmdVirtualMachines.Flags().BoolVarP(&listConfiguredVMs, "list-configured-vms", "l", false, "Lists currently configured VMs in the cloudini config file")
	cmdVirtualMachines.Flags().BoolVar(&runAgainstAllConfiguredVMs, "all", false, "Runs the command against all configured VMs")
	cmdVirtualMachines.Flags().StringVarP(&tenantName, "tenantName", "t", "", "Name of Tenant for configured auth")
	cmdVirtualMachines.MarkFlagRequired("tenantName")
	// cmdVirtualMachines.MarkFlagRequired("configured-vms")
	cmdVirtualMachines.MarkFlagsMutuallyExclusive("start", "stop", "deallocate", "status-info", "list-configured-vms")
	cmdVirtualMachines.MarkFlagsMutuallyExclusive("list-configured-vms", "all", "configured-vms")
	cmdVirtualMachines.MarkFlagsOneRequired("start", "stop", "deallocate", "status-info", "list-configured-vms")

	// resourcesCmd.Flags().StringVar(&getRoleDefByName, "getRoleDefByName", "", "Get  Role Definition details by Role Def name")
	// resourcesCmd.Flags().StringVar(&getRoleDefById, "getRoleDefById", "", "Get  Role Definition details by Role Def ID")
	// resourcesCmd.Flags().StringVarP(&tenantName, "tenantName", "n", "", "Tenant name to use configured auth. Defaults to Tenant of current active Az CLI subscription")
	// resourcesCmd.Flags().BoolVar(&onlyShowId, "onlyId", false, "Flag to only print the Role Def ID for 'getRoleDefByName'")
	// resourcesCmd.Flags().BoolVar(&onlyShowName, "onlyName", false, "Flag to only print the Role Def Name for 'getRoleDefById'")
	// resourcesCmd.MarkFlagsMutuallyExclusive("onlyId", "getRoleDefById")
	// resourcesCmd.MarkFlagsMutuallyExclusive("onlyName", "getRoleDefByName")
}
