// author by @xiaoyusilen

package repository

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name string
}

func Register(s *mgo.Session, user *User, db string) (bool, error) {

	session := s.Copy()

	defer session.Close()

	c := session.DB(db).C("user")

	err := c.Insert(bson.M{
		"name": user.Name,
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
