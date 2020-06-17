package main

import (
	"fmt"
	_ "github.com/alexbrainman/odbc" // google's odbc driver
	"github.com/go-xorm/xorm"
	"xorm.io/core"
	"github.com/axgle/mahonia"
	"time"
)

type Person struct {
	PersonId int    `xorm:"PERSONID"`
	Sex      string `xorm:"SEX"`
	Name     string `xorm:"NAME"`
	Email    string `xorm:"EMAIL"`
	Phone    string `xorm:"PHONE"`
}

func main() {

	Engine, err := xorm.NewEngine("odbc", "driver={DM7 ODBC DRIVER};server=127.0.0.1:5236;database=DM;uid=SYSDBA;pwd=SYSDBA;charset=utf8")
	if err != nil {
		fmt.Println("new engine got error:", err)
		return
	}
	if err := Engine.Ping(); err != nil {
		fmt.Println("ping got error:", err)
		return
	}

	Engine.SetTableMapper(core.SameMapper{})
	Engine.ShowSQL(true)
	Engine.SetMaxOpenConns(5)
	Engine.SetMaxIdleConns(5)

	total, err := Engine.Table("SYSDBA.PERSON").Count(&Person{})
	if nil != err {
		fmt.Println(`engine query got error:`, err.Error())
		return
	}
	fmt.Println("total count is:", total)
	result, err := Engine.Query("select sex,name from sysdba.person")
	if err != nil {
		panic(err)
	}
	for i, e := range result {
		fmt.Printf("%v\n", i)
		for k, v := range e {
			// 达梦数据库中文默认为gbk
			fmt.Printf("%v=%v\t", k, ConvertToString(string(v), "gbk", "utf-8"))
		}
		fmt.Printf("\n")
	}
	fmt.Println("----------")
	t1 := time.Now()
	var person []Person
	err = Engine.Table("SYSDBA.PERSON").Limit(5, 0).Find(&person)
	if err != nil {
		fmt.Println("查询出错")
		panic(err)
	}
	for i, e := range person {
		e.Name = ConvertToString(e.Name, "gbk", "utf-8")
		e.Sex = ConvertToString(e.Sex, "gbk", "utf-8")
		e.Email = ConvertToString(e.Email, "gbk", "utf-8")
		fmt.Printf("%v=%v\n", i, e)
	}
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))

}

// 字符串解码函数，处理中文乱码
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
