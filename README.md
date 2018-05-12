# Norcia
一个简单的静态博客框架

## 项目结构
 - 根目录下的 HTML \ CSS \ JavaScript 文件
 - `norcia` 文件夹存放是用来维护 `config.json` 的一个程序
 - `document` 文件夹用来存放博文 markdown 文件
 - config.json 作为静态博客的配置文件以及博客文章索引,该文件在初次设定好个人信息后可由 Norcia 程序自动更新与维护, 详情请看下文介绍
 
## config.json 自动更新
### 更新
运行 Norcia 程序就可以自动依照 document 文件夹里面的 markdown 文件的修改, 而自动维护更新 config.json 索引了

Norcia 为以下三个平台提供打包好的二进制程序

 - `Norcia_win_amd64` 适用于 64 位 windows 系统
 - `Norcia_drawin_amd64` 适用于 64 位 Mac OS 系统
 - `Norcia_linux_amd64` 适用于 64 位 windows 系统

### 格式和说明

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

### 使用示例

在工程目录根目录下运行下面命令就可以了
 - `Linux` 平台

        ./Norcia_linux_amd64

 - `Mac OS` 平台

        ./Norcia_drawin_amd64
        
 - `Windows` 平台
 
        ./Norcia_win_amd64.exe
        

如果运行正确的话, 会看到下面提示

    update 0 document(s), and create 0 documents(s)
