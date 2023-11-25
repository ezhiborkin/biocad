package er

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"github.com/unidoc/unioffice/document"
	"go.mongodb.org/mongo-driver/mongo"
)

// Error parse info
// @Description Error parsing file info
// @Description with filename, error and time
type ErrorOpenFile struct {
	FileName string    `json:"filename"` // File name
	Err      error     `json:"error"`    // Error
	Time     time.Time `json:"time"`     // Time
}

// Error response
// @Description Error response
type errorResponse struct {
	Message string `json:"message"` // Error message
}

func ErrorParsingFile(filePath string, client *mongo.Client, err error) error {
	log.Println(err)
	client.Database(viper.GetString("db.dbname")).Collection(viper.GetString("collectionName.errors")).InsertOne(context.Background(), ErrorOpenFile{
		FileName: filepath.Base(filePath),
		Err:      err,
		Time:     time.Now(),
	})
	doc := document.New()
	kek := fmt.Sprintf("Error: %s, File: %s, Time: %s", err.Error(), filepath.Base(filePath), time.Now().String())
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText(kek)
	doc.SaveToFile(fmt.Sprintf("%s/%s.docx", os.Getenv("ERRORS_DIRECTORY_PATH_DOCKER"), fmt.Sprintf("%s_%s_%s", filepath.Base(filePath), time.Now().String(), "error")))

	return err
}
