package server

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dp/dots/app/server/definition"
	"github.com/scryinfo/dp/dots/app/storage"
	DBDef "github.com/scryinfo/dp/dots/app/storage/definition"
	scry2 "github.com/scryinfo/dp/dots/binary/scry"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
	"math/rand"
	"net/http"
	"os/exec"
	"time"
)

// WSServer with optional leading article
type WSServer struct {
	CurUser    scry2.Client
	connParams *websocket.Conn
	funcMap    map[string]definition.PresetFunc
	SaltStr    []byte
	config     serverConfig
	DB         *storage.SQLite `dot:""`
}

type serverConfig struct {
	Port           string `json:"wsPort"`
	UIResourcesDir string `json:"uiResourcesDir"`
}

// WebSocketTypeId websocket type id
const WebSocketTypeId = "40ef6679-5cfc-4436-a1f6-7f39870bc5ef"

var _ Server = (*WSServer)(nil)

func newWebSocketDot(conf []byte) (dot.Dot, error) {
	var err error

	dConf := &serverConfig{}
	err = dot.UnMarshalConfig(conf, dConf)
	if err != nil {
		return nil, err
	}

	d := &WSServer{config: *dConf}

	return d, err
}

// WebSocketTypeLive add a dot component to dot.line with 'line.PreAdd()'
func WebSocketTypeLive() *dot.TypeLives {
	return &dot.TypeLives{
		Meta: dot.Metadata{
			TypeId: WebSocketTypeId,
			NewDoter: func(conf []byte) (dot.Dot, error) {
				return newWebSocketDot(conf)
			},
		},
	}
}

// Create dot.Create
func (ws *WSServer) Create(l dot.Line) error {
	rand.Seed(time.Now().Unix())

	ws.funcMap = make(map[string]definition.PresetFunc)

	ws.SaltStr = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	return nil
}

// ListenAndServe start a http server and listen given port
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
		var mi definition.MessageIn
		{
			// receive from client
			var reply []byte
			if err = websocket.Message.Receive(ws.connParams, &reply); err != nil {
				if err.Error() == "EOF" {
					logger.Infoln("> Client disconnect: " + conn.LocalAddr().String())
					return
				}
				logger.Errorln("Received failed. ", zap.NamedError("", err))
				continue
			}
			if err = json.Unmarshal(reply, &mi); err != nil {
				logger.Errorln("JSON unmarshal failed. ", zap.NamedError("", err))
				continue
			}

			logger.Debugln("received   :", zap.Any("message in: ", mi))

			if _, ok := ws.funcMap[mi.Name]; !ok {
				logger.Errorln("Unknown request name: " + mi.Name)
				continue
			}
		}

		// handle and send result back to js.
		{
			// handle
			var message *definition.MessageInPayload
			message, err = ws.calcPwdHash(mi.Payload)
			if err != nil {
				logger.Errorln("calculate password's hash failed. ", zap.NamedError("", err))
				continue
			}

			var payload interface{}
			payload, err = ws.funcMap[mi.Name](message)

			name := mi.Name + ".callback"
			if err != nil {
				name += ".error"
				payload = errors.Wrap(err, mi.Name+" failed. ").Error()
				logger.Errorln("", zap.Any("", payload))
			}

			// send
			if err = ws.SendMessage(name, payload); err != nil {
				logger.Errorln("", zap.NamedError(name+EventSendFailed, err))
			}
		}
	}
}

// SendMessage send message to client
func (ws *WSServer) SendMessage(name string, payload interface{}) error {
	if bs, ok := payload.([]byte); ok {
		payload = string(bs) // []byte will be base58 encode first, in json marshal
	}

	mo := definition.MessageOut{
		Name:    name,
		Payload: payload,
	}

	b, err := json.Marshal(mo)
	if err != nil {
		return errors.Wrap(err, "Json marshal failed. ")
	}

	dot.Logger().Debugln("before send:", zap.Any("message out: ", mo))

	return websocket.Message.Send(ws.connParams, string(b)) // avoid base58 encode
}

// PresetMsgHandleFuncs preset system functions' handler
func (ws *WSServer) PresetMsgHandleFuncs(name []string, presetFunc []definition.PresetFunc) error {
	if len(name) != len(presetFunc) {
		return errors.New("Quantities of name and function are not equal. ")
	}

	for i := range name {
		ws.funcMap[name[i]] = presetFunc[i]
	}

	return nil
}

func (ws *WSServer) calcPwdHash(bytes json.RawMessage) (result *definition.MessageInPayload, err error) {
	if string(bytes) == "\"\"" { // payload is nil, no pwd item at all
		return
	}

	result = &definition.MessageInPayload{}
	if err = json.Unmarshal(bytes, result); err != nil {
		return
	}
	dot.Logger().Debugln("show param :", zap.Any("result", result))

	if result.Password != "" {
		var userAddr string
		if ws.CurUser == nil && result.Address == "" { // msg: create new account, only generate salt
			result.Salt = ws.generateSalt(10)
		} else if ws.CurUser == nil && result.Address != "" { // msg: login verify, match curUser with addr input
			userAddr = result.Address
		} else {
			userAddr = ws.CurUser.Account().Addr
		}

		var salt string
		if result.Salt != "" { // create
			salt = result.Salt
		} else {
			if salt, err = ws.readSalt(userAddr); err != nil {
				return
			}
		}

		t := sha256.Sum256([]byte(result.Password + salt))
		result.Password = hex.EncodeToString(t[:])

		dot.Logger().Debugln("pwd hash   :", zap.String("pwd", result.Password),
			zap.String("pwd meta", result.Password),
			zap.String("salt", salt),
		)
	}

	return
}

func (ws *WSServer) generateSalt(length int) string {
	var num int
	for i := len(ws.SaltStr) - 1; i > 0; i-- {
		num = rand.Intn(i + 1)
		ws.SaltStr[i], ws.SaltStr[num] = ws.SaltStr[num], ws.SaltStr[i]
	}

	num = rand.Intn(53)

	return string(ws.SaltStr[num : num+length])
}

func (ws *WSServer) readSalt(addr string) (string, error) {
	var acc DBDef.Account
	num, err := ws.DB.Read(&acc, "", "address = ?", addr)
	if num != 1 || err != nil {
		err = errors.Wrap(err, "db read failed")
		return "", err
	}

	return acc.Salt, nil
}
