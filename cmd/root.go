/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/AndersonOdilo/fullcycle-stress-test/internal/usecase"
	"github.com/spf13/cobra"
)

var url string
var requests int
var concurrency int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fullcycle-stress-test",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {

		usecase.Executa(url, requests, concurrency)
  },
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
	rootCmd.PersistentFlags().StringVar(&url, "url", "", "URL do serviço a ser testado.")
	rootCmd.PersistentFlags().IntVar(&requests, "requests",  0, "Número total de requests.")
	rootCmd.PersistentFlags().IntVar(&concurrency, "concurrency", 0, "Número de chamadas simultâneas.")

	rootCmd.MarkPersistentFlagRequired("url") 
	rootCmd.MarkPersistentFlagRequired("requests") 
	rootCmd.MarkPersistentFlagRequired("concurrency") 

}


