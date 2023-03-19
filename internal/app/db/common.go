package db

import (
	"fengjx/go-web-quickstart/internal/app/applog"
	"fengjx/go-web-quickstart/internal/app/redis"
	"fmt"
	"reflect"
	"time"
)

func Add(bean interface{}) (bool, error) {
	affected, err := Default().Omit("ctime", "utime").Insert(bean)
	if err != nil {
		applog.Log.Errorf("[db] insert err - %s", err.Error())
		return false, err
	}
	return affected > 0, err
}

func DelById(id int64, bean interface{}) (bool, error) {
	model, ok := bean.(Model)
	if !ok {
		return false, ErrNotImplModel
	}
	cacheKey := buildModelCacheKey(id, model)
	affected, err := Default().ID(id).Delete(bean)
	if err != nil {
		applog.Log.Errorf("[db] delete by id err - %s", err.Error())
		return false, nil
	}
	ok = affected > 0
	if !ok {
		redis.Default().Del(redis.Ctx, cacheKey)
	}
	return ok, err
}

func GetById(id int64, bean interface{}) error {
	model, ok := bean.(Model)
	if !ok {
		return ErrNotImplModel
	}
	cacheKey := buildModelCacheKey(id, model)
	has, _ := redis.GetObj(cacheKey, bean)
	if has {
		return nil
	}
	ok, err := Default().ID(id).Get(bean)
	if err != nil {
		applog.Log.Errorf("get by id err - %s", err.Error())
		return err
	}
	if !ok {
		return nil
	}
	err = redis.SetObj(cacheKey, bean, 2*time.Minute)
	if err != nil {
		applog.Log.Errorf("set cache err - %s", err)
	}
	return err
}

func DelCache(id int64, bean interface{}) (bool, error) {
	model, ok := bean.(Model)
	if !ok {
		return false, ErrNotImplModel
	}
	cacheKey := buildModelCacheKey(id, model)
	ret, err := redis.Default().Del(redis.Ctx, cacheKey).Result()
	if err != nil {
		applog.Log.Errorf("del cache err - %s", err.Error())
		return false, err
	}
	return ret > 0, nil
}

func UpdateById(id int64, bean interface{}) (bool, error) {
	model, ok := bean.(Model)
	if !ok {
		return false, ErrNotImplModel
	}
	cacheKey := buildModelCacheKey(id, model)
	ret, err := redis.Default().Del(redis.Ctx, cacheKey).Result()
	if err != nil {
		applog.Log.Errorf("del cache err - %s", err.Error())
	}
	applog.Log.Debugf("clean model cache result: %t", ret)
	affected, err := Default().ID(id).Update(bean)
	if err != nil {
		applog.Log.Errorf("update by id err - $s", err.Error())
		return false, err
	}
	// delay double del
	go delayDel(cacheKey)
	return affected > 0, err
}

// 延迟删除
func delayDel(key string) {
	time.Sleep(500)
	redis.Default().Del(redis.Ctx, key)
}

func buildModelCacheKey(id int64, m Model) string {
	return fmt.Sprintf("model:%s:%s:%d", reflect.TypeOf(m).String(), m.Version(), id)
}
