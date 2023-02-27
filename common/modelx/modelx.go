package modelx

type Modelx interface {
	tableName() string
	Rows() string
}
