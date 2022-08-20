package presentation

type GetBankDataListResponse struct {
	Category string                        `json:"category"`
	Files    []GetBankDataListResponseFile `json:"files"`
}

type GetBankDataListResponseFile struct {
	Name string `json:"name"`
}
