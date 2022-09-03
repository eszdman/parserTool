package Errors

var Err error

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
