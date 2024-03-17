package database

import user "gostart/internal/models"

type ModelConstraint interface {
	user.User
}
