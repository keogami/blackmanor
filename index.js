var ws = new WebSocket("ws://localhost:8080/socketgateway")
ws.onmessage = console.log
ws.onopen = function () {
    console.log("open")
    ws.send("hi")
}