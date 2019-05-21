// Scry Info.  All rights reserved.
// license that can be found in the license file.

package websocket

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/scryinfo/dot/dot"
	app2 "github.com/scryinfo/dp/dots/app"
	"github.com/scryinfo/dp/dots/app/settings"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
	"net/http"
	"os/exec"
)

const EventSendFailed = " event send failed. "

var ( // todo: use goroutine handle read and write, use struct instance instead of global variable.
	funcMap    = make(map[string]settings.PresetFunc)
	connParams *websocket.Conn
)

type WSServer struct {
	Port string
}

func ConnectWithProtocolWebsocket(port string) error {
	ws := WSServer{
		Port: port,
	}
	return errors.Wrap(ws.start(), "Websocket connect failed. ")
}

func (ws *WSServer) start() error {
	dot.Logger().Infoln("> Start listening ... ")

	http.HandleFunc("/", bindHTMLFile)
	t := http.StripPrefix("/static/", http.FileServer(http.Dir(app2.GetGapp().ScryInfo.Config.UIResourcesDir+"/static")))
	http.Handle("/static/", t)
	http.Handle("/ws", websocket.Handler(ws.handleMessages))

	if err := exec.Command("cmd", "/c start http://127.0.0.1:"+ws.Port).Start(); err != nil {
		return errors.Wrap(err, "auto-open url failed, (origin: http://127.0.0.1:"+ws.Port+"). ")
	}

	dot.Logger().Infoln("> Listening at http://127.0.0.1:" + ws.Port)

	if err := http.ListenAndServe(":"+ws.Port, nil); err != nil {
		return errors.Wrap(err, "Listen and Server failed. ")
	}

	return nil
}
func bindHTMLFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, app2.GetGapp().ScryInfo.Config.UIResourcesDir+"/index.html")
}

func (ws *WSServer) handleMessages(conn *websocket.Conn) {
	logger := dot.Logger()
	connParams = conn

	logger.Infoln("> A new websocket connection: " + conn.RemoteAddr().String() + " -> " + conn.LocalAddr().String())

	var err error
	for {
		// receive message form js.
		var mi settings.MessageIn
		{
			var reply []byte
			if err = websocket.Message.Receive(conn, &reply); err != nil {
				logger.Errorln("", zap.NamedError("Received failed. ", err))
				continue
			}
			if err = json.Unmarshal(reply, &mi); err != nil {
				logger.Errorln("", zap.NamedError("JSON unmarshal failed. ", err))
				continue
			}
			if _, ok := funcMap[mi.Name]; !ok {
				logger.Errorln("Unknown method name: " + mi.Name)
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
				logger.Errorln("", zap.Any("", payload))
			}

			// send
			if err = sendMessage(name, payload); err != nil {
				logger.Errorln("", zap.NamedError(name+EventSendFailed, err))
			}
		}
	}
}

func sendMessage(name string, payload interface{}) error {
	mo := settings.MessageOut{
		Name:    name,
		Payload: payload,
	}

	b, err := json.Marshal(mo)
	if err != nil {
		return errors.Wrap(err, "Json marshal failed. ")
	}

	return websocket.Message.Send(connParams, string(b))
}

func addCallbackFunc(name string, presetFunc settings.PresetFunc) {
	funcMap[name] = presetFunc
}
