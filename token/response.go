package token

type TokenResponse struct {
	Token    string  `json:"token" xml:"Token"`
	ExpTime  int64   `json:"exp_time" xml:"ExpTime"`
	ExpireIn float64 `json:"expire_in" xml:"ExpireIn"`
}

type Response struct {
	Disk   string `json:"disk" xml:"Disk"`
	Bucket string `json:"bucket" xml:"Bucket"`
	Object string `json:"object" xml:"Object"`
}
