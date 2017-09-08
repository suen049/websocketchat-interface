package websocketchat

type mysql interface {
	getInitSqlConnection() error
}
