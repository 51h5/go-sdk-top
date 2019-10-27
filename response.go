package top

type Response interface {
    Success() bool
    Fix()
}
