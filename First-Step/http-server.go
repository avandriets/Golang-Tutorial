package main

/*
#cgo CFLAGS: -I ./indigo/include
#cgo LDFLAGS: -L ./indigo/lib -lindigo
#include <indigo.h>

#include <stdio.h>
#include <stdlib.h>

static void myprint(char* s) {
  printf("%s\n", s);
}
*/
import "C"
import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"unsafe"

	"github.com/gorilla/mux"
)

type SmilesResponse struct {
	smiles string
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html><html><head><title>Hello World</title></head><body>Hello World!</body></html>`,
	)
}

func smilesToMol(w http.ResponseWriter, r *http.Request) {
	smilesString := r.FormValue("smiles")

	smiles := C.CString(smilesString)
	molecule := C.indigoLoadSmartsFromString(smiles)
	molString := C.indigoMolfile(molecule)

	cstr := C.GoString(molString)

	respondWithJson(w, http.StatusOK, map[string]string{"mol": cstr})

	defer func() {
		C.free(unsafe.Pointer(smiles))
	}()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/smiles-to-mol", smilesToMol).Methods("GET")

	if err := http.ListenAndServe(":9000", r); err != nil {
		log.Fatal(err)
	}
}
