package utils

import (
    "51h5.com/sdk/top/internal/constants"
    "bytes"
    "crypto/md5"
    "fmt"
    "net/url"
    "sort"
    "strings"
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
    if signMethod == constants.SIGN_TYPE_HMAC {
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

// TODO: HmacSHA256 签名算法
func encryptHMACSHA256(s, secret string) string {
    return ""
}

// func SignToRequest(sysParams, apiParams url.Values) string {
//     if sysParams == nil {
//         return ""
//     }
//
//     vs := url.Values{}
//
//     // 业务参数
//     if apiParams != nil {
//         for k := range apiParams {
//             if apiParams.Get(k) != "" {
//                 vs.Set(k, apiParams.Get(k))
//             }
//         }
//     }
//
//     // 系统参数
//     for k := range sysParams {
//         if sysParams.Get(k) != "" {
//             vs.Set(k, sysParams.Get(k))
//         }
//     }
//
//     keys := make([]string, 0, len(vs))
//     for k := range vs {
//         keys = append(keys, k)
//     }
//
//     sort.Strings(keys)
//
//     var buf bytes.Buffer
//     for _, k := range keys {
//         // if buf.Len() > 0 {
//         //     buf.WriteString("&")
//         // }
//         buf.WriteString(k)
//         // buf.WriteString("=")
//         buf.WriteString(vs.Get(k))
//     }
//
//     return buf.String()
// }
