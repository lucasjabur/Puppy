package puppy

import (
	"fmt"

	dog "github.com/lucasjabur/Dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From11() {
	fmt.Println("I'm from version v1.1.0!")
}

func From12() {
	fmt.Println("I'm from version v1.2.0!")
}

func From13() {
	fmt.Println("I'm from version v1.3.0!")
}

// To see the tags on the repo's website on GitHub, you must use the following command:
// 'git push origin --tags'
// This will push all the tags created to the repository info.
// Now, you can see that is has 3 tags ('v1.1.0', 'v1.2.0' and 'v1.3.0') on the repo's page.
