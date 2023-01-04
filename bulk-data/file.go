package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pkg/profile"
)

func main() {
	//Trace
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	//CPU
	//defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	//Memory
	//defer profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.MemProfileRate(1)).Stop()

	basePath := "C:/Users/Fico/Desktop/truora/zinc/data_base/enron_mail_20110402/maildir"

	folders := mailDir(basePath)

	fmt.Println(folders)
	//single
	/**/
	for i, s := range folders {
		folderPath := basePath + "/" + s
		fmt.Printf("folder %d: %s\n", i, folderPath)
		readFolder(folderPath)

	}

	//GORUTINAS

	/*
		var wg sync.WaitGroup

		for i := 0; i < 4; i++ {
			wg.Add(1)

			for _, s := range folders {
				folderPath := basePath + "/" + s
				//fmt.Printf("folder %d: %s\n", i, folderPath)

				go func() {
					// Aquí va el código de la función
					readFolder(folderPath)
					wg.Done()
				}()

			}

		}

		wg.Wait()
	*/

	bulkZinc("json")

}

func mailDir(basePath string) []string {
	listFolder := []string{}

	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		//fmt.Println(file.Name(), file.IsDir())
		listFolder = append(listFolder, file.Name())
	}
	return listFolder
}

func readFolder(pathFolder string) {
	// Asume que existe un directorio llamado "dir" en el mismo directorio que este programa
	filepath.Walk(pathFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Si ocurre un error al acceder al archivo o directorio, imprimimos el error y devolvemos el error para detener el recorrido
			fmt.Println(err)
			return err
		}

		if !info.IsDir() {
			parseAndCreateJSON(info.Name(), path)
		}

		return nil
	})

}

func parseAndCreateJSON(nameFile string, pathFile string) {

	// Si encontramos un archivo, lo abrimos para lectura
	file, err := os.Open(pathFile)
	if err != nil {
		// Si ocurre un error al abrir el archivo, imprimimos el error y devolvemos el error para detenerel recorrido
		fmt.Println(err)

	}
	// Asegúrate de cerrar el archivo cuando hayas terminado de usarlo
	defer file.Close()

	// Creamos un nuevo Scanner para leer el archivo línea a línea
	scanner := bufio.NewScanner(file)

	re, _ := regexp.Compile("^[a-zA-Z]+(?:-[a-zA-Z]+)*:\\s")
	allBody := make(map[string]string)

	isBody := false
	prevKey := ""

	for scanner.Scan() {

		tmpLinea := scanner.Text()

		// Intenta hacer match con la expresión regular
		match := re.MatchString(tmpLinea)
		if match && !isBody {
			indexSeparator := strings.Index(tmpLinea, ": ")

			key := tmpLinea[0:indexSeparator]
			value := tmpLinea[indexSeparator+2:]
			allBody[key] = value
			prevKey = key
			continue

		}

		if !isBody {
			isBody = len(tmpLinea) == 0
		}

		if !isBody {
			allBody[prevKey] = allBody[prevKey] + tmpLinea

		} else {

			_, ok := allBody["body"]
			if !ok {
				allBody["body"] = tmpLinea
				prevKey = "body"
			} else {
				allBody["body"] = allBody["body"] + tmpLinea
			}
		}

	}

	// Revisamos si hubo algún error durante la lectura
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}

	createJSON(nameFile, allBody)
}

func createJSON(nameFile string, content map[string]string) {

	f, err := os.OpenFile("json/"+nameFile+".json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		// Si hay un error al abrir el archivo, lo manejamos aquí
		fmt.Println("Error al abrir el archivo:", err)
		return
	}

	header := Index{
		Index: map[string]string{
			"_index": "email",
		},
	}

	encoder := json.NewEncoder(f)
	// Escribe la primera línea de objeto JSON
	err = encoder.Encode(header)
	if err != nil {
		// Si hay un error al escribir el objeto, lo manejamos aquí
		fmt.Println("Error al escribir el objeto:", err)
		return
	}

	err = encoder.Encode(content)
	if err != nil {
		// Si hay un error al escribir el objeto, lo manejamos aquí
		fmt.Println("Error al escribir el objeto:", err)
		return
	}

	// Cierra el archivo
	err = f.Close()
	if err != nil {
		// Si hay un error al cerrar el archivo, lo manejamos aquí
		fmt.Println("Error al cerrar el archivo:", err)
		return
	}
}

func bulkZinc(pathFolder string) {

	folders := mailDir(pathFolder)
	url := "http://localhost:4080/api/_bulk"
	username := "admin"
	password := "Complexpass#123"

	for _, s := range folders {
		fileJson := pathFolder + "/" + s
		//fmt.Printf("folder %d: %s\n", i, folderPath)
		//readFolder(folderPath)
		data, err := ioutil.ReadFile(fileJson)
		if err != nil {
			fmt.Println(err)
			return
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		if err != nil {
			fmt.Println(err)
			return
		}
		req.SetBasicAuth(username, password)
		req.Header.Set("Content-Type", "application/octet-stream")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(body))

	}

	/**/

}

type Index struct {
	Index map[string]string `json:"index"`
}
