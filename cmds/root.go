package cmds

import (
	"github.com/spf13/cobra"
	"video-factory/cmds/video"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(
		video.GetCmd,
		video.ConcatCmd,
		video.CutCmd,
		video.RemoveAudioCmd,
		video.AddAudioCmd,
	)
	return rootCmd
}
