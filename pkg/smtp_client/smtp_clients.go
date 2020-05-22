package smtp_client

type SmtpClients struct {
	servers SmtpServerList
	counter int
}

func NewSmtpClients(configFile string) (*SmtpClients, error) {
	serverList := SmtpServerList{}
	if err := serverList.ReadFromFile(configFile); err != nil {
		return nil, err
	}

	sc := &SmtpClients{
		servers: serverList,
		counter: 0,
	}
	return sc, nil
}
