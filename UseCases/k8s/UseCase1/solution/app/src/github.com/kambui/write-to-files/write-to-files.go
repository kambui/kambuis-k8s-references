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
		log.Output(1, "Getting London time from worldtime.org")
		resp, errr := http.Get("http://worldtimeapi.org/api/timezone/Europe/London.txt")
		if errr != nil {
			// handle err
			log.Fatal(errr)
		}
		defer resp.Body.Close()
		log.Output(1, "Reading body data from worldtime.org")
		body, errrr := ioutil.ReadAll(resp.Body)
		if errrr != nil {
			//handle errrr
			log.Fatal(errr)
		}
		bs := string(body)
		log.Output(2, "Creating String Output")
		myvalidString := "Current Server Time:" + "[" + currentTime.String() + "]" + "\n Current London Time:\n" + bs + "\n"
		log.Output(3, "Writing to file")
		err := WriteToFile("/tmp/data/output", myvalidString)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * 60)
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
