package sir

import (
	"log"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal("[ERROR]", err)
	}
}
