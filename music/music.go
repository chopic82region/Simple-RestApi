package music

type Music struct {
	Title  string
	Author string

	Downloaded bool
}

func NewMusic(name string, author string) Music {
	return Music{
		Title:  name,
		Author: author,

		Downloaded: false,
	}

}
