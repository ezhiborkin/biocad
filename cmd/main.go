package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mas"
	"mas/configs"
	"mas/pkg/er"
	"mas/pkg/handler"
	"mas/pkg/repository"
	"mas/pkg/service"
	"mas/pkg/unit"
	"os"
	"path/filepath"
	"sync"
	"time"

	_ "mas/docs"

	"github.com/grailbio/base/tsv"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// @title File Parser API
// @version 1.0
// @description API Server for parsing files and getting data from MongoDB
// @contact github: @ezhiborkin

// @host localhost:8080
// @BasePath /api

// @contact.name   Evgenii Zhiborkin
// @contact.url    https://t.me/zyltrcuj
// @contact.email  zhiborkin_ei@mail.ru

func main() {
	var mutex sync.Mutex
	if err := initConfig(); err != nil {
		log.Fatalf("error occured while initializing configs: %s", err.Error())
	}

	conf, err := configs.GetUniDocCred()
	if err != nil {
		return
	}

	err = license.SetMeteredKey(conf.Key)
	if err != nil {
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error occured while loading env variables: %s", err.Error())
	}

	directoryPath := os.Getenv("DIRECTORY_PATH_DOCKER")

	fileQueue := make([]string, 0)

	log.Println(directoryPath)
	client, err := repository.NewMongoDB(repository.Config{
		Host: os.Getenv("MONGO_HOST"),
		Port: os.Getenv("MONGO_PORT"),
	})
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(client)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	go watchDirectory(&fileQueue, directoryPath, client)

	go processFiles(&fileQueue, &mutex, client)

	srv := new(mas.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func watchDirectory(fileQueue *[]string, directoryPath string, client *mongo.Client) {
	for {
		err := filepath.WalkDir(directoryPath, func(path string, info os.DirEntry, err error) error {
			if err != nil {
				log.Println(err)
				return err
			}
			if !info.IsDir() && filepath.Ext(path) == ".tsv" {
				var result bson.M
				filePath := filepath.Base(path)
				filter := bson.M{"filepath": filePath}
				err := client.Database(viper.GetString("db.dbname")).Collection(viper.GetString("collectionName.processedFiles")).FindOne(context.Background(), filter).Decode(&result)
				if err == nil {
					// log.Println("File already processed:", path)
				} else if err == mongo.ErrNoDocuments {
					*fileQueue = append(*fileQueue, path)
					_, err := client.Database(viper.GetString("db.dbname")).Collection(viper.GetString("collectionName.processedFiles")).InsertOne(context.Background(), filter)
					if err != nil {
						log.Println(err)
						return err
					}
					log.Println("Added to queue:", path)
				} else {
					return er.ErrorParsingFile(path, client, err)
				}
			}
			return nil
		})

		if err != nil {
			log.Println(err)
		}

		time.Sleep(10 * time.Second)
	}
}

func processFiles(fileQueue *[]string, mutex *sync.Mutex, client *mongo.Client) {
	for {
		if len(*fileQueue) != 0 {
			mutex.Lock()
			filePath := (*fileQueue)[0]
			*fileQueue = (*fileQueue)[1:]
			processFile(filePath, client)
			mutex.Unlock()
		} else {
			time.Sleep(5 * time.Second)
			log.Println("No files to process")
			continue
		}
	}

}

func processFile(filePath string, client *mongo.Client) error {
	file, err := os.Open(filePath)
	if err != nil {
		return er.ErrorParsingFile(filePath, client, err)
	}
	defer file.Close()

	reader := tsv.NewReader(file)

	strings, err := reader.ReadAll()
	if err != nil {
		return er.ErrorParsingFile(filePath, client, err)
	}

	for _, record := range strings[1:] {
		unit := unit.Unit{
			Number:    record[0],
			Mqtt:      record[1],
			Invid:     record[2],
			UnitGuid:  record[3],
			MessageID: record[4],
			Text:      record[5],
			Context:   record[6],
			Class:     record[7],
			Level:     record[8],
			Area:      record[9],
			Addr:      record[10],
			Block:     record[11],
			Type_:     record[12],
			Bit:       record[13],
			InvertBit: record[14],
		}

		var doc *document.Document

		if fileExists(fmt.Sprintf("%s/%s.docx", os.Getenv("PROCESSED_DIRECTORY_PATH_DOCKER"), unit.UnitGuid)) {
			doc, err = document.Open(fmt.Sprintf("%s/%s.docx", os.Getenv("PROCESSED_DIRECTORY_PATH_DOCKER"), unit.UnitGuid))
			if err != nil {
				log.Println(err)
			}
		} else if !fileExists(fmt.Sprintf("%s/%s.docx", os.Getenv("PROCESSED_DIRECTORY_PATH_DOCKER"), unit.UnitGuid)) {
			doc = document.New()
		} else {
			log.Println(err)
		}
		kek := fmt.Sprintf("Number: %s, Mqtt: %s, Invid: %s, UnitGuid: %s, MessageID: %s, Text: %s, Context: %s, Class: %s, Level: %s, Area: %s, Addr: %s, Block: %s, Type_: %s, Bit: %s, InvertBit: %s", unit.Number, unit.Mqtt, unit.Invid, unit.UnitGuid, unit.MessageID, unit.Text, unit.Context, unit.Class, unit.Level, unit.Area, unit.Addr, unit.Block, unit.Type_, unit.Bit, unit.InvertBit)
		para := doc.AddParagraph()
		run := para.AddRun()
		run.AddText(kek)
		doc.SaveToFile(fmt.Sprintf("%s/%s.docx", os.Getenv("PROCESSED_DIRECTORY_PATH_DOCKER"), unit.UnitGuid))

		_, err := client.Database(viper.GetString("db.dbname")).Collection(viper.GetString("collectionName.processedData")).InsertOne(context.Background(), unit)
		if err != nil {
			log.Println(err)
			return err
		}

		fullFilePath := filepath.Join(fmt.Sprintf("%s/", os.Getenv("PROCESSED_DIRECTORY_PATH_DOCKER")), unit.UnitGuid)

		outputFile, err := os.OpenFile(fullFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return er.ErrorParsingFile(filePath, client, err)
		}
		defer outputFile.Close()

		encoder := json.NewEncoder(outputFile)
		err = encoder.Encode(unit)
		if err != nil {
			return er.ErrorParsingFile(filePath, client, err)
		}

	}

	return nil
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
