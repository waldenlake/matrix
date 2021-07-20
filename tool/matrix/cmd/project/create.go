package project

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "matrix create a new project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create a new project...")
	},
}
