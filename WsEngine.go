package WsEngine

import "github.com/gorilla/websocket"


type TWsEngine struct {
	WsClient map[*websocket.Conn]string
	WsNotify chan TWsNotify
}

type TWsNotify struct {
	EntityID string `json:"entityid"`
	UserName string `json:"username"`
	Notified string `json:"notified"`
	ShiftUsr string `json:"shiftusr"`
	ShiftPwd string `json:"shiftpwd"`
}