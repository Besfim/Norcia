package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

/**

配置文件 json 格式

	{
		"head": "博客名称",
		"introduce": "博客介绍",
		"github": "github地址",
		"weibo": "weibo地址",
		"tag": "文章标签",
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
func main() {
	var jsonString = readJsonFile()
	fmt.Println(jsonString)
}

func readJsonFile() string{
	inputFile, inputError := os.Open("config.json")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
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
	head		string
	introduce 	string
	github		string
	tag			string
	article		Article
}

type Article struct {
	title		string
	tag			string
	create		string
	update		string
}
// 遍历 document 文件夹
// 获取配置文件
// 对比两者
// 查看是否有更新或者有不同
// 如果有不同的话那么就更新 config.json
// 生成简介

// 生成搜索索引