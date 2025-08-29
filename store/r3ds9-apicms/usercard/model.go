package usercard

import (
	commons "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	OID     = "_id"
	USERID  = "userId"
	OBJTYPE = "objType"
	SYSINFO = "sysinfo"
)

type UserCard struct {
	OId     primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	UserId  string             `json:"userId,omitempty" bson:"userId,omitempty"`
	ObjType string             `json:"objType,omitempty" bson:"objType,omitempty"`
	Sysinfo commons.SysInfo    `json:"sysinfo,omitempty" bson:"sysinfo,omitempty"`
}

func (s UserCard) IsZero() bool {
	/*
	   if s.OId == primitive.NilObjectID {
	       return false
	   }
	   if s.UserId == "" {
	       return false
	   }
	   if s.ObjType == "" {
	       return false
	   }
	   if s.Sysinfo.IsZero() {
	       return false
	   }
	       return true
	*/

	return s.OId == primitive.NilObjectID && s.UserId == "" && s.ObjType == "" && s.Sysinfo.IsZero()
}
