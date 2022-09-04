package main

import "myapp/gorp"

type Test struct {
	Title string
	Page  int
}

func main() {
	//author := Author.Author{Name: "田中太郎"}
	//result := author.ConnectAuthorDb()
	//fmt.Println(result)
	//result = author.ConnectAuthorDbDi()
	//fmt.Println(result)
	//test := NewTest("Go本", 10)
	//fmt.Println(test)
	//test2 := NewTest("Go本", 10)
	//fmt.Println(*test2)
	// 構造体を試す
	//structs := structs.Structs{}
	//_ = structs.Test()
	//_ = literal.Test()
	//_ = arrays.Test()
	//_ = slices.Test()
	//_ = databasesql.Test()
	_ = gorptest.Test()
}

func NewTest(title string, page int) *Test {
	return &Test{
		Title: title,
		Page:  page,
	}
}
