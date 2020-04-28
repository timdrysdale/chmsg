package chmsg

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//ype MessagerConf struct {
//	ExamName     string
//	FunctionName string
//	TaskName     string
//}
//
//// what we send in every message
//type MessageInfo struct {
//	MessageConf
//	Message string
//	Time    time.Duration
//}
//
//// the main struct created by New
//type Messager struct {
//	MessagerConf
//	Chan    chan MessageInfo
//	Timeout time.Duration
//}
//
//func New(info MessageInfo, infoChan chan MessageInfo, timeout time.Duration) *Messager {
//	return &Messager{
//		Conf:    info,
//		Chan:    infoChan,
//		Timeout: timeout,
//	}
//}
//
//// You had _one_ job!
//func (m *Messager) Send(msg string) error {
//	var msgInfo MessageInfo
//
//

func TestInstantiate(t *testing.T) {

	mc := MessagerConf{
		ExamName:     "ChMsg",
		FunctionName: "Test",
		TaskName:     "Chatter",
	}

	ch := make(chan MessageInfo)

	to := 100 * time.Millisecond

	m := New(mc, ch, to)

	text1 := "hello"
	text2 := "there"

	go func() {
		m.Send(text1)
		m.Send(text2)
	}()

	select {
	case msg := <-ch:
		assert.Equal(t, strings.Compare(msg.Message, text1), 0)
	case <-time.After(100 * time.Millisecond):
		t.Error("timeout on m1")
	}
	select {
	case msg := <-ch:
		assert.Equal(t, strings.Compare(msg.Message, text2), 0)
	case <-time.After(100 * time.Millisecond):
		t.Error("timeout on m2")
	}
}
