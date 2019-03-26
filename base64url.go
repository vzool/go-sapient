package sapient

// Original base64url Package copied from
// https://github.com/kalaspuffar/base64url @ 2018-03-24 11:39 PM
// Reasons:	- Make sapient package depends only on golang standard packages.
// 			- Offer extra code coverage for the whole package as general.
//			-

import (
	"encoding/base64"
	"errors"
	"strings"
)

// Base64UrlEncode is the unpadded alternate base64 encoding defined in RFC 4648.
// It is typically used in URLs and file names.
// This is the same as URLEncoding but omits padding characters.
func Base64UrlEncode(data []byte) string {
	str := base64.StdEncoding.EncodeToString(data)
	str = strings.Replace(str, "+", "-", -1)
	str = strings.Replace(str, "/", "_", -1)
	str = strings.Replace(str, "=", "", -1)
	return str
}

// Base64UrlDecode is the unpadded alternate base64 decoding defined in RFC 4648.
// It is typically used in URLs and file names.
// This is the same as URLEncoding but omits padding characters.
func Base64UrlDecode(str string) ([]byte, error) {
	if strings.ContainsAny(str, "+/") {
		return nil, errors.New("invalid base64url encoding")
	}
	str = strings.Replace(str, "-", "+", -1)
	str = strings.Replace(str, "_", "/", -1)
	for len(str)%4 != 0 {
		str += "="
	}
	return base64.StdEncoding.DecodeString(str)
}
