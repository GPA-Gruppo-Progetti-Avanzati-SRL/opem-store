package box

import (
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"

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
	Domain      UnsetMode
	Site        UnsetMode
	Bid         UnsetMode
	Et          UnsetMode
	Magazzino   UnsetMode
	Prodotto    UnsetMode
	FocalPoint  UnsetMode
	Info        UnsetMode
	Status      UnsetMode
	Recipient   UnsetMode
	Events      UnsetMode
	Notes       UnsetMode
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
func WithDomainUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Domain = m
	}
}
func WithSiteUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Site = m
	}
}
func WithBidUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Bid = m
	}
}
func WithEtUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Et = m
	}
}
func WithMagazzinoUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Magazzino = m
	}
}
func WithProdottoUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Prodotto = m
	}
}
func WithFocalPointUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.FocalPoint = m
	}
}
func WithInfoUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Info = m
	}
}
func WithStatusUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Status = m
	}
}
func WithRecipientUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Recipient = m
	}
}
func WithEventsUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Events = m
	}
}
func WithNotesUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Notes = m
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
func GetUpdateDocument(obj *Box, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	ud.setOrUnsetDomain(obj.Domain, uo.ResolveUnsetMode(uo.Domain))
	ud.setOrUnsetSite(obj.Site, uo.ResolveUnsetMode(uo.Site))
	ud.setOrUnset_bid(obj.Bid, uo.ResolveUnsetMode(uo.Bid))
	ud.setOrUnset_et(obj.Et, uo.ResolveUnsetMode(uo.Et))
	ud.setOrUnsetMagazzino(&obj.Magazzino, uo.ResolveUnsetMode(uo.Magazzino))
	ud.setOrUnsetProdotto(&obj.Prodotto, uo.ResolveUnsetMode(uo.Prodotto))
	ud.setOrUnsetFocal_point(&obj.FocalPoint, uo.ResolveUnsetMode(uo.FocalPoint))
	ud.setOrUnsetInfo(&obj.Info, uo.ResolveUnsetMode(uo.Info))
	ud.setOrUnsetStatus(&obj.Status, uo.ResolveUnsetMode(uo.Status))
	ud.setOrUnsetRecipient(&obj.Recipient, uo.ResolveUnsetMode(uo.Recipient))
	ud.setOrUnsetEvents(obj.Events, uo.ResolveUnsetMode(uo.Events))
	ud.setOrUnsetNotes(obj.Notes, uo.ResolveUnsetMode(uo.Notes))
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

// SetDomain No Remarks
func (ud *UpdateDocument) SetDomain(p string) *UpdateDocument {
	mName := fmt.Sprintf(DomainFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetDomain No Remarks
func (ud *UpdateDocument) UnsetDomain() *UpdateDocument {
	mName := fmt.Sprintf(DomainFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetDomain No Remarks
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

// @tpm-schematics:start-region("domain-field-update-section")
// @tpm-schematics:end-region("domain-field-update-section")

// SetSite No Remarks
func (ud *UpdateDocument) SetSite(p string) *UpdateDocument {
	mName := fmt.Sprintf(SiteFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetSite No Remarks
func (ud *UpdateDocument) UnsetSite() *UpdateDocument {
	mName := fmt.Sprintf(SiteFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetSite No Remarks
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

// @tpm-schematics:start-region("site-field-update-section")
// @tpm-schematics:end-region("site-field-update-section")

// Set_bid No Remarks
func (ud *UpdateDocument) Set_bid(p string) *UpdateDocument {
	mName := fmt.Sprintf(BidFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// Unset_bid No Remarks
func (ud *UpdateDocument) Unset_bid() *UpdateDocument {
	mName := fmt.Sprintf(BidFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnset_bid No Remarks
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

// @tpm-schematics:start-region("-bid-field-update-section")
// @tpm-schematics:end-region("-bid-field-update-section")

// Set_et No Remarks
func (ud *UpdateDocument) Set_et(p string) *UpdateDocument {
	mName := fmt.Sprintf(EtFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// Unset_et No Remarks
func (ud *UpdateDocument) Unset_et() *UpdateDocument {
	mName := fmt.Sprintf(EtFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnset_et No Remarks
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

// @tpm-schematics:start-region("-et-field-update-section")
// @tpm-schematics:end-region("-et-field-update-section")

// SetMagazzino No Remarks
func (ud *UpdateDocument) SetMagazzino(p *commons.BidTextPair) *UpdateDocument {
	mName := fmt.Sprintf(MagazzinoFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetMagazzino No Remarks
func (ud *UpdateDocument) UnsetMagazzino() *UpdateDocument {
	mName := fmt.Sprintf(MagazzinoFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetMagazzino No Remarks - here2
func (ud *UpdateDocument) setOrUnsetMagazzino(p *commons.BidTextPair, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetMagazzino(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetMagazzino()
		case SetData2Default:
			ud.UnsetMagazzino()
		}
	}
}

func UpdateWithMagazzino(p *commons.BidTextPair) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetMagazzino(p)
		} else {
			ud.UnsetMagazzino()
		}
	}
}

// @tpm-schematics:start-region("magazzino-field-update-section")
// @tpm-schematics:end-region("magazzino-field-update-section")

// SetProdotto No Remarks
func (ud *UpdateDocument) SetProdotto(p *commons.BidTextPair) *UpdateDocument {
	mName := fmt.Sprintf(ProdottoFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetProdotto No Remarks
func (ud *UpdateDocument) UnsetProdotto() *UpdateDocument {
	mName := fmt.Sprintf(ProdottoFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetProdotto No Remarks - here2
func (ud *UpdateDocument) setOrUnsetProdotto(p *commons.BidTextPair, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetProdotto(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetProdotto()
		case SetData2Default:
			ud.UnsetProdotto()
		}
	}
}

func UpdateWithProdotto(p *commons.BidTextPair) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetProdotto(p)
		} else {
			ud.UnsetProdotto()
		}
	}
}

// @tpm-schematics:start-region("prodotto-field-update-section")
// @tpm-schematics:end-region("prodotto-field-update-section")

// SetFocal_point No Remarks
func (ud *UpdateDocument) SetFocal_point(p *commons.BidTextPair) *UpdateDocument {
	mName := fmt.Sprintf(FocalPointFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetFocal_point No Remarks
func (ud *UpdateDocument) UnsetFocal_point() *UpdateDocument {
	mName := fmt.Sprintf(FocalPointFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetFocal_point No Remarks - here2
func (ud *UpdateDocument) setOrUnsetFocal_point(p *commons.BidTextPair, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetFocal_point(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetFocal_point()
		case SetData2Default:
			ud.UnsetFocal_point()
		}
	}
}

func UpdateWithFocal_point(p *commons.BidTextPair) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetFocal_point(p)
		} else {
			ud.UnsetFocal_point()
		}
	}
}

// @tpm-schematics:start-region("focal-point-field-update-section")
// @tpm-schematics:end-region("focal-point-field-update-section")

// SetInfo No Remarks
func (ud *UpdateDocument) SetInfo(p *Info) *UpdateDocument {
	mName := fmt.Sprintf(InfoFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetInfo No Remarks
func (ud *UpdateDocument) UnsetInfo() *UpdateDocument {
	mName := fmt.Sprintf(InfoFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetInfo No Remarks - here2
func (ud *UpdateDocument) setOrUnsetInfo(p *Info, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetInfo(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetInfo()
		case SetData2Default:
			ud.UnsetInfo()
		}
	}
}

func UpdateWithInfo(p *Info) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetInfo(p)
		} else {
			ud.UnsetInfo()
		}
	}
}

// @tpm-schematics:start-region("info-field-update-section")
// @tpm-schematics:end-region("info-field-update-section")

// SetStatus No Remarks
func (ud *UpdateDocument) SetStatus(p *Status) *UpdateDocument {
	mName := fmt.Sprintf(StatusFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetStatus No Remarks
func (ud *UpdateDocument) UnsetStatus() *UpdateDocument {
	mName := fmt.Sprintf(StatusFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetStatus No Remarks - here2
func (ud *UpdateDocument) setOrUnsetStatus(p *Status, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetStatus(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetStatus()
		case SetData2Default:
			ud.UnsetStatus()
		}
	}
}

func UpdateWithStatus(p *Status) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetStatus(p)
		} else {
			ud.UnsetStatus()
		}
	}
}

// @tpm-schematics:start-region("status-field-update-section")
// @tpm-schematics:end-region("status-field-update-section")

// SetRecipient No Remarks
func (ud *UpdateDocument) SetRecipient(p *commons.Address) *UpdateDocument {
	mName := fmt.Sprintf(RecipientFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetRecipient No Remarks
func (ud *UpdateDocument) UnsetRecipient() *UpdateDocument {
	mName := fmt.Sprintf(RecipientFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetRecipient No Remarks - here2
func (ud *UpdateDocument) setOrUnsetRecipient(p *commons.Address, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetRecipient(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetRecipient()
		case SetData2Default:
			ud.UnsetRecipient()
		}
	}
}

func UpdateWithRecipient(p *commons.Address) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetRecipient(p)
		} else {
			ud.UnsetRecipient()
		}
	}
}

// @tpm-schematics:start-region("recipient-field-update-section")
// @tpm-schematics:end-region("recipient-field-update-section")

// SetEvents No Remarks
func (ud *UpdateDocument) SetEvents(p []commons.Event) *UpdateDocument {
	mName := fmt.Sprintf(EventsFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetEvents No Remarks
func (ud *UpdateDocument) UnsetEvents() *UpdateDocument {
	mName := fmt.Sprintf(EventsFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetEvents No Remarks - here2
func (ud *UpdateDocument) setOrUnsetEvents(p []commons.Event, um UnsetMode) {
	if len(p) > 0 {
		ud.SetEvents(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetEvents()
		case SetData2Default:
			ud.UnsetEvents()
		}
	}
}

func UpdateWithEvents(p []commons.Event) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetEvents(p)
		} else {
			ud.UnsetEvents()
		}
	}
}

// @tpm-schematics:start-region("events-field-update-section")
// @tpm-schematics:end-region("events-field-update-section")

// SetNotes No Remarks
func (ud *UpdateDocument) SetNotes(p []commons.Note) *UpdateDocument {
	mName := fmt.Sprintf(NotesFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetNotes No Remarks
func (ud *UpdateDocument) UnsetNotes() *UpdateDocument {
	mName := fmt.Sprintf(NotesFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetNotes No Remarks - here2
func (ud *UpdateDocument) setOrUnsetNotes(p []commons.Note, um UnsetMode) {
	if len(p) > 0 {
		ud.SetNotes(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetNotes()
		case SetData2Default:
			ud.UnsetNotes()
		}
	}
}

func UpdateWithNotes(p []commons.Note) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetNotes(p)
		} else {
			ud.UnsetNotes()
		}
	}
}

// @tpm-schematics:start-region("notes-field-update-section")

func UpdateWithAddNote(p commons.Note) UpdateOption {
	return func(ud *UpdateDocument) {
		if !p.IsZero() {
			mName := fmt.Sprintf(NotesFieldName)
			ud.Push().Add(func() bson.E {
				return bson.E{Key: mName, Value: p}
			})
		}
	}
}

// @tpm-schematics:end-region("notes-field-update-section")

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
// @tpm-schematics:end-region("sys-info-field-update-section")

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
