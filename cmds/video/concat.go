package video

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"video-factory/services"
)

var ConcatCmd = &cobra.Command{
	Use:   "concat",
	Short: "concat short video to long video",
	Long:  "concat short video to long video",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("please notify source video and destination video duration...")
			return
		}
		name := args[0]
		duration, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = services.NewVideoService().ConcatVideo(name, duration)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}
