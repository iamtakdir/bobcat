package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "bobcat",
	Short: "bobcat is a new type of bobcat which doesn't bite",
	Long:  "A cli tool ",
	Run: func(cmd *cobra.Command, fileName []string) {
		if len(fileName) != 0 {
			content, err := os.ReadFile(fileName[0])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("\n", string(content))
		} else {
			fmt.Println("File Name can't be empty, bobcat<space>File Name ")
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Error:", err)
	}
}
