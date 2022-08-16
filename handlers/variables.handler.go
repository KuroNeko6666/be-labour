package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Validate
var validate = validator.New()

// Message
const ERROR = "error"
const SUCCESS = "success"

// Status

const BAD_REQUEST = http.StatusBadRequest
const OK = http.StatusOK
const NOT_FOUND = http.StatusNotFound
const SERVER_ERROR = http.StatusInternalServerError

// Error Data
const ERR_NOT_FOUND = "user not found"
const ERR_ENSURE_ID = "please ensure that id is an integer"
const ERR_EXIST = "email or username already exist"
const ERR_NOT_EXIST = "user does not exist"
const ERR_ACCESS = "you do not have access to do something this user"
const ERR_SERVER_ERROR = "internal server error"
const ERR_MATCH_PASS = "password does not match"

// Success Data

const SCS_DELETE = "Succesfully delete user"
