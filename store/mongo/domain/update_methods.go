package domain

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
	Bid         UnsetMode
	Et          UnsetMode
	Name        UnsetMode
	Description UnsetMode
	Langs       UnsetMode
	Members     UnsetMode
	Apps        UnsetMode
	SysInfo     UnsetMode
}

func (uo *UnsetOptions) ResolveUnsetMode(um UnsetMode) UnsetMode {
	if um == UnSpecified {
		um = uo.DefaultMode
	}
	return um
}

func WithDefaultUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.DefaultMode = m }
}
func WithOIdUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.OId = m }
}
func WithBidUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.Bid = m }
}
func WithEtUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.Et = m }
}
func WithNameUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.Name = m }
}
func WithDescriptionUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.Description = m }
}
func WithLangsUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.Langs = m }
}
func WithMembersUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.Members = m }
}
func WithAppsUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.Apps = m }
}
func WithSysInfoUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.SysInfo = m }
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

func GetUpdateDocument(obj *model.Domain, opts ...UnsetOption) UpdateDocument {
	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}
	ud := UpdateDocument{}
	ud.setOrUnset_bid(obj.Bid, uo.ResolveUnsetMode(uo.Bid))
	ud.setOrUnset_et(obj.Et, uo.ResolveUnsetMode(uo.Et))
	ud.setOrUnsetName(obj.Name, uo.ResolveUnsetMode(uo.Name))
	ud.setOrUnsetDescription(obj.Description, uo.ResolveUnsetMode(uo.Description))
	ud.setOrUnsetLangs(obj.Langs, uo.ResolveUnsetMode(uo.Langs))
	ud.setOrUnsetMembers(obj.Members, uo.ResolveUnsetMode(uo.Members))
	ud.setOrUnsetApps(obj.Apps, uo.ResolveUnsetMode(uo.Apps))
	ud.setOrUnsetSys_info(&obj.SysInfo, uo.ResolveUnsetMode(uo.SysInfo))
	return ud
}

func (ud *UpdateDocument) SetOId(p bson.ObjectID) *UpdateDocument {
	mName := fmt.Sprintf(OIdFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetOId() *UpdateDocument {
	mName := fmt.Sprintf(OIdFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
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

func (ud *UpdateDocument) Set_bid(p string) *UpdateDocument {
	mName := fmt.Sprintf(BidFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) Unset_bid() *UpdateDocument {
	mName := fmt.Sprintf(BidFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

func (ud *UpdateDocument) setOrUnset_bid(p string, um UnsetMode) {
	if p != "" {
		ud.Set_bid(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.Unset_bid()
		case SetData2Default:
			ud.Unset_bid()
		}
	}
}

func UpdateWith_bid(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.Set_bid(p)
		} else {
			ud.Unset_bid()
		}
	}
}

func (ud *UpdateDocument) Set_et(p string) *UpdateDocument {
	mName := fmt.Sprintf(EtFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) Unset_et() *UpdateDocument {
	mName := fmt.Sprintf(EtFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

func (ud *UpdateDocument) setOrUnset_et(p string, um UnsetMode) {
	if p != "" {
		ud.Set_et(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.Unset_et()
		case SetData2Default:
			ud.Unset_et()
		}
	}
}

func UpdateWith_et(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.Set_et(p)
		} else {
			ud.Unset_et()
		}
	}
}

func (ud *UpdateDocument) SetName(p string) *UpdateDocument {
	mName := fmt.Sprintf(NameFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetName() *UpdateDocument {
	mName := fmt.Sprintf(NameFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

func (ud *UpdateDocument) setOrUnsetName(p string, um UnsetMode) {
	if p != "" {
		ud.SetName(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetName()
		case SetData2Default:
			ud.UnsetName()
		}
	}
}

func UpdateWithName(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetName(p)
		} else {
			ud.UnsetName()
		}
	}
}

func (ud *UpdateDocument) SetDescription(p string) *UpdateDocument {
	mName := fmt.Sprintf(DescriptionFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetDescription() *UpdateDocument {
	mName := fmt.Sprintf(DescriptionFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

func (ud *UpdateDocument) setOrUnsetDescription(p string, um UnsetMode) {
	if p != "" {
		ud.SetDescription(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetDescription()
		case SetData2Default:
			ud.UnsetDescription()
		}
	}
}

func UpdateWithDescription(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetDescription(p)
		} else {
			ud.UnsetDescription()
		}
	}
}

func (ud *UpdateDocument) SetLangs(p string) *UpdateDocument {
	mName := fmt.Sprintf(LangsFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetLangs() *UpdateDocument {
	mName := fmt.Sprintf(LangsFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

func (ud *UpdateDocument) setOrUnsetLangs(p string, um UnsetMode) {
	if p != "" {
		ud.SetLangs(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetLangs()
		case SetData2Default:
			ud.UnsetLangs()
		}
	}
}

func UpdateWithLangs(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetLangs(p)
		} else {
			ud.UnsetLangs()
		}
	}
}

func (ud *UpdateDocument) SetMembers(p []model.Member) *UpdateDocument {
	mName := fmt.Sprintf(MembersFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetMembers() *UpdateDocument {
	mName := fmt.Sprintf(MembersFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

func (ud *UpdateDocument) setOrUnsetMembers(p []model.Member, um UnsetMode) {
	if len(p) > 0 {
		ud.SetMembers(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetMembers()
		case SetData2Default:
			ud.UnsetMembers()
		}
	}
}

func UpdateWithMembers(p []model.Member) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetMembers(p)
		} else {
			ud.UnsetMembers()
		}
	}
}

func (ud *UpdateDocument) SetApps(p []commons.App) *UpdateDocument {
	mName := fmt.Sprintf(AppsFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetApps() *UpdateDocument {
	mName := fmt.Sprintf(AppsFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

func (ud *UpdateDocument) setOrUnsetApps(p []commons.App, um UnsetMode) {
	if len(p) > 0 {
		ud.SetApps(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetApps()
		case SetData2Default:
			ud.UnsetApps()
		}
	}
}

func UpdateWithApps(p []commons.App) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetApps(p)
		} else {
			ud.UnsetApps()
		}
	}
}

func (ud *UpdateDocument) SetSys_info(p *commons.SysInfo) *UpdateDocument {
	mName := fmt.Sprintf(SysInfoFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetSys_info() *UpdateDocument {
	mName := fmt.Sprintf(SysInfoFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
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
