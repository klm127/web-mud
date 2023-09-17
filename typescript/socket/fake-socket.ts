type onCloseCb = (ev: CloseEvent) => any
type onErrorCb = (ev: Event) => any
type onMessageCb = (ev: MessageEvent) => any
type onOpenCb = (ev: Event) => any
/**
 *
 * FakeWebSocket is used when, for some reason, a real web socket can't be created. It uses HTTP and intervals to poll for updates.
 *
 */
class FakeWebSocket implements WebSocket {
	binaryType: BinaryType
	bufferedAmount: number
	extensions: string

	closeCbs: Set<onCloseCb>
	errorCbs: Set<onErrorCb>
	messageCbs: Set<onMessageCb>
	openCbs: Set<onOpenCb>

	protocol: string
	readyState: number
	url: string
	readonly CONNECTING: 0
	readonly OPEN: 1
	readonly CLOSING: 2
	readonly CLOSED: 3

	constructor(url: string) {
		this.binaryType = 'arraybuffer'
		this.bufferedAmount = 0
		this.extensions = ''
		this.closeCbs = new Set()
		this.errorCbs = new Set()
		this.messageCbs = new Set()
		this.openCbs = new Set()
		this.protocol = 'fake-http-socket'
		this.readyState = 0
		this.url = url
		this.CONNECTING = 0
		this.OPEN = 1
		this.CLOSING = 2
		this.CLOSED = 3
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
		throw new Error('Method not implemented.')
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

	emitStringMessage(s: any) {
		const e = new MessageEvent('server-message-event', {
			data: JSON.stringify(s),
		})
		for (let f of this.messageCbs) {
			f(e)
		}
	}

	emitErrorMessage(s: string) {
		const e = new MessageEvent('server-error-event', {
			data: s,
		})
		for (let f of this.errorCbs) {
			f(e)
		}
	}

	async poll() {
		let url = this.url
		let me = this
		fetch(url)
			.then((r) => {
				return r.json()
			})
			.then((r) => {
				for (let i of r) {
					me.emitStringMessage(i)
				}
			})
			.catch((r) => {
				me.emitErrorMessage(r)
			})
	}
}
