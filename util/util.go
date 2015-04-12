package util
import (
    "crypto/md5"
    "encoding/hex"
    "bytes"
)

func StringToMD5(s string) string {
    hasher := md5.New()
    hasher.Write([]byte(s))
    return hex.EncodeToString(hasher.Sum(nil))
}

func EqualHashes(hash1 []byte, hash2 []byte) bool {
    return bytes.Compare(hash1[:], hash2[:]) == 0
}