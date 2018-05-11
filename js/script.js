function loadNavigation() {
    document.getElementsByTagName("body")[0].innerHTML +=
        ``;
}

/**
 * 调用示例
 *
 *           ajax_post(
 *               login_API,                      // 接口地址
 *               [                               // 接口参数
 *                   ["n",mail],
 *                   ["p",pw],
 *               ],
 *               function (responseText) {       //请求成功的回调方法
     *
     *               },
 *               function (responseText) {       //请求失败成功的回调方法
     *
     *               }
 *           );
 *
 *
 * @param url 接口的 URL
 * @param params 参数
 * @param success_callback  成功的回调方法
 * @param fail_callback     失败的回调方法
 */
function ajax_post(url,params,success_callback,fail_callback) {
    let formData = new FormData();
    // 将参数塞到 formData 里面
    for (let i=0;i<params.length;i++){
        formData.append(params[i][0],params[i][1])
    }
    let xhr = new XMLHttpRequest();
    xhr.open('POST', url, true);
    xhr.send(formData);
    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                success_callback(xhr.responseText)
            } else {
                fail_callback(xhr.status,xhr.responseText)
            }
        }
    }
}


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
 if (r != null) return unescape(r[2]); return null; //返回参数值
}
