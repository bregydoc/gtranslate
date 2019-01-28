package gtranslate

import "errors"

var errBadNetwork = errors.New("bad network, please check your internet connection")
var errBadRequest = errors.New("bad request, request on google translate api isn't working")
