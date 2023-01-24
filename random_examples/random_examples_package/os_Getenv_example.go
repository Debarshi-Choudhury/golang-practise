package random_examples_package

import (
	"fmt"
	"os"
)

func PrintSomeEnvVariables(envVarNames ...string) {
	for _, envVarName := range envVarNames {
		envVar := os.Getenv(envVarName)
		fmt.Println(envVarName, "====>", envVar)
	}
}

/*
Call function like this:
exmpl.PrintSomeEnvVariables("GOPATH", "GOROOT")

Sample Output:
GOPATH ====> /home/debarshi/go
GOROOT ====> /usr/local/go
*/
