package main

import (
	"net/http"
	"main/logic"
	"fmt"
	"time"
)

func main()  {

	mux := http.NewServeMux()
	mux.HandleFunc("/book", logic.GetBasicInfoByName)
	mux.HandleFunc("/barcode", logic.GetBasicInfoByBarcode)
	mux.HandleFunc("/isbn", logic.GetBasicInfoByIsbn)
	mux.HandleFunc("/detail", logic.GetLocationInfo)
	mux.HandleFunc("/sort", logic.GetBasicInfoBySortedTime)

	timeStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(timeStr + " Server starts listening at port: 1024")
	http.ListenAndServe(":1024", mux)


}