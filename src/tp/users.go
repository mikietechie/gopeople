package tp

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
	Other  Gender = "other"
)

type CreateUserReqBody struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type UpdateUserReqBody struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Gender      Gender `json:"gender"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`
}

type ReadUsersReqQuery struct {
	ID          *int   `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Gender      Gender `json:"gender"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`
}

type UserItem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Gender      Gender `json:"gender"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`
}
