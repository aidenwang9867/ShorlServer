package utils

import (
	"github.com/catinello/base62"
	"hash/crc32"
)

func EncodeLink(link string) string {
	hash := crc32.ChecksumIEEE([]byte(link))
	return base62.Encode(int(hash))
}
