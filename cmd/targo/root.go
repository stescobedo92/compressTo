package targo

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var targoRootCmd = &cobra.Command{
	Use:     "targo",
	Version: "1.0.0",
	Short:   "targo - a simple CLI to create targ.gz files",
	Long: `targo is a super fancy CLI (kidding)
   
One can use targo to modify or inspect strings straight from the terminal`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := targoRootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
