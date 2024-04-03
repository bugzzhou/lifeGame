var index = 0
var interval = 1000

var ip = "http://192.168.3.200";
var port = "9292";
var apiEndpoint = "/fetch/testPost";

var full_ip = ip + ":" + port;

function init() {
    intervalStr = document.getElementById('interval').value;
    interval = parseInt(intervalStr)
}

function get2Slice() {
    fetch(full_ip+"/fetch/autofresh")
        .then((response) => response.json())
        .then((data) => {
            res = getTableBy2Slice(data)
            document.getElementById("change").innerHTML = res;
        })
}

function reset() {
    fetch(full_ip+"/fetch/reset")
        .then((response) => response.json())
        .then((data) => {
            res = getTableBy2Slice(data)
            document.getElementById("change").innerHTML = res;
        })
}

function getTableBy2Slice(twoSlice) {
    var res = ""
    res += "<table width='150' height='150' border='1'>"
    for (var i = 0; i < 10; i++) {
        res += "<tr>"
        for (var j = 0; j < 10; j++) {
            if (twoSlice[i][j] != 1)
                res += "<td bgcolor='" + "white" + "'></td>"
            else
                res += "<td bgcolor='" + "black" + "'></td>"
        }
        res += "</tr>"
    }
    res += "</table>"
    return res
}

function sendNum() {
    // 获取输入框中的数字  
    number = document.getElementById('number').value;

    // 创建一个包含数字的Map对象  
    var data = { "number": parseInt(number) }
    JSON.stringify(data)
    fetch(full_ip+"/fetch/testPost", {
        method: "POST",
        body: JSON.stringify(data),

    })
        .then(res => res.json())
        .then(data => {
            autoFresh(data)
        })
}

//将三维数组，进行循环输出
//n个   [10][10]int， n为迭代的次数
function autoFresh(social) {
    var index = 0
    const timer1 = setInterval(() => {
        if (index >= social.length) {
            clearInterval(timer1);  // 当index超过了social的范围，就停止定时器
            return
        }

        var tmpSocial = getTableBy2Slice(social[index])
        console.log(tmpSocial)
        document.getElementById("change11").innerHTML = tmpSocial
        index++


    }, interval);  // 每过三秒钟执行一次
}


// function countIndex() {
//     index++
//     console.log(index)
// }

// function getMap(i) {
//     fetch("http://127.0.0.1:9292/getMap")
//         .then((response) => response.json())
//         .then((data) => {
//             console.log(data)
//             if (index >= data["result"].length) {
//                 document.getElementById("change").innerHTML = "数据已经遍历完成，数组长度为：" + data["result"].length + ",最后一个数为：" + data["result"][data["result"].length - 1];
//             }
//             else {
//                 document.getElementById("change").innerHTML = data["result"][index];
//             }
//         })
// }
