package lib1

import (
	"fmt"

	"mbbm-vas.com/go/wsdemo/lib2"
)

func Text() string {
  return "workspace demo"
}

func TextFromLib2() string {
  return fmt.Sprintf("lib2 says: %d", lib2.Number())
}