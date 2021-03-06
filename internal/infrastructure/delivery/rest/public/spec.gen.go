// Package public provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package public

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xYzY7bNhB+FYPtUV1t0ptubbNogxZIsMmeAh+40thmliIVkvJiYejdCw4pWT+kLW9t",
	"YN2brSHn55vhN0PuSC7LSgoQRpNsR3S+gZLizz82kD/dbUGYL0CNvocfNWhjJZWSFSjDANflshb42bxU",
	"QDLChIE1KNIkZAXU1Motm0oVFU8hiRXBj5opKEj2zetfJu06+fgdcmMVTDzUlRQapi4qKCkTTKxnmNuv",
	"jZqUtXkNGmA9/fihJ9RGWUNNQioFK1AgcqdlqPPVME7crzWooAcjEFpXEx9Pt/MwJjH8XwkKK64BK1b0",
	"YNoDdwCwO6WkiqMFVhyMvASt6Rpm5A9V7DcEvbCuTq0XoHPFKsOkCPoAovhADfqwkqqkhmSkoAZ+MawE",
	"kszOoqAlBAXaUGVOMRHKB2pPBsEkhNMXWdssoQmCscSh6SeIcv5pRbJvO/KzghXJyE/pnjdTT5qpQ7RJ",
	"xpBymdMWz0P7/6KcT8LpNk8dXbauDqiPGSj1+ZyecwR6awNeBgD2H6hS9MX+x8gnpRgvEcitOT0I91CQ",
	"X9yGqelRIL5sOv0h3+/l8wmuSlWACpORtk3rhBComeu/M9paCAWB2iZR0C1lnD7yfiiPUnKg4ng7F3X5",
	"GAu1kpq1J2DcGniEqh3Mx7q1fEb25cEoZ04YnXddEH5r0kMkjGIejipaDko+z0+5rbSZGUe9IRe/svwJ",
	"AqmGlv3PR21DyWm0H+XwaUwn8GMAj4hb2p+HOSfQdvZjax90oM4wKge81xLy8MGrPy+cMUvxMeTVYR6M",
	"zNVjuGEdsuTrONA7rNbT9TnPJ4fLVohYSYSUGUuE5HP9yFm++O3zR5KQLSiN5U7e3dwiv1cgaMVIRn69",
	"ub15TxJSUbNB82nuB2PEVrrbwmC86kZngoqUa6LFUKDcVeN3Wbw4whSmPbpVxZmr/vS7dofQxXcs+vE1",
	"phnm0Kga8IMDFaN5f3t7AfM+a2h/CM2nvwl+W9Gam7NZHo7eAbO4wF0I6FrbcsZK2d9yyNIKUzzGCMwa",
	"Aom9B1MroRecabOQq0XlisjvGmf7TzB3reRioI8mxSuA3GPSBzzd+RtWcxT6PuRRxPG8KlqCAaWxC458",
	"s4sWeKFj9r893G1rynq3veHZSXrQjJl5eekU/28y7PgzTp7jM4b7Ft3EtnDTb5BZ949G0wI4Z54vxNzT",
	"V7nGM/h/KKxZbTP25BbspFdUe3ZmOZnN3aYAtTx4wcUO+nDmefNg9zBOd+5h7Dh/1y5NQXCP0bZdE2Vt",
	"/zL3Vkh7MIVfZSpT44b6WSldtIsjqf3aiY/z8hvM5Ph+cw0Jxfuv2rZA14qTjGyMqXSWpruN1Mbi3aTb",
	"d6m96iRkSxWzLRYRbOUu7T4ovJZzK7KuLDub7dNIy75N0n3pBuzet3/ci22zbP4NAAD//18XwzqpGgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
