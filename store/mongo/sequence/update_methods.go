package sequence

import (
	"fmt"
	"time"

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
	Domain      UnsetMode
	Site        UnsetMode
	Value       UnsetMode
	Format      UnsetMode
	Prefix      UnsetMode
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
func WithDomainUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.Domain = m }
}
func WithSiteUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.Site = m }
}
func WithValueUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.Value = m }
}
func WithFormatUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.Format = m }
}
func WithPrefixUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.Prefix = m }
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

func GetUpdateDocument(obj *model.Sequence, opts ...UnsetOption) UpdateDocument {
	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}
	ud := UpdateDocument{}
	ud.setOrUnset_bid(obj.Bid, uo.ResolveUnsetMode(uo.Bid))
	ud.setOrUnset_et(obj.Et, uo.ResolveUnsetMode(uo.Et))
	ud.setOrUnsetDomain(obj.Domain, uo.ResolveUnsetMode(uo.Domain))
	ud.setOrUnsetSite(obj.Site, uo.ResolveUnsetMode(uo.Site))
	ud.setOrUnsetValue(obj.Value, uo.ResolveUnsetMode(uo.Value))
	ud.setOrUnsetFormat(obj.Format, uo.ResolveUnsetMode(uo.Format))
	ud.setOrUnsetPrefix(obj.Prefix, uo.ResolveUnsetMode(uo.Prefix))
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

func (ud *UpdateDocument) SetDomain(p string) *UpdateDocument {
	mName := fmt.Sprintf(DomainFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetDomain() *UpdateDocument {
	mName := fmt.Sprintf(DomainFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

func (ud *UpdateDocument) setOrUnsetDomain(p string, um UnsetMode) {
	if p != "" {
		ud.SetDomain(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetDomain()
		case SetData2Default:
			ud.UnsetDomain()
		}
	}
}

func UpdateWithDomain(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetDomain(p)
		} else {
			ud.UnsetDomain()
		}
	}
}

func (ud *UpdateDocument) SetSite(p string) *UpdateDocument {
	mName := fmt.Sprintf(SiteFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetSite() *UpdateDocument {
	mName := fmt.Sprintf(SiteFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

func (ud *UpdateDocument) setOrUnsetSite(p string, um UnsetMode) {
	if p != "" {
		ud.SetSite(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetSite()
		case SetData2Default:
			ud.UnsetSite()
		}
	}
}

func UpdateWithSite(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetSite(p)
		} else {
			ud.UnsetSite()
		}
	}
}

func (ud *UpdateDocument) SetValue(p int32) *UpdateDocument {
	mName := fmt.Sprintf(ValueFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetValue() *UpdateDocument {
	mName := fmt.Sprintf(ValueFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

func (ud *UpdateDocument) setOrUnsetValue(p int32, um UnsetMode) {
	if p != 0 {
		ud.SetValue(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetValue()
		case SetData2Default:
			ud.UnsetValue()
		}
	}
}

func UpdateWithValue(p int32) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != 0 {
			ud.SetValue(p)
		} else {
			ud.UnsetValue()
		}
	}
}

func (ud *UpdateDocument) IncValue(p int32) *UpdateDocument {
	mName := fmt.Sprintf(ValueFieldName)
	ud.Inc().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func UpdateWithIncrementValue(increment int32) UpdateOption {
	return func(ud *UpdateDocument) {
		if increment != 0 {
			ud.IncValue(increment)
		}
	}
}

func (ud *UpdateDocument) SetFormat(p string) *UpdateDocument {
	mName := fmt.Sprintf(FormatFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetFormat() *UpdateDocument {
	mName := fmt.Sprintf(FormatFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

func (ud *UpdateDocument) setOrUnsetFormat(p string, um UnsetMode) {
	if p != "" {
		ud.SetFormat(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetFormat()
		case SetData2Default:
			ud.UnsetFormat()
		}
	}
}

func UpdateWithFormat(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetFormat(p)
		} else {
			ud.UnsetFormat()
		}
	}
}

func (ud *UpdateDocument) SetPrefix(p string) *UpdateDocument {
	mName := fmt.Sprintf(PrefixFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetPrefix() *UpdateDocument {
	mName := fmt.Sprintf(PrefixFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

func (ud *UpdateDocument) setOrUnsetPrefix(p string, um UnsetMode) {
	if p != "" {
		ud.SetPrefix(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetPrefix()
		case SetData2Default:
			ud.UnsetPrefix()
		}
	}
}

func UpdateWithPrefix(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetPrefix(p)
		} else {
			ud.UnsetPrefix()
		}
	}
}
