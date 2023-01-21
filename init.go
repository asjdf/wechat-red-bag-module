package atomLunarRedBag

import (
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"sync"
	"wechat-mp-server/hub"
)

type Mod struct {
}

func (m *Mod) GetModuleInfo() hub.ModuleInfo {
	return hub.ModuleInfo{
		ID:       hub.NewModuleID("atom", "LunarRedBag"),
		Instance: m,
	}
}

func (m *Mod) Init() {

}

func (m *Mod) PostInit() {

}

func (m *Mod) Serve(s *hub.Server) {
	s.MsgEngine.MsgText("^ATOM-REDBAG-XYH$", 1, func(msg *hub.Message) {
		msg.Reply = &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("支付宝红包口令：卷鸡祝您新年快乐XD")}
	})
}

func (m *Mod) Start(_ *hub.Server) {

}

func (m *Mod) Stop(_ *hub.Server, wg *sync.WaitGroup) {
	wg.Done()
}

func init() {
	hub.RegisterModule(&Mod{})
}
