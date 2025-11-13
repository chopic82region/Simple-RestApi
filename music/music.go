package music

type Music struct {
	Title  string
	Author string

	Downloaded bool
}

func NewMusic(name string, author string) Music {
	music := Music{}
}
