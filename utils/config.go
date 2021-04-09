package utils

import (
	"io/ioutil"
	"log"

	"github.com/spf13/viper"
)

func ValidateConfigFile() bool {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.Name() == "config.yaml" {
			if ReadUsername() != "" && ReadCommentStyle() != "" && ReadFileExtension() != "" {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func ReadUsername() string {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	username := viper.GetString("username")
	return username
}

func ReadCommentStyle() string {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	username := viper.GetString("comment-style")
	return username
}

func ReadFileExtension() string {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.GetString("file-extension")
}

func ReadPlaceholder() string {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.GetString("placeholder")
}
