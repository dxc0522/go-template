package typex

type ErrorBody struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
