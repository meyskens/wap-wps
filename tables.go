package main

var verbs = map[byte]string{
	0x01: "CONNECT",
	0x04: "REPLY",
	0x05: "DISCONNECT",
	0x06: "PUSH",
	0x08: "SUSPEND",
	0x09: "RESUME",
	0x40: "GET",
	0x60: "POST",
	0x61: "PUSH",
}

var headers = map[byte]string{
	0x00: "Accept",
	0x01: "Accept-Charset",
	0x02: "Accept-Encoding",
	0x03: "Accept-Language",
	0x04: "Accept-Ranges",
	0x0B: "Content-Encoding",
	0x0C: "Content-Language",
	0x0D: "Content-Length",
	0x10: "Content-Range",
	0x11: "Content-Type",
	0x12: "Date",
	0x1F: "Proxy-Authenticate",
	0x20: "Proxy-Authorization",
	0x26: "Server",
	0x29: "User-Agent",
}

var fileTypes = map[byte]string{
	0x01: "text/*",
	0x02: "text/html",
	0x03: "text/plain",
	0x08: "text/vnd.wap.wml",
	0x09: "text/vnd.wap.wmlscript",
	0x12: "application/x-www-form-urlencoded",
	0x14: "application/vnd.wap.wmlc",
	0x15: "application/vnd.wap.wmlscriptc",
	0x16: "application/vnd.wap.wta-eventc",
	0x17: "application/vnd.wap.uaprof",
	0x18: "application/vnd.wap.wtls-ca-certificate",
	0x19: "application/vnd.wap.wtls-user-certificate",
	0x1D: "image/gif",
	0x1E: "image/jpeg",
	0x21: "image/vnd.wap.wbmp",
	0x27: "application/xml",
	0x28: "text/xml",
	0x29: "application/vnd.wap.wbxml",
}

var charsets = map[byte]string{
	0x03: "iso-8859-1",
}
