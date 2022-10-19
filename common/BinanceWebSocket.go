package common

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"time"
)

var done chan interface{}
var interrupt chan os.Signal

func HandleReceive(conn *websocket.Conn) {
	defer close(done)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error in receive:", err)
			return
		}
		log.Printf("Received: %s\n", msg)
	}
}

func Client(streamName string) {

	// notice:每个到 stream.binance.com 的链接有效期不超过24小时，请妥善处理断线重连
	/*
		单一原始 streams 格式为 /ws/<streamName>
		组合streams的URL格式为 /stream?streams=<streamName1>/<streamName2>/<streamName3>
		每3分钟，服务端会发送ping帧，客户端应当在10分钟内回复pong帧，否则服务端会主动断开链接。
			允许客户端发送不成对的pong帧(即客户端可以以高于10分钟每次的频率发送pong帧保持链接)。
	*/
	socketUrl := "wss://stream.binance.com:9443/ws/" + streamName
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)

	if err != nil {
		color.Red("Error connecting to Websocket Server:", err)
	}
	defer conn.Close()

	go HandleReceive(conn)

	for {
		select {
		case <-time.After(time.Duration(1) * time.Second):
			msgType, p, err := conn.ReadMessage()
			if err != nil {
				return
			}
			fmt.Printf("messageType:%d,message:%s", msgType, p)
		case <-interrupt:
			// We received a SIGINT (Ctrl + C). Terminate gracefully...
			log.Println("Received SIGINT interrupt signal. Closing all pending connections")

			// Close our websocket connection
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error during closing websocket:", err)
				return
			}
		}

	}
}
