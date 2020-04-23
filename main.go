package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"./cmd"
	"./pcvolume"
	"./rgo"
	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
)

func router(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var err error
	var resp []byte

	if r.Method == "POST" && r.URL.Path == "/keyboard" {
		err = rgo.HandleKeyboardClick(r.Body)
	} else if r.Method == "POST" && r.URL.Path == "/mouse" {
		err = rgo.HandleMouseMovement(r.Body)
	} else if r.Method == "POST" && r.URL.Path == "/mouse_wheel" {
		err = rgo.HandleMouseWheel(r.Body)
	} else if r.Method == "POST" && r.URL.Path == "/controller" {
		err = cmd.HandleCMD(r.Body)
	} else if r.Method == "POST" && r.URL.Path == "/set_volume" {
		err = pcvolume.SetCurrentVolume(r.Body)
	} else if r.Method == "GET" && r.URL.Path == "/mute_volume" {
		err = pcvolume.UnmuteVolume()
	} else if r.Method == "GET" && r.URL.Path == "/unmute_volume" {
		err = pcvolume.MuteVolume()
	} else if r.Method == "GET" && r.URL.Path == "/get_current_volume" {
		resp, err = pcvolume.GetCurrentVolume()
	} else if r.Method == "GET" && r.URL.Path == "/check_connection" {

	} else {
		http.NotFound(w, r)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(resp)
}

func main() {
	port, err := getAvailablePort()
	if err != nil {
		fmt.Println("cant get any available port. Trying to do it with default 4040 port")
	}
	ip, err := getOutboundIP()
	if err != nil {
		fmt.Println("cant get your local ip address. Please type ipconfig command in cmd")
	}
	if ip != "" && port != "" {
		addr := ip + ":" + port
		fmt.Println("Your local ip address is: " + addr)
		fmt.Println("Please type it in settings page inside PCcontroll application, or just scan the QR below")
		generateQR(addr)
		fmt.Print("\n\n\n")
	}

	http.HandleFunc("/", router)
	http.ListenAndServe(":"+port, nil)
}

func getOutboundIP() (string, error) {
	conn, err := net.DialTimeout("tcp", "google.com:80", time.Duration(5)*time.Second)

	if err != nil {
		return "", err
	}
	ip := conn.LocalAddr().String()
	dblDotIdx := strings.Index(ip, ":")

	return ip[:dblDotIdx], nil
}

func generateQR(content string) {
	obj := qrcodeTerminal.New()
	obj.Get(content).Print()
}

func getAvailablePort() (string, error) {
	return "4040", nil
}
