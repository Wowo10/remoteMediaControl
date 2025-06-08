
const socket = new WebSocket("ws://" + window.location.hostname + ":" + window.location.port + "/ws");

// Connection opened
socket.addEventListener("open", () => {
    log("Connected to server")
});

// Handle messages from server
socket.addEventListener("message", (event) => {
    log(event.data)
});

// Send something manually (e.g., a code) later
function sendCode(code) {
    if (socket.readyState === WebSocket.OPEN) {
        try {
            socket.send(code);
        } catch (e) {
            log(e)
        }

    } else {
        log("Socket not ready " + socket.readyState)
    }
}

function log(message) {
    console.log(message);

    const p = document.createElement('p');
    p.textContent = message;
    result.appendChild(p);
}

document.getElementById('playpause').onclick = () => sendCode(100);
document.getElementById('next').onclick = () => sendCode(101);
document.getElementById('prev').onclick = () => sendCode(102);
