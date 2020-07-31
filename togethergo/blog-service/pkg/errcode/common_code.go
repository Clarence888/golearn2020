package errcode

var (
	Success = NewError(0,"成功")
	ServerError = NewError(1000000,"服务内部错误")
	InvalidParams = NewError(1000001,"入参错误")
	NotFound = NewError(1000002,"找不到信息")
	UnauthorizedAuthNotExist = NewError(1000003,"鉴权失败,找不到对应鉴权信息")
	UnauthorizedTokenError = NewError(1000004,"鉴权失败,token错误")
	UnauthorizedTokenTimeout = NewError(1000005,"鉴权失败,token超时")
	UnauthorizedTokenGenerate = NewError(1000006,"鉴权失败,token生成失败")
	TooManyRequests = NewError(1000007,"请求过多")
)



