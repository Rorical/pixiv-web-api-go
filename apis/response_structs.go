package apis

type IllustDetailBody struct {
	Alt           string `json:"alt"`
	BookmarkCount uint   `json:"bookmarkCount"`
}

type IllustDetailResponse struct {
	Message string           `json:"message"`
	Error   bool             `json:"error"`
	Body    IllustDetailBody `json:"body"`
}
