package main

import (
	"flag"
	"fmt"
	"os"

	myAst "github.com/BelehovEgor/go-fuzz-targets-search-engine/pkg/ast"
	metrics "github.com/BelehovEgor/go-fuzz-targets-search-engine/pkg/metrics"
	rank "github.com/BelehovEgor/go-fuzz-targets-search-engine/pkg/ranking"
)

const (
	Complexity    string = "complexity"
	Volnerability string = "volnerability"
)

func main() {
	folder := flag.String("folder", "", "Your name")
	packageName := flag.String("package", "", "Your package")
	algorithm := flag.String("algorithm", Volnerability, "Ranking algorithm")

	flag.Parse()

	if *algorithm != Complexity && *algorithm != Volnerability {
		fmt.Println("Invalid algorithm")
		os.Exit(1)
	}

	if *packageName == "" {
		fmt.Println("Empty package name")
		os.Exit(1)
	}

	pkg, err := myAst.GetPackage(*folder, *packageName)
	if err != nil {
		fmt.Printf("Error while reading package: %s", err)
		os.Exit(2)
	}

	var funcs []metrics.Rankable

	switch *algorithm {
	case Complexity:
		funcs, err = metrics.CalculateComplexitiesFromPackage(pkg)
	case Volnerability:
		funcs, err = metrics.CalculateVulnerabilities(pkg)
	}
	if err != nil {
		fmt.Printf("Error while calculating metrics: %s", err)
		os.Exit(3)
	}

	priorities := rank.Prioritize(funcs, 5)
	for _, prioritet := range priorities {
		fmt.Printf("Name: %s Priority: %d\n", prioritet.Name, prioritet.Priority)
	}
}
