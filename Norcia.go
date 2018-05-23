package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"encoding/json"
	"io/ioutil"
	"strings"
	"sort"
	"strconv"
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
				"mini": "文章缩略前300个字"
			},
			{
				"title": "文章标题",
				"tag": "文章标签",
				"create": "创作日期",
				"update": "更新日期"
				"mini": "文章缩略前300个字"
			}
		]
	}
*/

// md文件存储文件夹
const docDirName string = "document/"
// Norcia 多语言支持
var language = "cn"
//const languageMap = initLanguageMap()
func main() {
	//var oldArticleNum = 0
	printHeader()
	updateNum := 0
	createNum := 0
	//var deleteNum = 0
	blogconfig := parseConfigJson(readFileToString("config.json"))
	//oldArticleNum = len(blogconfig.Articles)
	//读取 document 文件
	files, _ := ioutil.ReadDir("document")
	var articles []Article
	for _, file := range files {
		var fileName = file.Name()                                 //文件名
		var title = strings.Replace(fileName, ".md", "", -1)       //title
		var docContent = readFileToString(docDirName + fileName)   //文档内容
		var miniDoc = substr(cleanMarkdownDoc(docContent), 0, 300) //文档缩略
		var updateTime = file.ModTime()                            //更新时间
		var temp Article
		articleFromConfig, successFlag := getArticleFromConfigByTitle(title, blogconfig)
		if successFlag == 1 {
			//如果能够找到旧的文件
			//最后修改时间没变
			if articleFromConfig.Update == substr(updateTime.String(), 0, 16) {
				temp = articleFromConfig
			} else {
				temp = Article{
					Title:  title,
					Tag:    articleFromConfig.Tag,
					Update: substr(updateTime.String(), 0, 16),
					Create: articleFromConfig.Create,
					Mini:   miniDoc,
				}
				updateNum++
			}
		} else if successFlag == 0 {
			//如果无法找到旧的文件,证明文件时新建的!
			temp = Article{
				Title:  title,
				Tag:    inputDocumentsTag(title,blogconfig),
				Update: substr(updateTime.String(), 0, 16),
				Create: substr(updateTime.String(), 0, 16),
				Mini:   miniDoc,
			}
			createNum++
		}
		articles = append(articles, temp)
	}
	sort.Sort(articleList(articles))
	blogconfig.Articles = articles
	outputNewBlogConfig(blogconfig)
	fmt.Printf("\n更新了 %d 个文档, 并且创建了 %d 个文档 \n\n", updateNum , createNum)
}

//用户输入标签，或者是从旧的标签里面选一个
func inputDocumentsTag(title string,config BlogConfig) string{
	tagMap := make(map[int64]string)
	var tagCount int64
	tagCount = 0
	for _,article:= range config.Articles{
		tagsTemp := strings.Split(article.Tag,",")
		for _,tag := range tagsTemp {
			flag := true
			for _,tagHaveTemp := range tagMap{
				if tagHaveTemp == tag {
					flag = false
					break
				}
			}
			if flag {
				tagMap[tagCount] = tag
				tagCount++
			}
		}
	}
	fmt.Println("\n以下为已有的标签及编号：")
	for i,tagTemp := range tagMap{
		fmt.Println("\t",i,".",tagTemp)
	}
	fmt.Printf("请输入文章 ' %s ' 的新标签名称，或者输入已有标签的序号，多个输入之间使用空格分隔 : \n",title)
	reader := bufio.NewReader(os.Stdin)
	input, _, _ := reader.ReadLine()
	res := ""
	inputTemp := strings.Split(string(input)," ")
	for i,tag := range inputTemp{
		flag,num :=  isInt(tag)
		if flag {
			if tagMap[num] == "" {
				res += tag
			}else {
				res += tagMap[num]
			}
		}else {
			res += tag
		}
		if i != len(inputTemp)-1 {
			res += ","
		}
	}
	return res
}

func isInt(str string)( bool,int64){
	num,err := strconv.ParseInt(str,0,64)
	if err != nil {
		return false,-1
	}else {
		return true,num
	}
}

func readFileToString(fileName string) string {
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
		if readerError == io.EOF {
			return res
		}
	}
	return res
}

type BlogConfig struct {
	Head      string    `json:"head"`
	Introduce string    `json:"introduce"`
	Github    string    `json:"github"`
	Mail      string    `json:"mail"`
	Articles  []Article `json:"articles"`
}

type Article struct {
	Title  string `json:"title"`
	Tag    string `json:"tag"`
	Create string `json:"create"`
	Update string `json:"update"`
	Mini   string `json:"mini"`
}

//排序 Article
type articleList []Article

func (I articleList) Len() int {
	return len(I)
}
func (I articleList) Less(i, j int) bool {
	return I[i].Create > I[j].Create
}
func (I articleList) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}

// 解析 配置 Json 的函数
func parseConfigJson(jsonString string) BlogConfig {
	var config BlogConfig
	json.Unmarshal([]byte(jsonString), &config)
	return config
}

//将 BlogConfig 对象变成 string
func generateConfigJson(config BlogConfig) string {
	str, _ := json.MarshalIndent(config, "", "\t")
	return string(str)
}

//将新的 json 写到 config.json 里面去
func outputNewBlogConfig(config BlogConfig) {
	writeStringToFile(generateConfigJson(config), "config.json")
}

func writeStringToFile(outputString string, fileName string) {
	outputFile, outputError := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
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

func getArticleFromConfigByTitle(title string, config BlogConfig) (Article, int) {
	for _, article := range config.Articles {
		if article.Title == title {
			return article, 1
		}
	}
	return nullArticle, 0
}

//去除 markdown 文档里面的 markdown 符号
func cleanMarkdownDoc(mkDoc string) string {
	mkDoc = strings.Replace(mkDoc, "#", "", -1)
	mkDoc = strings.Replace(mkDoc, "**", "", -1)
	mkDoc = strings.Replace(mkDoc, "-", "", -1)
	mkDoc = strings.Replace(mkDoc, "+", "", -1)
	mkDoc = strings.Replace(mkDoc, ">", "", -1)
	mkDoc = strings.Replace(mkDoc, "-", "", -1)
	mkDoc = strings.Replace(mkDoc, "|", "", -1)
	mkDoc = strings.Replace(mkDoc, "\r", " ", -1)
	mkDoc = strings.Replace(mkDoc, "\n", " ", -1)
	return mkDoc
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

func makeMap(lang []string) map[string]string{
	mapTemp := map[string]string{
		"zh":lang[0],
		"en":lang[1],
	}
	return mapTemp
}

func printHeader() {
	fmt.Println(" _   _                _       ")
	fmt.Println("| \\ | | ___  _ __ ___(_) __ _ ")
	fmt.Println("|  \\| |/ _ \\| '__/ __| |/ _` |")
	fmt.Println("| |\\  | (_) | | | (__| | (_| |")
	fmt.Println("|_| \\_|\\___/|_|  \\___|_|\\__,_|")
}

func getStringsLan(languageMap map[string]map[string]string,key string) string{
	return languageMap[key][language]
}


func initLanguageMap() map[string]map[string]string{
	var languageMap map[string]map[string]string
	languageMap = make(map[string](map[string]string))
	//载入多语言字符串
	languageMap["final"] = makeMap([]string{
		"更新了 %d 个文档, 并且创建了 %d 个文档",
		"update %d document(s), and create %d documents(s)",
	})
	languageMap["key_select"] = makeMap([]string{
		"请输入文章 ' %s ' 标签或者选择标签，多个标签之间使用 ',' 分隔",
		"Please enter a tag or select a tag for ' %s ' . Separate multiple tags by ','",
	})
	return languageMap
}