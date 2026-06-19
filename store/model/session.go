package model

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
)

type Session struct {
	OId        string          `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Userid     string          `json:"userid,omitempty" bson:"userid,omitempty" yaml:"userid,omitempty"`
	Nickname   string          `json:"nickname,omitempty" bson:"nickname,omitempty" yaml:"nickname,omitempty"`
	RemoteAddr string          `json:"remote_addr,omitempty" bson:"remote_addr,omitempty" yaml:"remote_addr,omitempty"`
	Flags      string          `json:"flags,omitempty" bson:"flags,omitempty" yaml:"flags,omitempty"`
	SysInfo    commons.SysInfo `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`
}

func (s Session) IsZero() bool {
	return s.OId == "" && s.Userid == "" && s.Nickname == "" && s.RemoteAddr == "" && s.Flags == "" && s.SysInfo.IsZero()
}

func (s Session) SessionId() string {
	return s.OId
}

type SessionQueryResult struct {
	Records int       `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []Session `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}
