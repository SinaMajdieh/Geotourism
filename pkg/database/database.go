package database

/*
import (
	"fmt"
	"github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/sqlite"
)
type Database interface {
	Insert(interface{})
	Load_page_comments(url string)[]MSG
}
type sql struct {
	name string
	addr string
}
type MSG struct {
	Name      string `db:"name" json:"name"`
	Text      string `db:"text" json:"text"`
	Date      string `db:"date" json:"date"`
	Url       string `db:"url" json:"url"`
	Confirmed int    `db:"confirmed" json:"confirmed"`
}
func (*MSG) TableName() string {
	return "comment"
}
func New_database(addr string , name string)Database{
	return sql{
		name: name,
		addr: addr,
	}
}

func (DB sql) Insert(st interface{}){
	db, err := godb.Open(sqlite.Adapter, DB.addr)
	if nil != err{
		fmt.Println(err)
	}
	err = db.Insert(st).Do()
	if nil != err{
		fmt.Println(err)
	}
	db.Close()

}
func (DB sql) Load_page_comments(url string)[]MSG  {
	db , err := godb.Open(sqlite.Adapter , DB.addr)
	if nil != err {
		fmt.Println(err)
	}
	sts := make([]MSG , 0 ,0)
	iter, err := db.SelectFrom("comment").
		Columns("name", "text", "date" , "url" , "confirmed").
		DoWithIterator()
	if nil != err {
		fmt.Println(err)
	}
	for iter.Next() {
		msg := MSG{}
		err := iter.Scan(&msg)
		fmt.Println(msg)
		if nil != err {
			fmt.Println(err)
		}
		if msg.Url == url && msg.Confirmed == 1{
			sts = append(sts , msg)
		}

	}
	db.Close()

	return sts
}
*/
