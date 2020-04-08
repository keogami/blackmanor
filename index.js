$ = document
var wsProtocol =  (location.protocol === "https:") ? "wss:" : "ws:"
var socketGatewayPath = `${wsProtocol}//${location.host}/socketgateway`;
var ws = new WebSocket(socketGatewayPath)
var box = $.querySelector("div.box")
ws.onmessage = function (message) {
    box.innerText += "\n"
    try {
        var mess = "He Said: " + JSON.parse(message.data).message
        box.innerText += mess
    } catch (e) {
        box.innerText += message.data
    }
}
ws.onclose = function () {
    box.innerText += "\n"
    box.innerText += "The connection was closed."
}
ws.onopen = function () {
    console.log("open")
    $.querySelector('input[type=button]').addEventListener("click", function () {
        var val = {
            message: $.querySelector('input[type=text]').value
        }
        if (val && val.message.length > 0) {
            ws.send(JSON.stringify(val))
            box.innerText += "\n"
            box.innerText += ("You Said: " + val.message)
            $.querySelector('input[type=text]').value = ""
        }
    })
}