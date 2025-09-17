package usercard

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	UserId      UnsetMode
	ObjType     UnsetMode
	Sysinfo     UnsetMode
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
func WithUserIdUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.UserId = m
	}
}
func WithObjTypeUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.ObjType = m
	}
}
func WithSysinfoUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Sysinfo = m
	}
}

type UpdateOption func(ud *UpdateDocument)
type UpdateOptions []UpdateOption

// GetUpdateDocument convenience method to create an update document from single updates instead of a whole object
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
func GetUpdateDocument(obj *UserCard, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	ud.setOrUnsetUserId(obj.UserId, uo.ResolveUnsetMode(uo.UserId))
	ud.setOrUnsetObjType(obj.ObjType, uo.ResolveUnsetMode(uo.ObjType))
	// if !obj.Sysinfo.IsZero() {
	//   ud.SetSysinfo ( obj.Sysinfo)
	// } else {
	ud.setOrUnsetSysinfo(obj.Sysinfo, uo.ResolveUnsetMode(uo.Sysinfo))
	// }

	return ud
}

//----- oId - object-id -  [oId]

// SetOId No Remarks
func (ud *UpdateDocument) SetOId(p primitive.ObjectID) *UpdateDocument {
	mName := fmt.Sprintf(OID)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetOId No Remarks
func (ud *UpdateDocument) UnsetOId() *UpdateDocument {
	mName := fmt.Sprintf(OID)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetOId No Remarks
func (ud *UpdateDocument) setOrUnsetOId(p primitive.ObjectID, um UnsetMode) {
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

//----- userId - string -  [userId]

// SetUserId No Remarks
func (ud *UpdateDocument) SetUserId(p string) *UpdateDocument {
	mName := fmt.Sprintf(USERID)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetUserId No Remarks
func (ud *UpdateDocument) UnsetUserId() *UpdateDocument {
	mName := fmt.Sprintf(USERID)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetUserId No Remarks
func (ud *UpdateDocument) setOrUnsetUserId(p string, um UnsetMode) {
	if p != "" {
		ud.SetUserId(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetUserId()
		case SetData2Default:
			ud.UnsetUserId()
		}
	}
}

func UpdateWithUserId(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetUserId(p)
		} else {
			ud.UnsetUserId()
		}
	}
}

//----- objType - string -  [objType]

// SetObjType No Remarks
func (ud *UpdateDocument) SetObjType(p string) *UpdateDocument {
	mName := fmt.Sprintf(OBJTYPE)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetObjType No Remarks
func (ud *UpdateDocument) UnsetObjType() *UpdateDocument {
	mName := fmt.Sprintf(OBJTYPE)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetObjType No Remarks
func (ud *UpdateDocument) setOrUnsetObjType(p string, um UnsetMode) {
	if p != "" {
		ud.SetObjType(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetObjType()
		case SetData2Default:
			ud.UnsetObjType()
		}
	}
}

func UpdateWithObjType(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetObjType(p)
		} else {
			ud.UnsetObjType()
		}
	}
}

// ----- sysinfo - ref-struct -  [sysinfo]
// SetSysinfo No Remarks
func (ud *UpdateDocument) SetSysinfo(p commons.SysInfo) *UpdateDocument {
	mName := fmt.Sprintf(SYSINFO)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetSysinfo No Remarks
func (ud *UpdateDocument) UnsetSysinfo() *UpdateDocument {
	mName := fmt.Sprintf(SYSINFO)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetSysinfo No Remarks - here2
func (ud *UpdateDocument) setOrUnsetSysinfo(p commons.SysInfo, um UnsetMode) {

	//----- ref-struct\n

	if !p.IsZero() {
		ud.SetSysinfo(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetSysinfo()
		case SetData2Default:
			ud.UnsetSysinfo()
		}
	}
}

func UpdateWithSysinfo(p commons.SysInfo) UpdateOption {
	return func(ud *UpdateDocument) {

		if !p.IsZero() {
			ud.SetSysinfo(p)
		} else {
			ud.UnsetSysinfo()
		}
	}
}
