package commands

import (
	"fmt"
	"os"
	"skeleton-services/adapters"

	"github.com/spf13/cobra"
)

// Define the makefile command
var makeHandlerCmd = &cobra.Command{
	Use:   "make:file [filename]",
	Short: "Generate a new Handler file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Nama file yang akan dibuat
		fileName := args[0] + ".go"
		content := `package main
				
		func main() {
			fmt.Println("Hello from ` + args[0] + `!")
		}`

		// Menulis file ke disk
		err := os.WriteFile(fileName, []byte(content), 0644)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

		fmt.Println("File", fileName, "has been created successfully.")
	},
}

func init() {
	// Tambahkan makeHandlerCmd ke root command
	adapters.RootCmd.AddCommand(makeHandlerCmd)
	// Jalankan root command
	adapters.Execute()
}
