!
    function(f, c) {
        if ("object" == typeof exports && "object" == typeof module) module.exports = c();
        else if ("function" == typeof define && define.amd) define([], c);
        else {
            c = c();
            for (var e in c)("object" == typeof exports ? exports : f)[e] = c[e]
        }
    }(window, function() {
        return function(f) {
            function c(b) {
                if (e[b]) return e[b].exports;
                var a = e[b] = {
                    i: b,
                    l: !1,
                    exports: {}
                };
                return f[b].call(a.exports, a, a.exports, c), a.l = !0, a.exports
            }
            var e = {};
            return c.m = f, c.c = e, c.d = function(b, a, d) {
                c.o(b, a) || Object.defineProperty(b, a, {
                    enumerable: !0,
                    get: d
                })
            }, c.r = function(b) {
                "undefined" != typeof Symbol && Symbol.toStringTag && Object.defineProperty(b, Symbol.toStringTag, {
                    value: "Module"
                });
                Object.defineProperty(b, "__esModule", {
                    value: !0
                })
            }, c.t = function(b, a) {
                if ((1 & a && (b = c(b)), 8 & a) || 4 & a && "object" == typeof b && b && b.__esModule) return b;
                var d = Object.create(null);
                if (c.r(d), Object.defineProperty(d, "default", {
                    enumerable: !0,
                    value: b
                }), 2 & a && "string" != typeof b) for (var g in b) c.d(d, g, function(h) {
                    return b[h]
                }.bind(null, g));
                return d
            }, c.n = function(b) {
                var a = b && b.__esModule ?
                    function() {
                        return b.
                            default
                    } : function() {
                        return b
                    };
                return c.d(a, "a", a), a
            }, c.o = function(b, a) {
                return Object.prototype.hasOwnProperty.call(b, a)
            }, c.p = "", c(c.s = 0)
        }([function(f, c, e) {
            function b(a) {
                var d = a.pingTimeout,
                    g = a.pongTimeout,
                    h = a.reconnectTimeout,
                    k = a.pingMsg,
                    l = a.repeatLimit;
                this.opts = {
                    url: a.url,
                    pingTimeout: void 0 === d ? 15E3 : d,
                    pongTimeout: void 0 === g ? 1E4 : g,
                    reconnectTimeout: void 0 === h ? 2E3 : h,
                    pingMsg: void 0 === k ? "heartbeat" : k,
                    repeatLimit: void 0 === l ? null : l
                };
                this.ws = null;
                this.repeat = 0;
                this.onclose = function() {};
                this.onerror = function() {};
                this.onopen = function() {};
                this.onmessage = function() {};
                this.onreconnect = function() {};
                this.createWebSocket()
            }
            Object.defineProperty(c, "__esModule", {
                value: !0
            });
            b.prototype.createWebSocket = function() {
                try {
                    this.ws = new WebSocket(this.opts.url), this.initEventHandle()
                } catch (a) {
                    throw this.reconnect(), a;
                }
            };
            b.prototype.initEventHandle = function() {
                var a = this;
                this.ws.onclose = function() {
                    a.onclose();
                    a.reconnect()
                };
                this.ws.onerror = function() {
                    a.onerror();
                    a.reconnect()
                };
                this.ws.onopen = function() {
                    a.repeat = 0;
                    a.onopen();
                    a.heartCheck()
                };
                this.ws.onmessage = function(d) {
                    a.onmessage(d);
                    a.heartCheck()
                }
            };
            b.prototype.reconnect = function() {
                var a = this;
                0 < this.opts.repeatLimit && this.opts.repeatLimit <= this.repeat || this.lockReconnect || this.forbidReconnect || (this.lockReconnect = !0, this.repeat++, this.onreconnect(), setTimeout(function() {
                    a.createWebSocket();
                    a.lockReconnect = !1
                }, this.opts.reconnectTimeout))
            };
            b.prototype.send = function(a) {
                this.ws.send(a)
            };
            b.prototype.heartCheck = function() {
                this.heartReset();
                this.heartStart()
            };
            b.prototype.heartStart = function() {
                var a = this;
                this.forbidReconnect || (this.pingTimeoutId = setTimeout(function() {
                    a.ws.send(a.opts.pingMsg);
                    a.pongTimeoutId = setTimeout(function() {
                        a.ws.close()
                    }, a.opts.pongTimeout)
                }, this.opts.pingTimeout))
            };
            b.prototype.heartReset = function() {
                clearTimeout(this.pingTimeoutId);
                clearTimeout(this.pongTimeoutId)
            };
            b.prototype.close = function() {
                this.forbidReconnect = !0;
                this.heartReset();
                this.ws.close()
            };
            "undefined" != typeof window && (window.WebsocketHeartbeatJs = b);
            c.
                default = b
        }])
    });