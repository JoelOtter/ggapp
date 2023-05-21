package ggapp

type PlayStatusCode int

const (
	PlayStatusNone = iota
	PlayStatusWantToPlay
	PlayStatusPlaying
	PlayStatusBeaten
	PlayStatusCompleted
	PlayStatusShelved
	PlayStatusAbandoned
)

var AllPlayStatuses = []PlayStatusCode{
	PlayStatusWantToPlay,
	PlayStatusPlaying,
	PlayStatusBeaten,
	PlayStatusCompleted,
	PlayStatusShelved,
	PlayStatusAbandoned,
}

func PlayStatusName(status PlayStatusCode) string {
	switch status {
	case PlayStatusNone:
		return "None"
	case PlayStatusWantToPlay:
		return "Want to play"
	case PlayStatusPlaying:
		return "Playing"
	case PlayStatusBeaten:
		return "Beaten"
	case PlayStatusCompleted:
		return "Completed"
	case PlayStatusShelved:
		return "Shelved"
	case PlayStatusAbandoned:
		return "Abandoned"
	}
	return ""
}
