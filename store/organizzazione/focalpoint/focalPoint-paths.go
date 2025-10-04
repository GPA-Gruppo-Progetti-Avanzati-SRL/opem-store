package focalpoint

// This file contains the paths to the field in the generated entity.
// A path is a string with all the identifiers from the root document to the single leaves.
// In case of maps and arrays place holder for the key (%s) or the index %d have been provided.

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

const (
	OIdFieldName                   = "_id"
	DomainFieldName                = "domain"
	SiteFieldName                  = "site"
	BidFieldName                   = "_bid"
	EtFieldName                    = "_et"
	OfficerNameFieldName           = "officer_name"
	ExtendedNameFieldName          = "extended_name"
	ReducedNameFieldName           = "reduced_name"
	AddressFieldName               = "address"
	Address_TypeFieldName          = "address.type"
	Address_CountyFieldName        = "address.county"
	Address_County_BidFieldName    = "address.county.bid"
	Address_Townhall_BidFieldName  = "address.townhall.bid"
	Address_County_TextFieldName   = "address.county.text"
	Address_Townhall_TextFieldName = "address.townhall.text"
	Address_TownhallFieldName      = "address.townhall"
	Address_ZipcodeFieldName       = "address.zipcode"
	Address_AddressFieldName       = "address.address"
	Address_AttnToFieldName        = "address.attn_to"
	SysInfoFieldName               = "sys_info"
	SysInfo_StatusFieldName        = "sys_info.status"
	SysInfo_CreatedAtFieldName     = "sys_info.created_at"
	SysInfo_ModifiedAtFieldName    = "sys_info.modified_at"
)

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
