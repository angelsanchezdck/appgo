package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/csepulveda/http_to_s3_test/common"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("ListeningAddr", "0.0.0.0:8080")
	viper.SetDefault("BucketName", "mybucketr2")
	handleRequests(viper.GetString("ListeningAddr"), viper.GetString("BucketName"))
}

func handleRequests(ListeningAddr, BucketName string) {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/putfile", putfile(BucketName)).Methods("POST")

	server := &http.Server{
		Addr:              ListeningAddr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           r,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func putfile(b string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		sess, err := createSession()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "err: unable to create session\n")
			return
		}

		content, err := common.ParseBody(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "err: unable to parse content\n")
			return
		}

		content.Bucket = b
		result, err := common.CreateFile(content, sess)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "err: unable to create file\n")
			return
		}

		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "err: unable to parse result\n")
			return
		}
	}
}

func createSession() (*session.Session, error) {
	viper.AutomaticEnv()
	viper.SetDefault("AWSRegion", "us-west-2")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(viper.GetString("AWSRegion"))},
	)
	if err != nil {
		fmt.Printf("Unable create session, %v", err)
		return nil, err
	}
	return sess, nil
}
