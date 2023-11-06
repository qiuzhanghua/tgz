package main

import (
	"fmt"
	"github.com/spf13/cobra"
)
import "github.com/qiuzhanghua/common/tgz"

var TgzCCmd = &cobra.Command{
	Use:   "c",
	Short: "Compress tar.gz",
	Long:  "Compress files into a tar.gz",
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil || len(args) < 1 {
			fmt.Println("Usage:\n     tgz c <tar.gz> [dir|file ...]")
			return
		}
		tgzName := args[0]
		tgz.Compress(tgzName, args[1:]...)
	},
}
