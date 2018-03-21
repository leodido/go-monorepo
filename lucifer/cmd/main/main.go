package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/leodido/go-monorepo/amenadiel"
	"github.com/leodido/go-monorepo/lucifer"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	a := lucifer.Fights()
	fmt.Fprintf(os.Stdout, "%s fights %s also if he is his %s\n", lucifer.GetName(), a.Name, amenadiel.RelationWithLucifer())
}
