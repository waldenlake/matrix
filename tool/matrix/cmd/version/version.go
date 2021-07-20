package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Version = "v0.1.0"
)
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "matrix version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("matrix Version: ", Version)
	},
}
