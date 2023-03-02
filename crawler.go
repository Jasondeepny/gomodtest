package main

//
//import (
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"regexp"
//)
//
////import (
////	"fmt"
////	"io/ioutil"
////	"net/http"
////	"regexp"
////)
//
////
////import (
////	"database/sql"
////	"fmt"
////	"log"
////	"net/http"
////	"github.com/PuerkitoBio/goquery"
////	_ "github.com/go-sql-driver/mysql"
////)
////
////func main()  {
////
////	sql_connect := "root:123456@tcp(127.0.0.1:3306)/test"
////	db, err := sql.Open("mysql", sql_connect)
////	if err != nil {
////		panic(err)
////	}
////	defer db.Close()
////
////	for {
////		resp, err := http.Get("https://www.nihaowua.com/home.html")
////		if err != nil {
////			log.Fatal(err)
////		}
////		defer resp.Body.Close()
////		doc, err := goquery.NewDocumentFromReader(resp.Body)
////		if err != nil {
////			fmt.Printf("解析错误%s", err)
////		}
////		doc.Find("section div").Each(func(i int, s *goquery.Selection) {
////			word := s.Text()
////			fmt.Printf("Review : 第%d条，文字为:%s\n",i, word)
////			sqlstr := "INSERT INTO chicken(word) VALUE (?)"
////			ret, err := db.Exec(sqlstr, word)
////			if err != nil {
////				fmt.Printf("插入失败，下次继续")
////			}
////			fmt.Printf("成功插入:%d\n", ret)
////		})
////	}
////
////
//////
////	GetEmail()
////}
//
//var (
//	// \d是数字
//	reQQEmail = `(\d+)@qq.com`
//)
//
//// 爬邮箱
//func GetEmail() {
//	// 1.去网站拿数据
//	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
//	HandleError(err, "http.Get url")
//	defer resp.Body.Close()
//	// 2.读取页面内容
//	pageBytes, err := ioutil.ReadAll(resp.Body)
//	HandleError(err, "ioutil.ReadAll")
//	// 字节转字符串
//	pageStr := string(pageBytes)
//	//fmt.Println(pageStr)
//	// 3.过滤数据，过滤qq邮箱
//	re := regexp.MustCompile(reQQEmail)
//	// -1代表取全部
//	results := re.FindAllStringSubmatch(pageStr, -1)
//	//fmt.Println(results)
//
//	// 遍历结果
//	for _, result := range results {
//		fmt.Println("email:", result[0])
//		fmt.Println("qq:", result[1])
//	}
//}
//
//// 处理异常
//func HandleError(err error, why string) {
//	if err != nil {
//		fmt.Println(why, err)
//	}
//}
