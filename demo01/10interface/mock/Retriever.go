package mock

type Retriever struct {
	Contents  string
	UserAgent string
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}
