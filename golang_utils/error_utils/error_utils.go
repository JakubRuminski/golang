package error_utils

import (
	"github.com/jakubruminski/golang/golang_webapp_utils/print_utils"
	"runtime/debug"
)

func CheckErrorButDoNotStop(errorGiven, outcomeExpected error) {
	if errorGiven != outcomeExpected {
		print_utils.LOGGER("---------------------------------")
		debug.PrintStack()
		print_utils.LOGGER("---------------------------------")
	}
}

func CheckErrorAndPanic(errorGiven, outcomeExpected error) {
	if errorGiven != outcomeExpected {
		panic(errorGiven)
	}
}
