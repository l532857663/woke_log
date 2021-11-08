package cache

import (
	"backend-gateway/library/cache"
	"backend-gateway/library/config"
	"errors"
	"fmt"
	"time"

	logger "github.com/cihub/seelog"
	"github.com/go-redis/redis/v7"
)

const (
	CACHE_ENGINE_REDIS      = "redis"
	CACHE_ENGINE_MEMCACHE   = "memcache"
	CACHE_ENGINE_GROUPCACHE = "groupcache"

	NO_EXPIRATION_TIME time.Duration = time.Duration(0) // 永不过期
)

// @Description DAO CACHE管理资源结构
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
type Cache struct {
	// redis
	Redis *redis.Client
}

// @Description 创建一个 DAO 并返回对象
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func New(conf *config.CacheConfig) *Cache {
	c := &Cache{}

	// 创建缓存连接对象
	switch conf.Engine {
	case CACHE_ENGINE_REDIS:
		if conf.Redis.Enable {
			c.Redis = cache.NewRedis(conf.Redis)
		}
		// TODO: 未来支持更多缓存引擎
		// case CACHE_ENGINE_MEMCACHE:
		// case CACHE_ENGINE_GROUPCACHE:
		// default:
	}

	return c
}

// @Description 关闭DAO CACHE创建的资源
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) Close() {
	if c.Redis != nil {
		c.Redis.Close()
	}
}

// http://godoc.org/github.com/go-redis/redis

// @Description 处理IntResult结果
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) getIntResult(intCmd *redis.IntCmd) error {
	if intCmd.Err() != nil {
		return errors.New(fmt.Sprint("Failed to save data into redis: ", intCmd.Err()))
	}
	return nil
}

// @Description 处理StatusResult结果
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) getStatusResult(statusCmd *redis.StatusCmd) error {
	if statusCmd.Err() != nil {
		return errors.New(fmt.Sprint("Failed to save data into redis: ", statusCmd.Err()))
	}
	return nil
}

// @Description 处理BoolResult结果
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) getBoolResult(boolCmd *redis.BoolCmd) error {
	if boolCmd.Err() != nil {
		return errors.New(fmt.Sprint("Failed to save data into redis: ", boolCmd.Err()))
	}
	return nil
}

// @Description 处理FloatResult结果
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) getFloatResult(floatCmd *redis.FloatCmd) error {
	if floatCmd.Err() != nil {
		return errors.New(fmt.Sprint("Failed to save data into redis: ", floatCmd.Err()))
	}
	return nil
}

// @Description 创建key-value缓存并带过期时间
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) SetString(key, value string, expiration time.Duration) (string, error) {
	logger.Debugf("Redis set string: %s - %s", key, value)

	return c.Redis.Set(key, value, expiration).Result()
}

// @Description 自减计数
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) DECR(key string) (int64, error) {
	logger.Debugf("Redis decr value: %s", key)

	return c.Redis.Decr(key).Result()
}

// @Description 获取key对应的value缓存
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) GetString(key string) (string, error) {
	logger.Debugf("Redis get string: %s", key)

	return c.Redis.Get(key).Result()
}

// @Description 设置集合缓存
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) SetAdd(key, member string) error {
	logger.Debugf("Redis set add: %s - %s", key, member)

	return c.getIntResult(c.Redis.SAdd(key, member))
}

// @Description 移除列表指定元素
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) ListPush(key, value string, removeDup bool) error {
	logger.Debugf("Redis list push: %s - %s", key, value)

	if removeDup {
		// NOTE: remove duplicate value
		c.Redis.LRem(key, 0, value)
	}

	return c.getIntResult(c.Redis.LPush(key, value))
}

// @Description 移除列表指定元素
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) ListRem(key, value string) error {
	logger.Debugf("Redis list rem: %s - %s", key, value)

	return c.getIntResult(c.Redis.LRem(key, 0, value))
}

// @Description 获取并移除列表第一个元素
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) ListPop(key string) (value string, err error) {
	logger.Debugf("Redis list pop: %s", key)

	return c.Redis.RPop(key).Result()
}

// @Description 获取所有符合给定模式的key
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) ListKeys(keyPattern string) (value []string, err error) {
	logger.Debugf("Redis list keys: %s", keyPattern)

	return c.Redis.Keys(keyPattern).Result()
}

// @Description 获取列表长度
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) ListSize(key string) (size int64, err error) {
	logger.Debugf("Redis list size: %s", key)

	return c.Redis.LLen(key).Result()
}

// @Description 移除列表的最后一个元素，并将该元素添加到另一个列表并返回
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) ListRPopLPush(srcKey, targetKey string) (result string, err error) {
	logger.Debugf("Redis list right pop left push: %s - %s", srcKey, targetKey)

	return c.Redis.RPopLPush(srcKey, targetKey).Result()
}

// @Description 在列表中添加一个值
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) ListRPush(srcKey, targetKey string) (result int64, err error) {
	logger.Debugf("Redis list right push: %s - %s", srcKey, targetKey)

	return c.Redis.RPush(srcKey, targetKey).Result()
}

// @Description 移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) ListBLPop(srcKey string, timeout time.Duration) (result []string, err error) {
	logger.Debugf("Redis list b left pop: %s - timeout: %v", srcKey, timeout)

	return c.Redis.BLPop(timeout, srcKey).Result()
}

// @Description 移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) ListBRPop(srcKey string, timeout time.Duration) (result []string, err error) {
	logger.Debugf("Redis list b right pop: %s - timeout: %v", srcKey, timeout)

	return c.Redis.BRPop(timeout, srcKey).Result()
}

// @Description 获取列表指定范围内的元素
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) ListRange(key string, from, to int64) (result []string, err error) {
	logger.Debugf("Redis list range: %s - %d:%d", key, from, to)

	return c.Redis.LRange(key, from, to).Result()
}

// @Description 将哈希表 key 中的字段 field 的值设为 value
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) HSet(key, field string, data string) error {
	logger.Debugf("Redis hash set: %s - %s:%s", key, field, data)

	return c.getIntResult(c.Redis.HSet(key, field, data))
}

// @Description 获取存储在哈希表中指定字段的值。
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) HGet(key, field string) (result string, err error) {
	logger.Debugf("Redis hash get: %s - %s", key, field)

	return c.Redis.HGet(key, field).Result()
}

// @Description 同时将多个 field-value (域-值)对设置到哈希表 key 中。
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) HMSet(key string, fields map[string]interface{}) error {
	logger.Debugf("Redis hash set: %s - %#v", key, fields)

	return c.getBoolResult(c.Redis.HMSet(key, fields))
}

// @Description 获取所有给定字段的值
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) HMGet(key string, fields ...string) ([]interface{}, error) {
	logger.Debugf("Redis hash get: %s - %s", key, fields)

	return c.Redis.HMGet(key, fields...).Result()
}

// @Description 删除一个或多个哈希表字段
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) HDel(key string, fields ...string) error {
	logger.Debugf("Redis hash del: %s - %s", key, fields)

	return c.getIntResult(c.Redis.HDel(key, fields...))
}

// @Description 给指定Key设置过期时间
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) ExpireKey(key string, expiration time.Duration) (bool, error) {
	logger.Debugf("Redis expire key: %s", key)

	return c.Redis.Expire(key, expiration).Result()
}

// @Description 删除指定Key
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) DelKey(key string) (int64, error) {
	logger.Debugf("Redis delete key: %s", key)

	return c.Redis.Del(key).Result()
}

// @Description 向集合添加一个或多个成员
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) SAdd(key string, data []string) error {
	logger.Debugf("Redis set members: %s - %v", key, data)

	return c.getIntResult(c.Redis.SAdd(key, data))
}

// @Description 返回集合中的所有成员
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) SMembers(key string) ([]string, error) {
	logger.Debugf("Redis get members: %s", key)

	return c.Redis.SMembers(key).Result()
}

// @Description 返回有序集合中的所有成员
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) ZRangeAll(key string) ([]string, error) {
	logger.Debugf("Redis get sorted set all members: %s", key)

	return c.Redis.ZRange(key, 0, -1).Result()
}

// @Description 返回有序集合中的指定位置内的成员
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) ZRange(key string, start, end int64) ([]string, error) {
	logger.Debugf("Redis get sorted set index members: %s - %v - %v", key, start, end)

	return c.Redis.ZRange(key, start, end).Result()
}

// @Description 返回有序集合中的所有成员的数量
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) ZCountAll(key string) (int64, error) {
	logger.Debugf("Redis get sorted set all members count: %s", key)

	return c.Redis.ZCount(key, "-inf", "+inf").Result()
}

// @Description 向有序集合添加一个或多个成员
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (c *Cache) ZAdd(key string, score int, member []byte) error {
	logger.Debugf("Redis sorted set add member: %s - %v - %v", key, score, member)

	return c.getIntResult(c.Redis.ZAdd(key, &redis.Z{
		Score:  float64(score),
		Member: member,
	}))
}
