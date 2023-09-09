import { CreateWebSocket } from './socket/socket.js'
import { Terminal } from './terminal/Terminal.js'

function RunGame() {
	const terminal = new Terminal(document.body)
	var socket: WebSocket
	try {
		socket = CreateWebSocket()
	} catch (e) {
		terminal.localError('Failed to connect socket! ' + e)
		return
	}
	socket.onopen = (event) => {
		terminal.enableInput()
		terminal.localMessage('Socket connected.')
	}

	socket.onmessage = (event) => {
		// terminal.serverMessage("Server: " +  event.data)
		terminal.parseServerMessage(event.data)
		terminal.enableInput()
	}

	socket.onclose = (event) => {
		terminal.serverMessage('Server connection closed.')
		terminal.disableInput()
	}

	terminal.onInput = (inputString) => {
		terminal.disableInput()
		socket.send(inputString)
	}
}

RunGame()
