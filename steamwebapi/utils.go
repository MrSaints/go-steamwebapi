package steamwebapi

// Sort by Match Id.
type ByMatchId Matches

func (m ByMatchId) Len() int           { return len(m) }
func (m ByMatchId) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m ByMatchId) Less(i, j int) bool { return m[i].Id < m[j].Id }
