export function CreateWebSocket() {
    const url = new URL(window.location.href);
    console.log(url.host);
    const socket = new WebSocket(`ws://${url.host}:80/sock/connect`);
    return socket;
}
