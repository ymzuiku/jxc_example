package user

type CheckSimCodeBody struct {
	Phone string `json:"phone" validate:"required,min=3,max=32"`
	Code  string `json:"code" validate:"required, min=4,max=6"`
}

type regiestSendSimBody struct {
	Phone string `json:"phone" validate:"required,min=3,max=32"`
}

type deleteBody struct {
	Phone string `json:"phone" validate:"required,min=6,max=32"`
}
