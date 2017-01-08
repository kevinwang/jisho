package client

import (
	"fmt"
)

func (f *Form) String() string {
	if f.Word != "" && f.Reading != "" {
		return fmt.Sprintf("%s 【%s】", f.Word, f.Reading)
	} else if f.Word != "" {
		return f.Word
	}
	return f.Reading
}
