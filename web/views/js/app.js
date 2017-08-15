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

// 接收api
ws.onmessage = function(m) {
    var data = JSON.parse(m.data)
    console.log(data);

    switch (data.operate) {
        // 初始化运行参数
        case "init":
            if (!data.initiative) {
                // window.location.href = window.location.href;
                location = location;
                return
            };
            // 设置当前运行模式
            mode = data.mode;
            // 打开软件界面
            var index = layer.open({
                type: 1,
                title: data.title,
                content: Html(data),
                // area: ['300px', '195px'],
                maxmin: false,
                scrollbar: false,
                move: false,
            });

            layer.full(index);
            $(".layui-layer-close1").attr("title", "退出").click(function() {
                Close();
            });

            $("#init").text(" 开  启 ").css({
                "background-color": "#337ab7",
                "border-color": "#2e6da4"
            });

            break;

        // 任务开始通知
        case "run":
            $("#btn-run").text("Stop").attr("data-type", "stop");

            if (data.mode == offline) {
                $("#btn-run").text("Stop").attr("data-type", "stop").addClass("btn-danger").removeClass("btn-primary");
                $("#btn-pause").text("Pause").removeAttr("disabled").show();
            };
            break;

        // 任务结束通知
        case "stop":
            $("#btn-pause").hide();
            $("#btn-run").text("Run").attr("data-type", "run").removeAttr("disabled");
            if (data.mode == offline) {
                $("#btn-run").text("Run").attr("data-type", "run").addClass("btn-primary").removeClass("btn-danger");
            };
            break;

        // 暂停与恢复
        case "pauseRecover":
            if ($("#btn-pause").text() == "Pause") {
                $("#btn-pause").text("Go on...").addClass("btn-info").removeClass("btn-warning");
            } else {
                $("#btn-pause").text("Pause").addClass("btn-warning").removeClass("btn-info");
            };
            break;

        case "exit":
            layer.closeAll();
            selectMode(unset);
    }
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