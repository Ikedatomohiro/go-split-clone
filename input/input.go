package input

type Input struct {
	Option      string
	OptionValue int64
	Uint        string
	FileName    string
	Prefix      string
}

type Exist struct {
	Option bool
	Prefix bool
}
