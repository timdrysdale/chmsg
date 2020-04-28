package chmsg

import (
	"errors"
	"time"
)

// arguments to New
type MessagerConf struct {
	ExamName     string
	FunctionName string
	TaskName     string
}

// what we send in every message
type MessageInfo struct {
	ExamName     string
	FunctionName string
	TaskName     string
	Message      string
	Time         time.Duration
}

// the main struct created by New
type Messager struct {
	Chan    chan MessageInfo
	Timeout time.Duration
	MessagerConf
}

func New(conf MessagerConf, infoChan chan MessageInfo, timeout time.Duration) *Messager {

	m := Messager{}
	m.ExamName = conf.ExamName
	m.FunctionName = conf.FunctionName
	m.TaskName = conf.TaskName
	m.Chan = infoChan
	m.Timeout = timeout
	return &m
}

// You had _one_ job!
func (m *Messager) Send(msg string) error {
	var msgInfo MessageInfo

	msgInfo = MessageInfo{}
	msgInfo.ExamName = m.ExamName
	msgInfo.FunctionName = m.FunctionName
	msgInfo.TaskName = m.TaskName
	//copy this in case we sending messages in parallel from same Messager instance
	msgInfo.Message = msg

	select {
	case <-time.After(m.Timeout):
		return errors.New("timeout")
	case m.Chan <- msgInfo:
		return nil
	}
}
