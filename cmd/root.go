package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"rest-api/routes"

	"rest-api/application/config"
)

var rootCmd = &cobra.Command{
	Use:   "GO-POSTGRESQL",
	Short: "Using golang in postgresql",
	Long:  `Using postgresql and some plugins`,

	Run: func(cmd *cobra.Command, args []string) {
		routes.Route()
	},
}

func init() {
	cobra.OnInitialize(splash)
	config.Init()
}
func splash() {
	fmt.Println(`
	running application            
  `)
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		log.Fatalf("Error running Execute")
	}
}
