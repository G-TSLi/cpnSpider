/**
 * Created by viruser on 2017/8/15.
 */

var wsUri = "ws://" + location.hostname + ":" + location.port + "/ws";
var ws = null;

if ('WebSocket' in window) {
    ws = new WebSocket(wsUri);
}

function selectMode(m) {
    switch (m) {
        case offline:
            $("#js_mode").text("单机模式");
            $("#step1 .js_port").hide();
            $("#step1 .js_ip").hide();
            $("#mode").val(offline);
            break;
        case server:
            $("#js_mode").text("服务端模式");
            $("#step1 .js_ip").hide();
            $("#step1 .js_port").show();
            $("#mode").val(server);
            break;
        case client:
            $("#js_mode").text("客户端模式");
            $("#step1 .js_ip").show();
            $("#step1 .js_port").show();
            $("#mode").val(client);
            break;
        default:
            $("#js_mode").text("运行模式");
            $("#step1 .js_port").hide();
            $("#step1 .js_ip").hide();
            $("#mode").val(unset);
            return;
    }
    $("#init").removeAttr("disabled");
}

// 执行入口
function home() {
    switch (parseInt($("#mode").val())) {
        case offline:
            $("#js_mode").text("单机模式");
            break;
        case server:
            $("#js_mode").text("服务端模式");
            break;
        case client:
            $("#js_mode").text("客户端模式");
            break;
        default:
            $("#init").attr("disabled", "disabled");
            return;
    }
}


// 发送api
ws.onsend = function(data) {
    var dataStr = JSON.stringify(data);
    ws.send(dataStr);
    console.log("send: " + dataStr);
}


// 按模式启动
function Open(operate) {
    $("#init").text(" 开  启 …").css({
        "background-color": "#286090",
        "border-color": "#204d74"
    }).attr("disabled", "disabled");

    var formJson = {
        'operate': operate,
        'mode': document.step1.elements['mode'].value,
        'port': document.step1.elements['port'].value,
        'ip': document.step1.elements['ip'].value,
    };
    console.log(formJson)
    ws.onsend(formJson);
    return false;
}