package top

import "net/url"

type Request interface {
    Check() (uint, error)
    Method() string
    Values() url.Values

    Body() []byte
    SetBody([]byte)

    TargetAppKey() string
    SetTargetAppKey(string)
}
