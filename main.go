package main

import (
  "fmt"
  "bufio"
  "os"
  "crypto/md5"
  "encoding/hex"
  "time"
  "./controller"
)

func main(){
    fmt.Println("Iniciando la Aplicación...")

    privateKey := "d3d7b6b49c1e0513be3a9e1a3c82a237314d4310"
    publicKey := "1333ae45bdc0a8266d4fc25e1cc578f4"
    ts := time.Now().Format("2006-01-02");

    hash_aux := md5.New()
    hash_aux.Write([]byte (ts + privateKey + publicKey))
    hash := hex.EncodeToString(hash_aux.Sum(nil))

    var opcion int
    //var superHeroe string

    fmt.Println("Opciones")
    fmt.Println("----------------------------------------------------------")
    fmt.Println("1. Buscar por nombre")
    fmt.Println("2. Listar")
    fmt.Println("----------------------------------------------------------")
    fmt.Println("Elija una opción: ")
    fmt.Scanf("%d\n", &opcion)

    switch opcion {
      case 1:
        reader := bufio.NewReader(os.Stdin)
        fmt.Println("Ingrese el nombre del Súper Héroe: ")
        superHeroe, _ := reader.ReadString('\n')
        fmt.Print("EL SUPER HEROE ES: ", superHeroe)
        //fmt.Println("TEXTO REEMPLAZADO: ", strings.Replace(superHeroe, " ", "%20", -1))
        controller.ConsultarPorNombres(ts, publicKey, hash, superHeroe)
      case 2:
        controller.Listar(ts, publicKey, hash)
      default:
        fmt.Println("La opción ingresada no es válida...")
    }

    fmt.Println("Terminando la Aplicación...")
}
