package video

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"video-factory/services"
)

var CutCmd = &cobra.Command{
	Use:   "cut",
	Short: "cut video with notify duration",
	Long:  "cut video with notify duration",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("please notify the video name and duration...")
			return
		}
		duration, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = services.NewVideoService().CutVideo(args[0], duration)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}
