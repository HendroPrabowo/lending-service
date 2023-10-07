package files

type Repository interface {
	Insert(entity File) (int, error)
	Get(queryParam FileQueryParam) (file File, err error)
}
