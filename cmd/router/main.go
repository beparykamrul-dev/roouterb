package main

import (
 "encoding/json"
 "log"
 "net/http"
 "os"
 "strings"
 "time"
)

type Protocol struct { ID string `json:"id"`; Name string `json:"name"`; Type string `json:"type"`; Status string `json:"status"` }

var protocols = []Protocol{
 {"amneziawg","AmneziaWG","tunnel","available"},
 {"wireguard","WireGuard","tunnel","available"},
 {"xray","Xray","proxy","available"},
 {"shadowsocks","Shadowsocks","proxy","available"},
 {"hysteria2","Hysteria2","tunnel","available"},
 {"openvpn","OpenVPN","tunnel","available"},
 {"gre","GRE","network","available"},
 {"ipsec","IPsec","network","available"},
}

func main() {
 addr := os.Getenv("ROUTERB_ADDR"); if addr=="" { addr=":8080" }
 mux := http.NewServeMux()
 mux.HandleFunc("/healthz", func(w http.ResponseWriter,r *http.Request){ json.NewEncoder(w).Encode(map[string]any{"status":"ok","service":"roouterb","time":time.Now().UTC()}) })
 mux.HandleFunc("/api/v1/protocols", func(w http.ResponseWriter,r *http.Request){
  if r.Method!="GET" { http.Error(w,"method not allowed",405); return }; json.NewEncoder(w).Encode(protocols)
 })
 mux.HandleFunc("/api/v1/node", func(w http.ResponseWriter,r *http.Request){
  json.NewEncoder(w).Encode(map[string]any{"name":hostname(),"domain":os.Getenv("ROUTERB_DOMAIN"),"protocols":len(protocols)})
 })
 mux.Handle("/", http.FileServer(http.Dir("./web")))
 log.Printf("roouterb multi-protocol control plane listening on %s",addr)
 log.Fatal(http.ListenAndServe(addr, logging(mux)))
}
func hostname() string { h,_:=os.Hostname(); return h }
func logging(next http.Handler) http.Handler { return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){ if !strings.HasPrefix(r.URL.Path,"/healthz") { log.Printf("%s %s",r.Method,r.URL.Path) }; next.ServeHTTP(w,r) }) }
