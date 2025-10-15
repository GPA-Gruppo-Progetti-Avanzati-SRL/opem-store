package session

import "go.mongodb.org/mongo-driver/v2/bson"
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type Session struct {
	OId        bson.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Userid     string          `json:"userid,omitempty" bson:"userid,omitempty" yaml:"userid,omitempty"`
	Nickname   string          `json:"nickname,omitempty" bson:"nickname,omitempty" yaml:"nickname,omitempty"`
	RemoteAddr string          `json:"remote_addr,omitempty" bson:"remote_addr,omitempty" yaml:"remote_addr,omitempty"`
	Flags      string          `json:"flags,omitempty" bson:"flags,omitempty" yaml:"flags,omitempty"`
	SysInfo    commons.SysInfo `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Session) IsZero() bool {
	return s.OId == bson.NilObjectID && s.Userid == "" && s.Nickname == "" && s.RemoteAddr == "" && s.Flags == "" && s.SysInfo.IsZero()
}

type QueryResult struct {
	Records int       `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []Session `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

// @tpm-schematics:start-region("bottom-file-section")

func (s Session) SessionId() string {
	if s.OId == bson.NilObjectID {
		return ""
	}

	return s.OId.Hex()
}

// @tpm-schematics:end-region("bottom-file-section")
