package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gitio",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func exitError(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func getShortenURL(arg string) (string, error) {
	v := url.Values{}
	v.Add("url", arg)
	resp, err := http.DefaultClient.PostForm("https://git.io/create", v)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if resp.Status[0] != '2' {
		return "", errors.New(string(body))
	}

	return "https://git.io/" + string(body), err
}

func Execute() {
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = rootCmd.Help()
			return
		}
		url, err := getShortenURL(args[0])
		if err != nil {
			exitError(err)
		}
		println(url)
	}

	if err := rootCmd.Execute(); err != nil {
		exitError(err)
	}
}
