package assert

// Assertions provides assertion methods around the
// TestingT interface.
type Assertions struct {
	t        TestingT
	getTfunc func() TestingT
}

// New makes a new Assertions object for the specified TestingT.
func New(t TestingT) *Assertions {
	return &Assertions{
		t: t,
	}
}

func (a *Assertions) T() TestingT {
	if a.getTfunc != nil {
		return a.getTfunc()
	}

	return a.t
}

func (a *Assertions) SetTFunc(callback func() TestingT) {
	a.getTfunc = callback
}

//go:generate sh -c "cd ../_codegen && go build && cd - && ../_codegen/_codegen -output-package=assert -template=assertion_forward.go.tmpl -include-format-funcs"
