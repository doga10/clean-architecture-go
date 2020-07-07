package account

type AddAccountParams struct {
	Name      string    `bson:"name" json:"name"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"password"`
}

type AddAccount interface {
	Add(accountData *AddAccountParams) (interface{}, error)
}