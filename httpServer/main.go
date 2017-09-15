package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Start Http Server...v2"))
}
func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye Http Server...v2"))
}

func main() {
	server := &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 2,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	route := http.NewServeMux()
	route.Handle("/", &myHandler{})
	route.HandleFunc("/bye", sayBye)
	server.Handler = route

	go func() {
		<-quit

		if err := server.Close(); err != nil {
			log.Fatal("Close Server:", err)
		}
	}()

	log.Println("Server start...v3")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("Server closed under request")
		} else {
			log.Print(err)
			log.Fatal("Server closed unexpeceted")
		}
	}

	log.Println("Server exit")
	// log.Fatal(server.ListenAndServe())

}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server start...v2"))
}
