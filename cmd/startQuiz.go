/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// startQuizCmd represents the startQuiz command
var startQuizCmd = &cobra.Command{
	Use:   "startQuiz",
	Short: "A simple quiz with a few questions",
	Long:  `A  simple quiz with a few questions and a few alternatives for each question`,
	Run: func(cmd *cobra.Command, args []string) {

		var name string
		var email string
		var input int

		fmt.Println("Type your full name and press enter:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		name = scanner.Text()

		fmt.Println("Type your e-mail and press enter:")
		fmt.Scanf("%s", &email)

		fmt.Println("Type 1 for register a new user or Type 2 for login using your e-mail")
		fmt.Scanf("%d", &input)

		switch input {
		case 1:
			
			values := map[string]string{"name": name, "email": email}
			json_data, err := json.Marshal(values)

			if err != nil {
				log.Fatal(err)
			}

			response, _ := http.Post("http://localhost:8080/user/", "application/json", bytes.NewBuffer(json_data))

			if response.StatusCode == http.StatusCreated {
				body, _ := io.ReadAll(response.Body)
				fmt.Println(string(body))
			}

		case 2:
			fmt.Println("Caso 2")
		}

	},
}

func init() {
	rootCmd.AddCommand(startQuizCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startQuizCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startQuizCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
