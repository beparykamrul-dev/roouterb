package main

import (
 "encoding/json"
 "log"
 "net/http"
 "os"
 "time"

 "github.com/go-chi/chi/v5"
 "github.com/go-chi/chi/v5/middleware"
)

type Protocol struct { Name string `json:"name"`; Status string `json:"status"`; Managed bool `json:"managed"` }
var protocols = []Protocol{
 {"amneziawg","available",true},{"wireguard","available",true},{"xray","available",true},
 {"shadowsocks","available",true},{"hysteria2","available",true},{"openvpn","available",true},
 {"gre","available",true},{"ipsec","available",true},
}
func main(){
 addr:=os.Getenv("ROUTERB_ADDR"); if addr=="" { addr=":8080" }
 r:=chi.NewRouter(); r.Use(middleware.RequestID,middleware.RealIP,middleware.Recoverer)
 r.Get("/health",func(w http.ResponseWriter,r *http.Request){ json.NewEncoder(w).Encode(map[string]any{"status":"ok","service":"routerb","time":time.Now().UTC()}) })
 r.Get("/api/v1/protocols",func(w http.ResponseWriter,r *http.Request){ w.Header().Set("Content-Type","application/json"); json.NewEncoder(w).Encode(protocols) })
 r.Get("/api/v1/info",func(w http.ResponseWriter,r *http.Request){ json.NewEncoder(w).Encode(map[string]any{"service":"FTN RouterB","version":"1.0.0","control_plane":"cloud.familytimenet.com","protocols":len(protocols)}) })
 log.Printf("RouterB listening on %s",addr); log.Fatal(http.ListenAndServe(addr,r))
}
