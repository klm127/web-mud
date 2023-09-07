
import { Dom } from "../util/dom.js";



export class Terminal {
    el: HTMLDivElement
    messageArea: HTMLDivElement
    inputArea: HTMLDivElement
    input: HTMLInputElement
    onInput: (inputString: string)=>void
    constructor(parent: HTMLElement) {
        this.el = Dom.el("div", "terminal")
        this.messageArea = Dom.el("div")
        this.inputArea = Dom.el("div")
        this.input = Dom.el("input")
        this.input.type = "text"
        this.inputArea.append(this.input)
        this.el.append(this.messageArea, this.inputArea)
        this.input.addEventListener("keydown", this.getInputKeypressCallback())
        this.onInput = ()=>{}
        parent.append(this.el)
    }

    disableInput() {
        this.input.disabled = true;
    }
    enableInput() {
        this.input.disabled = false;
        this.input.focus()
    }

    getInputKeypressCallback() {
        const me = this;
        return (e:KeyboardEvent)=>{
            if(e.key == "Enter") {
                me.inputSubmit()
            }
        }
    }

    inputSubmit() {
        const val = this.input.value
        if(val.length > 0) {
            this.localMessage(val)
            this.onInput(val)
            this.input.value = ""
        }
    }

    localError(s: string) {
        const div = Dom.textEl("div", s, ["terminal-msg", "terminal-local-error"])
        this.messageArea.append(div)
    }
    serverError(s: string) {
        const div = Dom.textEl("div", s, ["terminal-msg", "terminal-server-error"])
        this.messageArea.append(div)
    }
    serverMessage(s: string) {
        const div = Dom.textEl("div", s, ["terminal-msg", "terminal-server-msg"])
        this.messageArea.append(div)

    }
    localMessage(s: string) {
        const div = Dom.textEl("div", s, ["terminal-msg", "terminal-local-msg"])
        this.messageArea.append(div)

    }


}