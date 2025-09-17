package file

// This file contains the paths to the field in the generated entity.
// A path is a string with all the identifiers from the root document to the single leaves.
// In case of maps and arrays place holder for the key (%s) or the index %d have been provided.

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

const (
	OIdFieldName                = "_id"
	FnFieldName                 = "fn"
	DescrFieldName              = "descr"
	RoleFieldName               = "role"
	EntRefsFieldName            = "entRefs"
	EntRefs_IFieldName          = "entRefs.%d"
	MetadataFieldName           = "metadata"
	VrntsFieldName              = "vrnts"
	Vrnts_IFieldName            = "vrnts.%d"
	SysInfoFieldName            = "sysInfo"
	SysInfo_StatusFieldName     = "sysInfo.status"
	SysInfo_CreatedatFieldName  = "sysInfo.createdat"
	SysInfo_ModifiedatFieldName = "sysInfo.modifiedat"
)

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
