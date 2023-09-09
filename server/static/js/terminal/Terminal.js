import { Dom } from '../util/dom.js';
import { ScrollerToBottom, } from '../util/stickScroller.js';
export class Terminal {
    el;
    messageArea;
    outerScrollArea;
    inputArea;
    input;
    onInput;
    constructor(parent) {
        this.el = Dom.el('div', 'terminal');
        this.outerScrollArea = Dom.el('div');
        this.outerScrollArea.id = 'terminal-message-scroller';
        this.messageArea = Dom.el('div');
        this.outerScrollArea.append(this.messageArea);
        this.messageArea.id = 'terminal-message-area';
        this.inputArea = Dom.el('div');
        this.input = Dom.el('input');
        this.input.type = 'text';
        this.inputArea.append(this.input);
        this.el.append(this.outerScrollArea, this.inputArea);
        this.input.addEventListener('keydown', this.getInputKeypressCallback());
        this.onInput = () => { };
        parent.append(this.el);
    }
    disableInput() {
        this.input.disabled = true;
    }
    enableInput() {
        this.input.disabled = false;
        this.input.focus();
    }
    getInputKeypressCallback() {
        const me = this;
        return (e) => {
            if (e.key == 'Enter') {
                me.inputSubmit();
            }
        };
    }
    inputSubmit() {
        const val = this.input.value;
        if (val.length > 0) {
            this.localMessage(val);
            this.onInput(val);
            this.input.value = '';
        }
    }
    addMessageEl(el) {
        if (this.messageArea.firstChild != null) {
            this.messageArea.insertBefore(el, this.messageArea.firstChild);
        }
        else {
            this.messageArea.append(el);
        }
        ScrollerToBottom(this.outerScrollArea);
    }
    localError(s) {
        const div = Dom.textEl('div', s, ['terminal-msg', 'terminal-local-error']);
        this.addMessageEl(div);
    }
    serverError(s) {
        const div = Dom.textEl('div', s, ['terminal-msg', 'terminal-server-error']);
        this.addMessageEl(div);
    }
    serverMessage(s) {
        const div = Dom.textEl('div', s, ['terminal-msg', 'terminal-server-msg']);
        this.addMessageEl(div);
    }
    parseServerMessage(s) {
        const msg = JSON.parse(s);
        if (msg.Parts.length == 0) {
            this.addMessageEl(Dom.textEl('div', '* empty message *'));
        }
        else {
            const div = Dom.el('div');
            this.parseMessagePartRecurse(msg, 0, div);
            this.addMessageEl(div);
        }
    }
    parseMessagePartRecurse(s, index, lastParent) {
        const part = s.Parts[index];
        const partEl = Dom.el('span', undefined, { display: 'inline-block' });
        if (part.Color) {
            partEl.style.color = part.Color;
        }
        if (part.Indent) {
            partEl.style.paddingLeft = `${part.Indent}px`;
        }
        if (part.Link) {
            partEl.style.textDecoration = 'underline';
            partEl.style.cursor = 'pointer';
            let my = this;
            partEl.addEventListener('click', () => {
                my.input.value = part.Link;
            });
        }
        if (part.Text) {
            partEl.textContent = part.Text;
        }
        if (part.Css) {
            for (let c of part.Css) {
                partEl.classList.add(c);
            }
        }
        lastParent.append(partEl);
        let nextParent = lastParent;
        if (part.NewLine) {
            for (let i = 0; i < part.NewLine; i++) {
                nextParent = Dom.textEl('div', ' ');
                nextParent.style.minHeight = '14px';
                lastParent.append(nextParent);
            }
        }
        index++;
        if (index < s.Parts.length) {
            this.parseMessagePartRecurse(s, index, nextParent);
        }
    }
    localMessage(s) {
        const div = Dom.textEl('div', s, ['terminal-msg', 'terminal-local-msg']);
        this.addMessageEl(div);
    }
}
