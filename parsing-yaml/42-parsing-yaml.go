package main

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Deployment struct {
	// Name name of the function
	Metadata struct {
		Labels struct {
			App string `yaml:"app"`
		}
		Name string `yaml:"name"`
	}
	Spec Spec
}

type Spec struct {
	Replicas int                          `yaml:"replicas"`
	Selector map[string]map[string]string `yaml:"selector"`
}

func main() {
	bytesOut, err := os.ReadFile("../freeCodeCamp/data/nginx.yaml")
	if err != nil {
		panic(err)
	}
	dep := Deployment{}
	if err := yaml.Unmarshal(bytesOut, &dep); err != nil {
		panic(err)
	}
	fmt.Printf("App name: %s\n", dep.Metadata.Name)
	fmt.Printf("Replica: %d\n", dep.Spec.Replicas)
	fmt.Println(dep.Spec.Selector)
	fmt.Println(dep.Spec.Selector["matchLabels"]["app"])

}
