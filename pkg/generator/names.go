package generator

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	adjectives = []string{"amazing", "thankful", "joyful", "exciting", "dark", "fast", "soft", "sluggish", "slippery"}
	nouns      = []string{"zebra", "lion", "gopher", "fan", "bed", "chair", "mouse", "dongle", "dog", "cat"}
)

func pickRandom(l []string) string {
	rand.Seed(time.Now().UnixNano())
	return l[rand.Intn(len(l))]
}

func GenerateCodeName() string {
	return fmt.Sprintf("%s-%s", pickRandom(adjectives), pickRandom(nouns))
}
