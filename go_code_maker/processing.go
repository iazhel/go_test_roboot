package main

import (
	"access_api"
	"fmt"
	"gojson"
	"os"
	"strings"
	"templates"
)

// write to file, StdOut ...
func outputln(a ...interface{}) {
	fmt.Fprintln(output, a...)
}
func outputf(f string, a ...interface{}) {
	fmt.Fprintf(output, f, a...)
}

func procMainPage(newPackage string) (env []templates.GoApiEnvironment) {

	// write first package line
	outputln("package", newPackage)

	// get struct list
	jsonList, err := getJSON("", newPackage)
	if err != nil {
		fmt.Println("Main page reader:", err)
		os.Exit(2)
	}

	// parse url and golang names in list
	base, urlPath := doNamesAndPath(jsonList)

	// make types and methods names
	for i := range urlPath {
		/*
		if stringurlPath[i]
		*/
		e := templates.GoApiEnvironment{
			WideType:      base[i] + "Type",
			ExampleType:   base[i] + "",
			ExampleName:   base[i] + "Field",
			GoApiCallName: base[i] + "Call",
			ApiCall:       strings.Trim(urlPath[i], `"`),
			Base:          base[i],
		}
		outputln("const", e.GoApiCallName, "=", urlPath[i])
		env = append(env, e)
		//		break
	}

	// make structures from json output
	for i, e := range env {
		structure, err := getJSON(e.ApiCall+"/?format=json", e.WideType)
		if err != nil {
			err = fmt.Errorf("/* getJSON error %s%s : %s */", rest.BaseUrl, e.ApiCall, err)
			e.HasError = true
			e.Error = err.Error()
			env[i] = e
			continue
		}
		finalStructure := replaseComplex(structure, &e)
		// write structure
		env[i] = e
		outputln(finalStructure)

	}
	return env
}

func MakeEnv(env []templates.GoApiEnvironment) []templates.GoApiEnvironment {

	rest.PrintUrl(true)
	for i, e := range env {
		if e.HasError{
			continue
		}

		// get headers
		fmt.Println(" getting  ALLOW")
//		ha, err, _ := rest.GetHeaders(e.ApiCall + "/?format=json")// EPD format
		ha, err, _ := rest.GetHeaders(e.ApiCall /* + "/?format=json"*/)  // STELLUS format
		if err != nil {
			err = fmt.Errorf("/* getHeader %s%s : %s */", rest.BaseUrl, e.ApiCall, err)
			e.HasError = true
			e.Error = err.Error()
			env[i] = e
			continue
		}

		// get Allow header
		h := ha["Allow"]
		if len(h) == 0 {
			fmt.Println("Headers:", ha)
			continue
		}
		fmt.Println(" - ALLOW", h)

		// make functions names
		if strings.Contains(h[0], "GET") {
			e.GetMethod = "Get" + e.Base
			if e.IsComplex {
				e.GetOneMethod = "GetOne" + e.Base
				e.CommonTest = e.GetMethod
							} else {
				e.CommonTest = e.GetOneMethod
			}
			if e.CommonTest == ""{
				e.CommonTest = e.GetMethod
			}
		}

		if strings.Contains(h[0], "POST") {
			e.PostMethod = "Post" + e.Base
		}
		if strings.Contains(h[0], "PUT") {
			e.PutMethod = "Put" + e.Base
		}
		if strings.Contains(h[0], "PATCH") {
			e.PatchMethod = "Patch" + e.Base
		}
		if strings.Contains(h[0], "DELETE") {
			e.DelMethod = "Delete" + e.Base
		}
		if strings.Contains(h[0], "HEAD") {
			e.HeadMethod = "Head" + e.Base
		}
		if strings.Contains(h[0], "OPTIONS") {
			e.OptMethod = "Option" + e.Base
		}
		// add field composite stucrure
		if e.IsComplex {
			// TODO: input this fieald interactive
//			e.IdField = "You_Enter_FieldId"
//			e.PrimaryFieldName = "You_Enter_PrimaryFieldName"
		} else {
			e.ExampleType = e.WideType
		}
		env[i] = e
	}

	for _, e := range env {
		if e.HasError{
			fmt.Println(e.Error)
		}
	}
	return env
}

func replaseComplex(structure string, e *templates.GoApiEnvironment) (out string) {

	// s - structure without first line 'package ...'
	firstP := strings.Index(structure, "type")
	s := structure[firstP:]

	// check on Complex structure fields
	if !(strings.Contains(s, "Count") && strings.Contains(s, "Next") && strings.Contains(s, "Previous")) {
		e.GetRejected = true
		e.GetOneRejected = true
		return s
	}

	e.IsComplex = true
	e.GetSimpleRejected = true

	// look brackets pair
	resultP := strings.Index(s, "Results")
	n := findCloseBracket(s, resultP)
	if !(n > 0 && n < len(s)) {
		return s
	}
	// take out internal structure
	subStructure := s[resultP:n]
	end := ""

	// search []interface{} in subStructure
	if strings.Contains(subStructure, "[]interface{}") && len(subStructure) < 33 {
		end = `Results []struct {
					UnknownFields interface{}
				}
`
	} else {
		end = subStructure
	}

	out = strings.Replace(s, subStructure, e.ExampleName+" []"+e.ExampleType, 1)
	end = strings.Replace(end, "Results", e.ExampleType, 1)
	end = strings.Replace(end, "[]", " ", 1)

	return out + "type " + end
}

func findCloseBracket(s string, n int) int {
	var open, close int
	j := strings.Index(s[n:], "{")
	if j < 0 {
		return -1
	}
	open = 1
	o, c := byte("{"[0]), byte("}"[0])
	for i := j + n + 1; i < len(s); i++ {
		switch s[i] {
		case o:
			open++
		case c:
			close++
		}
		if open == close {
			return i + 1
		}
	}
	return -1
}

func doNamesAndPath(list string) (name []string, path []string) {
	for _, t := range strings.Split(list, "\n") {
		n := strings.Index(t, "string `json:")
		if n > 0 {
			first := strings.TrimRight(t[:n], " ")
			sec := strings.Replace(t[n:], "string `json:", "", 1)
			second := strings.Replace(sec, "`", "", 1)
			name = append(name, strings.TrimSpace(first))
			path = append(path, second)
		}
	}
	return
}

func connect() *access_api.RestApiClient {
	rest, err, _ := access_api.Login(baseUrl, login, pass)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return rest
}

func getJSON(path, name string) (string, error) {
	b, err, _ := rest.Get(path)
	r := strings.NewReader(string(b))
	// generate golang style names
	b, err = gojson.Generate(r, gojson.ParseJson, name, "gojson", []string{"json"}, false)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
