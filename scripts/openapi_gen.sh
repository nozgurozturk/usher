#!/bin/bash

function generate() {
  DIRCONFIG="api/openapi"
  DIRCONFIGADMIN="${DIRCONFIG}/admin.yaml"
  DIRCONFIGPUBLIC="${DIRCONFIG}/public.yaml"

  DIR="internal/infrastructure/delivery/rest"
  DIRADMIN="${DIR}/admin"
  DIRPUBLIC="${DIR}/public"

  [ -d "${DIRADMIN}" ] || mkdir -p "${DIRADMIN}"
  [ -d "${DIRPUBLIC}" ] || mkdir -p "${DIRPUBLIC}"


   oapi-codegen -generate chi-server -o "${DIRADMIN}/api.gen.go" -package "admin" ${DIRCONFIGADMIN}
   oapi-codegen -generate spec -o "${DIRADMIN}/spec.gen.go"  -package "admin" ${DIRCONFIGADMIN}
   oapi-codegen -generate types -o "${DIRADMIN}/types.gen.go"  -package "admin" ${DIRCONFIGADMIN}

   oapi-codegen -generate chi-server -o "${DIRPUBLIC}/api.gen.go" -package "public" ${DIRCONFIGPUBLIC}
   oapi-codegen -generate spec -o "${DIRPUBLIC}/spec.gen.go" -package "public" ${DIRCONFIGPUBLIC}
   oapi-codegen -generate types -o "${DIRPUBLIC}/types.gen.go" -package "public" ${DIRCONFIGPUBLIC}

}

generate