package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func PodcastTypeValidator(s string) (string, error) {
	var IsValid = regexp.MustCompile(`^(GUEST_LECTURE|CAREER_PATH)$`).MatchString
	if !IsValid(s) {
		return s, errors.New("invalid type")
	}
	return s, nil
}

func NameValidator(s string) (string, error) {
	var IsLetter = regexp.MustCompile(`^[a-zA-Z., ]+$`).MatchString
	if !IsLetter(s) {
		return s, fmt.Errorf("name should only contain characters")
	}
	name := cases.Lower(language.Und).String(s)
	name = cases.Title(language.Und).String(name)
	return name, nil
}

func NumericValidator(s string) (string, error) {
	var IsLetter = regexp.MustCompile(`^[0-9]*$`).MatchString
	if !IsLetter(s) {
		return s, fmt.Errorf("RollNo should only contain numbers")
	}
	name := cases.Lower(language.Und).String(s)
	name = cases.Title(language.Und).String(name)
	return name, nil
}

func EmailValidator(s string) (string, error) {
	var IsValid = regexp.MustCompile(`^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$`).MatchString
	if !IsValid(s) {
		return s, errors.New("invalid Email format")
	}
	return s, nil
}

func PasswordHasher(s string) string {
	hasher := sha1.New()
	hasher.Write([]byte(s))
	passwordHash := hex.EncodeToString(hasher.Sum(nil))
	return passwordHash
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func URLValidator(s string) (string, error) {
	var IsValid = regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`).MatchString
	if !IsValid(s) {
		return s, errors.New("invalid URL format")
	}
	return s, nil
}
