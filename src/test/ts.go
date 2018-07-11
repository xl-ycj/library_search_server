package test
//
//import (
//	"encoding/json"
//	"fmt"
//	//"log"
//	//"os"
//	"log"
//	"io/ioutil"
//)
//
//type Address struct {
//	Type    string
//	City    string
//	Country string
//}
//
//type VCard struct {
//	FirstName string
//	LastName  string
//	Addresses []*Address
//	Remark    string
//}
//
//type Config struct {
//
//	Password string
//	User string
//	Server string
//	Encrypt  string
//	Database  string
//	Port int
//}
//
//func main() {
//	pa := &Address{"private", "Aartselaar", "Belgium"}
//	wa := &Address{"work", "Boom", "Belgium"}
//	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
//	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
//	// JSON format:
//	js, _ := json.Marshal(vc)
//	fmt.Printf("JSON format: %s", js)
//	// using an encoder:
//	buf, err := ioutil.ReadFile("resource/config.json")
//	if err != nil{
//		log.Println("Error in encoding json")
//	}
//	fmt.Println(buf)
//	//file, _ := os.OpenFile("../resource/config.json", os.O_CREATE|os.O_WRONLY, 0)
//	//defer file.Close()
//	var jj Config
//	err = json.Unmarshal(buf, &jj)
//	fmt.Println(jj.Server)
//	if err != nil{
//		log.Println("Error in encoding json 2")
//	}
//	//enc := json.NewEncoder(file)
//	//var jj interface {}
//	//err := enc.Encode(jj)
//	//if err != nil {
//	//	log.Println("Error in encoding json")
//	//}
//}
