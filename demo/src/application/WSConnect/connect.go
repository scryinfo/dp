package WSConnect

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/scryInfo/dp/demo/src/application/definition"
	rlog "github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
	"net/http"
	"os/exec"
)

type WSServer struct {
	Port string
}

const ResourcesDir = "D:/EnglishRoad/workspace/Go/src/github.com/scryInfo/dp/demo/src/application/resources/app"

var (
	funcMap    = make(map[string]definition.PresetFunc)
	connParams *websocket.Conn
	err        error
)

func WebsocketConnect(port string) error {
	ws := WSServer{
		Port: port,
	}
	return errors.Wrap(ws.start(), "websocket WSConnect failed. ")
}

func (ws *WSServer) start() error {
	fmt.Println("> Start listening ... ")
	http.HandleFunc("/", bindHTMLFile)
	http.Handle("/ws", websocket.Handler(ws.handleMessages))
	spfsd := http.StripPrefix("/static/", http.FileServer(http.Dir(ResourcesDir+"/static")))
	http.Handle("/static/", spfsd)
	fmt.Println("> Listening at http://127.0.0.1:" + ws.Port)
	if err = exec.Command("cmd", "/c start http://127.0.0.1:"+ws.Port).Start(); err != nil {
		err = errors.Wrap(err, "auto-open url failed, port: "+ws.Port+". ")
	}
	if err = http.ListenAndServe(":"+ws.Port, nil); err != nil {
		err = errors.Wrap(err, "Listen and Server failed. ")
	}
	return err
}

func bindHTMLFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, ResourcesDir+"/index.html")
}

func (ws *WSServer) handleMessages(conn *websocket.Conn) {
	connParams = conn
	fmt.Printf("> A new ws connection: %s -> %s\n", conn.RemoteAddr().String(), conn.LocalAddr().String())
	for {

		// -
		fmt.Println("------------------------------------------------------------")
		// -

		var reply []byte
		if err = websocket.Message.Receive(conn, &reply); err != nil {
			rlog.Error(errors.Wrap(err, "Received failed. "))
			continue
		}

		var mi definition.MessageIn
		if err = json.Unmarshal(reply, &mi); err != nil {
			rlog.Error(errors.Wrap(err, "json unmarshal failed. "))
			continue
		}

		// -
		fmt.Printf("Received from browser: %s | %s\n", mi.Name, string(mi.Payload))
		// -

		var payload interface{}
		if _, ok := funcMap[mi.Name]; !ok {
			rlog.Error("Unknown method name: ", mi.Name)
			continue
		}
		payload, err = funcMap[mi.Name](&mi)

		var mo definition.MessageOut
		mo.Name = mi.Name + ".callback"
		if err != nil {
			mo.Name += ".error"
			payload = errors.Wrap(err, mi.Name+" failed. ")
		}
		mo.Payload = payload
		var b []byte
		if b, err = json.Marshal(mo); err != nil {
			rlog.Error(errors.Wrap(err, "json marshal failed. "))
		}

		// -
		fmt.Println("Revert message to js :", mo.Name, "|", mo.Payload)
		// -

		if err = websocket.Message.Send(conn, string(b)); err != nil {
			rlog.Error(errors.Wrap(err, "Send failed. "))
			continue
		}
	}
}

func sendMessage(name string, payload interface{}) error {
	for {
		var b []byte
		{
			var mo definition.MessageOut
			mo.Name = name
			mo.Payload = payload
			if b, err = json.Marshal(mo); err != nil {
				break
			}
		}
		if err = websocket.Message.Send(connParams, string(b)); err != nil {
			break
		}
		break
	}

	return err
}

func addCallbackFunc(name string, presetFunc definition.PresetFunc) {
	funcMap[name] = presetFunc
}
