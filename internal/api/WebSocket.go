package api

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func reader(conn *websocket.Conn, dataChannel chan interface{}, id string) {
	go readMessage(conn, id)
	for {
		select {
		case d := <-dataChannel:
			err := writeMessage(conn, d)
			if err != nil {
				closeSocketRoutine(id, err)
				return
			}
		}
	}
}

func writeMessage(conn *websocket.Conn, d interface{}) error {
	response := map[string]interface{}{"status": "log", "data": d}
	data, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		return err
	}
	err = conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return err
	}
	return nil
}

func readMessage(conn *websocket.Conn, id string) {
	x := 0
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			closeSocketRoutine(id, err)
			return
		}
		if x == 0 {
			x++
			response := map[string]string{"status": "Connection to Puffin KYC established"}
			data, _ := json.Marshal(response)
			if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
				closeSocketRoutine(id, err)
				return
			}
		}

		if string(msg) == "logs" {
			response := map[string]string{"status": "Logs enabled"}
			data, _ := json.Marshal(response)
			if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
				closeSocketRoutine(id, err)
				return
			}
		} else if string(msg) == "ping" {
			response := map[string]string{"status": "pong"}
			data, _ := json.Marshal(response)
			if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
				closeSocketRoutine(id, err)
				return
			}
		}

	}
}

func closeSocketRoutine(id string, err error) {
	log.Println(err)
	if SocketChannels[id] != nil {
		delete(SocketChannels, id)
	}
}