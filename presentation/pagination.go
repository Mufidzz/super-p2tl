package presentation

type Pagination struct {
	Offset int64 `json:"offset"`
	Count  int64 `json:"count"`
}
