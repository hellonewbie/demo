package common

//CollectEntry 要收集的日志的结构体
type CollectEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}
