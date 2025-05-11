package dto

type LambdaCallbackResponse struct {
	Status int `json:"status"`
	Res    any `json:"res"`
	Err    any `json:"err"`
}
