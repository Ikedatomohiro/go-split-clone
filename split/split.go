package split

import (
	"errors"
	"os"

	in "split-clone/input"
	op "split-clone/option"
)

func Split(in in.Input, file *os.File) error {
	switch in.Option {
	case "l":
		err := op.Lines(in, file)
		if err != nil {
			return err
		}
	case "n":
		err := op.Numbers(in, file)
		if err != nil {
			return err
		}
	case "b":
		err := op.Bytes(in, file)
		if err != nil {
			return err
		}
	default:
		return errors.New("Invalid option")
	}
	return nil
}
