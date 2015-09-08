package steamwebapi

// ByMatchID implements sort.Interface for Matches based on the ID field.
type ByMatchID Matches

func (m ByMatchID) Len() int           { return len(m) }
func (m ByMatchID) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m ByMatchID) Less(i, j int) bool { return m[i].ID < m[j].ID }
