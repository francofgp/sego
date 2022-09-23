/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var Verbose bool

type Response struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	BaseUrl     string `json:"base_url"`
	Description string `json:"description"`
}

func getRandom() {
	resp, err := http.Get("https://librex.beparanoid.de/api.php/api.php?q=gentoo&p=2&type=0")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}
func getRandomJoke() {
	url := "https://librex.beparanoid.de/api.php/api.php?q=gentoo&p=2&type=0"
	responseBytes := getJokeData(url)
	joke := Response{}

	var it Response
	err := json.Unmarshal([]byte(responseBytes), &it)
	if err != nil {
		panic(err)
	}
	fmt.Println(it)

	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
	}

	fmt.Println(string(joke.Title))
}

func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet, //method
		baseAPI,        //url
		nil,            //body
	)

	if err != nil {
		log.Printf("Could not request a dadjoke. %v", err)
	}

	request.Header.Add("Accept", "application/json")
	//request.Header.Add("User-Agent", "Dadjoke CLI (https://github.com/example/dadjoke)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return responseBytes
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
		fmt.Println(args)
		fmt.Println(len(args))
		fmt.Println("Echo: " + strings.Join(args, " "))
		fmt.Println(Verbose)
		getRandom()
	},
}

// https://librex.beparanoid.de/api.php/api.php?q=gentoo&p=2&type=0
func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
