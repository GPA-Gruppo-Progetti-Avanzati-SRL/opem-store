package filerow

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
	DefaultMode   UnsetMode
	OId           UnsetMode
	Domain        UnsetMode
	Site          UnsetMode
	Bid           UnsetMode
	Et            UnsetMode
	Status        UnsetMode
	RowData       UnsetMode
	RowNumber     UnsetMode
	RowDataFormat UnsetMode
	Errs          UnsetMode
	File          UnsetMode
	SysInfo       UnsetMode
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
func WithStatusUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Status = m
	}
}
func WithRowDataUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.RowData = m
	}
}
func WithRowNumberUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.RowNumber = m
	}
}
func WithRowDataFormatUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.RowDataFormat = m
	}
}
func WithErrsUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Errs = m
	}
}
func WithFileUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.File = m
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
func GetUpdateDocument(obj *FileRow, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	ud.setOrUnsetDomain(obj.Domain, uo.ResolveUnsetMode(uo.Domain))
	ud.setOrUnsetSite(obj.Site, uo.ResolveUnsetMode(uo.Site))
	ud.setOrUnset_bid(obj.Bid, uo.ResolveUnsetMode(uo.Bid))
	ud.setOrUnset_et(obj.Et, uo.ResolveUnsetMode(uo.Et))
	ud.setOrUnsetStatus(obj.Status, uo.ResolveUnsetMode(uo.Status))
	ud.setOrUnsetRow_data(obj.RowData, uo.ResolveUnsetMode(uo.RowData))
	ud.setOrUnsetRow_number(obj.RowNumber, uo.ResolveUnsetMode(uo.RowNumber))
	ud.setOrUnsetRow_data_format(obj.RowDataFormat, uo.ResolveUnsetMode(uo.RowDataFormat))
	ud.setOrUnsetErrs(obj.Errs, uo.ResolveUnsetMode(uo.Errs))
	ud.setOrUnsetFile(&obj.File, uo.ResolveUnsetMode(uo.File))
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

// SetStatus No Remarks
func (ud *UpdateDocument) SetStatus(p string) *UpdateDocument {
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

// setOrUnsetStatus No Remarks
func (ud *UpdateDocument) setOrUnsetStatus(p string, um UnsetMode) {
	if p != "" {
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

func UpdateWithStatus(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetStatus(p)
		} else {
			ud.UnsetStatus()
		}
	}
}

// @tpm-schematics:start-region("status-field-update-section")
// @tpm-schematics:end-region("status-field-update-section")

// SetRow_data No Remarks
func (ud *UpdateDocument) SetRow_data(p string) *UpdateDocument {
	mName := fmt.Sprintf(RowDataFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetRow_data No Remarks
func (ud *UpdateDocument) UnsetRow_data() *UpdateDocument {
	mName := fmt.Sprintf(RowDataFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetRow_data No Remarks
func (ud *UpdateDocument) setOrUnsetRow_data(p string, um UnsetMode) {
	if p != "" {
		ud.SetRow_data(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetRow_data()
		case SetData2Default:
			ud.UnsetRow_data()
		}
	}
}

func UpdateWithRow_data(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetRow_data(p)
		} else {
			ud.UnsetRow_data()
		}
	}
}

// @tpm-schematics:start-region("row-data-field-update-section")
// @tpm-schematics:end-region("row-data-field-update-section")

// SetRow_number No Remarks
func (ud *UpdateDocument) SetRow_number(p int32) *UpdateDocument {
	mName := fmt.Sprintf(RowNumberFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetRow_number No Remarks
func (ud *UpdateDocument) UnsetRow_number() *UpdateDocument {
	mName := fmt.Sprintf(RowNumberFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetRow_number No Remarks
func (ud *UpdateDocument) setOrUnsetRow_number(p int32, um UnsetMode) {
	if p != 0 {
		ud.SetRow_number(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetRow_number()
		case SetData2Default:
			ud.UnsetRow_number()
		}
	}
}

func UpdateWithRow_number(p int32) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != 0 {
			ud.SetRow_number(p)
		} else {
			ud.UnsetRow_number()
		}
	}
}

// @tpm-schematics:start-region("row-number-field-update-section")
// @tpm-schematics:end-region("row-number-field-update-section")

// SetRow_data_format No Remarks
func (ud *UpdateDocument) SetRow_data_format(p string) *UpdateDocument {
	mName := fmt.Sprintf(RowDataFormatFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetRow_data_format No Remarks
func (ud *UpdateDocument) UnsetRow_data_format() *UpdateDocument {
	mName := fmt.Sprintf(RowDataFormatFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetRow_data_format No Remarks
func (ud *UpdateDocument) setOrUnsetRow_data_format(p string, um UnsetMode) {
	if p != "" {
		ud.SetRow_data_format(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetRow_data_format()
		case SetData2Default:
			ud.UnsetRow_data_format()
		}
	}
}

func UpdateWithRow_data_format(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetRow_data_format(p)
		} else {
			ud.UnsetRow_data_format()
		}
	}
}

// @tpm-schematics:start-region("row-data-format-field-update-section")
// @tpm-schematics:end-region("row-data-format-field-update-section")

// SetErrs No Remarks
func (ud *UpdateDocument) SetErrs(p []string) *UpdateDocument {
	mName := fmt.Sprintf(ErrsFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetErrs No Remarks
func (ud *UpdateDocument) UnsetErrs() *UpdateDocument {
	mName := fmt.Sprintf(ErrsFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetErrs No Remarks - here2
func (ud *UpdateDocument) setOrUnsetErrs(p []string, um UnsetMode) {
	if len(p) > 0 {
		ud.SetErrs(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetErrs()
		case SetData2Default:
			ud.UnsetErrs()
		}
	}
}

func UpdateWithErrs(p []string) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetErrs(p)
		} else {
			ud.UnsetErrs()
		}
	}
}

// @tpm-schematics:start-region("errs-field-update-section")
// @tpm-schematics:end-region("errs-field-update-section")

// SetFile No Remarks
func (ud *UpdateDocument) SetFile(p *commons.BidTextPair) *UpdateDocument {
	mName := fmt.Sprintf(FileFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetFile No Remarks
func (ud *UpdateDocument) UnsetFile() *UpdateDocument {
	mName := fmt.Sprintf(FileFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetFile No Remarks - here2
func (ud *UpdateDocument) setOrUnsetFile(p *commons.BidTextPair, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetFile(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetFile()
		case SetData2Default:
			ud.UnsetFile()
		}
	}
}

func UpdateWithFile(p *commons.BidTextPair) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetFile(p)
		} else {
			ud.UnsetFile()
		}
	}
}

// @tpm-schematics:start-region("file-field-update-section")
// @tpm-schematics:end-region("file-field-update-section")

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
