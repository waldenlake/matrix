package cmd

import (
	"fmt"
	"github.com/go-matrix/matrix/tool/matrix/cmd/project"
	"github.com/go-matrix/matrix/tool/matrix/cmd/tool"
	"github.com/go-matrix/matrix/tool/matrix/cmd/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "matrix",
	Short: "matrix go代码工具集",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("matrix go代码工具集...")
	},
}

func init() {
	rootCmd.AddCommand(tool.ToolCmd)
	rootCmd.AddCommand(project.CreateCmd)
	rootCmd.AddCommand(version.VersionCmd)
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
