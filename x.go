package main

import (
	"fmt"
	"github.com/spf13/cobra"
)
import "github.com/qiuzhanghua/common/tgz"

var TgzXCmd = &cobra.Command{
	Use:   "x",
	Short: "Extract tar.gz",
	Long:  "Extract files from a tar.gz",
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil || len(args) < 1 {
			fmt.Println("Usage:\n     tgz x <tar.gz> [outDir]")
			return
		}
		tgzName := args[0]
		var outDir string
		if len(args) < 2 {
			outDir = "."
		} else {
			outDir = args[1]
		}
		tgz.Extract(tgzName, outDir)
	},
}
