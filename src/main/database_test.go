package main

//import (
//
//	"io/ioutil"
//	"log"
//	"encoding/json"
//	"fmt"
//	 _ "github.com/mattn/go-oci8"
//	"database/sql"
//	//_ "gopkg.in/rana/ora.v4"
//)
//
//type Connect struct {
//	db *sql.DB
//	err error
//}
//
//type Config struct {
//	Password string
//	User string
//	Server string
//	SID  string
//	Port string
//	Engine string
//}
//
//var cfg Config
//
//
//func connect() {
//	buf, err := ioutil.ReadFile("../../resource/config.json")
//	if err != nil {
//		log.Fatalln("Read config.json error: ", err.Error())
//	}
//
//	err = json.Unmarshal(buf, &cfg)
//	if err != nil {
//		log.Fatalln("json unmarshal error: ", err.Error())
//	}
//
//	fmt.Println(connString)
//	//  db.SetMaxOpenConns(1000)
//	if err != nil {
//		fmt.Println("ERR: ", err.Error())
//	}else {
//		fmt.Println("Conn Suc: ")
//	}
//	//err1 := conn.db.Ping()
//	//if err1 != nil {
//	//	fmt.Println("PING: ", err.Error())
//	//}else {
//	//	fmt.Println("Ping Suc!")
//	//}
//
//	rows, err := db.Query("select book_name, book_id from M_TRANSFORM_TAG m1 where m1.rowid =  where rownum <=20")
//	if err != nil {
//		fmt.Println("Query ERR: ",err.Error())
//		return
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var id string
//		var data string
//		rows.Scan(&id, &data)
//		fmt.Println(id, string(data))
//	}
//}
//
//func main()  {
//	connect()
//}
