package dao

import (
	"github.com/bernardigiri/restfulUserAuth/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func UserInsert(user *model.User, db *mgo.Database) (err error) {
	return db.C("users").Insert(user)
}
func UserByLogin(login string, db *mgo.Database) (user *model.User, err error) {
	err = db.C("users").Find(bson.M{"login": login}).One(user)
	return
}
func UserUpdate(user *model.User, db *mgo.Database) (err error) {
	err = db.C("users").Update(bson.M{"login": user.Login}, user)
	return
}
func UserDelete(login string, db *mgo.Database) (err error) {
	err = db.C("users").Remove(bson.M{"login": login})
	return
}
