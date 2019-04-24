package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

// ActivateProfile runs the profiling endpoint
func ActivateProfile() {
	log.Println("Start profiling")
	go http.ListenAndServe(":10201", http.DefaultServeMux)
}
