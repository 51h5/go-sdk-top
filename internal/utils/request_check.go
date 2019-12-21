package utils

import (
    "fmt"
)

const(
    errCodeArgumentsMissing = iota + 40
    errCodeArgumentsInvalid
)

func CheckNotEmpty(k string, v string) (code uint, err error) {
    if v == "" {
        code = errCodeArgumentsMissing
        err = fmt.Errorf("client-error:Missing required arguments: %s", k)
        return
    }
    return
}

func CheckNotEmptyList(k string, v int) (code uint, err error) {
    if v == 0 {
        code = errCodeArgumentsMissing
        err = fmt.Errorf("client-error:Missing required arguments: %s", k)
        return
    }
    return
}

func CheckMaxLength(k string, v interface{}) (code uint, err error) {
    // v: file | string
    return
}

func CheckMaxSize(k string, v interface{}) (code uint, err error) {
    // v: slice | array | map | string
    return
}

func CheckMaxValue(k string, v int64) (code uint, err error) {
    return
}

func CheckMinValue(k string, v int64) (code uint, err error) {
    return
}