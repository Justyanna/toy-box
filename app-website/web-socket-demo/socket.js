let socket

export function connect(url, options = {}) {
    const {
        onClose = () => { console.log("WS connection closed"); },
        onError = (e) => { console.log("WS error occurred:", e); }
    } = options

    console.log(`Openning websocket connected to ws://${url}/rooms/ws`);
    socket = new WebSocket(`ws://${url}/rooms/ws`)

    socket.onopen = () => {
        sendEvent("authenticate", {
            id: "u-1234",
            secret: "zaq1@WSX"
        })
    }

    socket.onmessage = (event) => routeEvent(event.data)
    socket.onerror = onError
    socket.onclose = onClose
}

export function sendEvent(type, payload) {
    socket.send(JSON.stringify({ type, payload }))
}

export function routeEvent(eventData) {
    const event = JSON.parse(eventData)

    switch (event.type) {
        case "test_response":
            console.log(event.payload)
            break;

        default:
            console.log("Unsupported message type:", event.type);
            console.log("Ignored data:", event.payload);
            break;
    }
}
