package being

// organization
type relation struct {
	bid   bid
	value int64
}

// summary for relation
type relationship struct {
	father   bid
	mother   bid
	children []bid
	friends  []bid
	lovers   []bid
}
