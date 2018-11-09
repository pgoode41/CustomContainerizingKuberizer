package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	//Pull JSON And Store It In A Variable
	url := "https://swapi.co/api/people"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//Parse Through JSON
	var character Character
	json.Unmarshal(body, &character)
	people := character.Results

	//Create the dockerAuto Directory Tree and Building It IF It Doesn't Exist.

	dockerAutoPath := "/opt/dockerAuto"
	dockerBaseImagePath := "/opt/dockerAuto/dockerBaseImage"
	dockerCustomImagePath_ROOT := "/opt/dockerAuto/dockerGenImages"
	//dockerApplicationPATH := "/opt/dockerAppliactions"

	if _, err := os.Stat(dockerAutoPath); os.IsNotExist(err) {
		os.Mkdir(dockerAutoPath, 0764)
		os.Chmod(dockerAutoPath, 0764)
		if _, err := os.Stat(dockerCustomImagePath_ROOT); os.IsNotExist(err) {
			os.Mkdir(dockerCustomImagePath_ROOT, 0764)
			os.Chmod(dockerCustomImagePath_ROOT, 0764)
		}
		if _, err := os.Stat(dockerBaseImagePath); os.IsNotExist(err) {
			os.Mkdir(dockerBaseImagePath, 0764)
			os.Chmod(dockerBaseImagePath, 0764)
		}

	} else if _, err := os.Stat(dockerCustomImagePath_ROOT); os.IsNotExist(err) {
		os.Mkdir(dockerCustomImagePath_ROOT, 0764)
		os.Chmod(dockerCustomImagePath_ROOT, 0764)
	} else if _, err := os.Stat(dockerBaseImagePath); os.IsNotExist(err) {
		os.Mkdir(dockerBaseImagePath, 0764)
		os.Chmod(dockerBaseImagePath, 0764)
	}

	//List Of Illegal Characters For Variables.(Allows Spaces)
	//Use formatedJsonNameNOSpaces On Variable Directly
	badCharacterList := []string{"(", ")", `\`, `\`, ",", ".", "!", "@", "#", "$", "%", "^", "*", "+", "=", "[", "]", ";",
		`:`, "<", ">", "{", "}", "\\", "?", "/", "|", "`", "~", "_"}

	//Change To dockerBaseImagePath Directory.
	if err := os.Chdir(dockerBaseImagePath); err != nil {
		panic(err)
	}
	//Creating Base Dockerfile from Ubuntu Image and Customizing.
	baseImageDockerFile, err := os.Create("Dockerfile")
	os.Chmod("Dockerfile", 0764)
	if err != nil {
		log.Fatal("Cannot create Dockerfile", err)
	}
	defer baseImageDockerFile.Close()
	//Creating Base Dockerfile from Ubuntu Image and Customizing.
	//Name Must Stay Dockerfile (docker syntax rule).
	fmt.Fprintln(baseImageDockerFile, "FROM ubuntu")
	fmt.Fprintln(baseImageDockerFile, "RUN apt update -y && apt upgrade -y")
	fmt.Fprintln(baseImageDockerFile, "RUN apt install golang-go -y")
	fmt.Fprintln(baseImageDockerFile, "RUN apt install nano -y")
	fmt.Fprintln(baseImageDockerFile, "ENV TESTVAR1='This is an ENV var from dockerfile!'")
	baseImageDockerFile.Close()
	//Changing to the generated base Image directory.
	if err := os.Chdir(dockerBaseImagePath); err != nil {
		panic(err)
	}

	dockerBaseImageNAME := "dockerbase-image"
	dockerBaseImageVERSION := "v1"

	//Creating a bash script that runs the docker commands to make the image.
	bashBaseImageBuildScript, err := os.Create("bashBaseImageBuildScript.sh")
	os.Chmod("bashBaseImageBuildScript.sh", 0764)
	if err != nil {
		log.Fatal("Cannot create bashBaseImageBuildScript.sh", err)
	}
	defer baseImageDockerFile.Close()
	//Creating a bash script that runs the docker commands to make the image.
	fmt.Fprintln(bashBaseImageBuildScript, "#!/bin/bash")
	fmt.Fprintln(bashBaseImageBuildScript, "docker build -t"+dockerBaseImageNAME+":"+dockerBaseImageVERSION, ".")
	fmt.Fprintln(bashBaseImageBuildScript, "exit 0")
	bashBaseImageBuildScript.Close()
	//Executing a bash script that runs the docker commands to make the image.
	BASHScriptEXE(dockerBaseImagePath + "/bashBaseImageBuildScript.sh")

	//Change To /opt/dockerTesting/dockerGenImages Directory.
	if err := os.Chdir(dockerCustomImagePath_ROOT); err != nil {
		panic(err)
	}

	//Loop Through JSON
	for i := 0; i < len(people); i++ {
		//Create Variables From JSON
		// Using Name Method From Custom Type
		jsonName := people[i].Name
		//Removing Illegal Characters From Variable
		for _, badCharacter := range badCharacterList {
			jsonName = strings.Replace(jsonName, badCharacter, "", -1)
		}
		//Stores Data That Doesn't Contain Illegal Characters, But Allows Spaces.
		formatedJsonNameSpaces := strings.Title(jsonName)
		//Stores Data That Doesn't Contain Illegal Characters, But DOES NOT Allow Spaces.
		formatedJsonNameNOSpaces := strings.Replace(formatedJsonNameSpaces, " ", "", -1)
		//##################################################################################################################
		// VARIABLES BEING GENERATED START
		//##################################################################################################################
		//password := "SYS"+formatedJsonNameNOSpaces+"admiN"
		watchman_group_name := formatedJsonNameSpaces
		namesync := formatedJsonNameNOSpaces
		namesyncLower := strings.ToLower(namesync)
		//##################################################################################################################
		// VARIABLES BEING GENERATED END
		//##################################################################################################################

		if err := os.Chdir("/opt/dockerAuto/dockerGenImages"); err != nil {
			panic(err)
		}
		//Creating The Root Directory Fot The Generated Custom Images
		dockerCustomImagePath_IMAGES := dockerCustomImagePath_ROOT + "/" + namesyncLower
		if _, err := os.Stat(dockerCustomImagePath_IMAGES); os.IsNotExist(err) {
			os.Mkdir(dockerCustomImagePath_IMAGES, 0764)
			os.Chmod(dockerCustomImagePath_IMAGES, 0764)

		}

		//Change To Directory Were Custom Images Will Be Stored.
		if err := os.Chdir(dockerCustomImagePath_IMAGES); err != nil {
			panic(err)
		}

		/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

		dockerApplication_goServer, err := os.Create("goServer2.go")
		os.Chmod("goServer2.go", 0764)
		if err != nil {
			log.Fatal("Cannot create goServer2.go", err)
		}
		defer dockerApplication_goServer.Close()

		fmt.Fprintln(dockerApplication_goServer, "package main")
		fmt.Fprintln(dockerApplication_goServer, "import (")
		fmt.Fprintln(dockerApplication_goServer, `"fmt"`)
		fmt.Fprintln(dockerApplication_goServer, `"net/http"`)
		fmt.Fprintln(dockerApplication_goServer, ")")
		fmt.Fprintln(dockerApplication_goServer, "func main() {")
		fmt.Fprintln(dockerApplication_goServer, "http.HandleFunc(`/`, func(w http.ResponseWriter, r *http.Request) {")
		fmt.Fprintln(dockerApplication_goServer, "fmt.Fprint(w, `")
		fmt.Fprintln(dockerApplication_goServer, "<!doctype html>")
		fmt.Fprintln(dockerApplication_goServer, "<html lang='en'>")
		fmt.Fprintln(dockerApplication_goServer, "<head>")
		fmt.Fprintln(dockerApplication_goServer, "<meta charset='utf-8'>")
		fmt.Fprintln(dockerApplication_goServer, "<meta name='viewport' content='width=device-width, initial-scale=1'>")
		fmt.Fprintln(dockerApplication_goServer, "<title>"+watchman_group_name, "MicroMDM</title>")
		fmt.Fprintln(dockerApplication_goServer, "<style>")
		fmt.Fprintln(dockerApplication_goServer, "body {")
		fmt.Fprintln(dockerApplication_goServer, "font-family: -apple-system, BlinkMacSystemFont, sans-serif;")
		fmt.Fprintln(dockerApplication_goServer, "}")
		fmt.Fprintln(dockerApplication_goServer, "</style>")
		fmt.Fprintln(dockerApplication_goServer, "</head>")
		fmt.Fprintln(dockerApplication_goServer, "<body>")
		fmt.Fprintln(dockerApplication_goServer, "<h3>Welcome to", watchman_group_name, "MicroMDM!</h3>")
		fmt.Fprintln(dockerApplication_goServer, "<h1 style='background-color:rgb(0, 255, 255);'>"+watchman_group_name, "MicroMDM</h1>")
		fmt.Fprintln(dockerApplication_goServer, "<img class='right' src='https://657cea1304d5d92ee105-33ee89321dddef28209b83f19f06774f.ssl.cf1.rackcdn.com/gophercloud-edf0a107430a35b63fae80ea5d465fe648e194637e78f52455482f49c543769d.png'>")
		fmt.Fprintln(dockerApplication_goServer, "</body>")
		fmt.Fprintln(dockerApplication_goServer, "<h1>Doing Stuff In The Cloud ;D</h1>")
		fmt.Fprintln(dockerApplication_goServer, "</html>`)")
		fmt.Fprintln(dockerApplication_goServer, "})")
		fmt.Fprintln(dockerApplication_goServer, "fmt.Println(http.ListenAndServe(`:8080`, nil))")
		fmt.Fprintln(dockerApplication_goServer, "}")
		dockerApplication_goServer.Close()

		/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		//Creating Custon Dockerfile file from JSON DATA.
		//Name Must be Dockerfile (Docker syntax rule).
		customImageDockerFile, err := os.Create("Dockerfile")
		os.Chmod("Dockerfile", 0764)
		if err != nil {
			log.Fatal("Cannot create Dockerfile", err)
		}
		defer customImageDockerFile.Close()
		//Creating Custom Docker File From Json Data.
		fmt.Fprintln(customImageDockerFile, "FROM", dockerBaseImageNAME+":"+dockerBaseImageVERSION)
		fmt.Fprintln(customImageDockerFile, "RUN mkdir /app")
		fmt.Fprintln(customImageDockerFile, "COPY goServer2.go /app/goServer2.go")
		fmt.Fprintln(customImageDockerFile, "WORKDIR /app")
		fmt.Fprintln(customImageDockerFile, `CMD ["go", "run", "goServer2.go"]`)
		customImageDockerFile.Close()

		dockerCustomImageNAME := namesyncLower + "-image"
		dockerCustomImageVERSION := "v2"

		bashCustomImageBuildScript, err := os.Create("bashCustomImageBuildScript.sh")
		os.Chmod("bashCustomImageBuildScript.sh", 0764)
		if err != nil {
			log.Fatal("Cannot create bashCustomImageBuildScript.sh", err)
		}
		defer bashCustomImageBuildScript.Close()
		//Creating a bash script that runs the docker commands to make the image.
		fmt.Fprintln(bashCustomImageBuildScript, "#!/bin/bash")
		fmt.Fprintln(bashCustomImageBuildScript, "docker build -t"+dockerCustomImageNAME+":"+dockerCustomImageVERSION, ".")
		fmt.Fprintln(bashCustomImageBuildScript, "exit 0")
		bashCustomImageBuildScript.Close()
		//Executing a bash script that runs the docker commands to make the image.
		BASHScriptEXE(dockerCustomImagePath_IMAGES + "/bashCustomImageBuildScript.sh")

		/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		//Building Cluster

		gcloudProject := "terraform-gcloud4"
		clusterName := namesyncLower + "-cluster"
		clusterZone := "us-central1-a"
		externalPort := "8080"
		containerAppPort := "8080"

		bashKuberizeScript, err := os.Create("bashKuberizeScript.sh")
		os.Chmod("bashKuberizeScript.sh", 0764)
		if err != nil {
			log.Fatal("Cannot create bashKuberizeScript.sh", err)
		}
		defer bashKuberizeScript.Close()
		//Creating a bash script that runs the docker commands to make the image.
		fmt.Fprintln(bashKuberizeScript, "#!/bin/bash")
		fmt.Fprintln(bashKuberizeScript, `gcloud container clusters create`, clusterName, `--zone`, clusterZone)
		fmt.Fprintln(bashKuberizeScript, `docker tag`, dockerCustomImageNAME+":"+dockerCustomImageVERSION, `gcr.io/`+gcloudProject+`/`+dockerCustomImageNAME+":"+dockerCustomImageVERSION)
		fmt.Fprintln(bashKuberizeScript, `docker push gcr.io/`+gcloudProject+`/`+dockerCustomImageNAME+":"+dockerCustomImageVERSION)
		fmt.Fprintln(bashKuberizeScript, `kubectl run`, clusterName, `--image=gcr.io/`+gcloudProject+`/`+dockerCustomImageNAME+":"+dockerCustomImageVERSION)
		fmt.Fprintln(bashKuberizeScript, `kubectl expose deployment`, clusterName, `--type LoadBalancer --port`, externalPort, `--target-port`, containerAppPort)
		fmt.Fprintln(bashKuberizeScript, "exit 0")
		bashKuberizeScript.Close()
		//Executing a bash script that runs the docker commands to make the image.
		BASHScriptEXE(dockerCustomImagePath_IMAGES + "/bashKuberizeScript.sh")

		/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	}
}
func BASHScriptEXE(bashScript string) {
	//BASHScriptEXE executes shell scripts, WITHOUT arguments.
	//The Job WILL FAIL IF YOU ADD Them!
	//EX. /Users/prestongoode/Documents/testFrom.sh.
	//This Will Execute The Script At The End of The Path.
	//YOU MUST SPECIFY THE ENTIRE PATH TO THE SCRIPT, INCLUDING THE SCRIPT!!
	cmd := exec.Command(bashScript)
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(out.String())
}

//JSON Struct For Unmarshal
type Characters struct {
	Characters []Character `json:"users"`
}

type Character struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Name      string   `json:"name"`
		Height    string   `json:"height"`
		Mass      string   `json:"mass"`
		HairColor string   `json:"hair_color"`
		SkinColor string   `json:"skin_color"`
		EyeColor  string   `json:"eye_color"`
		BirthYear string   `json:"birth_year"`
		Gender    string   `json:"gender"`
		Homeworld string   `json:"homeworld"`
		Films     []string `json:"films"`
		Species   []string `json:"species"`
		Vehicles  []string `json:"vehicles"`
		Starships []string `json:"starships"`
		//Created   time.Time `json:"created"`
		//Edited    time.Time `json:"edited"`
		URL string `json:"url"`
	} `json:"results"`
}
