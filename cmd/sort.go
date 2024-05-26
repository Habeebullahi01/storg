/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// sortCmd represents the sort command
var sortCmd = &cobra.Command{
	Use:   "sort",
	Short: "Sorts the files in the directory according to file type/extension",
	Long:  `This sorts the files in a directory according to their type or file extension. The files are not renamed. They will be organised into folders named by their type, e.g 'videos' for files with 'mp4' extensions and 'images' for files with 'png' and 'jpg' extensions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("-------------sort called---------------")
		sourceDir, _ := cmd.Flags().GetString("srcDir")
		targetDir, _ := cmd.Flags().GetString("tarDir")
		sortFunction(sourceDir, targetDir)
	},
}

func init() {
	rootCmd.AddCommand(sortCmd)

	// Here you will define your flags and configuration settings.
	sortCmd.Flags().String("srcDir", ".", "The directory where the files to be sorted are located.")
	sortCmd.Flags().String("tarDir", ".", "The directory where the sub-directory containing the sorted files should be placed.")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sortCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sortCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Sorts the files in a folder
func sortFunction(srcDir, tarDir string) {

	if srcDir == "." {
		srcDir, _ = os.Getwd()
	}
	if tarDir == "." {
		tarDir = srcDir
	}
	fmt.Println("----------------Starting sort function------------------")
	// fmt.Printf("The source directory is: %s \n", srcDir)
	// fmt.Printf("The target directory is: %s \n", tarDir)

	// ensure existence of source directory + gather all files that are inside srcDir
	dirEntries, err := os.ReadDir(srcDir)
	if err != nil {
		log.Fatal("The source directory does not exist!")
	}
	// else {
	// 	fmt.Println("Source exists!")
	// }

	// the files in the directory
	fileEntries := make([]fs.DirEntry, 0)

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fileEntries = append(fileEntries, entry)
		}
	}

	// create a folder for every extension if one doesn't already exist
	for _, entry := range fileEntries {
		// fmt.Printf("entry number %d is %s \n", index, entry.Name())
		ext := filepath.Ext(entry.Name())
		ext = strings.Trim(ext, ".")
		// create folder within targetDir if it doesn't already exist
		if _, err := os.ReadDir(tarDir + "/" + ext); err != nil {
			os.Mkdir(tarDir+"/"+ext, os.ModeDevice)
			fmt.Printf("%s folder created for %s files.\n", ext, ext)
			// then copy file into the folder
			fc, _ := os.ReadFile(entry.Name())
			if err := os.WriteFile(tarDir+"/"+ext+"/"+entry.Name(), fc, 0600); err != nil {
				log.Fatal("Unable to write file" + entry.Name())
			}
			fmt.Printf("%s added to %s folder", entry.Name(), ext)
		} else {
			// just copy file into existing folder
			fc, _ := os.ReadFile(entry.Name())
			if err := os.WriteFile(tarDir+"/"+ext+"/"+entry.Name(), fc, 0600); err != nil {
				log.Fatal("Unable to write file" + entry.Name())
			}
			fmt.Printf("%s added to %s folder", entry.Name(), ext)
		}
	}

	// concurrently?: copy each file into its corresponding extension folder
	fmt.Println("-----------------Sort function ended---------------")
}
