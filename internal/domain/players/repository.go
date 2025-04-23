package players

type Repository interface {
	Create(player *Player) error
	FindByHashedID(hashedID string) (*Player, error)
}
