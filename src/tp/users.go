package tp

type CreateUpdateUserReqBody struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type ReadUsersReqQuery struct {
	ID          *int   `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Gender      string `json:"gender"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`
}

type UserItem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Gender      string `json:"gender"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`
}
