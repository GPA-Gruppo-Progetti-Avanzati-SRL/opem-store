package kv

import "go.mongodb.org/mongo-driver/v2/bson"

type UpdateOperator string

const (
	Set         UpdateOperator = "$set"
	Unset       UpdateOperator = "$unset"
	Inc         UpdateOperator = "$inc"
	CurrentDate UpdateOperator = "$currentDate"
	AddToSet    UpdateOperator = "$addToSet"
	Pull        UpdateOperator = "$pull"
)

type UpdateEntry func() bson.E

type Updates struct {
	operator UpdateOperator
	list     []UpdateEntry
}

func (u *Updates) Add(e UpdateEntry) {
	u.list = append(u.list, e)
}

func (u *Updates) Build() bson.E {
	doc := bson.D{}
	for _, e := range u.list {
		doc = append(doc, e())
	}
	return bson.E{Key: string(u.operator), Value: doc}
}

type UpdateDocument struct {
	set         *Updates
	unset       *Updates
	inc         *Updates
	currentDate *Updates
	addToSet    *Updates
	pull        *Updates
}

func (ud *UpdateDocument) Set() *Updates {
	if ud.set == nil {
		ud.set = &Updates{operator: Set}
	}
	return ud.set
}

func (ud *UpdateDocument) Unset() *Updates {
	if ud.unset == nil {
		ud.unset = &Updates{operator: Unset}
	}
	return ud.unset
}

func (ud *UpdateDocument) Inc() *Updates {
	if ud.inc == nil {
		ud.inc = &Updates{operator: Inc}
	}
	return ud.inc
}

func (ud *UpdateDocument) CurrentDate() *Updates {
	if ud.currentDate == nil {
		ud.currentDate = &Updates{operator: CurrentDate}
	}
	return ud.currentDate
}

func (ud *UpdateDocument) AddToSet() *Updates {
	if ud.addToSet == nil {
		ud.addToSet = &Updates{operator: AddToSet}
	}
	return ud.addToSet
}

func (ud *UpdateDocument) Pull() *Updates {
	if ud.pull == nil {
		ud.pull = &Updates{operator: Pull}
	}
	return ud.pull
}

func (ud *UpdateDocument) Build() bson.D {
	doc := bson.D{}
	if ud.set != nil && len(ud.set.list) > 0 {
		doc = append(doc, ud.set.Build())
	}
	if ud.unset != nil && len(ud.unset.list) > 0 {
		doc = append(doc, ud.unset.Build())
	}
	if ud.inc != nil && len(ud.inc.list) > 0 {
		doc = append(doc, ud.inc.Build())
	}
	if ud.currentDate != nil && len(ud.currentDate.list) > 0 {
		doc = append(doc, ud.currentDate.Build())
	}
	if ud.addToSet != nil && len(ud.addToSet.list) > 0 {
		doc = append(doc, ud.addToSet.Build())
	}
	if ud.pull != nil && len(ud.pull.list) > 0 {
		doc = append(doc, ud.pull.Build())
	}
	return doc
}
