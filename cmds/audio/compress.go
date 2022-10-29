package audio

import (
	"fmt"
	"github.com/spf13/cobra"
	"video-factory/services"
)

var CompressCmd = &cobra.Command{
	Use:   "compressA",
	Short: "compress large audio to smaller audio...",
	Long:  "compress large audio to smaller audio...",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("please notify the audio's name and bitrate...")
			return
		}
		_, err := services.NewAudioService().Compress(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}
