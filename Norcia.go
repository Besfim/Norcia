package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"encoding/json"
	"io/ioutil"
	"strings"
)

/**
配置文件 json 格式
	{
		"head": "博客名称",
		"introduce": "博客介绍",
		"github": "github地址",
		"weibo": "weibo地址",
		"articles": [
			{
				"title": "文章标题",
				"tag": "文章标签",
				"create": "创作日期",
				"update": "更新日期"
			},
			{
				"title": "文章标题",
				"tag": "文章标签",
				"create": "创作日期",
				"update": "更新日期"
			}
		]
	}
*/
const docDirName string = "document/"
const miniDocDirName string = "document_mini/"
func main() {
	//var oldArticleNum = 0
	var updateNum = 0
	var createNum = 0
	//var deleteNum = 0
	blogconfig := parseConfigJson(readFileToString("config.json"))
	//oldArticleNum = len(blogconfig.Articles)
	//读取 document 文件
	files, _ := ioutil.ReadDir("document")
	var articles []Article
	for _,file := range files {
		var fileName = file.Name()                              //文件名
		var title = strings.Replace(fileName,".md","",-1)       //title
		var docContent = readFileToString(docDirName +fileName) //文档内容
		var docMini = substr(docContent,0,200)                  //文档的缩小部分
		var updateTime = file.ModTime()                         //更新时间
		var temp Article
		articleFromConfig,successFlag := getArticleFromConfigByTitle(title,blogconfig)
		if successFlag == 1{
			//如果能够找到旧的文件
			//最后修改时间没变
			if articleFromConfig.Update == substr(updateTime.String(),0,16) {
				temp = articleFromConfig
			}else {
				temp = Article{
					Title:  title,
					Tag:    "tag",
					Update: substr(updateTime.String(),0,16),
					Create: articleFromConfig.Create,
				}
				updateNum++
			}
		}else if successFlag == 0{
			//如果无法找到旧的文件,证明文件时新建的!
			temp = Article{
				Title:  title,
				Tag:    "tag",
				Update: substr(updateTime.String(),0,16),
				Create: substr(updateTime.String(),0,16),
			}
			createNum++
		}
		//每一次都会刷新头部缓存
		writeStringToFile(docMini,miniDocDirName+fileName)
		articles = append(articles, temp)
	}
	blogconfig.Articles = articles
	outputNewBlogConfig(blogconfig)
	fmt.Println("update",updateNum,"document(s), and create",createNum,"documents(s)")
}


func readFileToString(fileName string) string{
	inputFile, inputError := os.Open(fileName)
	if inputError != nil {
		fmt.Printf("")
		return "error"
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	var res string
	for {
		inputString, readerError := inputReader.ReadString('\n')
		res += inputString
		if readerError == io.EOF{
			return res
		}
	}
	return res
}

type BlogConfig struct {
	Head      string    `json:"head"`
	Introduce string    `json:"introduce"`
	Github    string    `json:"github"`
	Weibo     string    `json:"weibo"`
	Articles  []Article `json:"articles"`
}

type Article struct {
	Title		string	`json:"title"`
	Tag			string	`json:"tag"`
	Create		string	`json:"create"`
	Update		string	`json:"update"`
}

// 解析 配置 Json 的函数
func parseConfigJson(jsonString string) BlogConfig {
	var config BlogConfig
	json.Unmarshal([]byte(jsonString),&config)
	return config
}

//将 BlogConfig 对象变成 string
func generateConfigJson(config BlogConfig) string {
	str,_ := json.MarshalIndent(config, "", "\t")
	return string(str)
}

//将新的 json 写到 config.json 里面去
func outputNewBlogConfig(config BlogConfig) {
	writeStringToFile(generateConfigJson(config),"config.json")
}

func writeStringToFile(outputString string,fileName string) {
	outputFile, outputError := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputWriter.WriteString(outputString)
	outputWriter.Flush()
}

//从 BlogConfig 里面获取一个 Article
// 获取成功的话第二个参数是 1 , 否则是 0
// const nullArticle := Article{}
var nullArticle Article
func getArticleFromConfigByTitle(title string,config BlogConfig) (Article,int) {
	for _,article := range config.Articles{
		if article.Title == title{
			return article,1
		}
	}
	return nullArticle,0
}

//裁剪字符串
func substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}