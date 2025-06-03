package main

import (
    "fmt"
    "net"
    "net/http"
)

// ローカルIPを取得する関数
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

// POSTを受け取って処理するハンドラ
func handlePost(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    name := r.FormValue("name")
    message := r.FormValue("message")
    fmt.Printf("投稿: %s - %s\n", name, message)
    w.Write([]byte("投稿を受け付けました"))
}

func main() {
    http.HandleFunc("/api/post", handlePost)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})


    localIP := getLocalIP()
    fmt.Printf("Go API running on: http://%s:8080\n", localIP)

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("サーバー起動失敗:", err)
    }
}

