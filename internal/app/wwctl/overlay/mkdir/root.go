package mkdir

import (
	"github.com/spf13/cobra"
)

var (
	baseCmd = &cobra.Command{
		Use:                "mkdir [overlay name] [directory path]",
		Short:              "Create a new directory",
		Long:               "Create a new directory within a Warewulf overlay",
		RunE:				CobraRunE,
		Args: 				cobra.MinimumNArgs(2),
	}
	SystemOverlay bool
//	PermMode os.FileMode
)

func init() {
	baseCmd.PersistentFlags().BoolVarP(&SystemOverlay, "system", "s", false, "Show System Overlays as well")
	// This will be added back as soon as I figure out how to properly handle the FileMode case
//	baseCmd.PersistentFlags().Uint32VarP(&PermMode, "mode", "m", 0755, "Permission mode for directory")
}

// GetRootCommand returns the root cobra.Command for the application.
func GetCommand() *cobra.Command {
	return baseCmd
}
