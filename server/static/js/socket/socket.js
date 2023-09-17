export function CreateWebSocket() {
    const url = new URL(window.location.href);
    console.log(url.host);
    let socket;
    try {
        socket = new WebSocket(`ws://${url.host}:80/sock/connect`);
    }
    catch (e) {
        socket = new FakeWebSocket(`http://${url.host}:80/sock/connect-fake`);
    }
    return socket;
}
