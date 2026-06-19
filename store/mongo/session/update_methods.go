package session

import (
	"fmt"
	"time"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
	model "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func UpdateMethodsGoInfo() string {
	i := fmt.Sprintf("tpm_morphia query filter support generated for %s package on %s", "author", time.Now().String())
	return i
}

type UnsetMode int64

const (
	UnSpecified     UnsetMode = 0
	KeepCurrent               = 1
	UnsetData                 = 2
	SetData2Default           = 3
)

type UnsetOption func(uopt *UnsetOptions)

type UnsetOptions struct {
	DefaultMode UnsetMode
	OId         UnsetMode
	Userid      UnsetMode
	Nickname    UnsetMode
	RemoteAddr  UnsetMode
	Flags       UnsetMode
	SysInfo     UnsetMode
}

func (uo *UnsetOptions) ResolveUnsetMode(um UnsetMode) UnsetMode {
	if um == UnSpecified {
		um = uo.DefaultMode
	}
	return um
}

func WithDefaultUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.DefaultMode = m
	}
}
func WithOIdUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.OId = m
	}
}
func WithUseridUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Userid = m
	}
}
func WithNicknameUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Nickname = m
	}
}
func WithRemoteAddrUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.RemoteAddr = m
	}
}
func WithFlagsUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Flags = m
	}
}
func WithSysInfoUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.SysInfo = m
	}
}

type UpdateOption func(ud *UpdateDocument)
type UpdateOptions []UpdateOption

func GetUpdateDocumentFromOptions(opts ...UpdateOption) UpdateDocument {
	ud := UpdateDocument{}
	for _, o := range opts {
		o(&ud)
	}
	return ud
}

func GetUpdateDocument(obj *model.Session, opts ...UnsetOption) UpdateDocument {
	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}
	ud := UpdateDocument{}
	ud.setOrUnsetUserid(obj.Userid, uo.ResolveUnsetMode(uo.Userid))
	ud.setOrUnsetNickname(obj.Nickname, uo.ResolveUnsetMode(uo.Nickname))
	ud.setOrUnsetRemote_addr(obj.RemoteAddr, uo.ResolveUnsetMode(uo.RemoteAddr))
	ud.setOrUnsetFlags(obj.Flags, uo.ResolveUnsetMode(uo.Flags))
	ud.setOrUnsetSys_info(&obj.SysInfo, uo.ResolveUnsetMode(uo.SysInfo))
	return ud
}

func (ud *UpdateDocument) SetOId(p bson.ObjectID) *UpdateDocument {
	mName := fmt.Sprintf(OIdFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetOId() *UpdateDocument {
	mName := fmt.Sprintf(OIdFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

func (ud *UpdateDocument) setOrUnsetOId(p bson.ObjectID, um UnsetMode) {
	if !p.IsZero() {
		ud.SetOId(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetOId()
		case SetData2Default:
			ud.UnsetOId()
		}
	}
}

func (ud *UpdateDocument) SetUserid(p string) *UpdateDocument {
	mName := fmt.Sprintf(UseridFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetUserid() *UpdateDocument {
	mName := fmt.Sprintf(UseridFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

func (ud *UpdateDocument) setOrUnsetUserid(p string, um UnsetMode) {
	if p != "" {
		ud.SetUserid(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetUserid()
		case SetData2Default:
			ud.UnsetUserid()
		}
	}
}

func UpdateWithUserid(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetUserid(p)
		} else {
			ud.UnsetUserid()
		}
	}
}

func (ud *UpdateDocument) SetNickname(p string) *UpdateDocument {
	mName := fmt.Sprintf(NicknameFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetNickname() *UpdateDocument {
	mName := fmt.Sprintf(NicknameFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

func (ud *UpdateDocument) setOrUnsetNickname(p string, um UnsetMode) {
	if p != "" {
		ud.SetNickname(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetNickname()
		case SetData2Default:
			ud.UnsetNickname()
		}
	}
}

func UpdateWithNickname(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetNickname(p)
		} else {
			ud.UnsetNickname()
		}
	}
}

func (ud *UpdateDocument) SetRemote_addr(p string) *UpdateDocument {
	mName := fmt.Sprintf(RemoteAddrFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetRemote_addr() *UpdateDocument {
	mName := fmt.Sprintf(RemoteAddrFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

func (ud *UpdateDocument) setOrUnsetRemote_addr(p string, um UnsetMode) {
	if p != "" {
		ud.SetRemote_addr(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetRemote_addr()
		case SetData2Default:
			ud.UnsetRemote_addr()
		}
	}
}

func UpdateWithRemote_addr(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetRemote_addr(p)
		} else {
			ud.UnsetRemote_addr()
		}
	}
}

func (ud *UpdateDocument) SetFlags(p string) *UpdateDocument {
	mName := fmt.Sprintf(FlagsFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetFlags() *UpdateDocument {
	mName := fmt.Sprintf(FlagsFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

func (ud *UpdateDocument) setOrUnsetFlags(p string, um UnsetMode) {
	if p != "" {
		ud.SetFlags(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetFlags()
		case SetData2Default:
			ud.UnsetFlags()
		}
	}
}

func UpdateWithFlags(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetFlags(p)
		} else {
			ud.UnsetFlags()
		}
	}
}

func (ud *UpdateDocument) SetSys_info(p *commons.SysInfo) *UpdateDocument {
	mName := fmt.Sprintf(SysInfoFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetSys_info() *UpdateDocument {
	mName := fmt.Sprintf(SysInfoFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

func (ud *UpdateDocument) setOrUnsetSys_info(p *commons.SysInfo, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetSys_info(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetSys_info()
		case SetData2Default:
			ud.UnsetSys_info()
		}
	}
}

func UpdateWithSys_info(p *commons.SysInfo) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetSys_info(p)
		} else {
			ud.UnsetSys_info()
		}
	}
}

func (ud *UpdateDocument) SetSysinfoModifiedAtNow() *UpdateDocument {
	mName := fmt.Sprintf(commons.SYSINFO_MODIFIEDAT)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: bson.NewDateTimeFromTime(time.Now())}
	})
	return ud
}
