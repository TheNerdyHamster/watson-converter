package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type Json struct {
	Array [][]interface{}
}

func main() {

	if runtime.GOOS == "windows" {
		fmt.Println("🪟 Windows detected,  this platform is currently not supported")
		return
	}

	fileptr := flag.String("file", "", "file to watson file to parse")
	flag.Parse()

	if len(*fileptr) == 0 {
		fmt.Println("Please specify what file you want to parse with -file $file")
		return
	}

	data := parseFile(fileptr)

	for _, v := range data.Array {
		startTime := time.Unix(int64(v[0].(float64)), 0).Format("2006-01-02 15:04:05")
		endTime := time.Unix(int64(v[1].(float64)), 0).Format("2006-01-02 15:04:05")

		execute(startTime, endTime, v[2].(string))
	}

	fmt.Println("❤️  Thanks for using Watson-converter, this is a totally useless cli application.")
}

func execute(startTime string, endTime string, project string) {
	out, err := exec.Command("watson", "add", "-f", startTime, "-t", endTime, project).Output()

	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	fmt.Println(string(out))
}

func parseFile(path *string) *Json {

	file, err := os.Open(*path)
	if err != nil {
		log.Fatalf("Could not read file %v", err)
		return nil
	}

	defer file.Close()

	data, _ := ioutil.ReadAll(file)

	var res Json
	err = json.Unmarshal([]byte(data), &res.Array)
	if err != nil {
		log.Fatalf("Could not parse file %v", err)
		return nil
	}

	return &res
}
