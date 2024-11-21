/*
Copyright © 2024 Evan Colwell ercolwell@gmail.com
*/
package lib

import (
	"log"
	"time"
)

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
