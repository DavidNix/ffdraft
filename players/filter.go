package players

type filterFunc func(p Player) bool

func filter(ps []Player, f filterFunc) []Player {
    filtered := []Player{}
    for _, player := range ps {
        if f(player) {
            filtered = append(filtered, player)
        }
    }
    return filtered
}
