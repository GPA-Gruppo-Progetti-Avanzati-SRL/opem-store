package commons

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type App struct {
	Id           string `json:"id,omitempty" bson:"id,omitempty" yaml:"id,omitempty"`
	ObjType      string `json:"objType,omitempty" bson:"objType,omitempty" yaml:"objType,omitempty"`
	Name         string `json:"name,omitempty" bson:"name,omitempty" yaml:"name,omitempty"`
	Description  string `json:"description,omitempty" bson:"description,omitempty" yaml:"description,omitempty"`
	Path         string `json:"path,omitempty" bson:"path,omitempty" yaml:"path,omitempty"`
	RoleRequired bool   `json:"roleRequired,omitempty" bson:"roleRequired,omitempty" yaml:"roleRequired,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s App) IsZero() bool {
	return s.Id == "" && s.ObjType == "" && s.Name == "" && s.Description == "" && s.Path == "" && !s.RoleRequired
}

// @tpm-schematics:start-region("bottom-file-section")

type AppObjType string

const (
	AppObjTypeWWW     AppObjType = "app-www"
	AppObjTypeConsole AppObjType = "app-admin"
)

// AppId ids of the available apps. The actual app is configured in the database.
type AppId string

const (
	AppIdHome AppId = "app-home"
	AppIdSys  AppId = "app-sys"
)

var AppIdsWorld = []AppId{AppIdHome, AppIdSys}

func IsAppIdInCatalog(n string) bool {
	for _, a := range AppIdsWorld {
		if string(a) == n {
			return true
		}
	}

	return false
}

// @tpm-schematics:end-region("bottom-file-section")
