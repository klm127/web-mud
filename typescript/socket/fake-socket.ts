type onCloseCb = (ev: CloseEvent) => any
type onErrorCb = (ev: Event) => any
type onMessageCb = (ev: MessageEvent) => any
type onOpenCb = (ev: Event) => any

type transmitMessage = {
	type: 'open' | 'message' | 'close'
	data: string | ArrayBufferLike | Blob | ArrayBufferView
}

type receiveMessage = {
	event: 'open' | 'message' | 'close' | 'error'
	data: string
}
/**
 *
 * FakeWebSocket is used when, for some reason, a real web socket can't be created. It uses HTTP and intervals to poll for updates.
 *
 */
export default class FakeWebSocket implements WebSocket {
	binaryType: BinaryType
	bufferedAmount: number
	extensions: string

	closeCbs: Set<onCloseCb>
	errorCbs: Set<onErrorCb>
	messageCbs: Set<onMessageCb>
	openCbs: Set<onOpenCb>

	sendQueue: transmitMessage[]

	protocol: string
	readyState: number
	url: string
	interval_id: number
	jwt: string

	readonly CONNECTING: 0
	readonly OPEN: 1
	readonly CLOSING: 2
	readonly CLOSED: 3

	constructor(url: string) {
		console.log('constructing fake web socket')
		this.binaryType = 'arraybuffer'
		this.bufferedAmount = 0
		this.extensions = ''
		this.closeCbs = new Set()
		this.errorCbs = new Set()
		this.messageCbs = new Set()
		this.openCbs = new Set()
		this.sendQueue = []
		this.protocol = 'fake-http-socket'
		this.readyState = 0
		this.url = url
		this.CONNECTING = 0
		this.OPEN = 1
		this.CLOSING = 2
		this.CLOSED = 3
		this.jwt = ''
		this.poll = this.poll.bind(this)
		this.interval_id = setInterval(this.poll, 1000)
	}

	set onclose(cb: onCloseCb | null) {
		if (cb == null) {
			this.closeCbs = new Set()
			return
		}
		this.closeCbs.add(cb)
	}
	get onclose(): onCloseCb | null {
		let cbs = this.closeCbs
		if (cbs.size == 0) {
			return null
		}
		return (ev: CloseEvent) => {
			for (let i of cbs) {
				i(ev)
			}
		}
	}

	set onerror(cb: onErrorCb | null) {
		if (cb == null) {
			this.errorCbs = new Set()
			return
		}
		this.errorCbs.add(cb)
	}
	get onerror(): onErrorCb | null {
		let cbs = this.errorCbs
		if (cbs.size == 0) return null
		return (ev: Event) => {
			for (let f of cbs) {
				f(ev)
			}
		}
	}

	set onmessage(cb: onMessageCb | null) {
		if (cb == null) {
			this.messageCbs = new Set()
			return
		}
		this.messageCbs.add(cb)
	}
	get onmessage(): onMessageCb | null {
		let cbs = this.messageCbs
		if (cbs.size == 0) return null
		return (ev: MessageEvent<any>) => {
			for (let f of cbs) {
				f(ev)
			}
		}
	}

	set onopen(cb: onOpenCb | null) {
		if (cb == null) {
			this.openCbs = new Set()
			return
		}
		this.openCbs.add(cb)
	}
	get onopen(): onOpenCb | null {
		let cbs = this.openCbs
		if (cbs.size == 0) return null
		return (ev: Event) => {
			for (let f of cbs) {
				f(ev)
			}
		}
	}

	close(code?: number | undefined, reason?: string | undefined): void {
		throw new Error('Method not implemented.')
	}
	send(data: string | ArrayBufferLike | Blob | ArrayBufferView): void {
		this.sendQueue.push({
			type: 'message',
			data: data,
		})
	}
	addEventListener<K extends keyof WebSocketEventMap>(
		type: K,
		listener: (ev: WebSocketEventMap[K]) => any,
		options?: boolean | AddEventListenerOptions | undefined
	): void {
		if (type == 'close') {
			this.onclose = listener as any
		} else if (type == 'error') {
			this.onerror = listener as any
		} else if (type == 'message') {
			this.onmessage = listener as any
		} else if (type == 'open') {
			this.onmessage = listener as any
		}
	}

	dispatchEvent(event: Event): boolean {
		throw new Error('Method not implemented.')
	}

	removeEventListener<K extends keyof WebSocketEventMap>(
		type: K,
		listener: (ev: WebSocketEventMap[K]) => any,
		options?: boolean | EventListenerOptions | undefined
	) {
		if (type == 'close') {
			this.closeCbs.delete(listener as any)
		} else if (type == 'error') {
			this.errorCbs.delete(listener as any)
		} else if (type == 'message') {
			this.messageCbs.delete(listener as any)
		} else if (type == 'open') {
			this.openCbs.delete(listener as any)
		}
	}

	emitMessage(s: any) {
		const e = new MessageEvent('server-message-event', {
			data: JSON.stringify(s),
		})
		for (let f of this.messageCbs) {
			f(e)
		}
	}

	emitError(s: string) {
		const e = new MessageEvent('server-error-event', {
			data: s,
		})
		for (let f of this.errorCbs) {
			f(e)
		}
	}

	emitOpen(s: string) {
		const e = new MessageEvent('server-open-event', {
			data: s,
		})
		for (let f of this.openCbs) {
			f(e)
		}
	}

	emitClose(s: string) {
		const e = new CloseEvent('server-close-event')
		for (let f of this.closeCbs) {
			f(e)
		}
	}

	async poll() {
		console.log('polling')
		try {
			const response = await fetch(this.url, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					sdcmcook: this.jwt,
				},
				body: JSON.stringify(this.sendQueue),
			})
			const result = await response.json()
			this.sendQueue = []
			this.handleServerMessages(result)
			const ac = response.headers.get('sdcmcook')
			if (ac != null) {
				this.jwt = ac
			}
		} catch (e) {
			console.error('fake socket polling error: ', e)
		}
	}

	handleServerMessages(servMessages: Array<receiveMessage>) {
		for (let m of servMessages) {
			console.log(m)
			if (m.event == 'open') {
				this.emitOpen(m.data)
			} else if (m.event == 'close') {
				this.emitClose(m.data)
			} else if (m.event == 'message') {
				this.emitMessage(JSON.parse(m.data))
			} else if (m.event == 'error') {
				this.emitMessage(m.data)
			} else {
				console.error('unknown message event', m.data, 'from server')
			}
		}
	}
}
