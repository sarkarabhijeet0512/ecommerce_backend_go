package utils

import (
	"fmt"
	"mime/multipart"
	"reflect"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	AccessKeyEnv       = "AWS_ACCESS_KEY"
	SecretAccessKey    = "AWS_SECRET_KEY"
	SESAccessKeyEnv    = "SES_ACCESS_KEY"
	SESSecretAccessKey = "SES_SECRET_KEY"
	Region             = "AWS_REGION"
	BucketName         = "AWS_BUCKET"
	ProductRole        = "product_management"
)

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Generic is a generic function to set common fields of any struct
func SetGenericFieldValue(i interface{}) {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		panic("Generic: input is not a pointer to a struct")
	}

	// Get the actual struct value (dereference the pointer)
	v = v.Elem()

	// Set common fields like IsActive, CreatedAt, UpdatedAt
	setField(v, "IsActive", true)
	setField(v, "CreatedAt", time.Now())
	setField(v, "UpdatedAt", time.Now())
}

// setField sets the value of a field in a struct using reflection
func setField(v reflect.Value, fieldName string, value interface{}) {
	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return // Field doesn't exist in the struct
	}
	if !field.CanSet() {
		return // Field is unexported or read-only
	}
	fieldType := field.Type()
	val := reflect.ValueOf(value)
	if val.Type().ConvertibleTo(fieldType) {
		field.Set(val.Convert(fieldType))
	} else {
		panic(fmt.Sprintf("Generic: value of type %T cannot be assigned to field %s of type %s", value, fieldName, fieldType))
	}
}

func FileProcessing(file multipart.File, header *multipart.FileHeader, productID int) (filename, contentType string) {
	fileNameParts := strings.Split(header.Filename, ".")
	ext := ""
	if len(fileNameParts) == 2 {
		ext = fileNameParts[1]
	}
	filename = fmt.Sprint(productID, ".", ext)
	contentType = "image/jpeg"
	if ext == "pdf" {
		contentType = "application/pdf"
	}
	return
}
