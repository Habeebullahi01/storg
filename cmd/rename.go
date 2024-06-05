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
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename files in a directory.",
	Long:  `The 'rename' command is used to rename the files in a directory in lexicographical order using numbers and an optional prefix.`,
	Run: func(cmd *cobra.Command, args []string) {
		sourceDir, _ := cmd.Flags().GetString("srcDir")
		targetDir, _ := cmd.Flags().GetString("tarDir")
		prefix, _ := cmd.Flags().GetString("prefix")
		rename(sourceDir, targetDir, prefix)
		// fmt.Println("rename called")
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)

	// Here you will define your flags and configuration settings.
	renameCmd.Flags().StringP("srcDir", "s", ".", "The directory containing the files to be renamed.")
	renameCmd.Flags().StringP("tarDir", "t", ".", "The directory to store the sub-directory containing the renamed files.")
	renameCmd.Flags().StringP("prefix", "p", "", "An optional prefix to be added to the renamed files.")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func rename(srcDir, tarDir, prefix string) {
	fmt.Println("Renaming...")
	if srcDir == "." {
		srcDir, _ = os.Getwd()
	}
	// Retrieve contents of source directory
	directoryContent, readErr := os.ReadDir(srcDir)
	if readErr != nil {
		// log error if error was encountered while reading source directory
		log.Fatal(readErr)
	}
	sourceFiles := 0
	for _, entry := range directoryContent {
		if !entry.IsDir() {
			sourceFiles += 1
		}
	}

	if tarDir == "." || srcDir == tarDir { // -> renamed files should be in the source directory, create a sub-directory for them.

		tarDir = filepath.Join(srcDir, "renamed files")
		if _, err := os.ReadDir(tarDir); err != nil {
			// Unable to read the sub-directory for reanmed files, assume it does not exist
			fmt.Println("Creating new directory for renamed files...")
			if err := os.Mkdir(tarDir, fs.ModeDevice); err != nil { // -> create new directory inside the source directory
				// Unable to create directory
				log.Fatal("Error creating sub-directory for renamed files.")
			}
			fmt.Println("New directory created")
		}
	} else if _, err := os.ReadDir(tarDir); err != nil {
		// specified target directory does not exist or some other reason
		log.Fatal(err.Error())
	}

	fileNumber := 1
	renamedFiles := 0

	prefix = strings.TrimSpace(prefix)

	filepath.WalkDir(srcDir, func(path string, d fs.DirEntry, err error) error {

		// File has to meet 2 conditions:
		// 1. Not be a directory
		// 2. Reside as a direct child of source directory

		if !d.IsDir() && filepath.Dir(path) == srcDir {
			// construct the new file name
			newFileName := filepath.Join(tarDir, strings.TrimSpace(strings.Join([]string{prefix, strconv.Itoa(fileNumber)}, " ")+filepath.Ext(path)))

			changeFileNumber := func() {
				fileNumber += 1
				newFileName = filepath.Join(tarDir, strings.TrimSpace(strings.Join([]string{prefix, strconv.Itoa(fileNumber)}, " "))+filepath.Ext(path))
			}
			// ensure name is unique
			unique := testUnique(newFileName)
			for !unique {
				changeFileNumber()
				unique = testUnique(newFileName)
			}

			// copy file?
			if unique {
				if fileContent, err := os.ReadFile(path); err == nil { // -> read file successful
					if err := os.WriteFile(newFileName, fileContent, os.ModeAppend); err != nil { // -> write file unsuccessful
						fmt.Println(err.Error())
						log.Fatal("Error copying file.")
					} else {
						renamedFiles += 1
						fmt.Printf("Renamed file %d of %d \n", renamedFiles, sourceFiles)
					}
				} else {
					log.Fatalf("Unable to read file: %s", d.Name())
				}
			}
		}

		return nil
	})
	fmt.Printf("Storg successfully renamed %d files. \n", renamedFiles)

}
func testUnique(fname string) bool {
	if _, err := os.ReadFile(fname); err == nil {
		// file exists
		return false
	} else {
		return true
	}
}
