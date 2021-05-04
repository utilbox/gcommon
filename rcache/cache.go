// Package rcache provide a redis cache service.
// It wraps go-redis as underlying client.
// For more details about go-redis, please refer to
// https://github.com/go-redis/redis
package rcache

import (
	"github.com/go-redis/redis/v7"
	"time"
)

var client *redis.Client

// Init can initialize the redis cache.
func Init(options *redis.Options) {
	client = redis.NewClient(options)
}

func Pipeline() redis.Pipeliner {
	return client.Pipeline()
}

func Pipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return client.Pipelined(fn)
}

func TxPipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return client.TxPipelined(fn)
}

func TxPipeline() redis.Pipeliner {
	return client.TxPipeline()
}

func Command() *redis.CommandsInfoCmd {
	return client.Command()
}

func ClientGetName() *redis.StringCmd {
	return client.ClientGetName()
}

func Echo(message interface{}) *redis.StringCmd {
	return client.Echo(message)
}

func Ping() *redis.StatusCmd {
	return client.Ping()
}

func Quit() *redis.StatusCmd {
	return client.Quit()
}

func Del(keys ...string) *redis.IntCmd {
	return client.Del(keys...)
}

func Unlink(keys ...string) *redis.IntCmd {
	return client.Unlink(keys...)
}

func Dump(key string) *redis.StringCmd {
	return client.Dump(key)
}

func Exists(keys ...string) *redis.IntCmd {
	return client.Exists(keys...)
}

func Expire(key string, expiration time.Duration) *redis.BoolCmd {
	return client.Expire(key, expiration)
}

func ExpireAt(key string, tm time.Time) *redis.BoolCmd {
	return client.ExpireAt(key, tm)
}

func Keys(pattern string) *redis.StringSliceCmd {
	return client.Keys(pattern)
}

func Migrate(host, port, key string, db int, timeout time.Duration) *redis.StatusCmd {
	return client.Migrate(host, port, key, db, timeout)
}

func Move(key string, db int) *redis.BoolCmd {
	return client.Move(key, db)
}

func ObjectRefCount(key string) *redis.IntCmd {
	return client.ObjectRefCount(key)
}

func ObjectEncoding(key string) *redis.StringCmd {
	return client.ObjectEncoding(key)
}

func ObjectIdleTime(key string) *redis.DurationCmd {
	return client.ObjectIdleTime(key)
}

func Persist(key string) *redis.BoolCmd {
	return client.Persist(key)
}

func PExpire(key string, expiration time.Duration) *redis.BoolCmd {
	return client.PExpire(key, expiration)
}

func PExpireAt(key string, tm time.Time) *redis.BoolCmd {
	return client.PExpireAt(key, tm)
}

func PTTL(key string) *redis.DurationCmd {
	return client.PTTL(key)
}

func RandomKey() *redis.StringCmd {
	return client.RandomKey()
}

func Rename(key, newkey string) *redis.StatusCmd {
	return client.Rename(key, newkey)
}

func RenameNX(key, newkey string) *redis.BoolCmd {
	return client.RenameNX(key, newkey)
}

func Restore(key string, ttl time.Duration, value string) *redis.StatusCmd {
	return client.Restore(key, ttl, value)
}

func RestoreReplace(key string, ttl time.Duration, value string) *redis.StatusCmd {
	return client.RestoreReplace(key, ttl, value)
}

func Sort(key string, sort *redis.Sort) *redis.StringSliceCmd {
	return client.Sort(key, sort)
}

func SortStore(key, store string, sort *redis.Sort) *redis.IntCmd {
	return client.SortStore(key, store, sort)
}

func SortInterfaces(key string, sort *redis.Sort) *redis.SliceCmd {
	return client.SortInterfaces(key, sort)
}

func Touch(keys ...string) *redis.IntCmd {
	return client.Touch(keys...)
}

func TTL(key string) *redis.DurationCmd {
	return client.TTL(key)
}

func Type(key string) *redis.StatusCmd {
	return client.Type(key)
}

func Scan(cursor uint64, match string, count int64) *redis.ScanCmd {
	return client.Scan(cursor, match, count)
}

func SScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return client.SScan(key, cursor, match, count)
}

func HScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return client.HScan(key, cursor, match, count)
}

func ZScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return client.ZScan(key, cursor, match, count)
}

func Append(key, value string) *redis.IntCmd {
	return client.Append(key, value)
}

func BitCount(key string, bitCount *redis.BitCount) *redis.IntCmd {
	return client.BitCount(key, bitCount)
}

func BitOpAnd(destKey string, keys ...string) *redis.IntCmd {
	return client.BitOpAnd(destKey, keys...)
}

func BitOpOr(destKey string, keys ...string) *redis.IntCmd {
	return client.BitOpOr(destKey, keys...)
}

func BitOpXor(destKey string, keys ...string) *redis.IntCmd {
	return client.BitOpXor(destKey, keys...)
}

func BitOpNot(destKey string, key string) *redis.IntCmd {
	return client.BitOpNot(destKey, key)
}

func BitPos(key string, bit int64, pos ...int64) *redis.IntCmd {
	return client.BitPos(key, bit, pos...)
}

func Decr(key string) *redis.IntCmd {
	return client.Decr(key)
}

func DecrBy(key string, decrement int64) *redis.IntCmd {
	return client.DecrBy(key, decrement)
}

func Get(key string) *redis.StringCmd {
	return client.Get(key)
}

func GetBit(key string, offset int64) *redis.IntCmd {
	return client.GetBit(key, offset)
}

func GetRange(key string, start, end int64) *redis.StringCmd {
	return client.GetRange(key, start, end)
}

func GetSet(key string, value interface{}) *redis.StringCmd {
	return client.GetSet(key, value)
}

func Incr(key string) *redis.IntCmd {
	return client.Incr(key)
}

func IncrBy(key string, value int64) *redis.IntCmd {
	return client.IncrBy(key, value)
}

func IncrByFloat(key string, value float64) *redis.FloatCmd {
	return client.IncrByFloat(key, value)
}

func MGet(keys ...string) *redis.SliceCmd {
	return client.MGet(keys...)
}

func MSet(pairs ...interface{}) *redis.StatusCmd {
	return client.MSet(pairs...)
}

func MSetNX(pairs ...interface{}) *redis.BoolCmd {
	return client.MSetNX(pairs...)
}

func Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return client.Set(key, value, expiration)
}

func SetBit(key string, offset int64, value int) *redis.IntCmd {
	return client.SetBit(key, offset, value)
}

func SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	return client.SetNX(key, value, expiration)
}

func SetXX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	return client.SetXX(key, value, expiration)
}

func SetRange(key string, offset int64, value string) *redis.IntCmd {
	return client.SetRange(key, offset, value)
}

func StrLen(key string) *redis.IntCmd {
	return client.StrLen(key)
}

func HDel(key string, fields ...string) *redis.IntCmd {
	return client.HDel(key, fields...)
}

func HExists(key, field string) *redis.BoolCmd {
	return client.HExists(key, field)
}

func HGet(key, field string) *redis.StringCmd {
	return client.HGet(key, field)
}

func HGetAll(key string) *redis.StringStringMapCmd {
	return client.HGetAll(key)
}

func HIncrBy(key, field string, incr int64) *redis.IntCmd {
	return client.HIncrBy(key, field, incr)
}

func HIncrByFloat(key, field string, incr float64) *redis.FloatCmd {
	return client.HIncrByFloat(key, field, incr)
}

func HKeys(key string) *redis.StringSliceCmd {
	return client.HKeys(key)
}

func HLen(key string) *redis.IntCmd {
	return client.HLen(key)
}

func HMGet(key string, fields ...string) *redis.SliceCmd {
	return client.HMGet(key, fields...)
}

func HMSet(key string, values ...interface{}) *redis.BoolCmd {
	return client.HMSet(key, values)
}

func HSet(key string, values ...interface{}) *redis.IntCmd {
	return client.HSet(key, values...)
}

func HSetNX(key, field string, value interface{}) *redis.BoolCmd {
	return client.HSetNX(key, field, value)
}

func HVals(key string) *redis.StringSliceCmd {
	return client.HVals(key)
}

func BLPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	return client.BLPop(timeout, keys...)
}

func BRPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	return client.BRPop(timeout, keys...)
}

func BRPopLPush(source, destination string, timeout time.Duration) *redis.StringCmd {
	return client.BRPopLPush(source, destination, timeout)
}

func LIndex(key string, index int64) *redis.StringCmd {
	return client.LIndex(key, index)
}

func LInsert(key, op string, pivot, value interface{}) *redis.IntCmd {
	return client.LInsert(key, op, pivot, value)
}

func LInsertBefore(key string, pivot, value interface{}) *redis.IntCmd {
	return client.LInsertBefore(key, pivot, value)
}

func LInsertAfter(key string, pivot, value interface{}) *redis.IntCmd {
	return client.LInsertAfter(key, pivot, value)
}

func LLen(key string) *redis.IntCmd {
	return client.LLen(key)
}

func LPop(key string) *redis.StringCmd {
	return client.LPop(key)
}

func LPush(key string, values ...interface{}) *redis.IntCmd {
	return client.LPush(key, values...)
}

func LPushX(key string, value interface{}) *redis.IntCmd {
	return client.LPushX(key, value)
}

func LRange(key string, start, stop int64) *redis.StringSliceCmd {
	return client.LRange(key, start, stop)
}

func LRem(key string, count int64, value interface{}) *redis.IntCmd {
	return client.LRem(key, count, value)
}

func LSet(key string, index int64, value interface{}) *redis.StatusCmd {
	return client.LSet(key, index, value)
}

func LTrim(key string, start, stop int64) *redis.StatusCmd {
	return client.LTrim(key, start, stop)
}

func RPop(key string) *redis.StringCmd {
	return client.RPop(key)
}

func RPopLPush(source, destination string) *redis.StringCmd {
	return client.RPopLPush(source, destination)
}

func RPush(key string, values ...interface{}) *redis.IntCmd {
	return client.RPush(key, values...)
}

func RPushX(key string, value interface{}) *redis.IntCmd {
	return client.RPushX(key, value)
}

func SAdd(key string, members ...interface{}) *redis.IntCmd {
	return client.SAdd(key, members...)
}

func SCard(key string) *redis.IntCmd {
	return client.SCard(key)
}

func SDiff(keys ...string) *redis.StringSliceCmd {
	return client.SDiff(keys...)
}

func SDiffStore(destination string, keys ...string) *redis.IntCmd {
	return client.SDiffStore(destination, keys...)
}

func SInter(keys ...string) *redis.StringSliceCmd {
	return client.SInter(keys...)
}

func SInterStore(destination string, keys ...string) *redis.IntCmd {
	return client.SInterStore(destination, keys...)
}

func SIsMember(key string, member interface{}) *redis.BoolCmd {
	return client.SIsMember(key, member)
}

func SMembers(key string) *redis.StringSliceCmd {
	return client.SMembers(key)
}

func SMembersMap(key string) *redis.StringStructMapCmd {
	return client.SMembersMap(key)
}

func SMove(source, destination string, member interface{}) *redis.BoolCmd {
	return client.SMove(source, destination, member)
}

func SPop(key string) *redis.StringCmd {
	return client.SPop(key)
}

func SPopN(key string, count int64) *redis.StringSliceCmd {
	return client.SPopN(key, count)
}

func SRandMember(key string) *redis.StringCmd {
	return client.SRandMember(key)
}

func SRandMemberN(key string, count int64) *redis.StringSliceCmd {
	return client.SRandMemberN(key, count)
}

func SRem(key string, members ...interface{}) *redis.IntCmd {
	return client.SRem(key, members...)
}

func SUnion(keys ...string) *redis.StringSliceCmd {
	return client.SUnion(keys...)
}

func SUnionStore(destination string, keys ...string) *redis.IntCmd {
	return client.SUnionStore(destination, keys...)
}

func XAdd(a *redis.XAddArgs) *redis.StringCmd {
	return client.XAdd(a)
}

func XDel(stream string, ids ...string) *redis.IntCmd {
	return client.XDel(stream, ids...)
}

func XLen(stream string) *redis.IntCmd {
	return client.XLen(stream)
}

func XRange(stream, start, stop string) *redis.XMessageSliceCmd {
	return client.XRange(stream, start, stop)
}

func XRangeN(stream, start, stop string, count int64) *redis.XMessageSliceCmd {
	return client.XRangeN(stream, start, stop, count)
}

func XRevRange(stream string, start, stop string) *redis.XMessageSliceCmd {
	return client.XRevRange(stream, start, stop)
}

func XRevRangeN(stream string, start, stop string, count int64) *redis.XMessageSliceCmd {
	return client.XRevRangeN(stream, start, stop, count)
}

func XRead(a *redis.XReadArgs) *redis.XStreamSliceCmd {
	return client.XRead(a)
}

func XReadStreams(streams ...string) *redis.XStreamSliceCmd {
	return client.XReadStreams(streams...)
}

func XGroupCreate(stream, group, start string) *redis.StatusCmd {
	return client.XGroupCreate(stream, group, start)
}

func XGroupCreateMkStream(stream, group, start string) *redis.StatusCmd {
	return client.XGroupCreateMkStream(stream, group, start)
}

func XGroupSetID(stream, group, start string) *redis.StatusCmd {
	return client.XGroupSetID(stream, group, start)
}

func XGroupDestroy(stream, group string) *redis.IntCmd {
	return client.XGroupDestroy(stream, group)
}

func XGroupDelConsumer(stream, group, consumer string) *redis.IntCmd {
	return client.XGroupDelConsumer(stream, group, consumer)
}

func XReadGroup(a *redis.XReadGroupArgs) *redis.XStreamSliceCmd {
	return client.XReadGroup(a)
}

func XAck(stream, group string, ids ...string) *redis.IntCmd {
	return client.XAck(stream, group, ids...)
}

func XPending(stream, group string) *redis.XPendingCmd {
	return client.XPending(stream, group)
}

func XPendingExt(a *redis.XPendingExtArgs) *redis.XPendingExtCmd {
	return client.XPendingExt(a)
}

func XClaim(a *redis.XClaimArgs) *redis.XMessageSliceCmd {
	return client.XClaim(a)
}

func XClaimJustID(a *redis.XClaimArgs) *redis.StringSliceCmd {
	return client.XClaimJustID(a)
}

func XTrim(key string, maxLen int64) *redis.IntCmd {
	return client.XTrim(key, maxLen)
}

func XTrimApprox(key string, maxLen int64) *redis.IntCmd {
	return client.XTrimApprox(key, maxLen)
}

func BZPopMax(timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	return client.BZPopMax(timeout, keys...)
}

func BZPopMin(timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	return client.BZPopMin(timeout, keys...)
}

func ZAdd(key string, members ...*redis.Z) *redis.IntCmd {
	return client.ZAdd(key, members...)
}

func ZAddNX(key string, members ...*redis.Z) *redis.IntCmd {
	return client.ZAddNX(key, members...)
}

func ZAddXX(key string, members ...*redis.Z) *redis.IntCmd {
	return client.ZAddXX(key, members...)
}

func ZAddCh(key string, members ...*redis.Z) *redis.IntCmd {
	return client.ZAddCh(key, members...)
}

func ZAddNXCh(key string, members ...*redis.Z) *redis.IntCmd {
	return client.ZAddNXCh(key, members...)
}

func ZAddXXCh(key string, members ...*redis.Z) *redis.IntCmd {
	return client.ZAddXXCh(key, members...)
}

func ZIncr(key string, member *redis.Z) *redis.FloatCmd {
	return client.ZIncr(key, member)
}

func ZIncrNX(key string, member *redis.Z) *redis.FloatCmd {
	return client.ZIncrNX(key, member)
}

func ZIncrXX(key string, member *redis.Z) *redis.FloatCmd {
	return client.ZIncrXX(key, member)
}

func ZCard(key string) *redis.IntCmd {
	return client.ZCard(key)
}

func ZCount(key, min, max string) *redis.IntCmd {
	return client.ZCount(key, min, max)
}

func ZLexCount(key, min, max string) *redis.IntCmd {
	return client.ZLexCount(key, min, max)
}

func ZIncrBy(key string, increment float64, member string) *redis.FloatCmd {
	return client.ZIncrBy(key, increment, member)
}

func ZInterStore(destination string, store *redis.ZStore) *redis.IntCmd {
	return client.ZInterStore(destination, store)
}

func ZPopMax(key string, count ...int64) *redis.ZSliceCmd {
	return client.ZPopMax(key, count...)
}

func ZPopMin(key string, count ...int64) *redis.ZSliceCmd {
	return client.ZPopMin(key, count...)
}

func ZRange(key string, start, stop int64) *redis.StringSliceCmd {
	return client.ZRange(key, start, stop)
}

func ZRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	return client.ZRangeWithScores(key, start, stop)
}

func ZRangeByScore(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	return client.ZRangeByScore(key, opt)
}

func ZRangeByLex(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	return client.ZRangeByLex(key, opt)
}

func ZRangeByScoreWithScores(key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	return client.ZRangeByScoreWithScores(key, opt)
}

func ZRank(key, member string) *redis.IntCmd {
	return client.ZRank(key, member)
}

func ZRem(key string, members ...interface{}) *redis.IntCmd {
	return client.ZRem(key, members...)
}

func ZRemRangeByRank(key string, start, stop int64) *redis.IntCmd {
	return client.ZRemRangeByRank(key, start, stop)
}

func ZRemRangeByScore(key, min, max string) *redis.IntCmd {
	return client.ZRemRangeByScore(key, min, max)
}

func ZRemRangeByLex(key, min, max string) *redis.IntCmd {
	return client.ZRemRangeByLex(key, min, max)
}

func ZRevRange(key string, start, stop int64) *redis.StringSliceCmd {
	return client.ZRevRange(key, start, stop)
}

func ZRevRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	return client.ZRevRangeWithScores(key, start, stop)
}

func ZRevRangeByScore(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	return client.ZRevRangeByScore(key, opt)
}

func ZRevRangeByLex(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	return client.ZRevRangeByLex(key, opt)
}

func ZRevRangeByScoreWithScores(key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	return client.ZRevRangeByScoreWithScores(key, opt)
}

func ZRevRank(key, member string) *redis.IntCmd {
	return client.ZRevRank(key, member)
}

func ZScore(key, member string) *redis.FloatCmd {
	return client.ZScore(key, member)
}

func ZUnionStore(dest string, store *redis.ZStore) *redis.IntCmd {
	return client.ZUnionStore(dest, store)
}

func PFAdd(key string, els ...interface{}) *redis.IntCmd {
	return client.PFAdd(key, els...)
}

func PFCount(keys ...string) *redis.IntCmd {
	return client.PFCount(keys...)
}

func PFMerge(dest string, keys ...string) *redis.StatusCmd {
	return client.PFMerge(dest, keys...)
}

func BgRewriteAOF() *redis.StatusCmd {
	return client.BgRewriteAOF()
}

func BgSave() *redis.StatusCmd {
	return client.BgSave()
}

func ClientKill(ipPort string) *redis.StatusCmd {
	return client.ClientKill(ipPort)
}

func ClientKillByFilter(keys ...string) *redis.IntCmd {
	return client.ClientKillByFilter(keys...)
}

func ClientList() *redis.StringCmd {
	return client.ClientList()
}

func ClientPause(dur time.Duration) *redis.BoolCmd {
	return client.ClientPause(dur)
}

func ClientID() *redis.IntCmd {
	return client.ClientID()
}

func ConfigGet(parameter string) *redis.SliceCmd {
	return client.ConfigGet(parameter)
}

func ConfigResetStat() *redis.StatusCmd {
	return client.ConfigResetStat()
}

func ConfigSet(parameter, value string) *redis.StatusCmd {
	return client.ConfigSet(parameter, value)
}

func ConfigRewrite() *redis.StatusCmd {
	return client.ConfigRewrite()
}

func DBSize() *redis.IntCmd {
	return client.DBSize()
}

func FlushAll() *redis.StatusCmd {
	return client.FlushAll()
}

func FlushAllAsync() *redis.StatusCmd {
	return client.FlushAllAsync()
}

func FlushDB() *redis.StatusCmd {
	return client.FlushDB()
}

func FlushDBAsync() *redis.StatusCmd {
	return client.FlushDBAsync()
}

func Info(section ...string) *redis.StringCmd {
	return client.Info(section...)
}

func LastSave() *redis.IntCmd {
	return client.LastSave()
}

func Save() *redis.StatusCmd {
	return client.Save()
}

func Shutdown() *redis.StatusCmd {
	return client.Shutdown()
}

func ShutdownSave() *redis.StatusCmd {
	return client.ShutdownSave()
}

func ShutdownNoSave() *redis.StatusCmd {
	return client.ShutdownNoSave()
}

func SlaveOf(host, port string) *redis.StatusCmd {
	return client.SlaveOf(host, port)
}

func Time() *redis.TimeCmd {
	return client.Time()
}

func Eval(script string, keys []string, args ...interface{}) *redis.Cmd {
	return client.Eval(script, keys, args...)
}

func EvalSha(sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	return client.EvalSha(sha1, keys, args...)
}

func ScriptExists(hashes ...string) *redis.BoolSliceCmd {
	return client.ScriptExists(hashes...)
}

func ScriptFlush() *redis.StatusCmd {
	return client.ScriptFlush()
}

func ScriptKill() *redis.StatusCmd {
	return client.ScriptKill()
}

func ScriptLoad(script string) *redis.StringCmd {
	return client.ScriptLoad(script)
}

func DebugObject(key string) *redis.StringCmd {
	return client.DebugObject(key)
}

func Publish(channel string, message interface{}) *redis.IntCmd {
	return client.Publish(channel, message)
}

func PubSubChannels(pattern string) *redis.StringSliceCmd {
	return client.PubSubChannels(pattern)
}

func PubSubNumSub(channels ...string) *redis.StringIntMapCmd {
	return client.PubSubNumSub(channels...)
}

func PubSubNumPat() *redis.IntCmd {
	return client.PubSubNumPat()
}

func ClusterSlots() *redis.ClusterSlotsCmd {
	return client.ClusterSlots()
}

func ClusterNodes() *redis.StringCmd {
	return client.ClusterNodes()
}

func ClusterMeet(host, port string) *redis.StatusCmd {
	return client.ClusterMeet(host, port)
}

func ClusterForget(nodeID string) *redis.StatusCmd {
	return client.ClusterForget(nodeID)
}

func ClusterReplicate(nodeID string) *redis.StatusCmd {
	return client.ClusterReplicate(nodeID)
}

func ClusterResetSoft() *redis.StatusCmd {
	return client.ClusterResetSoft()
}

func ClusterResetHard() *redis.StatusCmd {
	return client.ClusterResetHard()
}

func ClusterInfo() *redis.StringCmd {
	return client.ClusterInfo()
}

func ClusterKeySlot(key string) *redis.IntCmd {
	return client.ClusterKeySlot(key)
}

func ClusterGetKeysInSlot(slot int, count int) *redis.StringSliceCmd {
	return client.ClusterGetKeysInSlot(slot, count)
}

func ClusterCountFailureReports(nodeID string) *redis.IntCmd {
	return client.ClusterCountFailureReports(nodeID)
}

func ClusterCountKeysInSlot(slot int) *redis.IntCmd {
	return client.ClusterCountKeysInSlot(slot)
}

func ClusterDelSlots(slots ...int) *redis.StatusCmd {
	return client.ClusterDelSlots(slots...)
}

func ClusterDelSlotsRange(min, max int) *redis.StatusCmd {
	return client.ClusterDelSlotsRange(min, max)
}

func ClusterSaveConfig() *redis.StatusCmd {
	return client.ClusterSaveConfig()
}

func ClusterSlaves(nodeID string) *redis.StringSliceCmd {
	return client.ClusterSlaves(nodeID)
}

func ClusterFailover() *redis.StatusCmd {
	return client.ClusterFailover()
}

func ClusterAddSlots(slots ...int) *redis.StatusCmd {
	return client.ClusterAddSlots(slots...)
}

func ClusterAddSlotsRange(min, max int) *redis.StatusCmd {
	return client.ClusterAddSlotsRange(min, max)
}

func GeoAdd(key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd {
	return client.GeoAdd(key, geoLocation...)
}

func GeoPos(key string, members ...string) *redis.GeoPosCmd {
	return client.GeoPos(key, members...)
}

func GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	return client.GeoRadius(key, longitude, latitude, query)
}


func GeoRadiusByMember(key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	return client.GeoRadiusByMember(key, member, query)
}

func GeoDist(key string, member1, member2, unit string) *redis.FloatCmd {
	return client.GeoDist(key, member1, member2, unit)
}

func GeoHash(key string, members ...string) *redis.StringSliceCmd {
	return client.GeoHash(key, members...)
}

func ReadOnly() *redis.StatusCmd {
	return client.ReadOnly()
}

func ReadWrite() *redis.StatusCmd {
	return client.ReadWrite()
}

func MemoryUsage(key string, samples ...int) *redis.IntCmd {
	return client.MemoryUsage(key, samples...)
}

