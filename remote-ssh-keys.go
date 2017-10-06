package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/plopix/remote-ssh-keys/clients"
	"github.com/plopix/remote-ssh-keys/config"
	"github.com/plopix/remote-ssh-keys/logger"
)

func main() {
	logger.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	logger.Info.Println("Remote SSH Key Looker")
	if len(os.Args) != 5+1 {
		logger.Error.Fatal("Not enough arguments.")
	}
	userName := os.Args[1]
	userDirectory := os.Args[2]
	keyType := os.Args[3]
	keyFingerprint := os.Args[4]
	key := os.Args[5]
	logger.Trace.Printf("User: %s\n", userName)
	logger.Trace.Printf("Home: %s\n", userDirectory)
	logger.Trace.Printf("KeyType: %s\n", keyType)
	logger.Trace.Printf("KeyFingerprint: %s\n", keyFingerprint)
	logger.Trace.Printf("Key: %s\n", key)

	var configurationFile = "/etc/remote-ssh-keys.conf"
	var configuration = config.Load(configurationFile)

	var githubClient clients.Github
	var urlClient clients.URL

	if login, ok := configuration.Account[userName]; ok {
		for _, plainKey := range login.Plain {
			fmt.Println(plainKey)
		}
		for _, alias := range login.Github {
			var keys = githubClient.GetKeysForLogin(alias)
			for _, key := range keys {
				fmt.Println(key)
			}
		}
		for _, url := range login.URL {
			var keys = urlClient.GetKeysForLogin(url, userName)
			for _, key := range keys {
				fmt.Println(key)
			}
		}
	}
}
