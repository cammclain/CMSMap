package superficial

import (
	"math/rand"
	"time"
)

func Select_ascii_art() string {
	//

	intro_messages := []string{"Hello, World!\n\n		- sooni dev", "Welcome to the world of exploitation", "Exp: Giving blue teamers PTSD since 2024", "Exp: A modern exploit ??? framework"}

	// seeding the random number generator
	rand.Seed(time.Now().UnixNano())
	// get the number of ascii art
	random_ascii_art := intro_messages[rand.Intn(len(intro_messages))]

	return random_ascii_art
}
