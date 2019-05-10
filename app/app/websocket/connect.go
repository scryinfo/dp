package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/scryInfo/dp/app/app"
	settings2 "github.com/scryInfo/dp/app/app/settings"
	rlog "github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
	"net/http"
	"os/exec"
)

const EventSendFailed = " event send failed. "

var ( // todo: think how to reduce global variables.
	funcMap    = make(map[string]settings2.PresetFunc)
	connParams *websocket.Conn
)

type WSServer struct {
	Port string
}

func ConnectWithProtocolWebsocket(port string) error {
	ws := WSServer{
		Port: port,
	}
	return errors.Wrap(ws.start(), "websocket WSConnect failed. ")
}

func (ws *WSServer) start() error {
	fmt.Println("> Start listening ... ")

	http.HandleFunc("/", bindHTMLFile)
	t := http.StripPrefix("/static/", http.FileServer(http.Dir(app.GetGapp().ScryInfo.Config.UIResourcesDir+"/static")))
	http.Handle("/static/", t)
	http.Handle("/ws", websocket.Handler(ws.handleMessages))

	if err := exec.Command("cmd", "/c start http://127.0.0.1:"+ws.Port).Start(); err != nil {
		return errors.Wrap(err, "auto-open url failed, (origin: http://127.0.0.1:"+ws.Port+"). ")
	}
	fmt.Println("> Listening at http://127.0.0.1:" + ws.Port)
	if err := http.ListenAndServe(":"+ws.Port, nil); err != nil {
		return errors.Wrap(err, "Listen and Server failed. ")
	}

	return nil
}
func bindHTMLFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r,  app.GetGapp().ScryInfo.Config.UIResourcesDir+ "/index.html")
}

func (ws *WSServer) handleMessages(conn *websocket.Conn) {
	connParams = conn
	fmt.Printf("> A new websocket connection: %s -> %s\n", conn.RemoteAddr().String(), conn.LocalAddr().String())
	var err error
	for {
		// receive reply form js.
		var mi settings2.MessageIn
		{
			var reply []byte
			if err = websocket.Message.Receive(conn, &reply); err != nil {
				rlog.Error(errors.Wrap(err, "Received failed. "))
				continue
			}
			if err = json.Unmarshal(reply, &mi); err != nil {
				rlog.Error(errors.Wrap(err, "json unmarshal failed. "))
				continue
			}
			if _, ok := funcMap[mi.Name]; !ok {
				rlog.Error("Unknown method name: ", mi.Name)
				continue
			}
		}

		// handle and send result back to js.
		{
			// handle
			var payload interface{}
			payload, err = funcMap[mi.Name](&mi)

			name := mi.Name + ".callback"
			if err != nil {
				name += ".error"
				payload = errors.Wrap(err, mi.Name+" failed. ")
				rlog.Error(payload)
			}

			// send
			if err = sendMessage(name, payload); err != nil {
				rlog.Error(errors.Wrap(err, name+EventSendFailed))
			}
		}
	}
}

func sendMessage(name string, payload interface{}) (err error) {
	var b []byte

	mo := settings2.MessageOut{
		Name: name,
		Payload: payload,
	}
	if b, err = json.Marshal(mo); err != nil {
		return
	}
	if err = websocket.Message.Send(connParams, string(b)); err != nil {
		return
	}

	return
}

func addCallbackFunc(name string, presetFunc settings2.PresetFunc) {
	funcMap[name] = presetFunc
}
