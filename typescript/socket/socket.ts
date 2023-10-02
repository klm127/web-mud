import FakeWebSocket from './fake-socket.js'

export function CreateWebSocket() {
	const url = new URL(window.location.href)
	console.log(url.host)
	let socket: WebSocket
	try {
		socket = new WebSocket(`ws://${url.host}:80/sock/connect`)
	} catch (e) {
		socket = new FakeWebSocket(`http://${url.host}:80/sock/connect-http`)
		console.error('got error', e)
	}
	return socket 
}
