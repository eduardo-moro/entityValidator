package request

import (
	"bytes"
	"crypto/tls"
	json "encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type CfnJson struct {
	Success bool `json:"success"`
	Data    []struct {
		Nome         string `json:"nome"`
		Registro     string `json:"registro"`
		Crn          int    `json:"crn"`
		DataCadastro string `json:"data_cadastro"`
		Situacao     string `json:"situacao"`
		TipoRegistro string `json:"tipo_registro"`
	} `json:"data"`
}


//nutricionistas
func GetCfnByCode(codigo string) []byte {
	request := makeRequest("https://cnn.cfn.org.br/application/front-resource/get",
		"{\"comando\":\"get-nutricionista\",\"options\":{\"geral\":true, \"registro\": \""+codigo+"\"}}")
	return ValidaNutricionista(request, codigo)
}

func ValidaNutricionista(request string, codigo string) []byte {
	var parsedJson CfnJson
	err := json.Unmarshal([]byte(request), &parsedJson)

	if err != nil {
		log.Fatal(err)
	}

	index := -1

	for i := range parsedJson.Data {
		if parsedJson.Data[i].Registro == codigo {
			index = i
		}
	}

	if index >= 0 {
		retorno, _ := json.Marshal(parsedJson.Data[index])
		return retorno
	} else {
		retorno, _ := json.Marshal(parsedJson.Data)
		return retorno
	}
}


func GetCfnByName(nome string) (string) {
	return makeRequest("https://cnn.cfn.org.br/application/front-resource/get",
		"{\"comando\":\"get-nutricionista\",\"options\":{\"geral\":true, \"nome\": \""+nome+"\"}}")
}


//educador fisico
func GetCref(cref string) (string) {
	url := "https://www.confef.org.br/confef/registrados/ssp.registrados.php?columns[0][data]=0&columns[0][searchable]=true&columns[0][orderable]=false&columns[0][search][value]=&columns[0][search][regex]=false&columns[1][data]=1&columns[1][name]=&columns[1][searchable]=true&columns[1][orderable]=true&columns[1][search][value]=&columns[1][search][regex]=false&columns[2][data]=2&columns[2][name]=&columns[2][searchable]=true&columns[2][orderable]=true&columns[2][search][value]=&columns[2][search][regex]=false&columns[3][data]=3&columns[3][name]=&columns[3][searchable]=true&columns[3][orderable]=true&columns[3][search][value]=&columns[3][search][regex]=false&columns[4][data]=4&columns[4][name]=&columns[4][searchable]=true&columns[4][orderable]=true&columns[4][search][value]=&columns[4][search][regex]=false&columns[5][data]=5&columns[5][name]=&columns[5][searchable]=true&columns[5][orderable]=true&columns[5][search][value]=&columns[5][search][regex]=false&columns[6][data]=6&columns[6][name]=&columns[6][searchable]=true&columns[6][orderable]=true&columns[6][search][value]=&columns[6][search][regex]=false&columns[7][data]=7&columns[7][name]=&columns[7][searchable]=true&columns[7][orderable]=true&columns[7][search][value]=&columns[7][search][regex]=false&order[0][column]=1&order[0][dir]=asc&start=0&length=10&search[value]=" + cref + "&search[regex]=false"
	return makeRequest(url, "")
}

func GetCrefPj(cref string) (string) {
	url := "https://www.confef.org.br/confef/pj-registradas/ssp.registrados.php?columns[0][data]=0&columns[0][searchable]=true&columns[0][orderable]=false&columns[0][search][value]=&columns[0][search][regex]=false&columns[1][data]=1&columns[1][name]=&columns[1][searchable]=true&columns[1][orderable]=true&columns[1][search][value]=&columns[1][search][regex]=false&columns[2][data]=2&columns[2][name]=&columns[2][searchable]=true&columns[2][orderable]=true&columns[2][search][value]=&columns[2][search][regex]=false&columns[3][data]=3&columns[3][name]=&columns[3][searchable]=true&columns[3][orderable]=true&columns[3][search][value]=&columns[3][search][regex]=false&columns[4][data]=4&columns[4][name]=&columns[4][searchable]=true&columns[4][orderable]=true&columns[4][search][value]=&columns[4][search][regex]=false&columns[5][data]=5&columns[5][name]=&columns[5][searchable]=true&columns[5][orderable]=true&columns[5][search][value]=&columns[5][search][regex]=false&columns[6][data]=6&columns[6][name]=&columns[6][searchable]=true&columns[6][orderable]=true&columns[6][search][value]=&columns[6][search][regex]=false&columns[7][data]=7&columns[7][name]=&columns[7][searchable]=true&columns[7][orderable]=true&columns[7][search][value]=&columns[7][search][regex]=false&order[0][column]=1&order[0][dir]=asc&start=0&length=10&search[value]=" + cref + "&search[regex]=false"
	return makeRequest(url, "")
}


//cnpj
func GetCnpj(cnpj string) (string) {
	fmt.Println(os.Getenv("CNPJ_TOKEN"))
	url := "http://ws.hubdodesenvolvedor.com.br/v2/cnpj/?cnpj=" + cnpj + "&token=" + os.Getenv("CNPJ_TOKEN")
	return makeRequest(url, "")
}

func makeRequest(url string, input string) (string) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	body := strings.NewReader(input)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respBody := buf.String()


	if resp.Status != "200 OK" {
		log.Fatalf(resp.Status)
	}

	return respBody
}
