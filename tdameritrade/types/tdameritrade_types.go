package tdamertrade

type Credentials struct {
	UserId      string `json:"userid`
	Token       string `json:"token"`
	Company     string `json:"company"`
	Segment     string `json:"segment"`
	CdDomain    string `json:"cddomain"`
	UserGroup   string `json:"usergroup"`
	AccessLevel string `json:"accesslevel"`
	Authorized  string `json:"authorized"`
	Timestamp   int64  `json:"timestamp"`
	AppID       string `json:"appID"`
	ACL         string `json:"acl"`
}

type StreamerInfo struct {
	StreamerBinaryUrl string `json:"streamerBinaryUrl"`
	StreamerSocketUrl string `json:"streamerSocketUrl"`
	Token             string `json:"token"`
	TokenTimestamp    string `json:"tokenTimestamp"`
	UserGroup         string `json:"userGroup"`
	AccessLevel       string `json:"accessLevel"`
	ACL               string `json:"acl"`
	AppID             string `json:"appID"`
}

type Quotes struct {
	IsNyseDelayed   bool `json:"isNyseDelayed"`
	IsNasdaqDelayed bool `json:"isNasdaqDelayed"`
	IsOpraDelayed   bool `json:"isOpraDelayed"`
	IsAmexDelayed   bool `json:"isAmexDelayed"`
	IsCmeDelayed    bool `json:"isCmeDelayed"`
	IsIceDelayed    bool `json:"isIceDelayed"`
	IsForexDelayed  bool `json:"isForexDelayed"`
}

type Key struct {
	Key string `json:"key"`
}

type StreamerSubscriptionKeys struct {
	Keys []Key `json:"keys"`
}

type ExchangeAgreements struct {
	NasdaqExchangeAgreement string `json:"NASDAQ_EXCHANGE_AGREEMENT"`
	NyseExchangeAgreement   string `json:"NYSE_EXCHANGE_AGREEMENT"`
	OpraExchangeAgreement   string `json:"OPRA_EXCHANGE_AGREEMENT"`
}

type Authorizations struct {
	Apex               bool   `json:"apex"`
	LevelTwoQuotes     bool   `json:"levelTwoQuotes"`
	StockTrading       bool   `json:"stockTrading"`
	MarginTrading      bool   `json:"marginTrading"`
	StreamingNews      bool   `json:"streamingNews"`
	OptionTradingLevel string `json:"optionTradingLevel"`
	ScottradeAccount   bool   `json:"scottradeAccount"`
	AutoPositionEffect bool   `json:"autoPositionEffect"`
}

type Account struct {
	AccountId         string         `json:"accountId"`
	DisplayName       string         `json:"displayName"`
	AccountCdDomainId string         `json:"accountCdDomainId"`
	Company           string         `json:"company"`
	Segment           string         `json:"segment"`
	Acl               string         `json:"acl"`
	Authorizations    Authorizations `json:"authorizations"`
}

type UserPrincipals struct {
	UserId                   string                   `json:"userId"`
	UserCdDomainId           string                   `json:"userCdDomainId"`
	PrimaryAccountId         string                   `json:"primaryAccountId"`
	LastLoginTime            string                   `json:"lastLoginTime"`
	TokenExpirationTime      string                   `json:"tokenExpirationTime"`
	LoginTime                string                   `json:"loginTime"`
	AccessLevel              string                   `json:"accessLevel"`
	StalePassword            bool                     `json:"stalePassword"`
	StreamerInfo             StreamerInfo             `json:"streamerInfo"`
	ProfessionalStatus       string                   `json:"professionalStatus"`
	Quotes                   Quotes                   `json:"quotes"`
	StreamerSubscriptionKeys StreamerSubscriptionKeys `json:"streamerSubscriptionKeys"`
	ExchangeAgreements       ExchangeAgreements       `json:"exchangeAgreements"`
	Accounts                 []Account                `json:"accounts"`
}

type RequestParameters struct {
	Credential string `json:"credential"`
	Token      string `json:"token"`
	Version    string `json:"version"`
}

type WebSocketRequest struct {
	Service    string            `json:"service"`
	Command    string            `json:"command"`
	RequestID  int               `json:"requestid"`
	Account    string            `json:"account"`
	Source     string            `json:"source"`
	Parameters map[string]string `json:"parameters"`
}

type WebSocketRequests struct {
	Requests []WebSocketRequest `json:"requests"`
}

type StreamContent struct {
	Key         string  `json:"key"`
	Zero        string  `json:"0"`
	One         float64 `json:"1"`
	Two         float64 `json:"2"`
	Three       float64 `json:"3"`
	Four        float64 `json:"4"`
	Five        float64 `json:"5"`
	Six         string  `json:"6"`
	Seven       string  `json:"7"`
	Eight       float64 `json:"8"`
	Nine        float64 `json:"9"`
	Ten         float64 `json:"10"`
	Eleven      float64 `json:"11"`
	Twelve      float64 `json:"12"`
	Thirteen    float64 `json:"13"`
	Fourteen    string  `json:"14"`
	Fifteen     float64 `json:"15"`
	Sixteen     string  `json:"16"`
	Seventeen   bool    `json:"17"`
	Eighteen    bool    `json:"18"`
	Nineteen    float64 `json:"19"`
	Twenty      float64 `json:"20"`
	TwentyOne   float64 `json:"21"`
	TwentyTwo   float64 `json:"22"`
	TwentyThree float64 `json:"23"`
	TwentyFour  float64 `json:"24"`
	TwentyFive  string  `json:"25"`
	TwentySix   string  `json:"26"`
	TwentySeven float64 `json:"27"`
	TwentyEight float64 `json:"28"`
	TwentyNine  float64 `json:"29"`
	Thirty      float64 `json:"30"`
	ThirtyOne   float64 `json:"31"`
	ThirtyTwo   float64 `json:"32"`
	ThirtyThree float64 `json:"33"`
	ThirtyFour  float64 `json:"34"`
	ThirtyFive  float64 `json:"35"`
	ThirtySix   float64 `json:"36"`
	ThirtySeven float64 `json:"37"`
	ThirtyEight float64 `json:"38"`
	ThirtyNine  string  `json:"39"`
	Forty       string  `json:"40"`
	FortyOne    bool    `json:"41"`
	FortyTwo    bool    `json:"42"`
	FortyThree  float64 `json:"43"`
	FortyFour   float64 `json:"44"`
	FortyFive   float64 `json:"45"`
	FortySix    float64 `json:"46"`
	FortySeven  float64 `json:"47"`
	FortyEight  string  `json:"48"`
	FortyNine   float64 `json:"49"`
	Fifty       float64 `json:"50"`
	FiftyOne    float64 `json:"51"`
	FiftyTwo    int64   `json:"52"`
}

type StreamData struct {
	Service   string          `json:"service"`
	Timestamp int64           `json:"timestamp"`
	Command   string          `json:"command"`
	Content   []StreamContent `json:"content"`
}

type StreamMessage struct {
	Data []StreamData `json:"data"`
}
