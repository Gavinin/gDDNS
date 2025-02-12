package enum

type InitialLogCacheErrorType string

const (
	InitialLogCacheInfo     = "INFO"
	InitialLogCacheWarrning = "WARNING"
	InitialLogCacheError    = "ERROR"
	InitialLogCachePanic    = "PANIC"
)
