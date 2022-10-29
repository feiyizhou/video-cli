package video

import (
	"fmt"
	"github.com/spf13/cobra"
	"video-factory/services"
)

var CompressCmd = &cobra.Command{
	Use:   "compressV",
	Short: "compress large video to smaller video...",
	Long:  "compress large video to smaller video...",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("please notify the name of video...")
			return
		}
		_, err := services.NewVideoService().Compress(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}
