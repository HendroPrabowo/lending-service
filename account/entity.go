package account

type Account struct {
	tableName struct{} `pg:"account"`
	Id        int
	Username  string
	Password  string
	Name      string
	Email     string
	CreatedAt string
}
