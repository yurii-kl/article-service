package errors

type DomainError string

func (e DomainError) Error() string {
	return string(e)
}

const (
	ErrArticleNotFound DomainError = "article not found"
)
