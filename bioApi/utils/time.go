package utils

//gettime now automatically
import (
	"time"
)

func GetTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
