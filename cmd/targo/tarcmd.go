package targo

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/stescobedo92/compressTo/cmp"
)

var tarCmd = &cobra.Command{
	Use:     "tar",
	Aliases: []string{"--tar-file"},
	Short:   "Create a tar file",
	Run:     createTarfile,
}

func createTarfile(cmd *cobra.Command, args []string) {

	if len(args) > 0 {
		// Files which to include in the tar.gz archive
		filesToIncludeInTar := args[1:]

		// Create output file
		out, err := os.Create(args[0] + ".tar.gz")
		if err != nil {
			log.Fatalln("Error writing archive:", err)
		}
		defer out.Close()

		// Create the archive and write the output to the "out" Writer
		err = cmp.CreateArchive(filesToIncludeInTar, out)
		if err != nil {
			log.Fatalln("Error creating archive:", err)
		}

		fmt.Println("Archive created successfully")
	} else {
		fmt.Fprintln(os.Stderr, "No files found. Please specify a valid files names.")
		return
	}
}

func init() {
	targoRootCmd.AddCommand(tarCmd)
}
