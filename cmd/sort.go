/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// sortCmd represents the sort command
var sortCmd = &cobra.Command{
	Use:   "sort",
	Short: "Sorts the files in the directory according to file type/extension",
	Long:  `This sorts the files in a directory according to their type or file extension. The files are not renamed. They will be organised into folders named by their type, e.g 'videos' for files with 'mp4' extensions and 'images' for files with 'png' and 'jpg' extensions.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("-------------sort called---------------")
		sourceDir, _ := cmd.Flags().GetString("srcDir")
		targetDir, _ := cmd.Flags().GetString("tarDir")
		sortFunction(sourceDir, targetDir)
	},
}

func init() {
	rootCmd.AddCommand(sortCmd)

	// Here you will define your flags and configuration settings.
	sortCmd.Flags().StringP("srcDir", "s", ".", "The directory where the files to be sorted are located.")
	sortCmd.Flags().StringP("tarDir", "t", ".", "The directory where the sub-directory containing the sorted files should be placed.")

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
	startTime := time.Now()
	fmt.Println("----------------Starting sort Command------------------")
	// fmt.Printf("The source directory is: %s \n", srcDir)
	// fmt.Printf("The target directory is: %s \n", tarDir)

	// ensure existence of source directory + gather all files that are inside srcDir
	dirEntries, err := os.ReadDir(srcDir)
	if err != nil {
		log.Fatal("The source directory does not exist!")
	}

	// the files in the directory
	fileEntries := make([]fs.DirEntry, 0)

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fileEntries = append(fileEntries, entry)
		}
	}

	for _, entry := range fileEntries {
		// fmt.Printf("entry number %d is %s \n", index, entry.Name())
		ext := filepath.Ext(entry.Name())
		// ext = strings.Trim(ext, ".")

		// determine file type
		fileType := mime.TypeByExtension(ext)

		// if fileType is empty: add file to 'other' subdirectory
		if fileType != "" {
			// media type is included in 'mime' package
			mediaType, fileExtension := strings.Split(fileType, "/")[0], strings.Split(fileType, "/")[1]

			// check whether sub-directory for file 'media type' exists
			if _, err := os.ReadDir(filepath.Join(tarDir, mediaType)); err != nil {
				// create it if it doesn't
				os.Mkdir(filepath.Join(tarDir, mediaType), os.ModeDevice)
				// fmt.Printf("other folder created for unknown file type.\n")

				// create sub-directory for file extension
				os.Mkdir(filepath.Join(tarDir, mediaType, fileExtension), os.ModeDevice)

				// then copy file into the folder
				fc, _ := os.ReadFile(entry.Name())
				if err := os.WriteFile(filepath.Join(tarDir, mediaType, fileExtension, entry.Name()), fc, 0600); err != nil {
					log.Fatal("Unable to write file" + entry.Name())
				}
				fmt.Printf("%s added to %s folder \n", entry.Name(), mediaType)
			} else {
				// check whether sub-directory for 'file extension' exists
				if _, err := os.ReadDir(filepath.Join(tarDir, mediaType, fileExtension)); err != nil {
					// create it if it does not
					os.Mkdir(filepath.Join(tarDir, mediaType, fileExtension), os.ModeDevice)

					// then copy file into the folder
					fc, _ := os.ReadFile(entry.Name())
					if err := os.WriteFile(filepath.Join(tarDir, mediaType, fileExtension, entry.Name()), fc, 0600); err != nil {
						log.Fatal("Unable to write file" + entry.Name())
					}
					fmt.Printf("%s added to %s folder \n", entry.Name(), mediaType)
				} else {

					// just copy file into existing folder
					fc, _ := os.ReadFile(entry.Name())
					if err := os.WriteFile(filepath.Join(tarDir, mediaType, fileExtension, entry.Name()), fc, 0600); err != nil {
						log.Fatal("Unable to write file" + entry.Name())
					}
					fmt.Printf("%s added to %s folder \n", entry.Name(), mediaType)
				}
			}
		} else {
			// For media types not found in the 'mime' package

			// check if 'other' sub-directory already exists
			if _, err := os.ReadDir(filepath.Join(tarDir, "other")); err != nil {
				// create 'other' sub-directory
				os.Mkdir(filepath.Join(tarDir, "other"), os.ModeDevice)
				fmt.Printf("other folder created for %s files.\n", strings.Trim(ext, "."))

				// then copy file into sub-directory
				fc, _ := os.ReadFile(entry.Name())
				if err := os.WriteFile(filepath.Join(tarDir, "other", entry.Name()), fc, 0600); err != nil {
					log.Fatal("Unable to write file: " + entry.Name())
				} else {
					fmt.Printf("%s added to 'other' folder \n", entry.Name())
				}
			} else {
				// just copy file into existing 'other' sub-directory
				fc, _ := os.ReadFile(entry.Name())
				if err := os.WriteFile(filepath.Join(tarDir, "other", entry.Name()), fc, 0600); err != nil {
					log.Fatal("Unable to write file" + entry.Name())
				} else {
					fmt.Printf("%s added to 'other' folder \n", entry.Name())
				}
			}
		}
	}

	// concurrently?: copy each file into its corresponding extension folder
	endTime := time.Now()
	fmt.Println("-----------------Sort Command ended---------------")
	fmt.Printf("Done in: %s", endTime.Sub(startTime))
}
