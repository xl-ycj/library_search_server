package main

//import (
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"net/url"
//)
//
//func main() {
//
//		response, _ := http.PostForm("http://127.0.0.1:1024/book",url.Values{"book": {"思想"}, "mode":{"0"}})
//		//response, _ := http.PostForm("http://127.0.0.1:1024/detail",url.Values{"detail": {"1234567890048"}})
//		//response, _ := http.PostForm("http://127.0.0.1:1024/isbn",url.Values{"isbn": {"1234567890048"}})
//		//response, _ := http.PostForm("http://127.0.0.1:1024/barcode",url.Values{"barcode": {"1697143"}})
//		//response, _ := http.PostForm("http://127.0.0.1:1024/sort",url.Values{"sort": {"思想"},"mode":{"1"}})
//
//		body, _ := ioutil.ReadAll(response.Body)
//		defer response.Body.Close()
//		fmt.Println(string(body))
//
//
//
//}