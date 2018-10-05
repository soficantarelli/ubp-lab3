package main

/*
	Application name	Ejemplo
	API key				db72af200b346629a1bdd90e9f44d8e1
	Shared secret		a1828aa9a9263f18c0ab655fa64dcf52
	Registered to		emikohmann
*/

// dependencias
import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type APIResponseArtist struct {
	Artist struct {
		Bio struct {
			Content string `json:"content"`
			Links   struct {
				Link struct {
					Text string `json:"#text"`
					Href string `json:"href"`
					Rel  string `json:"rel"`
				} `json:"link"`
			} `json:"links"`
			Published string `json:"published"`
			Summary   string `json:"summary"`
		} `json:"bio"`
		Image []struct {
			Text string `json:"#text"`
			Size string `json:"size"`
		} `json:"image"`
		Mbid    string `json:"mbid"`
		Name    string `json:"name"`
		Ontour  string `json:"ontour"`
		Similar struct {
			Artist []struct {
				Image []struct {
					Text string `json:"#text"`
					Size string `json:"size"`
				} `json:"image"`
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"artist"`
		} `json:"similar"`
		Stats struct {
			Listeners string `json:"listeners"`
			Playcount string `json:"playcount"`
		} `json:"stats"`
		Streamable string `json:"streamable"`
		Tags       struct {
			Tag []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"tag"`
		} `json:"tags"`
		URL string `json:"url"`
	} `json:"artist"`
}

type APIError struct {
	Error   int           `json:"error"`
	Message string        `json:"message"`
	Links   []interface{} `json:"links"`
}

// constantes
const (
	// token de autenticacion
	apiKey = "d2c6ad784006816d4f17a1721ee7baba"

	// ubicacion de la api a consumir
	baseUrl = "http://ws.audioscrobbler.com/2.0/"
)

var (
	params = map[string]interface{}{
		"api_key": apiKey,
		"format":  "json",
	}
)

// funcion principal
func main() {
	for true {
		fmt.Print("Ingresa un artista: ")

		var artistName string
		fmt.Scanf("%s", &artistName)

		params["method"] = "artist.getinfo"
		params["artist"] = artistName

		url := GetUrl(params)

		artistData, err := GetArtistData(url)
		if err != nil {
			panic(err)
		}

		fmt.Println(artistData.Artist.Bio.Content)
	}
}

func GetUrl(params map[string]interface{}) string {
	var buff bytes.Buffer
	buff.WriteString(baseUrl)
	buff.WriteString("?")

	i := 0
	for k, v := range params {
		if i > 0 {
			buff.WriteString("&")
		}
		buff.WriteString(k)
		buff.WriteString("=")

		// %v representa un valor
		// %s representa un string
		// %d representa un numero
		vString := fmt.Sprintf("%v", v)
		buff.WriteString(vString)
		i++
	}
	return buff.String()
}

func GetArtistData(url string) (*APIResponseArtist, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal permite mapear un json en una estructura de Go
	var apiResponse APIResponseArtist
	err = json.Unmarshal(data, &apiResponse)
	if err != nil {
		return nil, err
	}

	// Unmarshal permite mapear un json en una estructura de Go
	var apiErr APIError
	err = json.Unmarshal(data, &apiErr)
	if err != nil {
		return nil, err
	}

	if apiErr.Message != "" {
		return nil, errors.New("Artist not found")
	}

	return &apiResponse, nil
}
