package MqttServer

import (
	"bytes"
	"log"

	"github.com/mochi-co/mqtt/v2"
	"github.com/mochi-co/mqtt/v2/listeners"
	"github.com/mochi-co/mqtt/v2/packets"
)

var (
	mqttServer *mqtt.Server
)

type MqttHook struct {
	mqtt.HookBase
}

func (h *MqttHook) ID() string {
	return "events-example"
}

func (h *MqttHook) Provides(b byte) bool {
	return bytes.Contains([]byte{
		mqtt.OnConnect,
		mqtt.OnDisconnect,
		mqtt.OnSubscribed,
		mqtt.OnUnsubscribed,
		mqtt.OnPublished,
		mqtt.OnPublish,
		mqtt.OnConnectAuthenticate,
		mqtt.OnACLCheck,
	}, []byte{b})
}

func (h *MqttHook) Init(config interface{}) error {
	h.Log.Info().Msg("initialised")
	return nil
}

// 网络连接成功
func (h *MqttHook) OnConnect(cl *mqtt.Client, pk packets.Packet) {
	log.Println("client connected")
}

// 连接授权
func (h *MqttHook) OnConnectAuthenticate(cl *mqtt.Client, pk packets.Packet) bool {

	_username := string(pk.Connect.Username)
	_passwd := string(pk.Connect.Password)

	rows, err := SqlUtil.Query("select '1' from mqtt_user where user=? and passwd=? limit 1;", _username, _passwd)
	if err != nil {
		return false
	}

	defer rows.Close()

	isOk := ""
	for rows.Next() {
		rows.Scan(&isOk)
	}

	if isOk != "1" {
		return false
	}

	log.Println("client auth: ", cl.ID)
	return true
}

// 主题 权限 检查
func (h *MqttHook) OnACLCheck(cl *mqtt.Client, topic string, write bool) bool {
	h.Log.Info().Str("client", cl.ID).Msgf("client OnACLCheck: %s \r\n %b", topic, write)
	return true
}

// 正常断开
func (h *MqttHook) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	h.Log.Info().Str("client", cl.ID).Bool("expire", expire).Err(err).Msg("client disconnected")
}

// 订阅
func (h *MqttHook) OnSubscribed(cl *mqtt.Client, pk packets.Packet, reasonCodes []byte) {
	h.Log.Info().Str("client", cl.ID).Interface("filters", pk.Filters).Msgf("subscribed qos=%v", reasonCodes)
}

// 取消订阅
func (h *MqttHook) OnUnsubscribed(cl *mqtt.Client, pk packets.Packet) {
	h.Log.Info().Str("client", cl.ID).Interface("filters", pk.Filters).Msg("unsubscribed")
}

// 发布
func (h *MqttHook) OnPublish(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	h.Log.Info().Str("client", cl.ID).Str("payload", string(pk.Payload)).Msg("received from client")

	pkx := pk
	if string(pk.Payload) == "hello" {
		pkx.Payload = []byte("hello world")
		h.Log.Info().Str("client", cl.ID).Str("payload", string(pkx.Payload)).Msg("received modified packet from client")
	}

	return pkx, nil
}

// 发布完成
func (h *MqttHook) OnPublished(cl *mqtt.Client, pk packets.Packet) {
	h.Log.Info().Str("client", cl.ID).Str("payload", string(pk.Payload)).Msg("published to client")
}

// mqtt server entry
func MqttServerRun() error {

	// An example of configuring various server options...
	options := &mqtt.Options{
		// InflightTTL: 60 * 15, // Set an example custom 15-min TTL for inflight messages
	}

	mqttServer = mqtt.New(options)

	// For security reasons, the default implementation disallows all connections.
	// If you want to allow all connections, you must specifically allow it.
	_hook := new(MqttHook)
	err := mqttServer.AddHook(_hook, nil)
	if err != nil {
		return err
	}

	tcp := listeners.NewTCP("server-1", ":1883", nil)
	err = mqttServer.AddListener(tcp)
	if err != nil {
		return err
	}

	go func() {
		err := mqttServer.Serve()
		if err != nil {
			log.Fatal(err)
		}
	}()
	return nil
}

// 正常停止服务
func MqttServerClose() {
	mqttServer.Log.Warn().Msg("caught signal, stopping...")
	mqttServer.Close()
	mqttServer.Log.Info().Msg("mqtt server stop finished.")
}
