let socket;

function initWebSocket() {
    socket = new WebSocket("ws://localhost:8080/ws");

    socket.onopen = () => {
        console.log("WebSocket connected");
    };

    socket.onmessage = (event) => {
        console.log("Message from server:", event.data);
        const outputElement = document.getElementById('output');
        if (outputElement) {
            outputElement.innerText = event.data;
        }
    };

    socket.onerror = (error) => {
        console.error("WebSocket error:", error);
    };

    socket.onclose = () => {
        console.log("WebSocket closed. Reconnecting...");
        setTimeout(initWebSocket, 1000); // Reconnect after 1 second
    };
}

function sendMessage() {
    const messageInput = document.getElementById('messageInput');
    if (messageInput && socket && socket.readyState === WebSocket.OPEN) {
        const message = messageInput.value;
        if (message) {
            socket.send(message);
        }
    }
}

export { initWebSocket, sendMessage };