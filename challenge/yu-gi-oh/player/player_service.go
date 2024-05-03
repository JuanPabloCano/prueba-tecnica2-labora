package player

func CreatePlayer(name string) *Player {
	newPlayer := &Player{
		Name: name,
	}
	Players = append(Players, newPlayer.Name)
	return newPlayer
}
