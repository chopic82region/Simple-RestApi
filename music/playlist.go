package music

import "sync"

type Playlist struct {
	playlist map[string]Music
	mu       sync.RWMutex
}

// Конструктор структуры
func NewPlaylist() *Playlist {
	return &Playlist{
		playlist: make(map[string]Music),
	}
}

func (p *Playlist) AddMusic(m Music) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.playlist[m.Title] = m

}

func (p *Playlist) PlayMusic(title string) (Music, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	music, exist := p.playlist[title]
	if !exist {
		return Music{}, ErrMusicNotFound
	}

	return music, nil
}

func (p *Playlist) ShowPlaylist() map[string]Music {
	p.mu.RLock()
	defer p.mu.RUnlock()

	// Дублируем мапу для изоляции основной мапы
	temporaryMap := make(map[string]Music)

	for k, v := range p.playlist {
		temporaryMap[k] = v
	}

	return temporaryMap
}

func (p *Playlist) DownloadMusic(title string) (Music, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	music, exist := p.playlist[title]
	if !exist {
		return Music{}, ErrMusicNotFound
	}

	music.Downloaded = true

	p.playlist[title] = music

	return music, nil
}

func (p *Playlist) ShowDownloadedMusic() map[string]Music {
	p.mu.RLock()
	defer p.mu.RUnlock()

	downloadedPlaylist := make(map[string]Music)

	for k, v := range p.playlist {
		if v.Downloaded == true {
			downloadedPlaylist[k] = v
		}
	}

	return downloadedPlaylist
}

func (p *Playlist) DeleteMusic(title string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, exist := p.playlist[title]
	if !exist {
		return ErrMusicNotFound
	}

	delete(p.playlist, title)

	return nil
}
