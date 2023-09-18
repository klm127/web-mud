/**
 *
 * FakeWebSocket is used when, for some reason, a real web socket can't be created. It uses HTTP and intervals to poll for updates.
 *
 */
export default class FakeWebSocket {
    binaryType;
    bufferedAmount;
    extensions;
    closeCbs;
    errorCbs;
    messageCbs;
    openCbs;
    protocol;
    readyState;
    url;
    CONNECTING;
    OPEN;
    CLOSING;
    CLOSED;
    constructor(url) {
        this.binaryType = 'arraybuffer';
        this.bufferedAmount = 0;
        this.extensions = '';
        this.closeCbs = new Set();
        this.errorCbs = new Set();
        this.messageCbs = new Set();
        this.openCbs = new Set();
        this.protocol = 'fake-http-socket';
        this.readyState = 0;
        this.url = url;
        this.CONNECTING = 0;
        this.OPEN = 1;
        this.CLOSING = 2;
        this.CLOSED = 3;
    }
    set onclose(cb) {
        if (cb == null) {
            this.closeCbs = new Set();
            return;
        }
        this.closeCbs.add(cb);
    }
    get onclose() {
        let cbs = this.closeCbs;
        if (cbs.size == 0) {
            return null;
        }
        return (ev) => {
            for (let i of cbs) {
                i(ev);
            }
        };
    }
    set onerror(cb) {
        if (cb == null) {
            this.errorCbs = new Set();
            return;
        }
        this.errorCbs.add(cb);
    }
    get onerror() {
        let cbs = this.errorCbs;
        if (cbs.size == 0)
            return null;
        return (ev) => {
            for (let f of cbs) {
                f(ev);
            }
        };
    }
    set onmessage(cb) {
        if (cb == null) {
            this.messageCbs = new Set();
            return;
        }
        this.messageCbs.add(cb);
    }
    get onmessage() {
        let cbs = this.messageCbs;
        if (cbs.size == 0)
            return null;
        return (ev) => {
            for (let f of cbs) {
                f(ev);
            }
        };
    }
    set onopen(cb) {
        if (cb == null) {
            this.openCbs = new Set();
            return;
        }
        this.openCbs.add(cb);
    }
    get onopen() {
        let cbs = this.openCbs;
        if (cbs.size == 0)
            return null;
        return (ev) => {
            for (let f of cbs) {
                f(ev);
            }
        };
    }
    close(code, reason) {
        throw new Error('Method not implemented.');
    }
    send(data) {
        throw new Error('Method not implemented.');
    }
    addEventListener(type, listener, options) {
        if (type == 'close') {
            this.onclose = listener;
        }
        else if (type == 'error') {
            this.onerror = listener;
        }
        else if (type == 'message') {
            this.onmessage = listener;
        }
        else if (type == 'open') {
            this.onmessage = listener;
        }
    }
    dispatchEvent(event) {
        throw new Error('Method not implemented.');
    }
    removeEventListener(type, listener, options) {
        if (type == 'close') {
            this.closeCbs.delete(listener);
        }
        else if (type == 'error') {
            this.errorCbs.delete(listener);
        }
        else if (type == 'message') {
            this.messageCbs.delete(listener);
        }
        else if (type == 'open') {
            this.openCbs.delete(listener);
        }
    }
    emitStringMessage(s) {
        const e = new MessageEvent('server-message-event', {
            data: JSON.stringify(s),
        });
        for (let f of this.messageCbs) {
            f(e);
        }
    }
    emitErrorMessage(s) {
        const e = new MessageEvent('server-error-event', {
            data: s,
        });
        for (let f of this.errorCbs) {
            f(e);
        }
    }
    async poll() {
        let url = this.url;
        let me = this;
        fetch(url)
            .then((r) => {
            return r.json();
        })
            .then((r) => {
            for (let i of r) {
                me.emitStringMessage(i);
            }
        })
            .catch((r) => {
            me.emitErrorMessage(r);
        });
    }
}
