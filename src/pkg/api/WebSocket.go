package api

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"puffinverificationbackend/src/pkg/global"
)

func reader(conn *websocket.Conn) {
	enabled := map[string]bool{"logs": false}
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				global.SocketCount--
				log.Println(err)
				return
			}
			response := map[string]string{"status": "Connection to Puffin KYC established"}
			global.SocketCount++
			data, _ := json.Marshal(response)
			if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println(err)
				return
			}

			if string(msg) == "logs" {
				enabled["logs"] = true
				response = map[string]string{"status": "Logs enabled"}
				data, _ = json.Marshal(response)
				if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
					log.Println(err)
					return
				}
			}

		}
	}()
	for {
		select {
		case d := <-global.SocketChannel:
			go func() {
				response := map[string]interface{}{"status": "log", "data": d}
				data, err := json.Marshal(response)
				if err != nil {
					return
				}
				if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
					log.Println(err)
					return
				}
			}()

		}
	}

}