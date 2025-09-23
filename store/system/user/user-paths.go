package user

// This file contains the paths to the field in the generated entity.
// A path is a string with all the identifiers from the root document to the single leaves.
// In case of maps and arrays place holder for the key (%s) or the index %d have been provided.

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

const (
	OIdFieldName                     = "_id"
	NicknameFieldName                = "nickname"
	ObjTypeFieldName                 = "obj_type"
	FirstnameFieldName               = "firstname"
	LastnameFieldName                = "lastname"
	EmailFieldName                   = "email"
	PasswordFieldName                = "password"
	RolesFieldName                   = "roles"
	Roles_IFieldName                 = "roles.%d"
	SysInfoFieldName                 = "sys_info"
	SysInfo_StatusFieldName          = "sys_info.status"
	SysInfo_CreatedAtFieldName       = "sys_info.created_at"
	SysInfo_ModifiedAtFieldName      = "sys_info.modified_at"
	ProfilePictureFieldName          = "profile_picture"
	ProfilePicture_OIdFieldName      = "profile_picture._id"
	ProfilePicture_SrcSetFieldName   = "profile_picture.src_set"
	ProfilePicture_SrcSet_IFieldName = "profile_picture.src_set.%d"
)

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
