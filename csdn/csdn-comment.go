package csdn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// json转struct https://oktools.net/json2go
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Count      int `json:"count"`
		PageCount  int `json:"pageCount"`
		FloorCount int `json:"floorCount"`
		FoldCount  int `json:"foldCount"`
		List       []struct {
			Info struct {
				CommentID             int           `json:"commentId"`
				ArticleID             int           `json:"articleId"`
				ParentID              int           `json:"parentId"`
				PostTime              string        `json:"postTime"`
				Content               string        `json:"content"`
				UserName              string        `json:"userName"`
				Digg                  int           `json:"digg"`
				DiggArr               []interface{} `json:"diggArr"`
				ParentUserName        string        `json:"parentUserName"`
				ParentNickName        string        `json:"parentNickName"`
				Avatar                string        `json:"avatar"`
				NickName              string        `json:"nickName"`
				DateFormat            string        `json:"dateFormat"`
				Years                 int           `json:"years"`
				Vip                   bool          `json:"vip"`
				VipIcon               string        `json:"vipIcon"`
				CompanyBlog           bool          `json:"companyBlog"`
				CompanyBlogIcon       string        `json:"companyBlogIcon"`
				Flag                  bool          `json:"flag"`
				FlagIcon              string        `json:"flagIcon"`
				LevelIcon             string        `json:"levelIcon"`
				CommentFromTypeResult struct {
					Index int    `json:"index"`
					Key   string `json:"key"`
					Title string `json:"title"`
				} `json:"commentFromTypeResult"`
				IsTop   bool `json:"isTop"`
				IsBlack bool `json:"isBlack"`
			} `json:"info"`
			Sub            []interface{} `json:"sub"`
			PointCommentID interface{}   `json:"pointCommentId"`
		} `json:"list"`
	} `json:"data"`
}

func CsdnComment() {
	client := &http.Client{}
	reqSpider, err := http.NewRequest("POST", "https://blog.csdn.net/phoenix/web/v1/comment/list/120168880?page=1&size=53&fold=unfold&commentId=", nil)
	if err != nil {
		log.Fatal(err)
	}
	reqSpider.Header.Set("content-length", "0") // 书写请求头
	reqSpider.Header.Set("accept", "*/*")
	reqSpider.Header.Set("x-requested-with", "XMLHttpRequest")
	respSpider, err := client.Do(reqSpider)
	if err != nil {
		log.Fatal(err)
	}
	defer respSpider.Body.Close()
	bodyText, err := ioutil.ReadAll(respSpider.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result Result
	_ = json.Unmarshal(bodyText, &result)
	for _, res := range result.Data.List {
		fmt.Println(res.Info.UserName, res.Info.Content)
	}
}
