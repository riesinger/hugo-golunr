package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var mtx sync.Mutex
var wg sync.WaitGroup
var posts []Post

// baseURL should be parsed from the config.toml file in the hugo repo
func main() {
	initConfig()

	filepath.Walk("./content", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error while walking content directory: ", err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), ".md") {
			wg.Add(1)
			go ParsePost(path)
		}
		return nil
	})
	wg.Wait()
	output, err := json.Marshal(posts)
	if err != nil {
		fmt.Println("Could not marshal posts to JSON: ", err)
		return
	}
	err = ioutil.WriteFile("static/search_index.json", output, 0644)
	if err != nil {
		fmt.Println("Could not write file: ", err)
		return
	}
	fmt.Println("Done!")
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetDefault("baseurl", "")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Could not read site config:", err)
		return
	}
}
