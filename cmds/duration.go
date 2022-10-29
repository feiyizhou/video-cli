package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"video-factory/services"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "get the duration of material",
	Long:  "get the duration of material",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("please notify the material path ...")
			return
		}
		seconds, err := services.NewCommonService().GetMaterialDuration(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(seconds)
	},
}
