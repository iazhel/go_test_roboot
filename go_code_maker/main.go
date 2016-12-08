package main

import (
	"access_api"
	"fmt"
	"io"
	"os"
	"path"
	"templates"
)

var (
	baseUrl string = "http://epd-dev.devicepros.net:8010/"
	login   string = "i.azhel"
	pass    string = "12345"
/*
	baseUrl string = "http://stellus.devicepros.net:8005/"
	login   string = "admin"
	pass    string = "password"
*/

	output  io.Writer
	rest    *access_api.RestApiClient
)

func configureOutput() string {
	if len(os.Args) < 2 {
		fmt.Println(" enter package name")
		os.Exit(1)
	}
	newPackage := os.Args[1]
	fileName := newPackage + "/" + "model_" + newPackage + ".go"

	// create dir
	dir, _ := path.Split(fileName)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		os.Exit(1)
	}
	f, _ := os.Create(fileName)
	output = f
	fmt.Println(f)
	return newPackage
}

func main() {

	newPackage := configureOutput()
	rest = connect()

	// set print url requests
	rest.PrintUrl(true)
	variables := procMainPage(newPackage)
	variables = MakeEnv(variables)

	templates.Do(variables, templates. ApiParameters{
		BaseUrl: baseUrl,
		Login: login,
		Password : pass,
		PackageName: newPackage,
	})

}
