package add

import "github.com/spf13/cobra"

var (
	baseCmd = &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "add [OPTIONS] NODENAME",
		Short:                 "Add new node to Warewulf",
		Long:                  "This command will add a new node named NODENAME to Warewulf.",
		RunE:                  CobraRunE,
		Args:                  cobra.MinimumNArgs(1),
	}
	SetClusterName  string
	SetNetName      string
	SetNetDev       string
	SetIpaddr       string
	SetIpaddr6      string
	SetNetmask      string
	SetGateway      string
	SetHwaddr       string
	SetType         string
	SetDiscoverable bool
)

func init() {
	baseCmd.PersistentFlags().StringVarP(&SetClusterName, "cluster", "c", "", "Set the node's cluster name")
	baseCmd.PersistentFlags().StringVarP(&SetNetName, "netname", "n", "default", "Define the network name to configure")
	baseCmd.PersistentFlags().StringVarP(&SetNetDev, "netdev", "N", "", "Define the network device to configure")
	baseCmd.PersistentFlags().StringVarP(&SetIpaddr, "ipaddr", "I", "", "Set the node's network device IP address")
	baseCmd.PersistentFlags().StringVarP(&SetIpaddr6, "ipaddr6", "6", "", "Set the node's network device IPv6 address")
	baseCmd.PersistentFlags().StringVarP(&SetNetmask, "netmask", "M", "", "Set the node's network device netmask")
	baseCmd.PersistentFlags().StringVarP(&SetGateway, "gateway", "G", "", "Set the node's network device gateway")
	baseCmd.PersistentFlags().StringVarP(&SetHwaddr, "hwaddr", "H", "", "Set the node's network device HW address")
	baseCmd.PersistentFlags().StringVarP(&SetType, "type", "T", "", "Set the node's network device type")
	baseCmd.PersistentFlags().BoolVar(&SetDiscoverable, "discoverable", false, "Make this node discoverable")
}

// GetRootCommand returns the root cobra.Command for the application.
func GetCommand() *cobra.Command {
	return baseCmd
}
