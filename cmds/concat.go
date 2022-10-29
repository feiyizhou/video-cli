package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"video-factory/services"
)

var ConcatCmd = &cobra.Command{
	Use:   "concat",
	Short: "concat short material to long material",
	Long:  "concat short material to long material",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("please notify source material and destination material duration...")
			return
		}
		name := args[0]
		duration, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = services.NewCommonService().ConcatMaterial(name, duration)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}
