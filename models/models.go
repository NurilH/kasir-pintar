package models

type Response struct {
	Movie []Movie `json:"Search"`
}

type Movie struct {
	ImdbID string `json:"imdbID"`
	Title  string `json:"Title"`
	// Year   string `json:"Year"`
	// Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type MovieDetail struct {
	ImdbID     string `json:"imdbID"`
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Type       string `json:"Type"`
	Poster     string `json:"Poster"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Country    string `json:"Country"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
}

type VirtuanAcount struct {
	TimeStamp      string `json:"timeStamp"`
	IMid           string `json:"iMid"`
	PayMethod      string `json:"payMethod"`
	Currency       string `json:"currency"`
	Amt            string `json:"amt"`
	ReferenceNo    string `json:"referenceNo"`
	GoodsNm        string `json:"goodsNm"`
	BillingNm      string `json:"billingNm"`
	BillingPhone   string `json:"billingPhone"`
	BillingEmail   string `json:"billingEmail"`
	BillingAddr    string `json:"billingAddr"`
	BillingCity    string `json:"billingCity"`
	BillingState   string `json:"billingState"`
	BillingPostCd  string `json:"billingPostCd"`
	BillingCountry string `json:"billingCountry"`
	// DeliveryNm      string `json:"deliveryNm"`
	// DeliveryPhone   string `json:"deliveryPhone"`
	// DeliveryAddr    string `json:"deliveryAddr"`
	// DeliveryCity    string `json:"deliveryCity"`
	// DeliveryState   string `json:"deliveryState"`
	// DeliveryPostCd  string `json:"deliveryPostCd"`
	// DeliveryCountry string `json:"deliveryCountry"`
	// Description     string `json:"description"`
	DbProcessUrl  string `json:"dbProcessUrl"`
	MerchantToken string `json:"merchantToken"`
	ReqDomain     string `json:"reqDomain"`
	ReqServerIp   string `json:"reqServerIP"`
	UserIp        string `json:"userIP"`
	UserSessionID string `json:"userSessionID"`
	UserAgent     string `json:"userAgent"`
	UserLanguage  string `json:"userLanguage"`
	CartData      string `json:"cartData"`
	BankCd        string `json:"bankCd"`
	VaccValidDt   string `json:"vacctValidDt"`
	VaccValidTm   string `json:"vacctValidTm"`
	MerFixAcctId  string `json:"merFixAcctId"`
}

type ResVirtual struct {
	ResultCd    string `json:"resultCd"`
	ResultMsg   string `json:"resultMsg" `
	TXid        string `json:"tXid"`
	ReferenceNo string `json:"referenceNo"`
	PayMethod   string `json:"payMethod"`
	Amt         string `json:"amt"`
	TransDt     string `json:"transDt"`
	TransTm     string `json:"transTm"`
	BankCd      string `json:"bankCd"`
	VacctNo     string `json:"vacctNo"`
	MitraCd     string `json:"mitraCd"`
	PayNo       string `json:"payNo"`
	Currency    string `json:"currency"`
	GoodsNm     string `json:"goodsNm"`
	BillingNm   string `json:"billingNm"`
	VaccValidDt string `json:"vacctValidDt"`
	VaccValidTm string `json:"vacctValidTm"`
}
