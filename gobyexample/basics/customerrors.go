package main
import (
	"error"
	"fmt"

)

type argError struct{
	arg int
	message string
}