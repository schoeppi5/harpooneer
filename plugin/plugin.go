package plugin

import "github.com/schoeppi5/harpooneer/logging"

// Extension is a extension of the code-base available at compilation time
type Extension interface {
	Name() string
}

// NewExtension returns an implementation of HerpooneerExtension
type NewExtension func(logging.Logger) Extension
