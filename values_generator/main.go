package main

import (
	"bufio"
	"os"
)

func main() {
	propsDev := readProps("properties/dev.properties")
	propsHlg := readProps("properties/hlg.properties")
	propsPrd := readProps("properties/prd.properties")

	comunProps := getComunProps(propsDev, propsHlg, propsPrd)
	propsDev = removeComunProps(propsDev, comunProps)
	propsHlg = removeComunProps(propsHlg, comunProps)
	propsPrd = removeComunProps(propsPrd, comunProps)

	savePropsToFile("properties/comun.properties", comunProps)
	savePropsToFile("properties/dev.properties", propsDev)
	savePropsToFile("properties/hlg.properties", propsHlg)
	savePropsToFile("properties/prd.properties", propsPrd)

	// fmt.Println(comunProps)
	// fmt.Println(propsDev)
	// fmt.Println(propsHlg)
	// fmt.Println(propsPrd)
}

func getComunProps(propsSet1 map[string]bool, propsSet2 map[string]bool, propsSet3 map[string]bool) map[string]bool {
	comunProps12 := propsIntercection(propsSet1, propsSet2)
	comunProps123 := propsIntercection(comunProps12, propsSet3)

	return comunProps123
}

func readProps(fileName string) map[string]bool {
	props := make(map[string]bool)
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	propsScanner := bufio.NewScanner(file)

	for propsScanner.Scan() {
		prop := propsScanner.Text()
		props[prop] = true
		// fmt.Println(prop)
	}

	return props
}

func propsIntercection(propsSet1 map[string]bool, propsSet2 map[string]bool) map[string]bool {
	intersection := make(map[string]bool)

	for prop := range propsSet1 {
		if propsSet2[prop] {
			intersection[prop] = true
		}
	}

	return intersection
}

func savePropsToFile(fileName string, props map[string]bool) {
	var file *os.File
	var err error
	_, err = os.Stat(fileName)

	if os.IsExist(err) {
		file, err = os.Open(fileName)
	} else {
		file, err = os.Create(fileName)
	}

	if err != nil {
		panic(err)
	}

	for prop := range props {
		file.WriteString(prop + "\n")
	}

}

func removeComunProps(prospSet map[string]bool, comunPorpsSet map[string]bool) map[string]bool {
	for prop := range prospSet {
		if comunPorpsSet[prop] {
			delete(prospSet, prop)
		}
	}

	return prospSet
}
