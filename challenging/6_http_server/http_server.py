import http.server

class HelloHandler(http.server.BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header("Content-type", "text/html")
        self.end_headers()
        self.wfile.write(bytes("<html><head><title>Hello</title></head>", "utf-8"))
        self.wfile.write(bytes("<body><p>Hello, world!</p>", "utf-8"))
        self.wfile.write(bytes("</body></html>", "utf-8"))


server = http.server.HTTPServer(("", 8080), HelloHandler)
server.serve_forever()