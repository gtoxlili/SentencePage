package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

type ZawenContent struct {
	Title   template.HTML `json:"title"`
	Content template.HTML `json:"content"`
}

func timetoChar() string {
	chineseNum := []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	nowstring := time.Now().Format("2006年1月2日3点04分")
	var changeString string
	for _, v := range nowstring {
		fmt.Println(v)
		if v <= 57 && v >= 48 {
			changeString += chineseNum[v-48]
		} else {
			changeString += string(v)
		}
	}
	return changeString
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.GET("/", func(c *gin.Context) {
		f, _ := ioutil.ReadFile("data.json")
		var contents []ZawenContent
		json.Unmarshal(f, &contents)

		c.HTML(200, "index.html", gin.H{
			"title": timetoChar(), "result": contents,
		})
	})
	router.Run()
}
