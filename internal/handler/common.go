package handler

// 返回消息结构体
type MRsp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

// 接收请求id的结构体
type IdInfo struct {
	Id int `json:"id"`
}
