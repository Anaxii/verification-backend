package api

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"puffinverificationbackend/src/pkg/global"
)

func reader(conn *websocket.Conn, dataChannel chan interface{}, id string) {
	enabled := map[string]bool{"logs": false}
	x := 0
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				delete(global.SocketChannels, id)
				return
			}
			response := map[string]string{"status": "Connection to Puffin KYC established"}
			data, _ := json.Marshal(response)
			if x == 0 {
				x++
				if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
					log.Println(err)
					delete(global.SocketChannels, id)
					return
				}
			}

			if string(msg) == "logs" {
				enabled["logs"] = true
				response = map[string]string{"status": "Logs enabled"}
				data, _ = json.Marshal(response)
				if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
					log.Println(err)
					delete(global.SocketChannels, id)
					return
				}
			}
			if string(msg) == "ping" {
				response = map[string]string{"status": "pong"}
				data, _ = json.Marshal(response)
				if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
					log.Println(err)
					delete(global.SocketChannels, id)
					return
				}
			}

		}
	}()
	for {
		select {
		case d := <-dataChannel:
			response := map[string]interface{}{"status": "log", "data": d}
			data, err := json.Marshal(response)
			if err != nil {
				log.Println(err)
				return
			}
			if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
				delete(global.SocketChannels, id)
				log.Println(err)
				return
			}
		}
	}

}