/**
 * 简化写法
 * @param domId document id 参数
 */
function domId(domId) {
    return document.getElementById(domId)
}

/**
 * 获取url中的参数
 */
function getUrlParam(name) {
    let reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
    let r = window.location.search.substr(1).match(reg); //匹配目标参数
    if (r != null) return decodeURI(r[2]);
    return null; //返回参数值
}

/**
 * 将 Article 对象数据库数据与 index 页面卡片试图绑定
 * @param article
 * @returns {string}
 */
function bindIndexArticleCard(article) {
    let tags = article.tag;
    let tagHtml = bindArticleTags(tags,true);
    let head =
        `<div class="mdui-card mdui-typo mdui-m-t-2 mdui-m-b-2 mdui-hoverable gradient-wrapper">
        <div class="mdui-card-primary">
            <div class="mdui-card-primary-title ">
                <a href="javascript:void(0);" onclick="handleCardClick('blog.html?title=${article.title}')">${article.title}</a>
            </div>
            <div class="mdui-card-primary-subtitle">${article.create}</div>
        </div>
        <div class="mdui-card-content">
            ${article.mini} ......
        </div>
        <div class="mdui-card-actions mdui-m-t-2 mdui-m-b-2">
            ${tagHtml}
            <button class="mdui-btn mdui-btn-icon mdui-float-right"><i
                    class="mdui-icon material-icons">expand_more</i></button>
        </div>
    </div>`;

    return head;
}

function handleCardClick(uri) {
    window.location.href = uri;
    if (window.location.pathname.endsWith("index.html")){
        //记录浏览的位置
        sessionStorage.setItem("loadNum",articleCount);
        sessionStorage.setItem("scrollIndex",document.body.scrollTop);
    }
}

/**
 * 传入文章的标签，返回绑定好的视图
 * 输入的数据如下
 *  tag1,tag2,tag3,
 *
 * @param randomColor 是否开启随机颜色
 * @param tags
 */
function bindArticleTags(tags,randomColor) {
    let temp = tags.split(",");
    let tagHtml = "";
    let colors = [
        "mdui-color-deep-orange-400",
        "mdui-color-teal-400",
        "mdui-color-red-400",
        "mdui-color-lime",
        "mdui-color-pink-400",
        "mdui-color-blue-grey-400"
    ];
    let colorClass = "mdui-color-theme-accent";
    for (let i=0;i<temp.length;i++){
        if (randomColor){
            colorClass = colors[getRandom(0,colors.length-1)];
        }
        tagHtml += `<div class="mdui-chip ${colorClass}  mdui-m-x-1 none-text-transform mdui-text-color-white blog-tag">
                        <span class="mdui-chip-title">${temp[i]}  </span>
                    </div>`;
    }
    return tagHtml;
}

/**
 * 载入侧边导航栏
 * @param dom 载入的位置
 * @param close 是否默认关闭
 */
function loadNavigation(dom,close) {
    let closeClass = "";
    if (close){
        closeClass="mdui-drawer-close";
    }
    domId(dom).innerHTML +=
        `<div class="mdui-container mdui-appbar-with-toolbar">
            <div class="mdui-drawer	mdui-drawer-full-height mdui-color-white ${closeClass}" id="drawer">
                <div class="navigation-bar-header">
                    <img class="mdui-img-circle logo mdui-m-l-3 mdui-m-t-4" src="img/logo.jpg"/>
                    <div class="mdui-typo mdui-m-l-4">
                        <h4 class="mdui-text-color-white mdui-m-t-2 name-mail-line-height"><strong id="blogName">
                        My Blog</strong></h4>
                        <p class="mdui-text-color-white mdui-m-t-0 name-mail-line-height" id="mail"></p>
                    </div>
                </div>
                <ul class="mdui-list">
                    <br>
                <li class="mdui-list-item mdui-ripple" onclick="">
                    <i class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-black">home</i>
                    <div class="mdui-list-item-content" onclick="gotoIndex()">首页 Home</div>
                </li>

                <br>
                <li class="mdui-list-item mdui-ripple" onclick="gotoArchives()">
                    <i class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-black">archive</i>
                    <div class="mdui-list-item-content">归档 Archives</div>
                </li>

                <br>
                <li class="mdui-list-item mdui-ripple" onclick="gotoTags()">
                    <i class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-black">attach_file</i>
                    <div class="mdui-list-item-content">标签 Tag</div>
                </li>

                <br>
                <li class="mdui-subheader">链接</li>
                <li class="mdui-list-item mdui-ripple" onclick="gotoGithub()">
                    <i class="mdui-list-item-icon material-icons mdui-text-color-grey">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 496 512"><path d="M165.9 397.4c0 2-2.3 3.6-5.2 3.6-3.3.3-5.6-1.3-5.6-3.6 0-2 2.3-3.6 5.2-3.6 3-.3 5.6 1.3 5.6 3.6zm-31.1-4.5c-.7 2 1.3 4.3 4.3 4.9 2.6 1 5.6 0 6.2-2s-1.3-4.3-4.3-5.2c-2.6-.7-5.5.3-6.2 2.3zm44.2-1.7c-2.9.7-4.9 2.6-4.6 4.9.3 2 2.9 3.3 5.9 2.6 2.9-.7 4.9-2.6 4.6-4.6-.3-1.9-3-3.2-5.9-2.9zM244.8 8C106.1 8 0 113.3 0 252c0 110.9 69.8 205.8 169.5 239.2 12.8 2.3 17.3-5.6 17.3-12.1 0-6.2-.3-40.4-.3-61.4 0 0-70 15-84.7-29.8 0 0-11.4-29.1-27.8-36.6 0 0-22.9-15.7 1.6-15.4 0 0 24.9 2 38.6 25.8 21.9 38.6 58.6 27.5 72.9 20.9 2.3-16 8.8-27.1 16-33.7-55.9-6.2-112.3-14.3-112.3-110.5 0-27.5 7.6-41.3 23.6-58.9-2.6-6.5-11.1-33.3 2.6-67.9 20.9-6.5 69 27 69 27 20-5.6 41.5-8.5 62.8-8.5s42.8 2.9 62.8 8.5c0 0 48.1-33.6 69-27 13.7 34.7 5.2 61.4 2.6 67.9 16 17.7 25.8 31.5 25.8 58.9 0 96.5-58.9 104.2-114.8 110.5 9.2 7.9 17 22.9 17 46.4 0 33.7-.3 75.4-.3 83.6 0 6.5 4.6 14.4 17.3 12.1C428.2 457.8 496 362.9 496 252 496 113.3 383.5 8 244.8 8zM97.2 352.9c-1.3 1-1 3.3.7 5.2 1.6 1.6 3.9 2.3 5.2 1 1.3-1 1-3.3-.7-5.2-1.6-1.6-3.9-2.3-5.2-1zm-10.8-8.1c-.7 1.3.3 2.9 2.3 3.9 1.6 1 3.6.7 4.3-.7.7-1.3-.3-2.9-2.3-3.9-2-.6-3.6-.3-4.3.7zm32.4 35.6c-1.6 1.3-1 4.3 1.3 6.2 2.3 2.3 5.2 2.6 6.5 1 1.3-1.3.7-4.3-1.3-6.2-2.2-2.3-5.2-2.6-6.5-1zm-11.4-14.7c-1.6 1-1.6 3.6 0 5.9 1.6 2.3 4.3 3.3 5.6 2.3 1.6-1.3 1.6-3.9 0-6.2-1.4-2.3-4-3.3-5.6-2z"></path></svg>
                    </i>
                    <div class="mdui-list-item-content">Github</div>
                </li>
                </ul>
            </div>
        </div>
        `;
}

//载入博客的一些配置信息
function loadBlogConfig(norcia) {
    document.title = norcia.head;
    domId("mail").innerHTML = norcia.mail;
    domId("blogName").innerHTML = norcia.head;
    console.log(norcia);
    github = norcia.github;
}

function getRandom(begin,end){
    return Math.floor(Math.random()*(end-begin + 1)) + begin;
}

function gotoIndex() {
    window.location.href="index.html";
}

function gotoArchives() {
    window.location.href="archives.html"
}

function gotoTags() {
    window.location.href="tags.html"
}

let github = "";
function gotoGithub() {
    if (github !== ""){
        window.location.href = github;
    }
}

let nextPage = "";
let prePage = "";
function gotoPrePage() {
    if (prePage !== ""){
        window.location.href = "blog.html?title="+prePage;
    }
}
function gotoNextPage() {
    if (nextPage !== ""){
        window.location.href = "blog.html?title="+ nextPage;
    }
}
function transDateToCH(monthDate) {
    let temp = monthDate.split("-")
    if (temp[1].startsWith("0")){
        temp[1] = temp[1].substring(1,temp[1].length)
    }
    return temp[0]+" 年 "+temp[1]+" 月";
}

function bindArchiveArticle(article){
    return `<div class="mdui-col-xs-12 mdui-col-md-6 mdui-col-sm-12 mdui-m-t-1 mdui-p-b-1 gradient-wrapper">
            <div class="mdui-card mdui-typo mdui-hoverable">
                <div class="mdui-p-x-2">
                    <h4 class="mdui-typo">
                        <a href="blog.html?title=${article.title}">${article.title}</a>
                    </h4>
                    <div class="mdui-card-primary-subtitle">${article.create}</div>
                </div>
                <div class="mdui-card-actions mdui-m-b-2">
                    ${bindArticleTags(article.tag,true)}
                </div>
            </div>
        </div>`
}

function showShareCode() {
    mdui.dialog({
        title: '',
        content: `<img class="mdui-center mdui-m-t-2" src = ${jrQrcode.getQrBase64(window.location.href)} style="height:150px;width:150px" >
                        <div class="mdui-text-center mdui-m-t-2">扫一扫二维码即可分享</div>`,
        buttons: [
            {
                text: '取消'
            }
        ]
    });
}