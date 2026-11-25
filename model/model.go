package model

type ErrReg struct {
	Login    string
	ErrLogin string
	ErrPass  string
}

type Registration struct {
	ID    int
	Login string
	Pass  string
}

type ErrSign struct {
	Login    string
	ErrLogin string
	ErrPass  string
}
