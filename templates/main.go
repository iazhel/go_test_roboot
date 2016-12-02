package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"
)

func main() {
	APIsBeginer := "beginner_APIs"
	InputMethodsSource := "template_methods"
	OuputMethodsSource := "methods/APIs.go"

	testsBeginer := "beginner_tests"
	InputTestsSource := "template_tests"
	OuputTestsSource := "methods/auto_test.go"

	for i, api := range env {

		// check field names
		if len(api.IdField) == 0 || len(api.PrimaryFieldName) == 0 {
			env[i].FieldsNameMissed = true
		}
		// check metods existence
		if len(api.GetMethod) == 0 {
			env[i].GetRejected = true
		}
		if len(api.GetOneMethod) == 0 {
			env[i].GetOneRejected = true
		}
		if len(api.PostMethod) == 0 {
			env[i].PostRejected = true
		}
		if len(api.PatchMethod) == 0 {
			env[i].PatchRejected = true
		}
		if len(api.DelMethod) == 0 {
			env[i].DelRejected = true
		}
		if len(api.PutMethod) == 0 {
			env[i].PutRejected = true
		}
		if len(api.HeadMethod) == 0 {
			env[i].HeadRejected = true
		}
		if len(api.OptMethod) == 0 {
			env[i].OptionRejected = true
		}
		// If types Get and GetOne are different
		if len(api.OneGetType) == 0 {
			env[i].OneGetType = env[i].ExampleType
		}

	}

	_ = process(APIsBeginer, InputMethodsSource, OuputMethodsSource, 2)
	_ = process(testsBeginer, InputTestsSource, OuputTestsSource, 0)
}

func process(Beginner, InputMethodsSource, Ouput string, n int) (err error) {
	// create methods dir
	dir, _ := path.Split(Ouput)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Printf("'%s' %s", dir, err)
		return
	}

	// read methods template
	b, err := ioutil.ReadFile(InputMethodsSource)
	if err != nil {
		fmt.Println("Failed to read template source", InputMethodsSource)
		return
	}
	InputSourceStr := string(b)

	// process methods template
	templ, err := template.New(InputMethodsSource).Parse(InputSourceStr)
	if err != nil {
		fmt.Println("Failed to create New template", err)
		return
	}
	// create output file
	f, err := os.Create(Ouput)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// read methods beginer
	b, err = ioutil.ReadFile(Beginner)
	if err != nil {
		fmt.Println("Failed to read APIs.go")
		return
	}

	// write beginer
	_, err = f.WriteString(string(b))
	if err != nil {
		fmt.Println("Failed to WriteString", err)
		return
	}
	fmt.Println("WRITED:", Ouput)
	// process and write template
	for i, task := range env {
		if i == n {
			break
		}
		err = templ.Execute(f, task)
		if err != nil {
			fmt.Println("Failed to execute template!", err)
			return
		}
		//		return nil // BREAKE on first
	}
	return nil
}
