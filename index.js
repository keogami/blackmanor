(function ($) {
    $.querySelectorAll(".pane").forEach((it) => {
        it.style.zIndex = it.dataset.layer
    })
    var mainCanvas = $.querySelector("#main-canvas-con canvas")
    var canvasWidth = mainCanvas.width = innerWidth
    var canvasHeight = mainCanvas.height = innerHeight
    var mainContext = mainCanvas.getContext("2d")

    mainContext.textAlign = "center"
    mainContext.font = "3em sans-serif"
    mainContext.fillStyle = "white"
    mainContext.beginPath()
    mainContext.fillText("Welcome to Black Manor", canvasWidth/2, canvasHeight/3)
    mainContext.closePath()

    $.querySelector('.shifter#chat').addEventListener('click', function () {
        $.querySelector('.pane#chat-con').classList.toggle("active")
    })

    var wsProtocol = (location.protocol === "http:") ? "ws:" : "wss:"
    var socketGateway = `${wsProtocol}//${location.host}/socketgateway`
    var ws = new WebSocket(socketGateway)
    var messageView = $.querySelector('#message-view')
    ws.onmessage = function (message) {
        try {
            var mess = "Partner: " + JSON.parse(message.data).message
            messageView.innerHTML += mess
            messageView.innerHTML += "<hr>"
        } catch (e) {
            $.querySelector('#speaker').innerHTML = "Butler"
            $.querySelector('#dialog p').innerText = message.data
        }
    }
    ws.onclose = function () {
        messageView.innerHTML += "Your Partner has left the scene."
        messageView.innerHTML += "<hr>"
    }
    ws.onopen = function () {
        console.log("open")
        $.querySelector('input[type=text]').addEventListener("keypress", function (ev) {
            if (ev.keyCode === 13) {
                var val = {
                    message: $.querySelector('input[type=text]').value
                }
                if (val && val.message.length > 0) {
                    ws.send(JSON.stringify(val))
                    messageView.innerHTML += ("You: " + val.message)
                    $.querySelector('input[type=text]').value = ""
                    messageView.innerHTML += "<hr>"
                }
            }
        })
    }
}(document))