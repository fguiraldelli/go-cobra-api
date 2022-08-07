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
	"os/exec"

	"github.com/spf13/cobra"
)

var user model.Registred_user

// startQuizCmd represents the startQuiz command
var startQuizCmd = &cobra.Command{
	Use:   "startQuiz",
	Short: "A simple quiz with a few questions",
	Long:  `A  simple quiz with a few questions and a few alternatives for each question`,
	Run: func(cmd *cobra.Command, args []string) {

		clearScreen()
		startMenu()
	},
}

func init() {
	rootCmd.AddCommand(startQuizCmd)
}

//Show menu to create a new user, start the quiz with an existed user or exit the program.
func startMenu() {
	fmt.Println("Choose an option to continue:")
	fmt.Println("1: Register a new user")
	fmt.Println("2: Start quiz")
	fmt.Println("3: Exit quiz app")

	var userInput int
	fmt.Scanln(&userInput)
	switch userInput {
	case 1:
		registerUser()
	case 2:
		startQuiz()
	case 3:
		os.Exit(1)
	case 4:
		updateUserQuestions()
	default:
		fmt.Println(" Invalid option.")

	}
}

//Register a new user into the API.
func registerUser() {
	var name string
	var email string

	fmt.Println("Type your full name and press enter:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("Type your e-mail and press enter:")
	fmt.Scanf("%s", &email)

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

	clearScreen()
	startMenu()

}

//Start the quiz with an existed user.
func startQuiz() {
	var email string

	fmt.Println("Type your e-mail to start the quiz:")
	fmt.Scanf("%s", &email)
	getUserQuestions(email)
	clearScreen()
	for i, question := range user.Quiz {
		answer_sheet := make(map[int]string)
		if !question.Answered {
			fmt.Println(question.Question)
			for i, alternative := range question.Answers {
				fmt.Println(i+1, alternative)
				answer_sheet[i] = alternative
			}

			var choice int
			fmt.Println("Input a number between 1-5 to answer.")
			fmt.Scanln(&choice)
			fmt.Printf("Your choice was: %d\n\n", choice)
			if answer_sheet[choice-1] == question.Correct_answer {
				user.Quiz[i].Is_corrected = true
				user.Number_corrected_answers++
			}
			user.Quiz[i].Answered = true
		}

		clearScreen()
	}

	fmt.Println("")

	updateUserQuestions()

	fmt.Println("Number of corrected answers: ", user.Number_corrected_answers)
	fmt.Println("You scored higher than ", int(user.User_rated), "% of all quizzers")

	fmt.Println("Press enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	fmt.Println(user.Quiz)

	clearScreen()
	startMenu()
}

//Register a new user into the API.
func updateUserQuestions() {
	json_data, err := json.Marshal(user)

	if err != nil {
		log.Fatal(err)
	}

	response, _ := http.Post("http://localhost:8080/user/"+user.Email+"/questions", "application/json", bytes.NewBuffer(json_data))

	if response.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		err := json.Unmarshal([]byte(body), &user)
		if err != nil {
			panic(err)
		}
	}

	clearScreen()
}

//Cleans the Screen
func clearScreen() {
	cs := exec.Command("clear")
	cs.Stdout = os.Stdout
	cs.Run()
}

func getUserQuestions(email string) *model.Registred_user {
	response, _ := http.Get("http://localhost:8080/user/" + email + "/email")

	if response.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		fmt.Println(string(body))

		err := json.Unmarshal([]byte(body), &user)

		if err != nil {
			panic(err)
		}

		fmt.Printf("\n\n json object:::: %+v", string(body))

	}
	return &user
}
