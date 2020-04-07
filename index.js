$ = document
var ws = new WebSocket(`wss://${location.host}/socketgateway`)
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