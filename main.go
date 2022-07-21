package main

import (
	"flag"
	"fmt"
	"os"
        "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
        "golang.org/x/crypto/bcrypt"

)

/*
   como usar: debera compilar y generar el binario de go
   argumentos:

      --generate [requests "string" or requests file "json"]       genera un json con el token
      --file [path to json]            Se envia el json generado
      --search "name apellido_pat apellido_mat" busca usando un nombre



*/

var url_token string = "https://web.sintesis.com.bo/payment-gateway-v1/token"
var url_search string = "https://web.sintesis.com.bo/payment-gateway-v1/sintesis/client/search"
var hash_password string = "$2a$10$zLKPdNqPHMO2jlJWgATpwuB3CfR3ave5PjPwvTxLgODFsKKA1hT4m"


const global_generate = "None"
const global_file = "None"
const global_search = "None"

type rquests_token struct {
	Token      string `json:"token"`
	Username   string `json:"username"`
	PublicID   string `json:"publicId"`
	SessionID  string `json:"sessionId"`
	ModuleID   string `json:"moduleId"`
	OrgID      string `json:"orgId"`
	MerchantID string `json:"merchantId"`
	Logo       string `json:"logo"`
	Code       int    `json:"code"`
}
	
type data_request struct {
	Code  int `json:"code"`
	Items []struct {
		CodError int    `json:"codError"`
		Mensaje  string `json:"mensaje"`
		Cuenta   []struct {
			Cuenta       string `json:"cuenta"`
			DescServicio string `json:"descServicio"`
			Detalle      string `json:"detalle"`
			Moneda       int    `json:"moneda"`
			Nombre       string `json:"nombre"`
			Servicio     int    `json:"servicio"`
		} `json:"cuenta"`
		FechaOperativa int `json:"fechaOperativa"`
		NroOperacion   int `json:"nroOperacion"`
	} `json:"items"`
}

func main() {
   file := flag.String("file", global_file, "Envia la ruta de un archivo json alternativo")
   generate := flag.String("generate", global_generate, "Genera lo necesario para realizar las consultas [json or string]")
   search := flag.String("search", global_search, "Busca un cliente en la api \"Nombre apellido_pat apellido_mat\"")

   flag.Parse()
	// Convertir el apuntador a string
   file_string := *file
   generate_string := *generate
   search_string := *search

   file_bool := true
   generate_bool := true
   search_bool := true
   
   argument_array := [3]string{generate_string, file_string, search_string}

   for a := 0; a < len(argument_array); a++ {
      if argument_array[0] == "None" {
         generate_bool = false
      }
      if argument_array[1] == "None" {
         file_bool = false
      }
      if argument_array[2] == "None" {
         search_bool = false
      }

   }
   argument_bool_array := [3]bool{generate_bool, file_bool, search_bool} 
   
   if argument_bool_array[0] || argument_bool_array[1] || argument_bool_array[2] {

      if argument_bool_array[0] {
         if argument_bool_array[1] {
            if argument_bool_array[2] {
               fmt.Println("generate, file and search using")
               if argument_array[2] == "None" && argument_array[1] == "None" && argument_array[0] == "None" {
                  fmt.Println("[!] Argumento vacio [json and path and name]")
               } else {
                  if argument_array[0] == "json" {
                     generate_func(argument_array[0])
                     search_func(argument_array[2], argument_array[1])
                  } else {
                     fmt.Println("[!] requires -generate=json")
                  }

               }

            } else {
               fmt.Println("generate and file using")
               fmt.Println("[!] generate and file debe usarse junto a -search")
            }
         } else if argument_bool_array[2] {
            if argument_bool_array[1] {
               fmt.Println("generate, search and file using")
               if argument_array[2] == "None" && argument_array[1] == "None" && argument_array[0] == "None" {
                  fmt.Println("[!] Argumento vacio [json and path and name]")
               } else {
                  if argument_array[0] == "json" {
                     generate_func(argument_array[0])
                     search_func(argument_array[2], argument_array[1])
                  } else {
                     fmt.Println("[!] requires -generate=json")
                  }

               }

               
            } else {
               fmt.Println("generate and search using")
               if argument_array[2] == "None" && argument_array[0] == "None" {
                  fmt.Println("[!] Argumento vacio [json and path and name]")
               } else {
                  if argument_array[0] == "json" {
                     generate_func(argument_array[0])
                     search_func(argument_array[2], "token.json")
                  } else {
                     fmt.Println("[!] requires -generate=json")
                  }

               }

            }

         } else {
            fmt.Println("generate using")

            if argument_array[0] == "None" {
               fmt.Println("[!] Argumento vacio [json or string]")
            } else {
               generate_func(argument_array[0])
            }
         }





      } else if argument_bool_array[1] {
         if argument_bool_array[0] {
            if argument_bool_array[2] {
               fmt.Println("file, generate and search using")
               if argument_array[2] == "None" && argument_array[1] == "None" && argument_array[0] == "None" {
                  fmt.Println("[!] Argumento vacio [json and path and name]")
               } else {
                  if argument_array[0] == "json" {
                     generate_func(argument_array[0])
                     search_func(argument_array[2], argument_array[1])
                  } else {
                     fmt.Println("[!] requires -generate=json")
                  }

               }


            } else {
               fmt.Println("file and generate using")
               fmt.Println("[!] file and generate debe usarse junto a -search")

            }
         } else if argument_bool_array[2] {
            if argument_bool_array[0] {
               fmt.Println("file, search and generate using")
               if argument_array[2] == "None" && argument_array[1] == "None" && argument_array[0] == "None" {
                  fmt.Println("[!] Argumento vacio [json and path and name]")
               } else {
                  if argument_array[0] == "json" {
                     generate_func(argument_array[0])
                     search_func(argument_array[2], argument_array[1])
                  } else {
                     fmt.Println("[!] requires -generate=json")
                  }

               }

               
            } else {
               fmt.Println("file and search using")
               if argument_array[2] == "None" && argument_array[1] == "None" {
                  fmt.Println("[!] Argumento vacio [json and path]")
               } else {
                  search_func(argument_array[2], argument_array[1])
               }

            }
         } else {
            fmt.Println("file using\n[!] Debe usarse junto al argumento -search")
         }



      } else if argument_bool_array[2] {
         if argument_bool_array[0] {
            if argument_bool_array[1] {
               fmt.Println("search, generate and file using")
               if argument_array[2] == "None" && argument_array[1] == "None" && argument_array[0] == "None" {
                  fmt.Println("[!] Argumento vacio [json and path and name]")
               } else {
                  if argument_array[0] == "json" {
                     generate_func(argument_array[0])
                     search_func(argument_array[2], argument_array[1])
                  } else {
                     fmt.Println("[!] requires -generate=json")
                  }

               }

            } else {
               fmt.Println("search and generate using")
               if argument_array[2] == "None" && argument_array[0] == "None" {
                  fmt.Println("[!] Argumento vacio [json and path and name]")
               } else {
                  if argument_array[0] == "json" {
                     generate_func(argument_array[0])
                     search_func(argument_array[2], "token.json")
                  } else {
                     fmt.Println("[!] requires -generate=json")
                  }

               }


            }
         } else if argument_bool_array[1] {
            if argument_bool_array[0] {
               fmt.Println("search, file and generate using")
            } else {
               fmt.Println("search and file using")
               if argument_array[2] == "None" && argument_array[1] == "None" {
                  fmt.Println("[!] Argumento vacio [json and path]")
               } else {
                  search_func(argument_array[2], argument_array[1])
               }

            }
         }  else {
            fmt.Println("search using")
            
            if argument_array[2] == "None" {
               fmt.Println("[!] Argumento vacio [nombre apellido_pat apellido_mat]")
            } else {
               search_func(argument_array[2], "token.json")
            }

         }

      }
   }
}


// verifica si existe un archivo
func json_exist(ruta string) bool {
	if _, err := os.Stat(ruta); os.IsNotExist(err) {
		return false
	}
	return true
}


// funcion generate terminado
////////////////////////////////////////////////////////////////////
func generate_func(cmd string) {

   hash_byte := []byte(hash_password)
   var pass string
   fmt.Print("[!] Contraseña de acceso : ")
   fmt.Scanln(&pass)
   
   hash_pass := []byte(pass)
   error := bcrypt.CompareHashAndPassword(hash_byte, hash_pass)

   if error == nil {
      if cmd == "json" {
         if json_exist("token.json") {
            menu := "[Parece que ya tiene un archivo json local]\n\t[1] crear nuevo\n\t[2] dejarlo ahi"
            fmt.Println(menu)
            fmt.Print("[+] seleccion : ")
            var action int
            fmt.Scanln(&action)
         
            switch action {
            case 1:
               create_json(pass)
               fmt.Println("[+] Processing continue ...")
            case 2:
               fmt.Println("[+] Processing continue ...")
            default:
               fmt.Println("[!] Elija una opcion correcta")
            }
         } else {
            create_json(pass)
         }
      
      } else if cmd == "string" {
         file, err := ioutil.ReadFile("token.json")
         if err != nil {
            log.Fatal(err)
         }
         text := string(file)
         fmt.Println()
         fmt.Println()
         fmt.Println(text)

      } else {
         fmt.Println("-generate no puede recibir este argumento : ["+cmd+"]")
      }
   } else {
      fmt.Print("[!] La contraseña no es correcta! ["+pass+"]")
      os.Exit(0)
   }
}
////////////////////////////////////////////////////////////////////


// funcion search terminada junto a file
////////////////////////////////////////////////////////////////////
func search_func(cmd string, ruta string) {
   
   if !json_exist(ruta) {
      log.Fatal("[!] No existe tal archivo")
   }

   nombre_dat := strings.Split(cmd, " ")

   file, err := ioutil.ReadFile(ruta)
    if err != nil {
        log.Fatal(err)
    }

    info := rquests_token{}
    data_resq := data_request{}


   err = json.Unmarshal(file, &info)

   if err != nil {
      log.Fatal(err)
   }
   payload := strings.NewReader("{\"idOperativo\":\""+info.SessionID+"\",\"codModulo\":\"91\",\"codCriterio\":2,\"codigo\":[\""+nombre_dat[0]+"\",\""+nombre_dat[1]+"\",\""+nombre_dat[2]+"\"]}1")
   
   
   req, _ := http.NewRequest("POST", url_search, payload)

   req.Header.Add("user-agent", "vscode-restclient")
   req.Header.Add("authorization", "Bearer "+info.Token)
   req.Header.Add("accept", "application/json")
   req.Header.Add("content-type", "application/json")
   req.Header.Add("referer", referer_gen(info))
   
   res, _ := http.DefaultClient.Do(req)

   defer res.Body.Close()
   body, _ := ioutil.ReadAll(res.Body)
   
   err = json.Unmarshal(body, &data_resq)
   data, err := json.MarshalIndent(data_resq, "", "  ")

   if err != nil {
      log.Fatal(err)
   }

   fmt.Println()
   fmt.Println(string(data))


}
func referer_gen(info rquests_token) string {

   referer := "https://web.sintesis.com.bo/payment-gateway-v1/suite/?&token="+info.Token+"&username="+info.Username+"&publicId="+info.PublicID+"&sessionId="+info.SessionID+"&moduleId="+info.ModuleID+"&merchantId="+info.MerchantID+"&orgId="+info.OrgID+"&type=WS_SINTESIS&merchantCustomerId=undefined&criteriaCode=undefined&onlyOneSubCriteria=true&logo=logo-emi.png&code=undefined&etiqueta=0"

   return referer
}
func existError(err error) bool {
  if err != nil {
    fmt.Println(err.Error())
  }
  return (err != nil)
}
func create_json(password string){
   payload := strings.NewReader("{\"username\":\"emi\",\"password\":\""+password+"\"}")
   req, _ := http.NewRequest("POST", url_token, payload)
   req.Header.Add("user-agent", "vscode-restclient")
   req.Header.Add("content-type", "application/json")
   res, _ := http.DefaultClient.Do(req)
   defer res.Body.Close()
   body, _ := ioutil.ReadAll(res.Body)
   info := rquests_token{}
   err := json.Unmarshal(body, &info)       
   data, err := json.MarshalIndent(info, "", "  ")
   if err != nil {
      log.Fatal(err)
   } 
   b := []byte(data)
   err = ioutil.WriteFile("token.json", b, 0644)
   if err != nil {
      log.Fatal(err)
   }
   fmt.Println("[+] Archivo creado satisfactoriamente")
}
