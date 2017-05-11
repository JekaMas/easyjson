package jlexer

import (
	"github.com/JekaMas/easyjson/fasthash"
	"github.com/JekaMas/easyjson/hashmap/fastinteger"
)

func NewFastHashMap(hint uint64) *fastinteger.FastIntegerHashMap {
	return fastinteger.New(hint)
}

func GetHashString(item string) uint64 {
	const primeSeed = 380383699

	return fasthash.Hash64(primeSeed, StrToBytes(item))
}

func GetHash(item []byte) uint64 {
	const primeSeed = 380383699

	return fasthash.Hash64(primeSeed, item)
}