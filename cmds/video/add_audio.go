package video

import (
	"fmt"
	"github.com/spf13/cobra"
	"video-factory/services"
)

var AddAudioCmd = &cobra.Command{
	Use:   "add-audio",
	Short: "add an audio to video",
	Long:  "add an audio to video",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("please notify the names of video and audio...")
			return
		}
		_, err := services.NewVideoService().AddAudio(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}
