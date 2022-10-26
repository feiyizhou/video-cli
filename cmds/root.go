package cmds

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(getCmd, concatCmd)
	return rootCmd
}
