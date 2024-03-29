package cmds

import (
	"github.com/spf13/cobra"
	"video-factory/cmds/audio"
	"video-factory/cmds/video"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(
		GetCmd,
		ConcatCmd,
		CutCmd,
		video.RemoveAudioCmd,
		video.AddAudioCmd,
		video.CompressCmd,
		audio.CompressCmd,
	)
	return rootCmd
}
