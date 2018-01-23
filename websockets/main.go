package main


import (
	"net/http"
	"github.com/gorilla/websocket"
	//"encoding/json"
	//"io"
	//"io/ioutil"
	//"reflect"
	"fmt"
)

var upgrader = websocket.Upgrader{} //convert http protocol to ws

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/v1/ws", func(w http.ResponseWriter, r *http.Request){
		conn,_ := upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn){
			for {
				mType, msg, _:= conn.ReadMessage()
				conn.WriteMessage(mType, msg)
			}
		}(conn)
	})
	type MsgConnected struct {
		User string `json:"user"`
		IsConnected bool `json:"isConnected"`
	}
	http.HandleFunc("/v2/ws", func(w http.ResponseWriter, r *http.Request) {
		conn,_ := upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			for{
				var msgJ = MsgConnected{}
				//_, msg, _:= conn.ReadMessage() // return msg type, msg body, error
				//ioutil.ReadAll(msg)
				err := conn.ReadJSON(&msgJ)
				//err:= json.Unmarshal(msg, &msgJ) //use for ReadMessage
				if err != nil {
					panic(err)
				}
				fmt.Println(msgJ)

			}
		}(conn)
	})

	http.ListenAndServe(":3000", nil)
}