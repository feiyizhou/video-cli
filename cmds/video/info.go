package video

import (
	"fmt"
	"github.com/spf13/cobra"
	"video-factory/services"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "get the information of video",
	Long:  "get the information of video",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("please notify the video path ...")
			return
		}
		seconds, err := services.NewVideoService().GetVideoDuration(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(seconds)
	},
}
