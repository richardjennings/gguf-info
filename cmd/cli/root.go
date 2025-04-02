package cli

import (
	"fmt"
	"github.com/richardjennings/gguf_info"
	"github.com/spf13/cobra"
	"os"
)

var metadataKeyFlag string

var rootCmd = &cobra.Command{
	Use:   "gguf-info <file.gguf>",
	Short: "print out gguf info",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		g, err := gguf_info.New(os.Args[1])
		if err != nil {
			return err
		}
		if metadataKeyFlag == "" {
			return g.Out(os.Stdout)
		}
		return g.MetadataValue(metadataKeyFlag, os.Stdout)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&metadataKeyFlag, "metadata-key", "m", "", "metadata key")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
