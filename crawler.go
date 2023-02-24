package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {

	sql_connect := "root:123456@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", sql_connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for {
		resp, err := http.Get("https://www.nihaowua.com/home.html")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Printf("解析错误%s", err)
		}
		doc.Find("section div").Each(func(i int, s *goquery.Selection) {
			word := s.Text()
			fmt.Printf("Review : %s \n", word)
			sqlstr := "INSERT INTO chicken(word) VALUE (?)"
			ret, err := db.Exec(sqlstr, word)
			if err != nil {
				fmt.Printf("插入失败，下次继续")
			}
			fmt.Printf("成功插入:%d\n", ret)
		})
	}



}