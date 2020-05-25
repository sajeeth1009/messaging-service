package smtp_client

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type SmtpServerList struct {
	Servers []SmtpServer `yaml:"servers"`
	From    string       `yaml:"from"`
	Sender  string       `yaml:"sender"`
	ReplyTo []string     `yaml:"replyTo"`
}

type SmtpServer struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Connections int    `yaml:"connections"`
	AuthData    struct {
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
