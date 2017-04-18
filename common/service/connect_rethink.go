// author by @xiaoyusilen

package service

import (
	r "github.com/GoRethink/gorethink"
	log "github.com/Sirupsen/logrus"
)

func NewRethinkdb(url string) *r.Session {
	var err error

	session, err := r.Connect(r.ConnectOpts{
		Address:    url,
		InitialCap: 10,
		MaxOpen:    10,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Debug("rethinkdb connect success!")
	return session
}
