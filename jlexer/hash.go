package jlexer

import "github.com/JekaMas/easyjson/fasthash"

func GetHashString(item string) uint64 {
	const primeSeed = 380383699

	return fasthash.Hash64(primeSeed, StrToBytes(item))
}