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
    sendQueue;
    protocol;
    readyState;
    url;
    interval_id;
    jwt;
    CONNECTING;
    OPEN;
    CLOSING;
    CLOSED;
    constructor(url) {
        console.log('constructing fake web socket');
        this.binaryType = 'arraybuffer';
        this.bufferedAmount = 0;
        this.extensions = '';
        this.closeCbs = new Set();
        this.errorCbs = new Set();
        this.messageCbs = new Set();
        this.openCbs = new Set();
        this.sendQueue = [];
        this.protocol = 'fake-http-socket';
        this.readyState = 0;
        this.url = url;
        this.CONNECTING = 0;
        this.OPEN = 1;
        this.CLOSING = 2;
        this.CLOSED = 3;
        this.jwt = '';
        this.poll = this.poll.bind(this);
        this.interval_id = setInterval(this.poll, 1000);
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
        this.sendQueue.push({
            type: 'message',
            data: data,
        });
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
    emitMessage(s) {
        const e = new MessageEvent('server-message-event', {
            data: JSON.stringify(s),
        });
        for (let f of this.messageCbs) {
            f(e);
        }
    }
    emitError(s) {
        const e = new MessageEvent('server-error-event', {
            data: s,
        });
        for (let f of this.errorCbs) {
            f(e);
        }
    }
    emitOpen(s) {
        const e = new MessageEvent('server-open-event', {
            data: s,
        });
        for (let f of this.openCbs) {
            f(e);
        }
    }
    emitClose(s) {
        const e = new CloseEvent('server-close-event');
        for (let f of this.closeCbs) {
            f(e);
        }
    }
    async poll() {
        console.log('polling');
        try {
            const response = await fetch(this.url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    sdcmcook: this.jwt,
                },
                body: JSON.stringify(this.sendQueue),
            });
            const result = await response.json();
            this.sendQueue = [];
            this.handleServerMessages(result);
            const ac = response.headers.get('sdcmcook');
            if (ac != null) {
                this.jwt = ac;
            }
        }
        catch (e) {
            console.error('fake socket polling error: ', e);
        }
    }
    handleServerMessages(servMessages) {
        for (let m of servMessages) {
            console.log(m);
            if (m.event == 'open') {
                this.emitOpen(m.data);
            }
            else if (m.event == 'close') {
                this.emitClose(m.data);
            }
            else if (m.event == 'message') {
                this.emitMessage(JSON.parse(m.data));
            }
            else if (m.event == 'error') {
                this.emitMessage(m.data);
            }
            else {
                console.error('unknown message event', m.data, 'from server');
            }
        }
    }
}
