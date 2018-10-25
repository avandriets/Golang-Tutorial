package main

/*
#cgo CFLAGS: -I./indigo/include
#cgo LDFLAGS: -L./indigo/lib -Wl,-rpath=./indigo/lib -lindigo
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
	"log"
	"net/http"
	"unsafe"

	"github.com/gorilla/mux"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
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
		C.indigoFree(molecule)
	}()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/smiles-to-mol", smilesToMol).Methods("GET")

	if err := http.ListenAndServe(":9000", r); err != nil {
		log.Fatal(err)
	}
}
