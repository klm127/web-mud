import { CreateWebSocket } from "./socket/socket.js"

const el = document.getElementById('body')

if (el) {
	el.textContent = '!!😊'
}


const socket = CreateWebSocket();