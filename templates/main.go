package templates

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"
)
type ApiParameters struct {
	BaseUrl  string
	Login string
	Password string
	PackageName string
}

type GoApiEnvironment struct {

	Error    string
	HasError bool

	IsComplex bool
	Base      string

	ApiCall         string // url value  "/test_cases/"
	GoApiCallName   string
	GetMethod       string
	GetSimpleMethod string
	GetOneMethod    string
	PostMethod      string
	PutMethod       string
	DelMethod       string
	OptMethod       string
	PatchMethod     string
	HeadMethod      string

	WideType    string // Common type, is returned by GET all.
	ExampleType string // Is included in WideType as slice element.
	OneGetType  string // Uses If ExampleType != OneGetType.

	ExampleName      string // Field name of slice in getted structure.
	IdField          string
	PrimaryFieldName string

	CommonTest string

	GetRejected       bool
	GetSimpleRejected bool
	GetOneRejected    bool
	PostRejected      bool
	DelRejected       bool
	PatchRejected     bool
	PutRejected       bool
	HeadRejected      bool
	OptionRejected    bool

	FieldsNameMissed bool
}

var (
	env        []GoApiEnvironment
newPackage string
//redential string
)

func Do(enviroument []GoApiEnvironment,p ApiParameters) {


	newPackage = p.PackageName
	env = enviroument
	//	APIsBeginer := "/s/beginner_APIs"
	APIsBeginer := "/home/illia/GoDev/src/templates/beginner_APIs"
	InputMethodsSource := "/home/illia/GoDev/src/templates/template_methods"
	OuputMethodsSource := newPackage + "/methods_" + newPackage + ".go"

	testsBeginer := "/home/illia/GoDev/src/templates/beginner_tests"
	InputTestsSource := "/home/illia/GoDev/src/templates/template_tests"
	OuputTestsSource := newPackage + "/" + newPackage + "_test.go"

	for i, api := range env {

		//		check field names
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

	consts := `
import (
	"testing"
	"access_api"
)
	`
	consts += fmt.Sprintf("\nconst BaseUrl = \"%s\"\nconst login = \"%s\"\nconst password = \"%s\"\n",p.BaseUrl,p.Login, p.Password)

	_ = process(APIsBeginer, InputMethodsSource, OuputMethodsSource, "")
	_ = process(testsBeginer, InputTestsSource, OuputTestsSource, consts)
}

func process(Beginner, InputMethodsSource, Ouput string, consts string) (err error) {
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
		fmt.Println("Failed to read template source", InputMethodsSource, err)
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

	_, err = f.WriteString("package " + newPackage + consts)
	if err != nil {
		fmt.Println("Failed to WriteString", err)
		return
	}
	_, err = f.WriteString(string(b))
	fmt.Println("WRITED:", Ouput)
	// process and write template
	for _, task := range env {
		err = templ.Execute(f, task)
		if err != nil {
			fmt.Println("Failed to execute template!", err)
			return
		}
		//		return nil // BREAKE on first
	}
	return nil
}
