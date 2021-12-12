package response

type Response interface {
	GetStatus() string
	GetCode() int
	GetMessage() string
	GetItems() interface{}
}

type BaseResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Items   interface{} `json:"items"`
}

func (b *BaseResponse) GetStatus() string {
	return b.Status
}

func (b *BaseResponse) GetCode() int {
	return b.Code
}

func (b *BaseResponse) GetMessage() string {
	return b.Message
}

func (b *BaseResponse) GetItems() interface{} {
	return b.Items
}
