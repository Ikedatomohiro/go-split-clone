package input

type Input struct {
	Option       string
	OptionValue  int64
	SuffixLength int
	FileName     string
	Prefix       string
}

type ArgPosition struct {
	Option   int
	AOption  int
	Prefix   int
	FileName int
}
