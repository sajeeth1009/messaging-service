package smtp_client

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type SmtpServerList struct {
	Servers  []SmtpServer `yaml:"servers"`
	FromName string       `yaml:"fromName"`
	FromAddr string       `yaml:"from"`
}

type SmtpServer struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	AuthData struct {
		Username string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"auth"`
}

// Address URI to smtp server
func (s *SmtpServer) Address() string {
	return s.Host + ":" + s.Port
}

func (sl *SmtpServerList) ReadFromFile(fname string) (err error) {
	yamlFile, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Printf("ReadServerConfig: err #%v ", err)
		return err
	}
	err = yaml.UnmarshalStrict(yamlFile, &sl)
	return
}
