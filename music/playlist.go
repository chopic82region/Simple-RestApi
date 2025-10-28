package music

type Playlist struct {
	playlist map[string]Music
}

// Конструктор структуры
func NewPlaylist() *Playlist {
	return &Playlist{
		playlist: make(map[string]Music),
	}
}

func (p *Playlist) AddMusic(m Music) {

	p.playlist[m.Title] = m
}

func (p *Playlist) PlayMusic(title string) (Music, error) {

	music, exist := p.playlist[title]
	if !exist {
		return Music{}, ErrMusicNotFound
	}

	return music, nil
}

func (p *Playlist) ShowPlaylist() map[string]Music {

	// Дублируем мапу для изоляции основной мапы

	temporaryMap := make(map[string]Music)

	for k, v := range p.playlist {
		temporaryMap[k] = v
	}

	return temporaryMap
}

func (p *Playlist) DownloadMusic(title string) error {

	music, exist := p.playlist[title]
	if !exist {
		return ErrMusicNotFound
	}

	music.Downloaded = true

	p.playlist[title] = music

	return nil
}

func (p *Playlist) ShowDownloadedMusic() map[string]Music {

	downloadedPlaylist := make(map[string]Music)

	for k, v := range p.playlist {
		if v.Downloaded == true {
			downloadedPlaylist[k] = v
		}
	}

	return downloadedPlaylist
}

func (p *Playlist) DeleteMusic(title string) error {

	_, exist := p.playlist[title]
	if !exist {
		return ErrMusicNotFound
	}

	delete(p.playlist, title)

	return nil
}
