package account

type AddAccountParams struct {
	Name     string `bson:"name" json:"name" validate:"required"`
	Email    string `bson:"email" json:"email" validate:"required,email"`
	Password string `bson:"password" json:"password" validate:"required"`
}

type AddAccount interface {
	Add(accountData *AddAccountParams) (interface{}, error)
}
