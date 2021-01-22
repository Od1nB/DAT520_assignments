package cipher

import (
	"io"
)

/*
Task 3: Rot 13

This task is taken from http://tour.golang.org.

A common pattern is an io.Reader that wraps another io.Reader, modifying the
stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of
compressed data) and returns a *gzip.Reader that also implements io.Reader (a
stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader,
modifying the stream by applying the rot13 substitution cipher to all
alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing
its Read method.
*/

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	ind, errs := r.r.Read(p) 
	if errs != nil {
		return 0, errs
	}

	for i := 0; i < ind; i++ {
		val := p[i]
		//ascii value of alphabetical signs 65-90 for capital and 97-122 for lower case
		if val >= 65 && val <= 90 {
			val += 13
			if val > 90 {
				val -= 26
			}
		}else if val >= 97  && val <= 122{
			val += 13
			if val > 122 {
				val -= 26
			}
		}
		p[i] = val
	}
	return ind,nil
}