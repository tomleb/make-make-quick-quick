#!/bin/sh

mkdir -p pkg/generated
cat > pkg/generated/zz_generated.go <<EOF
package generated

import (
	"log"
)

func Init() {
	log.Println("Calling generated.Init()")
}
EOF
