package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sync"
	"bufio"
    "os"
	"io"
	"time"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readFileParallel(fileName string, searchString string, wg *sync.WaitGroup){
	defer wg.Done()

	file, err := os.Open(fileName)
	defer file.Close()
	check(err)

	// scanner := bufio.NewScanner(file)

	// for scanner.Scan() {
    //     scanner.Text()
    // }

	reader := bufio.NewReader(file)
	//input, _ := reader.ReadString('\n')
	//var count = 0
	var found bool = false
	for {
		//line, _, err := reader.ReadLine()
		//str := string(line)
		line, err := reader.ReadString('\n')
		if strings.Contains(line, searchString) {
			//fmt.Println(line)
			found = true
		}
		if err == io.EOF {
			break
		}
		//count++
	}
	if found {
		fmt.Println(fileName)
	}
	//fmt.Printf(" %d\n", count)
}

func readFile(fileName string){

	file, err := os.Open(fileName)
	defer file.Close()
	check(err)

	reader := bufio.NewReader(file)

	for {
		_, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
	}
}

func main() {

	currentTime := time.Now()
	fmt.Println("The time is", currentTime)

	var wg sync.WaitGroup

	subDirToSkip := ".git"

	// err_parent := filepath.Walk("C:\\projetos\\tables_bigquerys\\path_dados", func(path string, info fs.FileInfo, err error) error {
	// 	if err != nil {
	// 		fmt.Printf("Error %q: %v\n", path, err)
	// 		return err
	// 	}
	// 	if ! info.IsDir() {
	// 		fileName := filepath.Base(path)
	// 		fileNameStr := strings.Replace(fileName, ".txt", "", -1)
	// 		fmt.Printf("Table name: %q\n", fileNameStr)

	// 		err := filepath.Walk("C:\\projetos\\repositorios\\repos_google\\", func(pathInside string, infoInside fs.FileInfo, errInside error) error {
	// 			if errInside != nil {
	// 				fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", pathInside, errInside)
	// 				return errInside
	// 			}
	// 			if infoInside.IsDir() && infoInside.Name() == subDirToSkip {
	// 				//fmt.Printf("skipping a dir without errors: %+v \n", infoInside.Name())
	// 				return filepath.SkipDir
	// 			}
	// 			if ! infoInside.IsDir() {
	// 				//fmt.Printf("File repo: %q\n", pathInside)
	// 				wg.Add(1)
	// 				go readFileParallel(pathInside, fileNameStr, &wg)
	// 				//readFile(path)
	// 			}
	// 			return nil
	// 		})

	// 		if err != nil {
	// 			fmt.Printf("error walking the path: %v\n", err)
	// 			return err
	// 		}
	// 	}
	// 	return nil
	// })

	err := filepath.WalkDir("C:\\projetos\\repositorios\\repos_google\\", func(pathInside string, infoInside fs.DirEntry, errInside error) error {
		if errInside != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", pathInside, errInside)
			return errInside
		}
		if infoInside.IsDir() && infoInside.Name() == subDirToSkip && infoInside.Name() == ".idea" {
			//fmt.Printf("skipping a dir without errors: %+v \n", infoInside.Name())
			return filepath.SkipDir
		}
		if ! infoInside.IsDir() {
			ext := filepath.Ext(infoInside.Name())
			if ext == ".java" || ext == ".sql" || ext == ".xml" || ext == ".hql" {
				//fmt.Printf("File repo: %q\n", pathInside)
				wg.Add(1)
				go readFileParallel(pathInside, "dw_r_recarga_oferta_dia", &wg)
			}
			//readFile(path)
		}
		return nil
	})

	wg.Wait()

	endTime := time.Now()
	difference := endTime.Sub(currentTime)

	fmt.Println("The Finish time is", difference)
	// if err_parent != nil {
	// 	fmt.Printf("error walking the path: %v\n", err_parent)
	// 	return
	// }

	if err != nil {
		fmt.Printf("error walking the path: %v\n", err)
		return
	}
}
