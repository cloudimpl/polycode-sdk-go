package polycode

import "github.com/cloudimpl/polycode-sdk-go/errors"

var ErrNotFound = errors.DefineError("polycode.sdk", 0, "not found")
var ErrAlreadyExist = errors.DefineError("polycode.sdk", 1, "already exist")
var ErrConflict = errors.DefineError("polycode.sdk", 2, "conflict")
