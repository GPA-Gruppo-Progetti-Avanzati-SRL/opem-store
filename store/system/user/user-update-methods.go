package user

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
	DefaultMode    UnsetMode
	OId            UnsetMode
	Nickname       UnsetMode
	Et             UnsetMode
	Firstname      UnsetMode
	Lastname       UnsetMode
	Email          UnsetMode
	Password       UnsetMode
	Roles          UnsetMode
	SysInfo        UnsetMode
	ProfilePicture UnsetMode
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
func WithNicknameUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Nickname = m
	}
}
func WithEtUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Et = m
	}
}
func WithFirstnameUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Firstname = m
	}
}
func WithLastnameUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Lastname = m
	}
}
func WithEmailUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Email = m
	}
}
func WithPasswordUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Password = m
	}
}
func WithRolesUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Roles = m
	}
}
func WithSysInfoUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.SysInfo = m
	}
}
func WithProfilePictureUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.ProfilePicture = m
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
func GetUpdateDocument(obj *User, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	ud.setOrUnsetNickname(obj.Nickname, uo.ResolveUnsetMode(uo.Nickname))
	ud.setOrUnset_et(obj.Et, uo.ResolveUnsetMode(uo.Et))
	ud.setOrUnsetFirstname(obj.Firstname, uo.ResolveUnsetMode(uo.Firstname))
	ud.setOrUnsetLastname(obj.Lastname, uo.ResolveUnsetMode(uo.Lastname))
	ud.setOrUnsetEmail(obj.Email, uo.ResolveUnsetMode(uo.Email))
	ud.setOrUnsetPassword(obj.Password, uo.ResolveUnsetMode(uo.Password))
	ud.setOrUnsetRoles(obj.Roles, uo.ResolveUnsetMode(uo.Roles))
	ud.setOrUnsetSys_info(&obj.SysInfo, uo.ResolveUnsetMode(uo.SysInfo))
	ud.setOrUnsetProfile_picture(&obj.ProfilePicture, uo.ResolveUnsetMode(uo.ProfilePicture))

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

// SetNickname No Remarks
func (ud *UpdateDocument) SetNickname(p string) *UpdateDocument {
	mName := fmt.Sprintf(NicknameFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetNickname No Remarks
func (ud *UpdateDocument) UnsetNickname() *UpdateDocument {
	mName := fmt.Sprintf(NicknameFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetNickname No Remarks
func (ud *UpdateDocument) setOrUnsetNickname(p string, um UnsetMode) {
	if p != "" {
		ud.SetNickname(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetNickname()
		case SetData2Default:
			ud.UnsetNickname()
		}
	}
}

func UpdateWithNickname(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetNickname(p)
		} else {
			ud.UnsetNickname()
		}
	}
}

// @tpm-schematics:start-region("nickname-field-update-section")
// @tpm-schematics:end-region("nickname-field-update-section")

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

// SetFirstname No Remarks
func (ud *UpdateDocument) SetFirstname(p string) *UpdateDocument {
	mName := fmt.Sprintf(FirstnameFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetFirstname No Remarks
func (ud *UpdateDocument) UnsetFirstname() *UpdateDocument {
	mName := fmt.Sprintf(FirstnameFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetFirstname No Remarks
func (ud *UpdateDocument) setOrUnsetFirstname(p string, um UnsetMode) {
	if p != "" {
		ud.SetFirstname(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetFirstname()
		case SetData2Default:
			ud.UnsetFirstname()
		}
	}
}

func UpdateWithFirstname(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetFirstname(p)
		} else {
			ud.UnsetFirstname()
		}
	}
}

// @tpm-schematics:start-region("firstname-field-update-section")
// @tpm-schematics:end-region("firstname-field-update-section")

// SetLastname No Remarks
func (ud *UpdateDocument) SetLastname(p string) *UpdateDocument {
	mName := fmt.Sprintf(LastnameFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetLastname No Remarks
func (ud *UpdateDocument) UnsetLastname() *UpdateDocument {
	mName := fmt.Sprintf(LastnameFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetLastname No Remarks
func (ud *UpdateDocument) setOrUnsetLastname(p string, um UnsetMode) {
	if p != "" {
		ud.SetLastname(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetLastname()
		case SetData2Default:
			ud.UnsetLastname()
		}
	}
}

func UpdateWithLastname(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetLastname(p)
		} else {
			ud.UnsetLastname()
		}
	}
}

// @tpm-schematics:start-region("lastname-field-update-section")
// @tpm-schematics:end-region("lastname-field-update-section")

// SetEmail No Remarks
func (ud *UpdateDocument) SetEmail(p string) *UpdateDocument {
	mName := fmt.Sprintf(EmailFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetEmail No Remarks
func (ud *UpdateDocument) UnsetEmail() *UpdateDocument {
	mName := fmt.Sprintf(EmailFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetEmail No Remarks
func (ud *UpdateDocument) setOrUnsetEmail(p string, um UnsetMode) {
	if p != "" {
		ud.SetEmail(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetEmail()
		case SetData2Default:
			ud.UnsetEmail()
		}
	}
}

func UpdateWithEmail(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetEmail(p)
		} else {
			ud.UnsetEmail()
		}
	}
}

// @tpm-schematics:start-region("email-field-update-section")
// @tpm-schematics:end-region("email-field-update-section")

// SetPassword No Remarks
func (ud *UpdateDocument) SetPassword(p string) *UpdateDocument {
	mName := fmt.Sprintf(PasswordFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetPassword No Remarks
func (ud *UpdateDocument) UnsetPassword() *UpdateDocument {
	mName := fmt.Sprintf(PasswordFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetPassword No Remarks
func (ud *UpdateDocument) setOrUnsetPassword(p string, um UnsetMode) {
	if p != "" {
		ud.SetPassword(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetPassword()
		case SetData2Default:
			ud.UnsetPassword()
		}
	}
}

func UpdateWithPassword(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetPassword(p)
		} else {
			ud.UnsetPassword()
		}
	}
}

// @tpm-schematics:start-region("password-field-update-section")
// @tpm-schematics:end-region("password-field-update-section")

// SetRoles No Remarks
func (ud *UpdateDocument) SetRoles(p []commons.UserRole) *UpdateDocument {
	mName := fmt.Sprintf(RolesFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetRoles No Remarks
func (ud *UpdateDocument) UnsetRoles() *UpdateDocument {
	mName := fmt.Sprintf(RolesFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetRoles No Remarks - here2
func (ud *UpdateDocument) setOrUnsetRoles(p []commons.UserRole, um UnsetMode) {
	if len(p) > 0 {
		ud.SetRoles(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetRoles()
		case SetData2Default:
			ud.UnsetRoles()
		}
	}
}

func UpdateWithRoles(p []commons.UserRole) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetRoles(p)
		} else {
			ud.UnsetRoles()
		}
	}
}

// @tpm-schematics:start-region("roles-field-update-section")
// @tpm-schematics:end-region("roles-field-update-section")

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

// SetProfile_picture No Remarks
func (ud *UpdateDocument) SetProfile_picture(p *commons.FileReference) *UpdateDocument {
	mName := fmt.Sprintf(ProfilePictureFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetProfile_picture No Remarks
func (ud *UpdateDocument) UnsetProfile_picture() *UpdateDocument {
	mName := fmt.Sprintf(ProfilePictureFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetProfile_picture No Remarks - here2
func (ud *UpdateDocument) setOrUnsetProfile_picture(p *commons.FileReference, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetProfile_picture(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetProfile_picture()
		case SetData2Default:
			ud.UnsetProfile_picture()
		}
	}
}

func UpdateWithProfile_picture(p *commons.FileReference) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetProfile_picture(p)
		} else {
			ud.UnsetProfile_picture()
		}
	}
}

// @tpm-schematics:start-region("profile-picture-field-update-section")
// @tpm-schematics:end-region("profile-picture-field-update-section")

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
