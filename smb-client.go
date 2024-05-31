package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hirochachacha/go-smb2"
)

type SMBQueryRequest string

type SMBClient struct {
	ServerHost string
	ServerPort string
	Username   string
	Password   string
	JobChannel chan SMBQueryRequest
}

func NewSMBClient(serverHost string, serverPort string, username string, password string, jobChannel chan SMBQueryRequest) *SMBClient {
	return &SMBClient{
		ServerHost: serverHost,
		ServerPort: serverPort,
		Username:   username,
		Password:   password,
		JobChannel: jobChannel,
	}
}

// Internal method to connect to
func (sc *SMBClient) connect() (net.Conn, *smb2.Session, error) {
	serverAddress := fmt.Sprintf("%s:%s", sc.ServerHost, sc.ServerPort)
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't dial smb server at %s: %v", serverAddress, err)
	}

	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     sc.Username,
			Password: sc.Password,
		},
	}

	s, err := d.Dial(conn)
	if err != nil {
		conn.Close()
		return nil, nil, fmt.Errorf("couldn't acquire smb server session: %w", err)
	}

	return conn, s, nil
}

// Internal method to diconnect from
func (sc *SMBClient) disconnect(conn net.Conn, s *smb2.Session) {
	conn.Close()
	s.Logoff()
}

func (sc *SMBClient) RunSession() {
	conn, s, err := sc.connect()
	if err != nil {
		log.Fatalln(err)
	}
	defer sc.disconnect(conn, s)

	for {

	}
}

func (sc *SMBClient) TestMount(shareName string) {

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", sc.ServerHost, sc.ServerPort))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     sc.Username,
			Password: sc.Password,
		},
	}
	s, err := d.Dial(conn)
	if err != nil {
		panic(err)
	}
	defer s.Logoff()

	fs, err := s.Mount(shareName)
	if err != nil {
		panic(err)
	}
	defer fs.Umount()

	dir, err := fs.ReadDir("./")
	if err != nil {
		panic(err)
	}

	for _, item := range dir {
		log.Println(item.Name())
	}
}
