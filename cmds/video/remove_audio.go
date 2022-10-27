package video

import (
	"fmt"
	"github.com/spf13/cobra"
	"video-factory/services"
)

var RemoveAudioCmd = &cobra.Command{
	Use:   "remove-audio",
	Short: "remove audio from video",
	Long:  "remove audio from video",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("please notify the name of video...")
			return
		}
		_, err := services.NewVideoService().RemoveAudio(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}
