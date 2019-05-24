// Scry Info.  All rights reserved.
// license that can be found in the license file.

package connection

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dp/dots/app/settings"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
	"net/http"
	"os/exec"
)

type WSServer struct {
	port           string
	connParams     *websocket.Conn
	funcMap        map[string]settings.PresetFunc
	uiResourcesDir string
}

const EventSendFailed = " event send failed. "

var _ Connection = (*WSServer)(nil)

func CreateConnetion(port, dir string) *WSServer {
	return &WSServer{
		port:           port,
		uiResourcesDir: dir,
		funcMap:        make(map[string]settings.PresetFunc),
	}
}

func (ws *WSServer) Connect() error {
	return errors.Wrap(ws.start(), "ws connect failed. ")
}

func (ws *WSServer) start() error {
	dot.Logger().Infoln("> Start listening ... ")

	http.HandleFunc("/", ws.bindHTMLFile)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(ws.uiResourcesDir+"/static"))))
	http.Handle("/ws", websocket.Handler(ws.handleMessages))

	if err := exec.Command("cmd", "/c start http://127.0.0.1:"+ws.port).Start(); err != nil {
		return errors.Wrap(err, "auto-open url failed, (origin: http://127.0.0.1:"+ws.port+"). ")
	}

	dot.Logger().Infoln("> Listening at http://127.0.0.1:" + ws.port)

	go func()  {
		if err := http.ListenAndServe(":"+ws.port, nil); err != nil {
			dot.Logger().Errorln("Listen and Server failed. ", zap.NamedError("", err))
		}
	}()

	return nil
}
func (ws *WSServer) bindHTMLFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, ws.uiResourcesDir+"/index.html")
}

func (ws *WSServer) handleMessages(conn *websocket.Conn) {
	logger := dot.Logger()
	ws.connParams = conn

	logger.Infoln("> A new connection connected: " + conn.RemoteAddr().String() + " -> " + conn.LocalAddr().String())

	var err error
	for {
		// receive message and check if handle function is exist.
		var mi settings.MessageIn
		{
			var reply []byte
			if err = websocket.Message.Receive(ws.connParams, &reply); err != nil {
				if err.Error() == "EOF" {
					logger.Warnln("> Client disconnect. ")
					return
				}
				logger.Errorln("Received failed. ", zap.NamedError("", err))
				continue
			}
			if err = json.Unmarshal(reply, &mi); err != nil {
				logger.Errorln("JSON unmarshal failed. ", zap.NamedError("", err))
				continue
			}
			if _, ok := ws.funcMap[mi.Name]; !ok {
				logger.Errorln("Unknown method name: " + mi.Name)
				continue
			}
		}

		// handle and send result back to js.
		{
			// handle
			var payload interface{}
			payload, err = ws.funcMap[mi.Name](&mi)

			name := mi.Name + ".callback"
			if err != nil {
				name += ".error"
				payload = errors.Wrap(err, mi.Name+" failed. ")
				logger.Errorln("", zap.Any("", payload))
			}

			// send
			if err = ws.SendMessage(name, payload); err != nil {
				logger.Errorln("", zap.NamedError(name+EventSendFailed, err))
			}
		}
	}
}

func (ws *WSServer) SendMessage(name string, payload interface{}) error {
	mo := settings.MessageOut{
		Name:    name,
		Payload: payload,
	}

	b, err := json.Marshal(mo)
	if err != nil {
		return errors.Wrap(err, "Json marshal failed. ")
	}

	return websocket.Message.Send(ws.connParams, string(b))
}

func (ws *WSServer) AddCallbackFunc(name string, presetFunc settings.PresetFunc) {
	ws.funcMap[name] = presetFunc
}

func (ws *WSServer) GetPort() string {
	return ws.port
}
