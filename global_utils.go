package jslike

import "encoding/hex"

// !
// ! all copy by https://github.com/wangluozhe/requests
// ! https://github.com/wangluozhe/requests/blob/2403fb5b60b83a9c29235cd6568c447e63cf54f1/utils/code.go#L76
// !
// !

func hexEncode(byte_s []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(byte_s)))
	n := hex.Encode(dst, byte_s)
	return dst[:n]
}
