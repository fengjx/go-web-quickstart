package user

import (
	"fengjx/go-web-quickstart/internal/app/applog"
	"fengjx/go-web-quickstart/internal/app/db"
)

func GetByIdCache(uid int64) (*User, error) {
	user := New()
	err := db.GetByIdCache(uid, user, Version, func() (interface{}, error) {
		has, err := db.Default().ID(uid).Get(user)
		if err != nil {
			applog.Log.Errorf("get user from db err: %s", err.Error())
			return nil, err
		}
		if has {
			applog.Log.Warnf("get user from db empty: %d, %s", uid, err.Error())
			return nil, nil
		}
		return user, err
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}
