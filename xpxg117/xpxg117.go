package xpxg117

// TODO - This is just a copy of the 116 structs.  Need to update for 11.7
import "encoding/xml"

// WSAD - Openedge 11.6 WSAD
type WSAD struct {
	// XMLName xml.Name `xml:"WSAD"`
	XMLNS     string    `xml:"xmlns,attr"`
	XSI       string    `xml:"xmlns:xsi,attr"`
	AppObject AppObject // `xml:"AppObject"`
}

// PGVersion - pxpg Structure
type PGVersion struct {
	PXGVersion string // `xml:"PXGVersion"`
}

// InitPGVersion - pxpg Structure
func InitPGVersion() PGVersion {
	return PGVersion{PXGVersion: "03"}
}

type AppObject struct {
	// XMLName xml.Name `xml:"AppObject"`
	PGVersion        PGVersion // `xml:"PGVersion"`
	PGGenInfo        PGGenInfo `xml:"PGGenInfo"`
	DeploymentWizard DeploymentWizard
}

type PGGenInfo struct {
	XMLName                                     xml.Name `xml:"PGGenInfo"`
	SessionFree                                 bool     `xml:"isSessionFree,attr"`
	Author                                      string
	VersionString                               string
	VersionNumber                               int
	Package                                     string
	Service                                     string `xml:"Service"`
	WorkDir                                     string
	VerboseLogging                              bool
	LeaveProxyFiles                             bool
	SaveBeforeGen                               bool
	DotNetClient                                bool
	JavaClient                                  bool
	WebServicesClient                           bool
	ESBClient                                   bool
	ESBClient2                                  bool
	SilverlightClient                           bool
	RESTClient                                  bool
	UserDefinedAppService                       bool
	JavaCompilerType                            int
	JavaCompiler                                string
	JavaSwitches                                string
	JavaClasspathSwitch                         string
	JavaClasspath                               string
	DotNetCompilerType                          int
	DotNetCompiler                              string
	DotNetSwitches                              string
	DotNetXSDGenerator                          string
	DotNetNamespace                             string
	DotNetTitle                                 string
	DotNetVersion                               string
	DotNetDesc                                  string
	DotNetCompany                               string
	DotNetProduct                               string
	DotNetPublicKey                             string
	DotNetDelaySign                             bool
	DotNetRuntime                               string
	DotNetDataSetNamespace                      string
	DNUseDefaultDSNamespace                     bool
	DNUseNamespaceAsDirs                        bool
	DotNetUseNullableTypes                      bool
	SilverlightDomainServiceNamespace           string
	SilverlightUseDefaultDomainServiceNamespace bool
	SilverlightEntityNamespace                  bool
	SilverlightUseDefaultEntityNamespace        bool
}

type DeploymentWizard struct {
	CurrentEncoding     int
	SoapEndpointURL     string
	WebServiceNamespace WebServiceNamespace
	SoapAction          SoapAction
	ConnectReturnString ConnectReturnString
	TestWSDL            TestWSDL
	ESBEncoding         int
	EntryEndpoint       EntryEndPoint
	FaultEndpoint       FaultEndPoint
	RejectedEndpoint    RejectedEndPoint
	Container           string
	AppserverURL        string
	FileName            FileName
	ResourceDirectory   string
}

type WebServiceNamespace struct {
	XMLName            xml.Name `xml:"WebServiceNamespace"`
	Value              string   `xml:",chardata"`
	UserDefined        bool     `xml:"userDefined,attr"`
	SessionFree        bool     `xml:"isSessionFree,attr"`
	AppendToSoapAction bool     `xml:"appendToSoapAction,attr"`
}

type SoapAction struct {
	XMLName   xml.Name `xml:"SoapAction"`
	Value     string   `xml:",chardata"`
	OverWrite bool     `xml:"overWrite,attr"`
}

type FileName struct {
	XMLName     xml.Name `xml:"FileName"`
	Value       string   `xml:",chardata"`
	UserDefined bool     `xml:"userDefined,attr"`
}

type TestWSDL struct {
	XMLName    xml.Name `xml:"TestWSDL"`
	Value      string   `xml:",chardata"`
	BGen       bool     `xml:"bGen,attr"`
	RPCEncoded bool     `xml:",attr"`
	RPCLiteral bool     `xml:",attr"`
	DocLiteral bool     `xml:",attr"`
}

type EntryEndPoint struct {
	XMLName     xml.Name `xml:"EntryEndPoint"`
	Value       string   `xml:",chardata"`
	UserDefined bool     `xml:"userDefined,attr"`
}

type ConnectReturnString struct {
	XMLName     xml.Name `xml:"ConnectReturnString"`
	Value       string   `xml:",chardata"`
	UserDefined bool     `xml:"userDefined,attr"`
}

type FaultEndPoint struct {
	XMLName     xml.Name `xml:"FaultEndPoint"`
	Value       string   `xml:",chardata"`
	UserDefined bool     `xml:"userDefined,attr"`
}

type RejectedEndPoint struct {
	XMLName     xml.Name `xml:"RejectedEndPoint"`
	Value       string   `xml:",chardata"`
	UserDefined bool     `xml:"userDefined,attr"`
}

type SonicAppService struct {
	UseFileSystem     bool `xml:"userFileSystem,attr"`
	DeployToDS        bool `xml:"deployToDS,attr"`
	CreateXAR         bool `xml:"createXAR,attr"`
	ResourceDirectory string
	OverWrite         bool `xml:"overWrite,attr"`
	XARName           string
}
