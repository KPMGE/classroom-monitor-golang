package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type EnvObject struct {
	CourseId string
}

func GetEnvObject() *EnvObject {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	courseId := os.Getenv("COURSE_ID")

	return &EnvObject{
		CourseId: courseId,
	}
}
