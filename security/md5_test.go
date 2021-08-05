package security

import (
	"log"
	"testing"
)

func Test_MD5(t *testing.T) {
	result := MD5("hello world")
	log.Println(result)
}
