package handler

type Request struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

type IdInfo struct {
	Id int `json:"id"`
}
