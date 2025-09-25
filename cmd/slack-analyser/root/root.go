package root

import (
	"fmt"
	"log"

	"github.com/jswiss/slack-analyser/internal/cli"
	"github.com/spf13/cobra"
)

var (
	channelDir string
	prompt     string
)

var rootCmd = &cobra.Command{
	Use:   "slack-analyser",
	Short: "Analyse Slack exports with natural language",
	RunE: func(cmd *cobra.Command, args []string) error {
		if channelDir == "" || prompt == "" {
			return fmt.Errorf("missing required flags: --channel-dir and --prompt")
		}
		out, err := cli.Run(channelDir, prompt)
		if err != nil {
			return err
		}
		fmt.Println(out)
		return nil
	},
}

func init() {
	rootCmd.Flags().StringVar(&channelDir, "channel-dir", "", "path to Slack channel export directory")
	rootCmd.Flags().StringVar(&prompt, "prompt", "", "natural language request (e.g. 'average thread length')")
	rootCmd.MarkFlagRequired("channel-dir")
	rootCmd.MarkFlagRequired("prompt")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
