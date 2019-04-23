package main

import (
	"bufio"
	"buildxpxg/xpxg116"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// This is a copy of the header from the xml package with standalone added
// TODO - is there a better way to do this.
const (
	xmlHeader = `<?xml version="1.0" encoding="utf-8" standalone="no"?>` + "\n"
)

func pprint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

type Config struct {
	Name                    string `json:"Name"`
	ProPath                 string
	HelpString              string
	ODLHelpString           string
	ObjectType              bool
	AllowUnknown            bool
	IsTTResultSet           bool
	WriteDatasetBeforeImage bool
}

type Bird struct {
	Species      string
	Description  string
	AnotherField string
}

func main() {

	var configFileFlag string
	flag.StringVar(&configFileFlag, "c", "", "Config File")
	flag.StringVar(&configFileFlag, "config", "", "Config File")

	flag.Parse()

	println(configFileFlag)

	var bird Bird
	birdFile, err := os.Open("data/birds.json")

	byteBirdFile, err := ioutil.ReadAll(birdFile)
	fmt.Println(string(byteBirdFile))
	json.Unmarshal(byteBirdFile, &bird)
	fmt.Println(pprint(bird))

	var config Config

	configFile, err := os.Open("data/Rhapsody2.json")
	decodeError := json.NewDecoder(configFile).Decode(&config)
	if decodeError != nil {
		log.Fatalln(decodeError)
	}

	fmt.Println("Decode : " + pprint(config))
	if err != nil {
		log.Fatalln(err)
	}
	defer configFile.Close()

	byteConfigFile, err := ioutil.ReadAll(configFile)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(byteConfigFile))
	fmt.Println(byteConfigFile)

	json.Unmarshal(byteConfigFile, &config)

	fmt.Println(pprint(config))

	os.Exit(0)

	// TODO - Need a better way to handle multiple OpenEdge versions

	wsad := new(xpxg116.WSAD)

	wsad.XMLNS = "urn:schemas-progress-com:WSAD:0011"
	wsad.XSI = "http://www.w3.org/2001/XMLSchema-instance"
	wsad.AppObject.PGVersion.PXGVersion = "03"
	wsad.AppObject.PGGenInfo.Service = "wsHPMqa"
	wsad.AppObject.DeploymentWizard.WebServiceNamespace.Value = "Test"
	// TODO - Will ProxyGen with unix slashes instead of dos
	wsad.AppObject.PGGenInfo.WorkDir = "S:\\progress\\wsa\\prod\\hpm\\"
	wsad.AppObject.PGGenInfo.VerboseLogging = true
	wsad.AppObject.PGGenInfo.LeaveProxyFiles = true
	wsad.AppObject.PGGenInfo.SaveBeforeGen = true
	wsad.AppObject.PGGenInfo.WebServicesClient = true
	wsad.AppObject.PGGenInfo.UserDefinedAppService = true
	wsad.AppObject.PGGenInfo.JavaCompilerType = 1
	wsad.AppObject.PGGenInfo.JavaCompiler = "<Install Dir>\\jdk\\bin\\javac"
	wsad.AppObject.PGGenInfo.JavaClasspathSwitch = "-classpath"
	wsad.AppObject.PGGenInfo.JavaClasspath = "<Proxy Dir>;<Install Dir>\\java\\progress.jar;<Install Dir>\\java\\messages.jar;;<Default classes.zip>"
	wsad.AppObject.PGGenInfo.DotNetCompilerType = 4
	wsad.AppObject.PGGenInfo.DotNetCompiler = "csc"
	wsad.AppObject.PGGenInfo.DotNetXSDGenerator = "xsd"
	wsad.AppObject.Name = "Rhapsody"
	wsad.AppObject.ProPath = "S:\\progress\\edi\\;S:\\progress\\mcs\\"
	wsad.AppObject.HelpString = ""
	wsad.AppObject.ODLHelpString = ""
	wsad.AppObject.ObjectType = false
	wsad.AppObject.AllowUnknown = false
	wsad.AppObject.IsTTResultSet = false
	wsad.AppObject.WriteDatasetBeforeImage = false
	wsad.AppObject.PGGenInfo.DotNetRuntime = "digitally signed"
	wsad.AppObject.PGGenInfo.DotNetDataSetNamespace = "StrongTypesNS"
	wsad.AppObject.PGGenInfo.DNUseDefaultDSNamespace = true
	wsad.AppObject.PGGenInfo.DNUseNamespaceAsDirs = true
	wsad.AppObject.PGGenInfo.DotNetUseNullableTypes = true
	wsad.AppObject.PGGenInfo.SilverlightDomainServiceNamespace = "Web.Services"
	wsad.AppObject.PGGenInfo.SilverlightUseDefaultDomainServiceNamespace = true
	wsad.AppObject.PGGenInfo.SilverlightEntityNamespace = "Web"
	wsad.AppObject.PGGenInfo.SilverlightUseDefaultEntityNamespace = true

	wsad.AppObject.DeploymentWizard.CurrentEncoding = 3
	wsad.AppObject.DeploymentWizard.SoapEndpointURL = "http://mcs01.caidan.local/hpm/wsa1"
	wsad.AppObject.DeploymentWizard.WebServiceNamespace.Value = "urn:hpmich-local-Rhapsody"
	wsad.AppObject.DeploymentWizard.WebServiceNamespace.UserDefined = true
	wsad.AppObject.DeploymentWizard.SoapAction.Value = ""
	wsad.AppObject.DeploymentWizard.SoapAction.AppendToSoapAction = false
	wsad.AppObject.DeploymentWizard.ConnectReturnString.UserDefined = false
	wsad.AppObject.DeploymentWizard.TestWSDL.BGen = true
	wsad.AppObject.DeploymentWizard.TestWSDL.RPCEncoded = false
	wsad.AppObject.DeploymentWizard.TestWSDL.RPCLiteral = false
	wsad.AppObject.DeploymentWizard.TestWSDL.DocLiteral = true
	wsad.AppObject.DeploymentWizard.ESBEncoding = 3
	wsad.AppObject.DeploymentWizard.EntryEndpoint.UseDefault = true
	wsad.AppObject.DeploymentWizard.FaultEndpoint.UseDefault = true
	wsad.AppObject.DeploymentWizard.RejectedEndpoint.UseDefault = true
	wsad.AppObject.DeploymentWizard.AppserverURL = "AppServerDC://localhost:3091"
	wsad.AppObject.DeploymentWizard.FileName.Value = "Rhapsody"
	wsad.AppObject.DeploymentWizard.FileName.OverWrite = false
	wsad.AppObject.DeploymentWizard.ResourceDirectory = "/Resources/OpenEdgeServices"
	wsad.AppObject.DeploymentWizard.SonicConnectionInfo.SonicURL = "tcp://localhost:2506"
	wsad.AppObject.DeploymentWizard.SonicConnectionInfo.SonicDomain = "Domain1"
	wsad.AppObject.DeploymentWizard.SonicConnectionInfo.User = "Administrator"

	wsad.AppObject.Procedures = make([]xpxg116.Procedure, 0)

	fmt.Println(pprint(wsad))

	file, err := os.Open("data/filelist.txt")

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		baseDirectory := "s:\\progress"
		pathSlice := strings.Split(scanner.Text(), "\\")
		l := len(pathSlice)
		if l < 1 {
			log.Fatalln(fmt.Sprintf("Invalid Program Name %s", scanner.Text()))
		}
		fullProgramName := pathSlice[l-1]
		programParts := strings.Split(fullProgramName, ".")
		if len(programParts) != 2 {
			log.Fatalln("Invalid Program Name", fullProgramName)
		}
		wsad.AppObject.Procedures = append(wsad.AppObject.Procedures,
			xpxg116.Procedure{UseFullName: false,
				IsPersistent: false,
				Name:         programParts[0],
				ProcPath:     "",
				ProPath:      strings.Join(append([]string{baseDirectory}, pathSlice[0:l-1]...), "\\"),
				ProcExt:      programParts[1]})
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	output, err := xml.MarshalIndent(wsad, "  ", "    ")

	if err != nil {
		log.Fatalln(fmt.Printf("error: %v\n", err))
	}
	output = []byte(xmlHeader + string(output))
	// os.Stdout.Write(output)
}
