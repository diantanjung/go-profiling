package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"time"
)

func Sample(w http.ResponseWriter, r *http.Request) {
	ProcessBigArray()
	Pingpong()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hi!"))
}

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/hi", Sample)

	// Add Profiling Endpoint
	handler.HandleFunc("/debug/pprof/", pprof.Index)
	handler.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	handler.HandleFunc("/debug/pprof/profile", pprof.Profile)
	handler.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	handler.HandleFunc("/debug/pprof/trace", pprof.Trace)

	log.Println("Server running on port: 9090")
	log.Fatal(http.ListenAndServe(":9090", handler))

}

func ProcessBigArray() {
	for i := 0; i < 500; i++ {
		arr := bigArray()
		if arr == nil {
			fmt.Println("Array is Nil")
		}
	}
}

func bigArray() *[]int {
	s := make([]int, 1000000)
	return &s
}

func Pingpong(){
	table := make(chan *ball)
	go player("dian", table)
	go player("dias", table)
	table <- new(ball)
	time.Sleep(100* time.Millisecond)
	<- table
}

type ball struct{
	hits int
}

func player(name string, table chan *ball)  {
	for{
		ball := <- table
		ball.hits++
		log.Println(name, "hit the ball", ball.hits)
		time.Sleep(50* time.Millisecond)
		table <- ball
	}
}