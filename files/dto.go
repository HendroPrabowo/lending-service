package files

type FileQueryParam struct {
	Id        int    `json:"id"`
	AccountId int    `json:"account_id"`
	Type      string `json:"type"`
}

type UploadResponseDto struct {
	IdFile    int    `json:"id_file"`
	AccountId int    `json:"account_id"`
	Type      string `json:"type"`
}
