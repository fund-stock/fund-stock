package response

type WebLoginUser struct {
	IsWorkSubmit       int    `json:"isWorkSubmit"`
	ItemCode           string `json:"itemCode"`
	LastWorkSubmitDate string `json:"lastWorkSubmitDate"`
	Nation             string `json:"nation"`
	RoleCode           string `json:"roleCode"`
	SystemUserId       string `json:"systemUserId"`
	UserNick           string `json:"userNick"`
}
