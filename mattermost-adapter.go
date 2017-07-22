package main

import (
"bytes"
"encoding/json"
"fmt"
"io"
"log"
"net/http"
"os"

"github.com/gorilla/mux"
"github.com/spf13/viper"
)

type AdapterConfig struct {
	MattermostServer string
	Channel string
	Username string
	IconUrl string
	Listen string
}

type PackagerJson struct {
	Event string `json:"event"`
	RepositoryUUID string `json:"repository_uuid"`
	RepositorySlug string `json:"repository_slug"`
	Filename string `json:"filename"`
	Commit string `json:"commit"`
	Branch string `json:"branch"`
	Tag string `json:"tag"`
	Tagged bool `json:"tagged"`
	RealTag string `json:"real_tag"`
	Distribution string `json:"distribution"`
	PackageURL string `json:"package_url"`
	UpstreamURL string `json:"upstream_url"`
	BuildURL string `json:"build_url"`
}

type MattermostJson struct {
	Channel string `json:"channel"`
	Text string `json:"text"`
	Username string `json:"username"`
	IconUrl string `json:"icon_url"`
}

func HandlePackagerPost(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "POST":

		request := fmt.Sprintf("Request from: %s\n", req.RemoteAddr)
		fmt.Printf("%s\n", request)

		dec := json.NewDecoder(req.Body)

		packageJson := new(PackagerJson)
		err := dec.Decode(&packageJson)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(packageJson)

		retjs, err := json.Marshal(packageJson)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		text := fmt.Sprintf("build succeeded! Branch: %s - Tag: %s - %s - %s - %s", packageJson.Branch,
																							packageJson.Tag,
																							packageJson.Commit,
																							packageJson.Distribution,
																							packageJson.PackageURL)

		packageMattermost := MattermostJson{Channel: config.Channel, Username: config.Username,
		                                    IconUrl: config.IconUrl,
											Text: text}

		payload := new(bytes.Buffer)
		json.NewEncoder(payload).Encode(packageMattermost)
		res, _ := http.Post(config.MattermostServer,
			                "application/json; charset=utf-8", payload)
		io.Copy(os.Stdout, res.Body)

		fmt.Fprintln(rw, string(retjs))
	}
}

var config AdapterConfig

func main() {

	viper.SetConfigName("app")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found or error parsing\n\n: %s", err)
	} else {
		config.Channel  = viper.GetString("general.channel")
		config.Username = viper.GetString("general.username")
		config.IconUrl  = viper.GetString("general.iconurl")
		config.MattermostServer = viper.GetString("general.mattermost")
		config.Listen     = viper.GetString("general.listen")

		fmt.Printf("\nUsing config:\n mattermost = %s\n channel = %s\n" +
			" username = %s\n" +
			" iconurl = %s\n\n" +
			"Listening on port: %s\n",
			config.MattermostServer,
			config.Channel,
			config.Username,
			config.IconUrl,
		    config.Listen)
	}


	router := mux.NewRouter()
	router.HandleFunc("/hook", HandlePackagerPost).Methods("POST")
	log.Fatal(http.ListenAndServe(config.Listen, router))
}
