
const socket = new WebSocket("ws://" + window.location.hostname + ":" + window.location.port + "/ws")

socket.addEventListener("open", () => {
    log("Connected to server")
})

socket.addEventListener("message", (event) => {
    log(event.data)
})

socket.addEventListener("close", (event) => {
    log("Disconnected from server: " + event.reason)

    document.querySelectorAll("#volume-control button, #track-control button").forEach(button => {
        button.disabled = true
        button.classList.add("disabled")
    })
})

function sendCode(code) {
    if (socket.readyState === WebSocket.OPEN) {
        try {
            socket.send(code)
        } catch (e) {
            log(e)
        }

    } else {
        log("Socket not ready " + socket.readyState)
    }
}

function log(message) {
    console.log(message)

    const p = document.createElement('p')
    p.textContent = message
    result.appendChild(p)
}

document.getElementById('playpause').onclick = () => sendCode(100)
document.getElementById('next').onclick = () => sendCode(101)
document.getElementById('prev').onclick = () => sendCode(102)
document.getElementById('vol-up').onclick = () => sendCode(103)
document.getElementById('vol-down').onclick = () => sendCode(104)
