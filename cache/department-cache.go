package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/haroldpacha/customer-management/entity"
)

type DepartmentCache interface {
	Set(key string, value []entity.Department)
	Get(key string) []entity.Department
	Del(key string)
}

type departmentCache struct {
	cache   *redis.Client
	expires time.Duration
}

func NewDepartmentCache(cache *redis.Client, expires time.Duration) DepartmentCache {
	return &departmentCache{
		cache:   cache,
		expires: expires,
	}
}

func (rc *departmentCache) Set(key string, value []entity.Department) {
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	rc.cache.Set(context.TODO(), "department:"+key, json, rc.expires*time.Second)
}

func (rc *departmentCache) Get(key string) []entity.Department {
	val, errGet := rc.cache.Get(context.TODO(), "department:"+key).Result()

	if errGet == redis.Nil {
		return nil
	} else if errGet != nil {
		panic(errGet)
	}

	var department []entity.Department
	err := json.Unmarshal([]byte(val), &department)
	if err != nil {
		panic(err)
	}

	return department
}

func (rc *departmentCache) Del(key string) {
	err := rc.cache.Del(context.TODO(), "department:"+key)

	log.Println(err)
}
