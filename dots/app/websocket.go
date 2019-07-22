package app

import (
    "encoding/json"
    "github.com/pkg/errors"
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dp/dots/app/server"
    "go.uber.org/zap"
    "golang.org/x/net/websocket"
    "net/http"
    "os/exec"
)

type WSServer struct {
    connParams *websocket.Conn
    funcMap    map[string]server.PresetFunc
    config     serverConfig
}

type serverConfig struct {
    Port           string `json:"wsPort"`
    UIResourcesDir string `json:"uiResourcesDir"`
}

const WebSocketTypeId = "40ef6679-5cfc-4436-a1f6-7f39870bc5ef"

var _ server.Server = (*WSServer)(nil)

func newWebSocketDot(conf interface{}) (dot.Dot, error) {
    var err error
    var bs []byte
    if bt, ok := conf.([]byte); ok {
        bs = bt
    } else {
        return nil, dot.SError.Parameter
    }

    dConf := &serverConfig{}
    err = dot.UnMarshalConfig(bs, dConf)
    if err != nil {
        return nil, err
    }

    d := &WSServer{config: *dConf}

    return d, err
}

func WebSocketTypeLive() *dot.TypeLives {
    return &dot.TypeLives{
        Meta: dot.Metadata{
            TypeId: WebSocketTypeId,
            NewDoter: func(conf interface{}) (dot.Dot, error) {
                return newWebSocketDot(conf)
            },
        },
    }
}

func (ws *WSServer) Create(l dot.Line) error {
    ws.funcMap = make(map[string]server.PresetFunc)

    return nil
}

func (ws *WSServer) ListenAndServe() error {
    return errors.Wrap(ws.start(), "web serve start failed. ")
}

func (ws *WSServer) start() error {
    dot.Logger().Infoln("> Start listening ... ")

    http.HandleFunc("/", ws.bindHTMLFile)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(ws.config.UIResourcesDir+"/static"))))
    http.Handle("/ws", websocket.Handler(ws.handleMessages))

    if err := exec.Command("cmd", "/c start http://127.0.0.1:"+ws.config.Port).Start(); err != nil {
        return errors.Wrap(err, "auto-open url failed, (origin: http://127.0.0.1:"+ws.config.Port+"). ")
    }

    dot.Logger().Infoln("> Listening at http://127.0.0.1:" + ws.config.Port)

    go func() {
        if err := http.ListenAndServe(":"+ws.config.Port, nil); err != nil {
            dot.Logger().Errorln("Listen and Server failed. ", zap.NamedError("", err))
        }
    }()

    return nil
}
func (ws *WSServer) bindHTMLFile(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, ws.config.UIResourcesDir+"/index.html")
}
func (ws *WSServer) handleMessages(conn *websocket.Conn) {
    logger := dot.Logger()
    ws.connParams = conn

    logger.Infoln("> A new connection connected: " + conn.RemoteAddr().String() + " -> " + conn.LocalAddr().String())

    var err error
    for {
        // receive message and check if handle function is exist.
        var mi server.MessageIn
        {
            // receive from client
            var reply []byte
            if err = websocket.Message.Receive(ws.connParams, &reply); err != nil {
                if err.Error() == "EOF" {
                    logger.Warnln("> Client disconnect: " + conn.LocalAddr().String())
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
                logger.Errorln("", zap.NamedError(name+server.EventSendFailed, err))
            }
        }
    }
}

func (ws *WSServer) SendMessage(name string, payload interface{}) error {
    if bs, ok := payload.([]byte); ok {
        payload = string(bs)
    }
    mo := server.MessageOut{
        Name:    name,
        Payload: payload,
    }

    b, err := json.Marshal(mo)
    if err != nil {
        return errors.Wrap(err, "Json marshal failed. ")
    }

    return websocket.Message.Send(ws.connParams, string(b))
}

func (ws *WSServer) PresetMsgHandleFuncs(name []string, presetFunc []server.PresetFunc) error {
    if len(name) != len(presetFunc) {
        return errors.New("Quantities of name and function are not equal. ")
    }

    for i := range name {
        ws.funcMap[name[i]] = presetFunc[i]
    }

    return nil
}
