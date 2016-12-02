package main

import (
	"epd_api_model/templates/methods"
	"fmt"
)

func main() {
	r, err, d := methods.Login("http://epd-dev.devicepros.net:8010/", "i.azhel", "12345")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("##########################", d)
	r.PrintUrl(true)
	r.PrintCurlCommand(true)

	_, err, _ = r.GetAuditHistoryView()
	if err != nil {
		fmt.Println("ERROR!  GetAuditHistoryView() ", err)
	}
	_, err, _ = r.OptAuditHistoryView()
	if err != nil {
		fmt.Println("ERROR!  OptAuditHistoryView() ", err)
	}
	_, err, _ = r.GetChoicesBlockSize()
	if err != nil {
		fmt.Println("ERROR!  GetChoicesBlockSize() ", err)
	}
	_, err, _ = r.OptChoicesBlockSize()
	if err != nil {
		fmt.Println("ERROR!  OptChoicesBlockSize() ", err)
	}
}
