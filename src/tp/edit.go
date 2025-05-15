package tp

type EditReqBody struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}
