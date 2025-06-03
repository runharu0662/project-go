package main

import (
    "fmt"
    "net"
    "net/http"
)

func getLocalIP() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return "IP取得失敗"
    }
    for _, addr := range addrs {
        if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
            return ipnet.IP.String()
        }
    }
    return "ローカルIPなし"
}

func handlePost(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")
    message := r.FormValue("message")
    fmt.Printf("投稿: %s - %s\n", name, message)
    w.Write([]byte("投稿を受け付けました"))
}

func main() {
    http.HandleFunc("/api/post", handlePost)

    localIP := getLocalIP()
    fmt.Printf("Go API running on: http://%s:8080\n", localIP)

    http.ListenAndServe(":8080", nil)
}

