// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	UserName string `form:"userName",options`
	Password string `form:"password"`
	Mobile   string `form:"mobile"`
}

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}

type UserInfoResponse struct {
	UserName string `json:"userName"`
	Addr     string `json:"addr"`
	Id       uint   `json:"id"`
}
