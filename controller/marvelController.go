package controller

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func ConsultarPorNombres(ts, publicKey, hash, superHeroe string) {
  r := strings.NewReplacer(" ", "%20", "\n", "")
  superHeroe = r.Replace(superHeroe)

  url := "http://gateway.marvel.com/v1/public/characters?name=" + superHeroe + "&apikey=" + publicKey + "&ts=" + ts + "&hash=" + hash

  response, err := http.Get(url)
  if err !=  nil {
    fmt.Println("Falló la petición HTTP \n", err)
  } else {
    data, _ := ioutil.ReadAll(response.Body)
    fmt.Println(string(data))
  }
}

func Listar(ts, publicKey, hash string) {
  url := "http://gateway.marvel.com/v1/public/characters?orderBy=name&apikey=" + publicKey + "&ts=" + ts + "&hash=" + hash

  response, err := http.Get(url)
  if err !=  nil {
    fmt.Println("Falló la petición HTTP \n", err)
  } else {
    data, _ := ioutil.ReadAll(response.Body)
    fmt.Println(string(data))
  }
}
