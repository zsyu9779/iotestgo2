package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

/*
	go run . -mode=server -proto=http
	go run . -mode=server -proto=udp
	go run . -mode=client -proto=udp
*/
func main() {
	mode := flag.String("mode", "server", "Mode: server or client")
	proto := flag.String("proto", "http", "Protocol: http or udp")
	addr := flag.String("addr", ":8080", "Address to listen on or connect to")
	flag.Parse()

	switch *mode {
	case "server":
		runServer(*proto, *addr)
	case "client":
		runClient(*proto, *addr)
	default:
		fmt.Println("Unknown mode. Use -mode=server or -mode=client")
		os.Exit(1)
	}
}

func runServer(proto, addr string) {
	switch proto {
	case "http":
		log.Printf("Starting HTTP server on %s", addr)
		mux := http.NewServeMux()
		mux.HandleFunc("/hello", Hello)
		mux.HandleFunc("/slow", Slow)
		// Apply middleware
		handler := Logging(mux)
		if err := http.ListenAndServe(addr, handler); err != nil {
			log.Fatalf("HTTP server failed: %v", err)
		}
	case "udp":
		log.Printf("Starting UDP server on %s", addr)
		if err := RunUDPServer(addr); err != nil {
			log.Fatalf("UDP server failed: %v", err)
		}
	default:
		log.Fatalf("Unknown protocol for server: %s", proto)
	}
}

func runClient(proto, addr string) {
	switch proto {
	case "udp":
		if err := RunUDPClient(addr); err != nil {
			log.Fatalf("UDP client failed: %v", err)
		}
	default:
		log.Fatalf("Unknown protocol for client or not implemented: %s", proto)
	}
}
