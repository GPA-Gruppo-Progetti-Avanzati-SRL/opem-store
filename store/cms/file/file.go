package file

import "go.mongodb.org/mongo-driver/v2/bson"
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type File struct {
	OId      bson.ObjectID         `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Fn       string                `json:"fn,omitempty" bson:"fn,omitempty" yaml:"fn,omitempty"`
	Descr    string                `json:"descr,omitempty" bson:"descr,omitempty" yaml:"descr,omitempty"`
	Role     string                `json:"role,omitempty" bson:"role,omitempty" yaml:"role,omitempty"`
	EntRefs  []EntRefStruct        `json:"ent_refs,omitempty" bson:"ent_refs,omitempty" yaml:"ent_refs,omitempty"`
	Metadata bson.M                `json:"metadata,omitempty" bson:"metadata,omitempty" yaml:"metadata,omitempty"`
	Vrnts    []commons.FileVariant `json:"vrnts,omitempty" bson:"vrnts,omitempty" yaml:"vrnts,omitempty"`
	SysInfo  commons.SysInfo       `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s File) IsZero() bool {
	return s.OId == bson.NilObjectID && s.Fn == "" && s.Descr == "" && s.Role == "" && len(s.EntRefs) == 0 && len(s.Metadata) == 0 && len(s.Vrnts) == 0 && s.SysInfo.IsZero()
}

type QueryResult struct {
	Records int    `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []File `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

// @tpm-schematics:start-region("bottom-file-section")

func (s File) FileReference() commons.FileReference {

	fr := commons.FileReference{OId: s.OId}
	for _, v := range s.Vrnts {
		fr.SrcSet = append(fr.SrcSet, commons.FileVariant{
			Ct:   v.Ct,
			Wd:   v.Wd,
			Ht:   v.Ht,
			Bln:  v.Bln,
			Url:  v.Url,
			Role: v.Role,
		})
	}

	return fr
}

// @tpm-schematics:end-region("bottom-file-section")
