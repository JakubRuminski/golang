package print_utils

import "log"

var logger_ON bool = true

var debugger_ON bool = true

var errors_ON bool = false

func LOGGER(message string) {
	if logger_ON {
		log.Println("[LOG]: " + message)
	}
}

func DEBUGGER(message string) {
	if debugger_ON {
		log.Println("[DEBUG]: " + message)
	}
}

func ERROR(message string) {
	if errors_ON {
		log.Println("[ERROR]: " + message)
	}
}
