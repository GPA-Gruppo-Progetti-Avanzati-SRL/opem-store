package box

// This file contains the paths to the field in the generated entity.
// A path is a string with all the identifiers from the root document to the single leaves.
// In case of maps and arrays place holder for the key (%s) or the index %d have been provided.

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

const (
	OIdFieldName                     = "_id"
	DomainFieldName                  = "domain"
	SiteFieldName                    = "site"
	BidFieldName                     = "_bid"
	EtFieldName                      = "_et"
	MagazzinoFieldName               = "magazzino"
	Magazzino_BidFieldName           = "magazzino.bid"
	Prodotto_BidFieldName            = "prodotto.bid"
	FocalPoint_BidFieldName          = "focal_point.bid"
	Recipient_County_BidFieldName    = "recipient.county.bid"
	Recipient_Townhall_BidFieldName  = "recipient.townhall.bid"
	Magazzino_TextFieldName          = "magazzino.text"
	Prodotto_TextFieldName           = "prodotto.text"
	FocalPoint_TextFieldName         = "focal_point.text"
	Recipient_County_TextFieldName   = "recipient.county.text"
	Recipient_Townhall_TextFieldName = "recipient.townhall.text"
	ProdottoFieldName                = "prodotto"
	FocalPointFieldName              = "focal_point"
	InfoFieldName                    = "info"
	Info_BoxCodeFieldName            = "info.box_code"
	Info_PackageCodeFieldName        = "info.package_code"
	StatusFieldName                  = "status"
	Status_NumInFieldName            = "status.num_in"
	Status_NumOutFieldName           = "status.num_out"
	Status_StatusFieldName           = "status.status"
	RecipientFieldName               = "recipient"
	Recipient_TypeFieldName          = "recipient.type"
	Recipient_CountyFieldName        = "recipient.county"
	Recipient_TownhallFieldName      = "recipient.townhall"
	Recipient_ZipcodeFieldName       = "recipient.zipcode"
	Recipient_AddressFieldName       = "recipient.address"
	Recipient_AttnToFieldName        = "recipient.attn_to"
	EventsFieldName                  = "events"
	Events_IFieldName                = "events.%d"
	NotesFieldName                   = "notes"
	Notes_IFieldName                 = "notes.%d"
	SysInfoFieldName                 = "sys_info"
	SysInfo_StatusFieldName          = "sys_info.status"
	SysInfo_CreatedAtFieldName       = "sys_info.created_at"
	SysInfo_ModifiedAtFieldName      = "sys_info.modified_at"
)

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
