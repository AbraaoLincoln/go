#!/bin/bash

# yq --input-format yaml --output-format props sample1.yaml

yamlFiles=$(ls yaml)

for file in $yamlFiles; do
    # echo $file
    env=$(echo $file | awk -F '.' '{print $1}')
    # echo $env
    yq --input-format yaml --output-format props "yaml/$file" > "properties/$env.properties"
done

go run main.go

propsFiles=$(ls properties)

for file in $propsFiles; do
    # echo $file
    env=$(echo $file | awk -F '.' '{print $1}')
    # echo $env
    yq --input-format props --output-format yaml "properties/$file" > "values/values-$env.yaml"
done