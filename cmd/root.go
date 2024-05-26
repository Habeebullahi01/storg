/*
Copyright © 2024 Lawal Habeebullahi <Lawalhabeebullahi008@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "storg",
	Short: "This program organizes files in a directory by sorting them into subfolders based on their types and extensions. It can also rename files in the directory by numbering them sequentially and adding an optional prefix",
	Long: `This program helps users manage and organize files within a specified directory by performing two main tasks. Firstly, it scans the directory and sorts files into subfolders based on their types (such as images, documents, or audio files) and specific extensions (like .jpg, .pdf, or .mp3). Each subfolder is named according to the file type or extension, resulting in a well-organized directory structure.

	Secondly, the program can rename all files in the directory by sequentially numbering them, starting from one, and it allows users to add a custom prefix to each filename if desired. This renaming feature is useful for batch processing files, making them easier to reference and manage. The program is designed to be user-friendly, enabling efficient and customizable file organization and renaming without requiring extensive user input.s`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.storg.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
