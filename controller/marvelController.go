package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//Estructura que contendrá el Response de la API de Marvel
type MarvelResponse struct {
	Code            int
	Status          string
	Copyright       string
	AttributionText string
	AttributionHTML string
	Etag            string
	Data            Data
}

type Data struct {
	Offset  int
	Limit   int
	Total   int
	Count   int
	Results []Results
}

type Results struct {
	Id          int
	Name        string
	Description string
	Modified    string
}

/**
* Método o Función que realiza una solicitud por nombre a la API de marvel (OPCION 1)
* @Parámetros
* ts - timestamp
* publicKey - llave pública
* hash - md5 de la concatenación de ts + privateKey + publicKey
* superHeroe - Super heroe ingresado por el usuario
 */
func ConsultarPorNombres(ts, publicKey, hash, superHeroe string) {
	r := strings.NewReplacer(" ", "%20", "\n", "")
	superHeroe = r.Replace(superHeroe)

	url := "http://gateway.marvel.com/v1/public/characters?name=" + superHeroe + "&apikey=" + publicKey + "&ts=" + ts + "&hash=" + hash

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Falló la petición HTTP \n", err)
	} else {
		var marvelResponse MarvelResponse

		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(data, &marvelResponse)

		fmt.Println("******************* DATOS DEL SÚPER HÉROE *********************")
		fmt.Println("ID: ", marvelResponse.Data.Results[0].Id)
		fmt.Println("NAME: ", marvelResponse.Data.Results[0].Name)
		fmt.Println("DESCRIPTION: ", marvelResponse.Data.Results[0].Description)
		fmt.Println("MODIFIED: ", marvelResponse.Data.Results[0].Modified)
	}
}

/**
* Método o Función que realiza una solicitud a la API de Marvel de los 20
primeros super heroes ordenados por nombre (OPCION 2)
* @Parámetros
* ts - timestamp
* publicKey - llave pública
* hash - md5 de la concatenación de ts + privateKey + publicKey
* superHeroe - Super heroe ingresado por el usuario
*/
func Listar(ts, publicKey, hash string) {
	url := "http://gateway.marvel.com/v1/public/characters?orderBy=name&limit=20&apikey=" + publicKey + "&ts=" + ts + "&hash=" + hash

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Falló la petición HTTP \n", err)
	} else {
		var marvelResponse MarvelResponse

		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(data, &marvelResponse)

		fmt.Println("***************** DATOS DE LOS SÚPER HÉROES *******************")
		for i := 0; i < marvelResponse.Data.Limit; i++ {
			fmt.Println("Súper Héroe # ", i+1)
			fmt.Println("--------------------")
			fmt.Println("ID: ", marvelResponse.Data.Results[i].Id)
			fmt.Println("NAME: ", marvelResponse.Data.Results[i].Name)
			fmt.Println("DESCRIPTION: ", marvelResponse.Data.Results[i].Description)
			fmt.Println("MODIFIED: ", marvelResponse.Data.Results[i].Modified)
			fmt.Println("")
		}
		//fmt.Println(string(data))
	}
}
