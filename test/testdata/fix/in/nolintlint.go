//golangcitest:args -Enolintlint -Elll
//golangcitest:expected_linter nolintlint
//golangcitest:config linters-settings.nolintlint.allow-leading-space=false
package p

import "fmt"

func nolintlint() {
	fmt.Println() // nolint:bob // leading space should be dropped
	fmt.Println() //  nolint:bob // leading spaces should be dropped

	// note that the next lines will retain trailing whitespace when fixed
	fmt.Println() //nolint // nolint should be dropped
	fmt.Println() //nolint:lll // nolint should be dropped

	fmt.Println() //nolint:alice,lll // we don't drop individual linters from lists
}
