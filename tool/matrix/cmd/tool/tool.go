package tool

import (
	"fmt"
	"github.com/go-matrix/matrix/tool/matrix/cmd/tool/protoc"
	"github.com/spf13/cobra"
)

var ToolCmd = &cobra.Command{
	Use:   "tool",
	Short: "matrix 代码生成工具集",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("matrix 代码生成工具集")
	},
}

func init() {
	ToolCmd.AddCommand(protoc.ProtocCmd)
}
