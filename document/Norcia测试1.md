## 前言
Norcia 是一个简单的静态博客框架，通过 JavaScript 使用 Ajax 请求本地的 markdown 文档文件来获得博客数据，并将其动态的展现在 Html 页面当中。

## 项目结构
 - 根目录下的 HTML \ CSS \ JavaScript 文件
 - `document` 文件夹用来存放博文 markdown 文件
 - `config.json` 作为静态博客的配置文件以及博客文章索引,该文件在初次设定好个人信息后可由 Norcia 程序自动更新与维护, 详情请看下文介绍


## **`config.json`** 说明
### 自动更新
运行 `Norcia` 程序就可以自动依照 `document` 文件夹里面的 markdown 文件的修改, 而自动维护、更新

### 格式和说明

	{
		"head": "博客名称",
		"introduce": "博客介绍",
		"github": "github地址",
		"mail": "邮箱",
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

## **`Norcia`** 程序说明
### **多平台支持**
Norcia 为以下三个平台提供打包好的二进制程序

 - `Norcia_win_amd64` 适用于 64 位 windows 系统
 - `Norcia_drawin_amd64` 适用于 64 位 Mac OS 系统
 - `Norcia_linux_amd64` 适用于 64 位 linux 系统

### **运行命令**
运行的时候在工程目录根目录下运行下面命令就可以了

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

### **开启预览服务**
**`Norcia`** 程序支持本地开启一个静态文件服务器以供博客效果预览，只需要在运行时候带上 `p` 参数就可以了，例如

    ./Norcia_linux_amd64 -p
 
然后就可以从 **`http://localhost:8666/index.html`** 预览到博客效果了

## **`Norcia.js`** 说明
`Norcia.js` 是　`Norcia`的　js 工具包，封装好了 Norcia 一些前端需要用到的函数，使用示例如下
    
    let norciaConfig;
    // 载入 config
    ajaxGetConfig(function (config) {
        norciaConfig = config;
        console.log(norciaConfig);
    });


更多的示例请查看 Norcia JavaScript API 的说明

## **编译Norcia**
感谢 `Golang` 的语言特性，我们能够很简单的编译 `Norcia`， 并且 `Norcia` 还提供了一个简单的 `Bash` 脚本，帮助你直接编译出各个平台适用的二进制文件
运行一下命令，将会自动构建适用于 Linux, Mac OS, Windows 的程序

    sh build.sh
    
如果编译成功，你将会将会有以下显示

        _   _                _       
       | \ | | ___  _ __ ___(_) __ _ 
       |  \| |/ _ \| '__/ __| |/ _  |
       | |\  | (_) | | | (__| | (_| |
       |_| \_|\___/|_|  \___|_|\__,_|
    ----------build script--------------
    
    start build for linux_amd64
    build for linux_amd64 successful
    ------------------------------------
    
    start build for drawin_amd64
    build for drawin_amd64 successful 
    ------------------------------------
    
    start build for window_amd64
    build for window_amd64 successful 
    ------------------------------------
    
    finish build in 
             ./Norcia_linux_amd64
             ./Norcia_drawin_amd64
             ./Norcia_win_amd64.exe
    ------------------------------------
