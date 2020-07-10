package scenariosdecoder

import "errors"

// Decoder Errors:
var ErrWrongFormatOfBackgroundTask error = errors.New("Wrong format of Background Task items")
var ErrWrongFormatOfForegroundTask error = errors.New("Wrong format of Foreground Task items")
