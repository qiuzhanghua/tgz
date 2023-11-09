package main

import (
	"fmt"
	"github.com/spf13/cobra"
)
import "github.com/qiuzhanghua/common/tgz"

var TgzTCmd = &cobra.Command{
	Use:   "t",
	Short: "List tar.gz",
	Long:  "List files in tar.gz",
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil || len(args) < 1 {
			fmt.Println("Usage:\n     tgz t <tar.gz>")
			return
		}
		tgzName := args[0]
		list, err := tgz.List(tgzName)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, s := range list {
			fmt.Println(s)
		}
		fmt.Printf("Total %d files\n", len(list))
	},
}
