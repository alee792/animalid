package animalid

import (
	"log"
	"regexp"
	"strings"
	"testing"
)

func TestDelimitedName(t *testing.T) {
	name := GetDelimitedAnimalID("_")
	log.Print(name)
	if !strings.Contains(name, "_") {
		t.Fatalf("Generated name does not contain an underscore")
	}
	if strings.ContainsAny(name, "0123456789") {
		t.Fatalf("Generated name contains numbers")
	}
}

func TestTitleName(t *testing.T) {
	name := GetDelimitedAnimalID("")
	log.Print(name)
	if match, _ := regexp.Match("^[A-Za-z]+$", []byte(name)); !match {
		t.Fatalf("Generated name contains illegal characters")
	}
}
