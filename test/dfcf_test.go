package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test(t *testing.T) {
	Client("002261")
	Detail("002261")
}

func Client(code string) {
	url := fmt.Sprintf(`https://push2.eastmoney.com/api/qt/stock/details/get?fields1=f1,f2,f3,f4&fields2=f51,f52,f53,f54,f55&fltt=2&pos=-11&secid=0.%v&ut=fa5fd1943c7b386f172d6893dbfba10b&_=1695021427324`, code)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var resp StockDetailResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	fmt.Println(resp)
}

func Detail(code string) {
	url := fmt.Sprintf(`https://np-anotice-stock.eastmoney.com/api/security/ann?page_size=5&page_index=1&market_code=0&stock_list=%v`, code)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var resp StockDetail2Resp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	fmt.Println(resp)
}

type StockDetailResp struct {
	Rc     int    `json:"rc"`
	Rt     int    `json:"rt"`
	Svr    int    `json:"svr"`
	Lt     int    `json:"lt"`
	Full   int    `json:"full"`
	Dlmkts string `json:"dlmkts"`
	Data   struct {
		Code     string   `json:"code"`
		Market   int      `json:"market"`
		Decimal  int      `json:"decimal"`
		PrePrice float64  `json:"prePrice"`
		Details  []string `json:"details"`
	} `json:"data"`
}

type StockDetail2Resp struct {
	Data struct {
		List []struct {
			ArtCode string `json:"art_code"`
			Codes   []struct {
				AnnType    string `json:"ann_type"`
				InnerCode  string `json:"inner_code"`
				MarketCode string `json:"market_code"`
				ShortName  string `json:"short_name"`
				StockCode  string `json:"stock_code"`
			} `json:"codes"`
			Columns []struct {
				ColumnCode string `json:"column_code"`
				ColumnName string `json:"column_name"`
			} `json:"columns"`
			DisplayTime string `json:"display_time"`
			EiTime      string `json:"eiTime"`
			Language    string `json:"language"`
			NoticeDate  string `json:"notice_date"`
			ProductCode string `json:"product_code"`
			SortDate    string `json:"sort_date"`
			SourceType  string `json:"source_type"`
			Title       string `json:"title"`
			TitleCh     string `json:"title_ch"`
			TitleEn     string `json:"title_en"`
		} `json:"list"`
		PageIndex int `json:"page_index"`
		PageSize  int `json:"page_size"`
		TotalHits int `json:"total_hits"`
	} `json:"data"`
	Error   string `json:"error"`
	Success int    `json:"success"`
}
