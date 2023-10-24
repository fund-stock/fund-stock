package k780

type List struct {
	Success string `json:"success"`
	Result  struct {
		Totline string `json:"totline"`
		Disline string `json:"disline"`
		Page    string `json:"page"`
		Lists   []struct {
			Symbol string `json:"symbol"`
			Sname  string `json:"sname"`
		} `json:"lists"`
	} `json:"result"`
}
