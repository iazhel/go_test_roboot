package main

// use cmd stdout to write in to file

import (
//	epd "epd_api_model"
	fs "file_to_string_slice"
	"fmt"
	"gojson"
	"strings"
	"access_api"
)

func DoSlise() {
	fn := "/home/illia/dev/epd_api_model/model.go"
	s, e := fs.GetSlice(fn)
	fmt.Println(e)
	fmt.Println(len(s))
	for _, t := range s {
		n := strings.Index(t, " = ")
		if n > 0 {
			sec := strings.Replace(t[n:], " = ", "", 1)
			second := strings.Replace(sec, "`", "", 1)
			fmt.Println(second + ",")
		} else {
			fmt.Println(t)
		}
	}

}

func DoNamesStruct() {
	fn := "/home/illia/dev/src/epd_api_model/model.go"
	s, e := fs.GetSlice(fn)
	fmt.Println(e)
	fmt.Println(len(s))
	for _, t := range s {
		n := strings.Index(t, " = ")
		if n > 0 {
			firs := strings.TrimSpace(t[:n])
			first := strings.Replace(firs, "Call", "Type", 1)
			fmt.Println(`"` + first + `",`)
		} else {
			fmt.Println(t)
		}
	}

}

func DoConst() {
	fn := "/home/illia/dev/epd_api_model/model.go"
	s, e := fs.GetSlice(fn)
	fmt.Println(e)
	fmt.Println(len(s))
	for _, t := range s {
		n := strings.Index(t, "string `json:")
		if n > 0 {
			first := strings.TrimRight(t[:n], " ")
			sec := strings.Replace(t[n:], "string `json:", "", 1)
			second := strings.Replace(sec, "`", "", 1)
			fmt.Println(first + "Call = " + second)
		} else {
			fmt.Println(t)
		}
	}

}
func connect() (*access_api.RestApiClient, error) {
	rest, err, _ := access_api.Login("http://epd-dev.devicepros.net:8010", "i.azhel", "12345")
	if err != nil {
		panic(err)
		return rest, err
	}
	return rest, err
}
func DoStructures() {
	Types := ReturnApiTypes()
	apiSuffix := ReturnApiSuffixes()
	if len(Types) != len(apiSuffix) {
		fmt.Println("Check slice lenght!!!")
		return
	}
	rest, _ := connect()
	rest.PrintUrl(true)
	for i, api := range apiSuffix {
		b, err, _ := rest.Get(api)
		if err != nil {
			fmt.Println(err)
			continue
		}
		r := strings.NewReader(string(b))
		if x, err := gojson.Generate(r, gojson.ParseJson, Types[i], "gojson", []string{"json"}, false); err != nil {
			fmt.Println("Generate() error:", err)
		} else {
			fmt.Println(string(x))
		}

	}
}

func MakeEnv() {
	rest, err := connect()
	if err != nil {
		panic(err)
	}
	rest.PrintCurlCommand(true)
	API := ReturnApiSuffixes()
	Name := ReturnApiTypes()

	for i, api := range API {
		//
		b, err, _ := rest.Get(api +"/?format=json")
		if err != nil {
			fmt.Println(`/*`, err, `*/`)
			fmt.Println(`/*`, access_api.LastHistory(), `*/`)
			continue
		}
		// get header
		ha, err, _ := rest.GetHeaders(api + "/?format=json")
		if err != nil {
			fmt.Println(`/*`, err, `*/`)
			fmt.Println(`/*`, access_api.LastHistory(), `*/`)
			continue
		}

		h := ha["Allow"]
		if len(h) == 0 {
			continue
		}

		base := strings.TrimSuffix(Name[i], "Type")

		fmt.Printf("{\n	ApiCall:\"%s\",\n", api)
		fmt.Printf("	GoApiCallName:\"%s%s\",\n", base, "Call")
		fmt.Printf("	WideType:\"%s\",\n", Name[i])

		if strings.Contains(h[0], "GET") {
			fmt.Printf("	GetMethod:\"%s%s\",\n", "Get", base)
		}
		if strings.Contains(h[0], "POST") {
			fmt.Printf("	PostMethod:\"%s%s\",\n", "Post", base)
		}
		if strings.Contains(h[0], "PUT") {
			fmt.Printf("	PutMethod:\"%s%s\",\n", "Put", base)
		}
		if strings.Contains(h[0], "PATCH") {
			fmt.Printf("	PatchMethod:\"%s%s\",\n", "Patch", base)
		}
		if strings.Contains(h[0], "DELETE") {
			fmt.Printf("	DelMethod:\"%s%s\",\n", "Del", base)
		}
		if strings.Contains(h[0], "HEAD") {
			fmt.Printf("	HeadMethod:\"%s%s\",\n", "Head", base)
		}
		if strings.Contains(h[0], "OPTIONS") {
			fmt.Printf("	OptMethod:\"%s%s\",\n", "Opt", base)
		}
		if strings.Contains(string(b), "previous") {
			fmt.Printf("	ExampleType:\"%s\",\n", base)
			fmt.Printf("	ExampleName:\"%s%s\",\n", base, "Fields")
			fmt.Printf("	IdField:\"%s\",\n", "EnterFieldId")
			fmt.Printf("	PrimaryFieldName:\"%s\",\n", "EnterPrimaryFieldName")

		}

		fmt.Println("},")
	}
}

func main() {
	//	DoStructures()
  DoSlise()
	//	DoNamesStruct()
	//	GetAllHeaders()
//		MakeEnv()
	//	DoStructures()
}
