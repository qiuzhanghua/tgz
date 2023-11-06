package main

import "github.com/spf13/cobra"

func main() {
	var rootCmd = &cobra.Command{
		Use:   "tgz",
		Short: "TGZ",
		Long:  "Use tar and gzip to compress/extract files",
	}
	rootCmd.AddCommand(TgzXCmd, TgzCCmd)
	cobra.CheckErr(rootCmd.Execute())

}
