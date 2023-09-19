import FakeWebSocket from './fake-socket.js'

export function CreateWebSocket() {
	const url = new URL(window.location.href)
	console.log(url.host)
	let socket: WebSocket
	try {
		socket = new FakeWebSocket(`http://${url.host}:80/sock/connect-http`)
	} catch (e) {
		socket = new WebSocket(`ws://${url.host}:80/sock/connect`)
		console.error('got error', e)
	}
	return socket
}
