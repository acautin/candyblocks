package game

import "errors"

// ErrGridAlreadyHasActiveCandy is an error thrown when trying to add a candy when the grid has an active one already
var ErrGridAlreadyHasActiveCandy = errors.New("The grid already has an active candy")
