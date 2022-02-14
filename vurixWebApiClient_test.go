package vurixwebapiclient

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestVurixWebClient(t *testing.T) {
	opt := NewOptVurixWebApiClient()
	opt.Host = "172.16.31.202"
	opt.Port = 8080
	opt.User = "admin"
	opt.Pass = "admin"
	opt.License = "licNormalClient"

	vc := NewVurixWebApiClient(opt)
	logger, _ := zap.NewDevelopment()
	vc.SetLogger(logger.Sugar())
	vc.SetDebug(true)
	vc.Login()

	// 이벤트를 붙이는 작업을 해야 한다.
	optVER := OptVurixEventReceiver{
		DeviceEvent:     true,
		MonitoringEvent: true,
		SystemEvent:     true,
		UserEvent:       true,
	}
	vc.SetEventHandler(CallbackFunc, optVER)

	time.Sleep(60 * time.Second)
	vc.Logout()

	time.Sleep(1 * time.Second)

}

func CallbackFunc(msg interface{}) {
	jsonmsg, _ := json.Marshal(msg)
	fmt.Println("recv : ", time.Now(), string(jsonmsg))
}
