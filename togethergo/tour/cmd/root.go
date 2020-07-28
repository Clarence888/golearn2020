package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "",
	Short: "",
	Long: "",
}

func Execute() error  {
	return rootCmd.Execute()
}

func init()  {
	//每个子命令都需要到rootCmd注册
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(sqlCmd)
}


