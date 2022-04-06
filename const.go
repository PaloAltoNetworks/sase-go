package sase

// Logging constants.
const (
	LogQuiet = 1 << (iota + 1)
	LogGet
	LogPost
	LogPut
	LogDelete
	LogPath
	LogSend
	LogReceive
)
