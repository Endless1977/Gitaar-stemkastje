// Stemkastje launcher voor Windows: serveert index.html op localhost en opent
// de standaardbrowser, zodat de microfoon werkt. Stopt zodra het tabblad dicht is.
package main

import (
	_ "embed"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

//go:embed index.html
var page string

const heartbeat = `<script>setInterval(()=>fetch("/ping").catch(()=>{}),2000);fetch("/ping").catch(()=>{});</script>`

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		os.Exit(1)
	}
	url := fmt.Sprintf("http://%s/", ln.Addr().String())
	html := strings.Replace(page, "</body>", heartbeat+"</body>", 1)

	var mu sync.Mutex
	last := time.Now()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, html)
	})
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		last = time.Now()
		mu.Unlock()
		w.WriteHeader(204)
	})
	go http.Serve(ln, mux)

	exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()

	for {
		time.Sleep(3 * time.Second)
		mu.Lock()
		idle := time.Since(last)
		mu.Unlock()
		if idle > 20*time.Second {
			os.Exit(0)
		}
	}
}
