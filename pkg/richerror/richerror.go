package richerror

type Kind int

type Op string

const (
	KindInvalid Kind = iota + 1
	KindForbidden
	KindNotFound
	KindUnExepted
)

type RichError struct {
	operation    Op
	wrappedError error
	message      string
	kind         Kind
	meta         map[string]interface{}
}

func New(pm Op) RichError {
	return RichError{
		operation: pm,
	}
}

func (r RichError) Error() string {
	return r.message
}

func (r RichError) WithMassage(massage string) RichError {
	r.message = massage
	return r
}

func (r RichError) WithError(err error) RichError {
	r.wrappedError = err

	return r
}

func (r RichError) WithKind(kind Kind) RichError {
	r.kind = kind

	return r
}

func (r RichError) WithMeta(meta map[string]interface{}) RichError {
	r.meta = meta

	return r
}

func (r RichError) Kind() Kind {
	if r.kind != 0 {
		return r.kind
	}
	re, ok := r.wrappedError.(RichError)
	if !ok {
		return 0
	}
	return re.kind
}

func (r RichError) Massage() string {
	if r.message != "" {
		return r.message
	}
	re, ok := r.wrappedError.(RichError)
	if !ok {
		return r.wrappedError.Error()
	}
	return re.message
}
