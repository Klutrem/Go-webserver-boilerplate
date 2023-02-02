package ws

import (
	"encoding/json"
	"fmt"
	"main/lib"
	"net/http"
	"time"

	"main/api/kubes"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	corev1 "k8s.io/api/core/v1"
)

type Ws struct {
	upgrader websocket.Upgrader
	logger   lib.Logger
}

func NewWs(logger lib.Logger) Ws {
	wsupgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return Ws{
		upgrader: wsupgrader,
		logger:   logger,
	}
}

func (w Ws) MessageHandler(c *gin.Context) {
	connection, err := w.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		w.logger.Fatal(err.Error())
	}
	defer connection.Close()

	_, pod_name, err := connection.ReadMessage()
	if err != nil {
		w.logger.Fatal(err.Error())
	}
	k := kubes.NewKubeRequest(w.logger)
	go func() {
		events := k.GetEvents("default")
		Watcher := events.ResultChan()
		fmt.Println(Watcher)
		for event := range Watcher {
			item := event.Object.(*corev1.Event)
			v, err := json.Marshal(item)
			if err != nil {
				w.logger.Panic(err.Error())
			}
			err = connection.WriteMessage(websocket.TextMessage, v)
			if err != nil {
				w.logger.Panic(err.Error())
			}
			fmt.Println(v)
		}
		events.Stop()

	}()

	for {
		response := k.GetCurrentPodStatusRequest(string(pod_name))
		err = connection.WriteMessage(websocket.TextMessage, response)
		time.Sleep(time.Second * 1)
		if err != nil {
			w.logger.Fatal(err.Error())
			break
		}
		defer connection.Close()
	}

}
