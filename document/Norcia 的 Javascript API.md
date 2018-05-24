## 前言
Norcia 将 JavaScript 对本地的 document 和 config.json 的读取封装成了一个 JavaScript 文件 `norcia.js`，引入这个文件之后，通过封装好 ajax 的一些函数，就能够拿到本地的文档数据和配置文件数据了，方便 Norcia 开发新主题时候为将静态的页面转换为动态页面。

## 封装对象
### Article
对应 config.json 当中 articles 数组里面的每篇文章信息

 - `title`
 - `tag`
 - `create`
 - `update`
 - `mini`
 - `content`


### NorciaConfig
对应 config 当中的各项配置，articles 是 Article 对象数组

 - `head`
 - `introduce`
 - `github`
 - `mail`
 - `articles`

## 封装函数
### **`ajaxGetConfig(callback)`**
 - 说明:
 该函数用以获取 `config.json` 配置文件并封装成 `NorciaConfig` 对象，回调函数将在成功获取 config 对象之后被调用
 - 示例:

        let norciaConfig;
        // 载入 config
        ajaxGetConfig(function (config) {
            norciaConfig = config;
            console.log(norciaConfig);
        });

### **`Article().loadContent(callback)`**
 - 说明:
 该函数用来获取获取文章的详细内容
 - 示例:

        norciaConfig.articles[i].loadContent(function (mdDocument){
            console.log(mdDocument);
        });


