package ggapp

type PlayStatus struct {
	Status PlayStatusCode `json:"id"`
	Title  string         `json:"title"`
}

type CurrentUserGamePlayStatus struct {
	Id         int        `json:"id"`
	PlayStatus PlayStatus `json:"playStatus"`
}

type Game struct {
	Name                      string                    `json:"name"`
	Slug                      string                    `json:"slug"`
	Description               string                    `json:"string"`
	CoverPath                 string                    `json:"coverPath"`
	CurrentUserGamePlayStatus CurrentUserGamePlayStatus `json:"currentUserGamePlayStatus"`
}

type GameWithStatus struct {
	Game       Game       `json:"game"`
	PlayStatus PlayStatus `json:"playStatus"`
}
