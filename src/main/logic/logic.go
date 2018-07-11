/*------------------------------------------------------------------------------*/
//version:1.0
//Author: xulin
//description: this is the server of library search system
//
/*------------------------------------------------------------------------------*/

package logic

import (
	"encoding/json"
	"net/http"
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-oci8"
	"log"
	"io/ioutil"
	."main/protocol"

	"time"
	"strings"
)


/*------------------------------------------------------------------------------*/
const BOOKNAME  = 0
const ISBN  	= 1
const BARCODE   = 2
const SORT 		= 3
const ADVANCE   = 4
var conn Connect // database connection
/*------------------------------------------------------------------------------*/

/*------------------------------------------------------------------------------*/

type Connect struct {
	db *sql.DB
	err error
}

/**
 connection function
 */
func (con *Connect) getDB(){
	var cfg Config

	buf, err := ioutil.ReadFile("resource/config.json")
	if err != nil{
		log.Fatalln(" Read config.json error: ", err.Error())
	}

	err = json.Unmarshal(buf, &cfg)
	if err != nil{
		log.Fatalln(" json unmarshal error: ", err.Error())
	}

	connString := fmt.Sprintf("%s/%s@%s:%s/%s",
		 cfg.User, cfg.Password, cfg.Server, cfg.Port, cfg.ServiceName)
	con.db, con.err = sql.Open(cfg.Engine, connString)
}

func (con *Connect) closeDB(){
	if con.db == nil{
		log.Fatalln(" error: con.db is nil you can not close")
		return
	}
	con.db.Close()
}
/*------------------------------------------------------------------------------*/

/**
  pattern int
  para string
 */
func genSqlStat(pattern int, para string) string {
	p 		:= pattern
    sqlStat := ""

	switch p {
	case BOOKNAME:
		sqlStat = fmt.Sprintf("SELECT BOOK_NAME, AUTHOR, PUBLISHING_HOUSE, BOOK_ISBN FROM BOOK_INFO WHERE BOOK_NAME LIKE '%s' ","%"+para + "%")
	case ISBN:
		sqlStat = fmt.Sprintf("SELECT BOOK_NAME, AUTHOR, PUBLISHING_HOUSE, BOOK_ISBN FROM BOOK_INFO WHERE BOOK_ISBN = %s ",para)
	case BARCODE:
		sqlStat = fmt.Sprintf("SELECT BOOK_NAME, AUTHOR, PUBLISHING_HOUSE, BOOK_ISBN FROM BOOK_INFO WHERE BOOK_ID = %s ", para)
	case SORT:
		list := strings.Split(para,"/")
		para  = list[0]
		if list[1] == "1" {
			sqlStat = fmt.Sprintf("SELECT BOOK_NAME, AUTHOR, PUBLISHING_HOUSE, BOOK_ISBN FROM BOOK_INFO WHERE BOOK_NAME  LIKE '%s' ORDER BY PUBLISHING_DATETIME","%"+para + "%")
		}else {
			sqlStat = fmt.Sprintf("SELECT BOOK_NAME, AUTHOR, PUBLISHING_HOUSE, BOOK_ISBN FROM BOOK_INFO WHERE BOOK_NAME  LIKE '%s' ORDER BY PUBLISHING_DATETIME DESC","%"+para + "%")
		}
	case ADVANCE:
		keys := strings.Split(para, " ")
		tmp  := ""
		for _, v := range keys{
			if v != " "{
				tmp += "%"+ v
			}
		}
		tmp += "%"
		sqlStat = fmt.Sprintf("SELECT BOOK_NAME, AUTHOR, PUBLISHING_HOUSE, BOOK_ISBN FROM BOOK_INFO WHERE BOOK_NAME LIKE '%s' ", tmp)
	default:
		sqlStat = fmt.Sprintf("SELECT BOOK_NAME, AUTHOR, PUBLISHING_HOUSE, BOOK_ISBN FROM BOOK_INFO WHERE BOOK_NAME LIKE '%s' ","%"+para + "%")

	}

	return sqlStat
}

/**
  get book basic information like name,author,etc
  input: pattern , mark
  return: books' basic information list
 */
func getBasicInfoList(pattern int, mark string) []BasicInfo {

	if conn.db == nil {
		conn.getDB()
	}

	if conn.err != nil {

		log.Fatal(" getBasicInfoList connect to database fail: ", conn.err.Error())
	}

	sqlStat   := genSqlStat(pattern, mark)
	rows, err := conn.db.Query(sqlStat)
	defer rows.Close()
	if err != nil{
		log.Fatal(" getBasicInfoList DB query fail: ", err.Error())
	}

	var basicInfoList []BasicInfo
	basicInfoMap := make(map[string]int)

	for rows.Next(){
		var bi BasicInfo
		err := rows.Scan(&bi.BookName, &bi.Author, &bi.Press, &bi.ISBN)
		if err != nil {
			timeStr := time.Now().Format("2006-01-02 15:04:05")
			fmt.Println(timeStr + " getBasicInfoList scan rows fail: ",err.Error())
			continue
		}
		if basicInfoMap[bi.ISBN] == 2{ // value 2 is non-meanable
			continue
		}
		basicInfoList = append(basicInfoList, bi)
		basicInfoMap[bi.ISBN] = 2
	}

	return basicInfoList
}


func GetBasicInfoByName(w http.ResponseWriter, r *http.Request) {
	//go func(){
		r.ParseForm()

		bookName := r.FormValue("book")
		mode 	 := r.FormValue("mode")
		fmt.Println(bookName)

		if bookName == "" {
			w.WriteHeader(403)
			return
		}

		p := 0
		if mode == "1"{
			p = 4
		}

		list := getBasicInfoList(p, bookName)
		fmt.Println(list)
		//book := string([]rune(bookName))
		//fmt.Println(book)

		reply := &ReplyBasic{Res:list}
		re, _ := json.Marshal(reply)
		rr    := string(re)
		fmt.Fprintln(w, rr)
	//}()

}

func GetBasicInfoBySortedTime(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	bookName := r.FormValue("sort")
	sortMode := r.FormValue("mode")
	bookName  = bookName + "/"+ string(sortMode)
	
	fmt.Println(bookName)
	fmt.Println(string(sortMode))

	list := getBasicInfoList(3, bookName)
	//fmt.Println(list)
	//book := string([]rune(bookName))
	//fmt.Println(book)

	reply := &ReplyBasic{Res:list}
	re, _ := json.Marshal(reply)
	rr    := string(re)
	fmt.Fprintln(w, rr)
}


func GetBasicInfoByBarcode(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	barCode := r.FormValue("barcode")
	fmt.Println(barCode)
	if barCode == ""{
		w.WriteHeader(403)
		return
	}

	list  := getBasicInfoList(2, barCode)
	reply := &ReplyBasic{Res:list}
	re, _ := json.Marshal(reply)
	rr    := string(re)

	w.Write([]byte(rr))
}


func GetBasicInfoByIsbn(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	isbn := r.FormValue("isbn")
	fmt.Println(isbn)
	if isbn == ""{
		w.WriteHeader(403)
		return
	}

	list  := getBasicInfoList(1, isbn)
	reply := &ReplyBasic{Res:list}
	re, _ := json.Marshal(reply)
	rr    := string(re)

	w.Write([]byte(rr))
}


func getDetailInfoList(isbn string) []Loc {

	if conn.db == nil {
		conn.getDB()
	}

	if conn.err != nil {
		log.Fatal( " getDetailInfoList connect to database fail: ", conn.err.Error())
	}

	sqlStat   := fmt.Sprintf("SELECT BOOK_INDEX, CURRENT_LIBRARY, AREANO, IS_ON_THE_SHELF, SHELFNO, LAYERNO FROM BOOK_INFO WHERE BOOK_ISBN = %s", isbn)
	rows, err := conn.db.Query(sqlStat)
	defer rows.Close()
	if err != nil{

		log.Fatal( " getDetailInfoList DB query fail: ", err.Error())
	}

	var detailInfoList []Loc
	for rows.Next(){
		var loc Loc
		err := rows.Scan(&loc.Index, &loc.House, &loc.Area, &loc.IsLoaned, &loc.ShelfId, &loc.Layer)
		if err != nil {
			fmt.Println( " getDetailInfoList scan rows fail: ",err.Error())
			continue
		}
		detailInfoList = append(detailInfoList, loc)
	}

	return detailInfoList
}


func GetLocationInfo(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	isbn := r.FormValue("detail")
	fmt.Println(isbn)
	if isbn == ""{
		w.WriteHeader(403)
		return
	}

	list  := getDetailInfoList(isbn)
	reply := &Reply{Res:list}
	re, _ := json.Marshal(reply)
	rr    := string(re)
	fmt.Fprintln(w, rr)
}


