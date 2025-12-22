package netbasics

import (
    "io"
    "net"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
)

func TestHTTPHello(t *testing.T) {
    mux := http.NewServeMux()
    mux.HandleFunc("/hello", Hello)
    srv := httptest.NewServer(Logging(mux))
    defer srv.Close()
    resp, err := http.Get(srv.URL + "/hello")
    if err != nil {
        t.Fatalf("http get: %v", err)
    }
    b, _ := io.ReadAll(resp.Body)
    resp.Body.Close()
    if string(b) == "" {
        t.Fatalf("empty body")
    }
}

func TestTCPEcho(t *testing.T) {
    ln, err := net.Listen("tcp", "127.0.0.1:0")
    if err != nil { t.Fatalf("listen: %v", err) }
    addr := ln.Addr().String()
    done := make(chan struct{})
    go func() {
        for {
            c, err := ln.Accept()
            if err != nil { close(done); return }
            go func(conn net.Conn){
                defer conn.Close()
                writeLine(conn, "echo:start")
            }(c)
        }
    }()
    c, err := net.DialTimeout("tcp", addr, time.Second)
    if err != nil { t.Fatalf("dial: %v", err) }
    defer c.Close()
    resp, err := readLine(c, 1024)
    if err != nil { t.Fatalf("read: %v", err) }
    if resp == "" { t.Fatalf("empty resp") }
    ln.Close()
    <-done
}

func TestUDP(t *testing.T) {
    pc, err := net.ListenPacket("udp", "127.0.0.1:0")
    if err != nil { t.Fatalf("listen: %v", err) }
    addr := pc.LocalAddr().String()
    go func(){
        buf := make([]byte, 256)
        n, a, _ := pc.ReadFrom(buf)
        pc.WriteTo([]byte("pong:"+string(buf[:n])), a)
    }()
    c, err := net.DialTimeout("udp", addr, time.Second)
    if err != nil { t.Fatalf("dial: %v", err) }
    defer c.Close()
    c.Write([]byte("ping"))
    b := make([]byte, 256)
    n, err := c.Read(b)
    if err != nil { t.Fatalf("read: %v", err) }
    if string(b[:n]) == "" { t.Fatalf("empty") }
    pc.Close()
}
