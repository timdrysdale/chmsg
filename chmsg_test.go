package chmsg

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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
