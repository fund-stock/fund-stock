package params

type Data struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	MaterialInfo struct {
		ProductId string `json:"productId"`
		FundCode  string `json:"fundCode"`
		FundType  string `json:"fundType"`
		TitleInfo struct {
			FundLimit         string `json:"fundLimit"`
			NetValue          string `json:"netValue"`
			NetValueDate      string `json:"netValueDate"`
			ProfitSevenDays   string `json:"profitSevenDays"`
			ProfitTenThousand string `json:"profitTenThousand"`
			DayOfGrowth       string `json:"dayOfGrowth"`
			LastWeek          string `json:"lastWeek"`
			RiskEvaluation    string `json:"riskEvaluation"`
			EstablishmentDate string `json:"establishmentDate"`
			AssetSize         string `json:"assetSize"`
			FundManagerName   string `json:"fundManagerName"`
		} `json:"titleInfo"`
		FundBrief struct {
			FundNameAbbr      string `json:"fundNameAbbr"`
			FundName          string `json:"fundName"`
			FundCode          string `json:"fundCode"`
			EstablishmentDate string `json:"establishmentDate"`
			ShareSize         string `json:"shareSize"`
			AssetSize         string `json:"assetSize"`
			FundManagerName   string `json:"fundManagerName"`
			SaleStatus        string `json:"saleStatus"`
			FundCompanyName   string `json:"fundCompanyName"`
			TrusteeName       string `json:"trusteeName"`
			ManageRate        string `json:"manageRate"`
			TrusteeRate       string `json:"trusteeRate"`
			PurchaseMinMount  string `json:"purchaseMinMount"`
			RedeemMinMount    string `json:"redeemMinMount"`
			PurchaseRatio     string `json:"purchaseRatio"`
			RedeemRatio       string `json:"redeemRatio"`
			GeneralInfo       struct {
				FundName              string `json:"fundName"`
				EstablishmentDate     string `json:"establishmentDate"`
				FundCode              string `json:"fundCode"`
				AssetSize             string `json:"assetSize"`
				FundCompanyName       string `json:"fundCompanyName"`
				TrusteeName           string `json:"trusteeName"`
				FundManagerBackground string `json:"fundManagerBackground"`
				FundManagerInfoList   []struct {
					Key        string `json:"key"`
					FundName   string `json:"fundName"`
					OfficeDate string `json:"officeDate"`
					Earnings   string `json:"earnings"`
				} `json:"fundManagerInfoList"`
				InvestPhilosophy string `json:"investPhilosophy"`
				InvestStrategy   string `json:"investStrategy"`
			} `json:"generalInfo"`
		} `json:"fundBrief"`
	} `json:"materialInfo"`
	IsLogin         bool   `json:"isLogin"`
	Csrf            string `json:"csrf"`
	IsCloseEstimate bool   `json:"isCloseEstimate"`
	PageName        string `json:"pageName"`
	UriBroker       struct {
		FaviconIcoUrl   string `json:"favicon.ico.url"`
		App404Url       string `json:"app.404.url"`
		ZdrmdataRestUrl string `json:"zdrmdata.rest.url"`
		AppErrorpageUrl string `json:"app.errorpage.url"`
		AuthcenterUrl   string `json:"authcenter.url"`
		AppGotoUrl      string `json:"app.goto.url"`
		BumngUrl        string `json:"bumng.url"`
		OmeoCheckUrl    string `json:"omeo.check.url"`
		OmeoGetUrl      string `json:"omeo.get.url"`
		AssetsUrl       string `json:"assets.url"`
	} `json:"uriBroker"`
}

type HistoryData struct {
	Success bool `json:"success"`
	List    []struct {
		Key           int    `json:"key"`
		NetValueDate  string `json:"netValueDate"`
		NetValue      string `json:"netValue"`
		TotalNetValue string `json:"totalNetValue"`
		DayOfGrowth   string `json:"dayOfGrowth"`
	} `json:"list"`
	TotalPages  int `json:"totalPages"`
	CurrentPage int `json:"currentPage"`
	TotalItems  int `json:"totalItems"`
}
