package mqttrapper

import (
	"context"
	"errors"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
	mainGlobal "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/config"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/global"
	"go.uber.org/zap"
)

func on_connect_succeed(cm *autopaho.ConnectionManager, connAck *paho.Connack) {
	mainGlobal.GVA_LOG.Info("mqtt connection up")
	if _, err := cm.Subscribe(context.Background(), &paho.Subscribe{
		Subscriptions: map[string]paho.SubscribeOptions{
			"/svc/basic/system": {QoS: 0},
		},
	}); err != nil {
		mainGlobal.GVA_LOG.Error("failed to subscribe!!", zap.Error(err))
		return
	}
	mainGlobal.GVA_LOG.Info("/svc/basic/system" + " :Subscribe made")
	global.MqttConnection = cm
}

func on_connect_failed(err error) {
	mainGlobal.GVA_LOG.Error("error whilst attempting connection:", zap.Error(err))
}

func on_client_err(err error) {
	mainGlobal.GVA_LOG.Error("server requested disconnect:", zap.Error(err))
}

func on_server_disconnect(d *paho.Disconnect) {
	if d.Properties != nil {
		mainGlobal.GVA_LOG.Error("server requested disconnect:", zap.Error(errors.New(d.Properties.ReasonString)))
	} else {
		mainGlobal.GVA_LOG.Error("server requested disconnect:", zap.Error(errors.New(d.Properties.ReasonString)))
		//fmt.Printf("server requested disconnect; reason code: %d\n", d.ReasonCode)
	}
}

func MqttNew(param *config.Config, onMessage func(m *paho.Publish)) {
	mqtturl, _ := url.Parse(config.MqttDns(&param.Mqtt))
	cliCfg := autopaho.ClientConfig{
		BrokerUrls:        []*url.URL{mqtturl},
		KeepAlive:         param.Mqtt.KeepAlive,
		ConnectRetryDelay: param.Mqtt.ConnectRetryDelay,
		OnConnectionUp:    on_connect_succeed,
		OnConnectError:    on_connect_failed,
		ClientConfig: paho.ClientConfig{
			ClientID:           param.Mqtt.ClientID,
			Router:             paho.NewSingleHandlerRouter(onMessage),
			OnClientError:      on_client_err,
			OnServerDisconnect: on_server_disconnect,
		},
	}
	cliCfg.SetUsernamePassword(param.Mqtt.Username, []byte(param.Mqtt.Password))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cm, err := autopaho.NewConnection(ctx, cliCfg)
	if err != nil {
		mainGlobal.GVA_LOG.Panic("Connection Err: ", zap.Error(err))
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	<-sig
	mainGlobal.GVA_LOG.Info("signal caught - exiting")

	// We could cancel the context at this point but will call Disconnect instead (this waits for autopaho to shutdown)
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_ = cm.Disconnect(ctx)

	mainGlobal.GVA_LOG.Info("shutdown complete")
}

func MqttPublish(p *paho.Publish) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	global.MqttConnection.Publish(ctx, p)
}

func MqttSubscribe(topic string, qos byte) {
	if _, err := global.MqttConnection.Subscribe(context.Background(), &paho.Subscribe{
		Subscriptions: map[string]paho.SubscribeOptions{
			topic: {QoS: qos},
		},
	}); err != nil {
		mainGlobal.GVA_LOG.Error("failed to subscribe!!", zap.Error(err))
		return
	}
	mainGlobal.GVA_LOG.Info(topic + " :Subscribe made")
}
