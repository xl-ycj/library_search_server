package protocol

//import "database/sql"

/* 第三层返回的信息 书的位置详细信息 */
type Reply struct {
	Res []Loc `json:"res"`
}

/* 第二层返回的信息 书的基本信息 */
type ReplyBasic struct {
	Res []BasicInfo `json:"res"`
}

type Loc struct{
	Index string `json:"index"` // 索书号 即中图法索引
	House string `json:"house"` // 在哪个馆
	Area string `json:"area"` //在哪个区域
	IsLoaned int `json:"is_loaned"` // 是否在架
	ShelfId string `json:"shelfId"` // 书架编号
	Layer int `json:"layer"` // 书架层号
	Number int `json:"number"` // 第几位置编号
}

//type Loc struct{
//	Index sql.NullString `json:"index"` // 索书号 即中图法索引
//	House sql.NullString `json:"house"` // 在哪个馆
//	Area sql.NullString `json:"area"` //在哪个区域
//	IsLoaned sql.NullBool `json:"is_loaned"` // 是否在架
//	ShelfId sql.NullString `json:"shelfId"` // 书架编号
//	Layer sql.NullInt64 `json:"layer"` // 书架层号
//  Number sql.NullInt64 `json:"number"` // 第几位置编号
//}

type BasicInfo struct{
	BookName string `json:"book_name"` // 书名
	Author string `json:"author"` // 作者
	Press string `json:"press"` // 出版社
	ISBN string `json:"isbn"` // isbn
}

//type BasicInfo struct{
//BookName sql.NullString `json:"book_name"` // 书名
//Author sql.NullString `json:"author"` // 作者
//Press sql.NullString `json:"press"` // 出版社
//ISBN sql.NullString `json:"isbn"` // isbn
//}

type Config struct {
	Password string
	User string
	Server string
	ServiceName  string
	Port string
	Engine string
}


