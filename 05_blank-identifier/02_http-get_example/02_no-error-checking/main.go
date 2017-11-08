package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("http://www.bridge.ee")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}

// pmst sisu on vist selles, et kuna nii Get kui ka ReadAll funktsioonid
// tahavad kahte parameetrit, aga mul on anda ainult yks, siis ma kasutan
// siinkohal blank identifierit _ mis lubab mul yhe parameetri 2ra j2tta
