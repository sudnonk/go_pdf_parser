package main

type Xref struct {
	Offset int
	Number int
	Str    string
}

func (x *Xref) IsUsing() bool {
	return x.Str == "n"
}
