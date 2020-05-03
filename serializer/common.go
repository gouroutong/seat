package serializer

// Response 基础序列化器
type Response struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	ErrMsg string      `json:"errMsg"`
}
type DataList struct {
	List interface{} `json:"list"`
}

//列表序列化器
func ListResponse(list interface{}, err error) Response {
	if err != nil {
		return Response{
			Code:   500,
			Result: "",
			ErrMsg: err.Error(),
		}
	}
	return Response{
		Code: 200,
		Result: DataList{
			List: list,
		},
		ErrMsg: "",
	}
}
func GetResponse(result interface{}, err error) Response {
	if err != nil {
		return Response{
			Code:   500,
			Result: "",
			ErrMsg: err.Error(),
		}
	}
	return Response{
		Code:   200,
		Result: result,
		ErrMsg: "",
	}

}
