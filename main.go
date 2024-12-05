package main

import (
	"flag"
	"fmt"
	"os"

	myAst "github.com/BelehovEgor/go-fuzz-targets-search-engine/pkg/ast"
	metrics "github.com/BelehovEgor/go-fuzz-targets-search-engine/pkg/metrics"
	rank "github.com/BelehovEgor/go-fuzz-targets-search-engine/pkg/ranking"
	"golang.org/x/tools/go/packages"
)

const (
	Complexity    string = "complexity"
	Volnerability string = "volnerability"
)

func main() {
	folder := flag.String("folder", "", "Your name")
	packageName := flag.String("package", "", "Your package")
	algorithm := flag.String("algorithm", Volnerability, "Ranking algorithm")
	top := flag.Uint("top", 5, "Count target funcs")

	flag.Parse()

	if *algorithm != Complexity && *algorithm != Volnerability {
		fmt.Println("Invalid algorithm")
		os.Exit(1)
	}

	var err error
	var funcs []metrics.Rankable
	if *packageName != "" {
		var pkg *packages.Package
		pkg, err = myAst.GetPackage(*folder, *packageName)
		if err == nil {
			funcs, err = getFuncs(*algorithm, pkg)
		}
	} else {
		var pkgs []*packages.Package
		pkgs, err = myAst.GetPackages(*folder)
		if err == nil {
			funcs, err = getAllFuncs(*algorithm, pkgs)
		}
	}

	if err != nil {
		fmt.Printf("Error while reading packages: %s", err)
		os.Exit(3)
	}

	priorities := rank.Prioritize(funcs, *top)
	for _, prioritet := range priorities {
		fmt.Printf("Package: %s, Name: %s, Priority: %d\n", prioritet.Package, prioritet.Name, prioritet.Priority)
	}
}

func getAllFuncs(algorithm string, pkgs []*packages.Package) ([]metrics.Rankable, error) {
	funcs := make([]metrics.Rankable, 0)
	for _, pkg := range pkgs {
		pkgFuncs, err := getFuncs(algorithm, pkg)
		if err != nil {
			return nil, err
		}

		funcs = append(funcs, pkgFuncs...)
	}

	return funcs, nil
}

func getFuncs(algorithm string, pkg *packages.Package) ([]metrics.Rankable, error) {
	var funcs []metrics.Rankable
	var err error

	switch algorithm {
	case Complexity:
		funcs, err = metrics.CalculateComplexitiesFromPackage(pkg)
	case Volnerability:
		funcs, err = metrics.CalculateVulnerabilities(pkg)
	}

	return funcs, err
}
