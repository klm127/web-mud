import { Dom } from '../util/dom.js'
import {
	ScrollerToBottom,
	StickScrollerOnInterval,
} from '../util/stickScroller.js'

type ServMessage = {
	Parts: ServMessagePart[]
}
type ServMessagePart = {
	Text?: string
	Color?: string
	NewLine?: number
	Indent?: number
	Link?: string
	Css?: string[]
}

export class Terminal {
	el: HTMLDivElement
	messageArea: HTMLDivElement
	outerScrollArea: HTMLElement
	inputArea: HTMLDivElement
	input: HTMLInputElement
	onInput: (inputString: string) => void
	constructor(parent: HTMLElement) {
		this.el = Dom.el('div', 'terminal')
		this.outerScrollArea = Dom.el('div')
		this.outerScrollArea.id = 'terminal-message-scroller'
		this.messageArea = Dom.el('div')
		this.outerScrollArea.append(this.messageArea)
		this.messageArea.id = 'terminal-message-area'
		this.inputArea = Dom.el('div', ['terminal-input'])
		this.input = Dom.el('input', ['terminal-input'])
		this.input.type = 'text'
		this.inputArea.append(this.input)
		this.el.append(this.outerScrollArea, this.inputArea)
		this.input.addEventListener('keydown', this.getInputKeypressCallback())
		this.onInput = () => {}
		parent.append(this.el)
	}

	disableInput() {
		this.input.disabled = true
	}
	enableInput() {
		this.input.disabled = false
		this.input.focus()
	}

	getInputKeypressCallback() {
		const me = this
		return (e: KeyboardEvent) => {
			if (e.key == 'Enter') {
				me.inputSubmit()
			}
		}
	}

	inputSubmit() {
		const val = this.input.value
		if (val.length > 0) {
			this.localMessage(val)
			this.onInput(val)
			this.input.value = ''
		}
	}

	addMessageEl(el: HTMLElement) {
		if (this.messageArea.firstChild != null) {
			this.messageArea.insertBefore(el, this.messageArea.firstChild)
		} else {
			this.messageArea.append(el)
		}
		ScrollerToBottom(this.outerScrollArea)
	}

	localError(s: string) {
		const div = Dom.textEl('div', s, ['terminal-msg', 'terminal-local-error'])
		this.addMessageEl(div)
	}
	serverError(s: string) {
		const div = Dom.textEl('div', s, ['terminal-msg', 'terminal-server-error'])
		this.addMessageEl(div)
	}
	serverMessage(s: string) {
		const div = Dom.textEl('div', s, ['terminal-msg', 'terminal-server-msg'])
		this.addMessageEl(div)
	}
	parseServerMessage(s: string) {
		const msg = JSON.parse(s) as ServMessage
		if (msg.Parts.length == 0) {
			this.addMessageEl(Dom.textEl('div', '* empty message *'))
		} else {
			const div = Dom.el('div')
			div.classList.add('terminal-server-msg')
			this.parseMessagePartRecurse(msg, 0, div)
			this.addMessageEl(div)
		}
	}
	parseMessagePartRecurse(
		s: ServMessage,
		index: number,
		lastParent: HTMLElement
	) {
		const part = s.Parts[index]
		const partEl = Dom.el('span', ['terminal-server-msg'], {
			display: 'inline-block',
		})
		if (part.Color) {
			partEl.style.color = part.Color
		}
		if (part.Indent) {
			partEl.style.paddingLeft = `${part.Indent}px`
		}
		if (part.Link) {
			partEl.style.textDecoration = 'underline'
			partEl.style.cursor = 'pointer'
			partEl.classList.add('link-el')
			let my = this
			partEl.addEventListener('click', () => {
				my.input.value = part.Link!
			})
		}
		if (part.Text) {
			partEl.textContent = part.Text
		}
		if (part.Css) {
			for (let c of part.Css) {
				partEl.classList.add(c)
			}
		}
		lastParent.append(partEl)
		let nextParent = lastParent
		if (part.NewLine) {
			for (let i = 0; i < part.NewLine; i++) {
				nextParent = Dom.textEl('div', ' ')
				nextParent.style.minHeight = '14px'
				lastParent.append(nextParent)
			}
		}
		index++
		if (index < s.Parts.length) {
			this.parseMessagePartRecurse(s, index, nextParent)
		}
	}
	localMessage(s: string) {
		const div = Dom.textEl('div', s, ['terminal-msg', 'terminal-local-msg'])
		this.addMessageEl(div)
	}
}
