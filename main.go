package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pelletier/go-toml"
)

func main() {
	parseToml()
	parseJson()
}

func parseToml() {
	config, _ := toml.LoadFile("config.toml")

	configTree := config.Get("postgres").(*toml.TomlTree)
	user := configTree.Get("user").(string)
	pw := configTree.Get("password").(string)
	port := configTree.Get("port").(int64)

	fmt.Sprintf("%s %s %d", user, pw, port)
}

func parseJson() {
	file, _ := ioutil.ReadFile("config.json")

	var conf Config
	json.Unmarshal(file, &conf)

	fmt.Sprintf("%s %s %d", conf.Posgres.User, conf.Posgres.Password, conf.Posgres.Port)
}

type Config struct {
	Posgres posgres `json:"posgres"`
}

type posgres struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Port     int    `json:"port"`
}
