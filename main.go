package main

import (
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	"gopkg.in/antage/eventsource.v1"
)

func main() {
	es := eventsource.New(nil, nil)
	defer es.Close()
	http.Handle("/events", es)
	go func() {
		id := 1
		for {
			es.SendEventMessage("tick", "tick-event", strconv.Itoa(id))
			id++
			time.Sleep(2 * time.Second)
		}
	}()
	Openbrowser("http://localhost:8080/events")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Openbrowser : Opens default web browser to specified url
func Openbrowser(url string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "linux":
		cmd = "chromium-browser"
		args = []string{""}

	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
