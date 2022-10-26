package cmds

import "github.com/spf13/cobra"

var concatCmd = &cobra.Command{
	Use:   "concat",
	Short: "concat short video to long video",
	Long:  "concat short video to long video",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
