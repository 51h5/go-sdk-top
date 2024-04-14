package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/51h5/go-sdk-top/internal/constants"
)

func SignToRequest(params url.Values, body []byte, secret, signMethod string) string {
	if params == nil {
		return ""
	}

	// 1. 复制副本, 用于后续排序 (顺便过滤掉空值参数)
	vs := url.Values{}
	for k := range params {
		if params.Get(k) != "" {
			vs.Set(k, params.Get(k))
		}
	}

	// 2. 参数排序拼接
	keys := make([]string, 0, len(vs))
	for k := range vs {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var buf bytes.Buffer

	if signMethod == constants.SIGN_TYPE_MD5 {
		buf.WriteString(secret)
	}

	for _, k := range keys {
		buf.WriteString(k)
		buf.WriteString(vs.Get(k))
	}

	// 3. 拼接请求主体
	if body != nil && len(body) > 0 {
		buf.Write(body)
	}

	// 4. 计算签名
	if signMethod == constants.SIGN_TYPE_HMAC_SHA256 {
		return encryptHMACSHA256(buf.Bytes(), []byte(secret))
	} else if signMethod == constants.SIGN_TYPE_HMAC {
		return encryptHMAC(buf.String(), secret)
	} else {
		buf.WriteString(secret)
		return encryptMD5(buf.String())
	}
}

func encryptMD5(s string) string {
	return strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(s))))
}

// TODO: Hmac 签名算法
func encryptHMAC(s, secret string) string {
	return ""
}

func encryptHMACSHA256(b, secret []byte) string {
	h := hmac.New(sha256.New, secret)
	h.Write(b)
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
