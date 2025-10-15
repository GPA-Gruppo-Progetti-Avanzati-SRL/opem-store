package session

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
)

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

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

// GetUpdateDocumentFromOptions convenience method to create an update document from single updates instead of a whole object
func GetUpdateDocumentFromOptions(opts ...UpdateOption) UpdateDocument {
	ud := UpdateDocument{}
	for _, o := range opts {
		o(&ud)
	}

	return ud
}

// GetUpdateDocument
// Convenience method to create an Update Document from the values of the top fields of the object. The convenience is in the handling
// the unset because if I pass an empty struct to the update it generates an empty object anyway in the db. Handling the unset eliminates
// the issue and delete an existing value without creating an empty struct.
func GetUpdateDocument(obj *Session, opts ...UnsetOption) UpdateDocument {

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

// SetOId No Remarks
func (ud *UpdateDocument) SetOId(p bson.ObjectID) *UpdateDocument {
	mName := fmt.Sprintf(OIdFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetOId No Remarks
func (ud *UpdateDocument) UnsetOId() *UpdateDocument {
	mName := fmt.Sprintf(OIdFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetOId No Remarks
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

// @tpm-schematics:start-region("o-id-field-update-section")
// @tpm-schematics:end-region("o-id-field-update-section")

// SetUserid No Remarks
func (ud *UpdateDocument) SetUserid(p string) *UpdateDocument {
	mName := fmt.Sprintf(UseridFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetUserid No Remarks
func (ud *UpdateDocument) UnsetUserid() *UpdateDocument {
	mName := fmt.Sprintf(UseridFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetUserid No Remarks
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

// @tpm-schematics:start-region("userid-field-update-section")
// @tpm-schematics:end-region("userid-field-update-section")

// SetNickname No Remarks
func (ud *UpdateDocument) SetNickname(p string) *UpdateDocument {
	mName := fmt.Sprintf(NicknameFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetNickname No Remarks
func (ud *UpdateDocument) UnsetNickname() *UpdateDocument {
	mName := fmt.Sprintf(NicknameFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetNickname No Remarks
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

// @tpm-schematics:start-region("nickname-field-update-section")
// @tpm-schematics:end-region("nickname-field-update-section")

// SetRemote_addr No Remarks
func (ud *UpdateDocument) SetRemote_addr(p string) *UpdateDocument {
	mName := fmt.Sprintf(RemoteAddrFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetRemote_addr No Remarks
func (ud *UpdateDocument) UnsetRemote_addr() *UpdateDocument {
	mName := fmt.Sprintf(RemoteAddrFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetRemote_addr No Remarks
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

// @tpm-schematics:start-region("remote-addr-field-update-section")
// @tpm-schematics:end-region("remote-addr-field-update-section")

// SetFlags No Remarks
func (ud *UpdateDocument) SetFlags(p string) *UpdateDocument {
	mName := fmt.Sprintf(FlagsFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetFlags No Remarks
func (ud *UpdateDocument) UnsetFlags() *UpdateDocument {
	mName := fmt.Sprintf(FlagsFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetFlags No Remarks
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

// @tpm-schematics:start-region("flags-field-update-section")
// @tpm-schematics:end-region("flags-field-update-section")

// SetSys_info No Remarks
func (ud *UpdateDocument) SetSys_info(p *commons.SysInfo) *UpdateDocument {
	mName := fmt.Sprintf(SysInfoFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetSys_info No Remarks
func (ud *UpdateDocument) UnsetSys_info() *UpdateDocument {
	mName := fmt.Sprintf(SysInfoFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetSys_info No Remarks - here2
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

// @tpm-schematics:start-region("sys-info-field-update-section")

func (ud *UpdateDocument) SetSysinfoModifiedAtNow() *UpdateDocument {
	mName := fmt.Sprintf(commons.SYSINFO_MODIFIEDAT)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: bson.NewDateTimeFromTime(time.Now())}
	})
	return ud
}

// @tpm-schematics:end-region("sys-info-field-update-section")

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
