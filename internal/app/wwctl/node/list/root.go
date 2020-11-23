package list

import "github.com/spf13/cobra"

var (
	baseCmd = &cobra.Command{
		Use:                "list",
		Short:              "List all nodes",
		Long:               "List all nodes",
		RunE:				CobraRunE,
	}
	ShowNet bool
	ShowIpmi bool
	ShowAll bool
)

func init() {
	baseCmd.PersistentFlags().BoolVarP(&ShowNet, "net", "n", false, "Show node network configurations")
	baseCmd.PersistentFlags().BoolVarP(&ShowIpmi, "ipmi", "i", false, "Show node IPMI configurations")
	baseCmd.PersistentFlags().BoolVarP(&ShowAll, "all", "a", false, "Show all node configurations")
}

// GetRootCommand returns the root cobra.Command for the application.
func GetCommand() *cobra.Command {
	return baseCmd
}