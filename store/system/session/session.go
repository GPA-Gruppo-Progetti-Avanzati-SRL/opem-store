package session

import "go.mongodb.org/mongo-driver/bson/primitive"
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type Session struct {
	OId        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Userid     string             `json:"userid,omitempty" bson:"userid,omitempty" yaml:"userid,omitempty"`
	Nickname   string             `json:"nickname,omitempty" bson:"nickname,omitempty" yaml:"nickname,omitempty"`
	Remoteaddr string             `json:"remoteaddr,omitempty" bson:"remoteaddr,omitempty" yaml:"remoteaddr,omitempty"`
	Flags      string             `json:"flags,omitempty" bson:"flags,omitempty" yaml:"flags,omitempty"`
	SysInfo    commons.SysInfo    `json:"sysInfo,omitempty" bson:"sysInfo,omitempty" yaml:"sysInfo,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Session) IsZero() bool {
	return s.OId == primitive.NilObjectID && s.Userid == "" && s.Nickname == "" && s.Remoteaddr == "" && s.Flags == "" && s.SysInfo.IsZero()
}

// @tpm-schematics:start-region("bottom-file-section")

func (s Session) SessionId() string {
	if s.OId == primitive.NilObjectID {
		return ""
	}

	return s.OId.Hex()
}

// @tpm-schematics:end-region("bottom-file-section")
