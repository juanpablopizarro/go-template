package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	viper.SetEnvPrefix("APP")
	viper.BindEnv("ID")

	id := viper.Get("ID")
	fmt.Println(id)
}
