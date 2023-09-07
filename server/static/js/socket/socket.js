export function CreateWebSocket() {
    const url = new URL(window.location.href);
    console.log(url.host);
    const socket = new WebSocket(`ws://${url.host}:80/sock/connect`);
    socket.onopen = (event) => {
        console.log("Got socket onOpen event.", event);
    };
    socket.onmessage = (event) => {
        console.log("Socket got message.", event);
    };
    socket.onclose = (event) => {
        console.log("Socket got close event.", event);
    };
    return socket;
}
