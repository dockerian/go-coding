package stock

// ==== byExchangeID ====

// byExchangeID implements sort.Interface for Exchange
type byExchangeID []*Exchange

// Len implements Len of sort.Interface
func (ex byExchangeID) Len() int {
	return len(ex)
}

// Less implements Less of sort.Interface
func (ex byExchangeID) Less(i, j int) bool {
	return ex[i].ID < ex[j].ID
}

// Swap implements Swap of sort.Interface
func (ex byExchangeID) Swap(i, j int) {
	ex[i], ex[j] = ex[j], ex[i]
}
