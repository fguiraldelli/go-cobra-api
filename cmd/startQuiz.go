/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fasttrack_api/model"
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
			response, _ := http.Get("http://localhost:8080/user/" + email + "/email")

			if response.StatusCode == http.StatusOK {
				body, _ := io.ReadAll(response.Body)
				fmt.Println(string(body))

				var user model.Registred_user
				err := json.Unmarshal([]byte(body), &user)

				if err != nil {
					panic(err)
				}

				fmt.Printf("\n\n json object:::: %+v", string(body))
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(startQuizCmd)
}

//Show menu to create a new user, start the quiz with an existed user or exit the program.
func startMenu() {}

//Register a new user into the API.
func registerUser() {}

//Start the quiz with an existed user.
func startQuiz() {}

//Cleans the Screen
func clearScreen() {}
