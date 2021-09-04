package main

import (
    "os"
    "log"
    "errors"
    "net/http"

    "github.com/pedrazzani/golang-file2api/api/handle"
    "github.com/pedrazzani/golang-file2api/usercase/product"
    "github.com/pedrazzani/golang-file2api/infrastructure/repository"
)

func check(e error) {
    if e != nil {
        os.Exit(1)
    }
}

func handleParams(pos int) (string, error) {
    if len(os.Args) < 2 {
        return "", errors.New("Invalid options")
    }
    return os.Args[pos], nil
}

func main() {

    txitensPath, err := handleParams(1)

    if err != nil {
        log.Fatal(err.Error())
    }

    fileRepository := repository.NewFileRepository(txitensPath);
    productService := product.NewService(fileRepository);
    handle.ProductHandles(productService)

    http.ListenAndServe(":8087", nil)
}