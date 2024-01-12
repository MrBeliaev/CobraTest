package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/MrBeliaev/CobraTest/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Cobra CLI application",
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

var rateCmd = &cobra.Command{
	Use:                "rate",
	Short:              "Get rates",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			arg := strings.Split(args[0], "=")
			pair := arg[1]
			dataStr, err := api.GetBinancePrice(&pair)
			if err != nil {
				fmt.Println(err.Error())
			}
			result := make(map[string]float64)
			err = json.Unmarshal([]byte(dataStr), &result)
			fmt.Println(result[pair])
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(rateCmd)
	rateCmd.PersistentFlags().StringP("pair", "p", "", "get rate")
}

func startServer() {
	http.HandleFunc("/api/v1/rates", api.GetPrice)

	fmt.Println("Listening at port 3001...")

	err := http.ListenAndServe(":3001", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
