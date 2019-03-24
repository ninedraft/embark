package embark

import "io"

type PackageManager interface {
	InitPackage(log io.Writer) error
}
