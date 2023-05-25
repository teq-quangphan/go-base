package util

import "time"

func ShardingHashKey(id int, g int) uint {
	currTime := uint(time.Now().UnixMilli()) << 23
	shardId := uint(g%2) << 10
	seqId := uint(id) << 0
	return currTime | shardId | seqId
}
func Decode(id int) (int, int) {
	shardId := (uint(id) >> 10) & 0x1FFFFFFFFFF
	userId := (uint(id) >> 0) & 0x3FF
	return int(userId), int(shardId)
}
