# Norcia
一个简单的静态博客框架

## 项目结构
 - 根目录下的 HTML \ CSS \ JavaScript 文件
 - `document` 文件夹用来存放博文 markdown 文件
 - config.json 作为静态博客的配置文件以及博客文章索引,该文件在初次设定好个人信息后可由 Norcia 程序自动更新与维护, 详情请看下文介绍
 
## config.json 自动更新
### 更新
运行 Norcia 程序就可以自动依照 document 文件夹里面的 markdown 文件的修改, 而自动维护更新 config.json 索引了

Norcia 为以下三个平台提供打包好的二进制程序

 - `Norcia_win_amd64` 适用于 64 位 windows 系统
 - `Norcia_drawin_amd64` 适用于 64 位 Mac OS 系统
 - `Norcia_linux_amd64` 适用于 64 位 linux 系统

### 格式和说明

	{
    	"head": "Html Title 和博客题目",
    	"introduce": "博客简介",
    	"github": "Github 链接",
    	"mail": "邮箱",
    	"articles": [
    		{
    			"title": "文章标题1",
    			"tag": "tag1,tag2",
    			"create": "2018-05-11 21:52",
    			"update": "2018-05-12 09:41",
    			"mini": "文章的250字缩略..."
    		},
    		{
    			"title": "文章标题2",
    			"tag": "tag1,tag3",
    			"create": "2018-05-11 18:15",
    			"update": "2018-05-12 09:41",
    			"mini": "文章的250字缩略..."
    		}
    	]
    }

### 使用示例

在工程目录根目录下运行下面命令就可以了
 - `Linux` 平台

        ./Norcia_linux_amd64

 - `Mac OS` 平台

        ./Norcia_drawin_amd64
        
 - `Windows` 平台
 
        ./Norcia_win_amd64.exe
        

如果运行正确的话, 会看到下面提示

     _   _                _       
    | \ | | ___  _ __ ___(_) __ _ 
    |  \| |/ _ \| '__/ __| |/ _` |
    | |\  | (_) | | | (__| | (_| |
    |_| \_|\___/|_|  \___|_|\__,_|
    
    更新了 0 个文档, 并且创建了 0 个文档 
    