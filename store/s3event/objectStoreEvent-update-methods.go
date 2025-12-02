package s3event

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/aws/aws-lambda-go/events"
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
	Bid           UnsetMode
	Et            UnsetMode
	Rip           UnsetMode
	Status        UnsetMode
	StatusReason  UnsetMode
	S3EventRecord UnsetMode
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
func WithRipUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Rip = m
	}
}
func WithStatusUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Status = m
	}
}
func WithStatusReasonUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.StatusReason = m
	}
}
func WithS3EventRecordUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.S3EventRecord = m
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
func GetUpdateDocument(obj *ObjectStoreEvent, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	ud.setOrUnset_bid(obj.Bid, uo.ResolveUnsetMode(uo.Bid))
	ud.setOrUnset_et(obj.Et, uo.ResolveUnsetMode(uo.Et))
	ud.setOrUnset_rip(obj.Rip, uo.ResolveUnsetMode(uo.Rip))
	ud.setOrUnset_status(obj.Status, uo.ResolveUnsetMode(uo.Status))
	ud.setOrUnset_status_reason(obj.StatusReason, uo.ResolveUnsetMode(uo.StatusReason))
	ud.setOrUnsetS3EventRecord(&obj.Event, uo.ResolveUnsetMode(uo.S3EventRecord))

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

// Set_rip No Remarks
func (ud *UpdateDocument) Set_rip(p bson.DateTime) *UpdateDocument {
	mName := fmt.Sprintf(RipFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// Unset_rip No Remarks
func (ud *UpdateDocument) Unset_rip() *UpdateDocument {
	mName := fmt.Sprintf(RipFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnset_rip No Remarks
func (ud *UpdateDocument) setOrUnset_rip(p bson.DateTime, um UnsetMode) {
	if p != 0 {
		ud.Set_rip(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.Unset_rip()
		case SetData2Default:
			ud.Unset_rip()
		}
	}
}

func UpdateWith_rip(p bson.DateTime) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != 0 {
			ud.Set_rip(p)
		} else {
			ud.Unset_rip()
		}
	}
}

// @tpm-schematics:start-region("-rip-field-update-section")
// @tpm-schematics:end-region("-rip-field-update-section")

// Set_status No Remarks
func (ud *UpdateDocument) Set_status(p string) *UpdateDocument {
	mName := fmt.Sprintf(StatusFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// Unset_status No Remarks
func (ud *UpdateDocument) Unset_status() *UpdateDocument {
	mName := fmt.Sprintf(StatusFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnset_status No Remarks
func (ud *UpdateDocument) setOrUnset_status(p string, um UnsetMode) {
	if p != "" {
		ud.Set_status(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.Unset_status()
		case SetData2Default:
			ud.Unset_status()
		}
	}
}

func UpdateWith_status(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.Set_status(p)
		} else {
			ud.Unset_status()
		}
	}
}

// @tpm-schematics:start-region("-status-field-update-section")
// @tpm-schematics:end-region("-status-field-update-section")

// Set_status_reason No Remarks
func (ud *UpdateDocument) Set_status_reason(p string) *UpdateDocument {
	mName := fmt.Sprintf(StatusReasonFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// Unset_status_reason No Remarks
func (ud *UpdateDocument) Unset_status_reason() *UpdateDocument {
	mName := fmt.Sprintf(StatusReasonFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnset_status_reason No Remarks
func (ud *UpdateDocument) setOrUnset_status_reason(p string, um UnsetMode) {
	if p != "" {
		ud.Set_status_reason(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.Unset_status_reason()
		case SetData2Default:
			ud.Unset_status_reason()
		}
	}
}

func UpdateWith_status_reason(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.Set_status_reason(p)
		} else {
			ud.Unset_status_reason()
		}
	}
}

// @tpm-schematics:start-region("-status-reason-field-update-section")
// @tpm-schematics:end-region("-status-reason-field-update-section")

// SetS3EventRecord No Remarks
func (ud *UpdateDocument) SetS3EventRecord(p *events.S3EventRecord) *UpdateDocument {
	mName := fmt.Sprintf(S3EventRecordFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetS3EventRecord No Remarks
func (ud *UpdateDocument) UnsetS3EventRecord() *UpdateDocument {
	mName := fmt.Sprintf(S3EventRecordFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetS3EventRecord No Remarks - here2
func (ud *UpdateDocument) setOrUnsetS3EventRecord(p *events.S3EventRecord, um UnsetMode) {
	if p != nil {
		ud.SetS3EventRecord(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetS3EventRecord()
		case SetData2Default:
			ud.UnsetS3EventRecord()
		}
	}
}

func UpdateWithS3EventRecord(p *events.S3EventRecord) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil {
			ud.SetS3EventRecord(p)
		} else {
			ud.UnsetS3EventRecord()
		}
	}
}

// @tpm-schematics:start-region("s3-event-record-field-update-section")
// @tpm-schematics:end-region("s3-event-record-field-update-section")

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
