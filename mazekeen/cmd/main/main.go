package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/leodido/go-monorepo/lucifer"
	uuid "github.com/satori/go.uuid"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Fprintf(os.Stdout, "Mazekeen works for %s\n", lucifer.GetName())
	fmt.Fprintf(os.Stdout, "And always will, regardless the UUID: %s.\n", uuid.NewV4())
}
