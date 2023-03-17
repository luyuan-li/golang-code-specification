package resp

type (
	UserResp struct {
		ID     uint   `json:"id"`
		Name   string `json:"name"`
		Age    int64  `json:"age"`
		Mobile string `json:"mobile"`
	}

	UserListResp []UserResp
)
