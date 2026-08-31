package main

import (
 "encoding/json"
 "log"
 "net/http"
 "os"
 "time"
)

type Protocol struct { Name string `json:"name"`; Status string `json:"status"`; Managed bool `json:"managed"` }
var protocols = []Protocol{{"amneziawg","available",true},{"wireguard","available",true},{"xray","available",true},{"shadowsocks","available",true},{"hysteria2","available",true},{"openvpn","available",true},{"gre","available",true},{"ipsec","available",true}}
func jsonOK(w http.ResponseWriter,v any){w.Header().Set("Content-Type","application/json"); _=json.NewEncoder(w).Encode(v)}
func main(){
 addr:=os.Getenv("ROUTERB_ADDR"); if addr=="" {addr=":8080"}
 mux:=http.NewServeMux()
 mux.HandleFunc("/health",func(w http.ResponseWriter,r *http.Request){jsonOK(w,map[string]any{"status":"ok","service":"routerb","time":time.Now().UTC()})})
 mux.HandleFunc("/healthz",func(w http.ResponseWriter,r *http.Request){jsonOK(w,map[string]string{"status":"ok"})})
 mux.HandleFunc("/api/v1/protocols",func(w http.ResponseWriter,r *http.Request){jsonOK(w,protocols)})
 mux.HandleFunc("/api/v1/node",func(w http.ResponseWriter,r *http.Request){jsonOK(w,map[string]any{"name":hostname(),"domain":os.Getenv("ROUTERB_DOMAIN"),"service":"routerb"})})
 mux.HandleFunc("/api/v1/info",func(w http.ResponseWriter,r *http.Request){jsonOK(w,map[string]any{"service":"FTN RouterB","version":"1.0.0","control_plane":"cloud.familytimenet.com","protocols":len(protocols)})})
 mux.Handle("/",http.FileServer(http.Dir("./web")))
 log.Printf("RouterB listening on %s",addr); log.Fatal(http.ListenAndServe(addr,mux))
}
func hostname() string { h,e:=os.Hostname(); if e!=nil || h=="" {return "routerb-node"}; return h }
