package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	for true {
		currentTime := time.Now()

		resp, errr := http.Get("http://worldtimeapi.org/api/timezone/Europe/London.txt")
		if errr != nil {
			// handle err
		}
		defer resp.Body.Close()
		body, errrr := ioutil.ReadAll(resp.Body)
		if errrr != nil {
			//handle errrr
		}
		bs := string(body)

		myvalidString := "Current Server Time:" + "[" + currentTime.String() + "]" + "\n Current London Time:" + bs + "\n"
		err := WriteToFile("/tmp/data/output", myvalidString)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * 120)
	}
}

// WriteToFile will print any string of text to a file safely by
// checking for errors and syncing at the end.
func WriteToFile(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
