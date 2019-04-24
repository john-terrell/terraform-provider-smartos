package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"path"
	"regexp"

	"github.com/google/uuid"
	"golang.org/x/crypto/ssh"
)

type SmartOSClient struct {
	host   string
	user   string
	client *ssh.Client
}

func (c *SmartOSClient) Connect() error {
	var err error = nil

	if c.client != nil {
		return nil
	}

	log.Println("Creating client")
	user, err := user.Current()
	if err != nil {
		return err
	}

	keyPath := path.Join(user.HomeDir, ".ssh", "id_rsa")
	log.Println("Loading private key from ", keyPath)
	keyBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return err
	}

	log.Println("Parsing private key")
	signer, err := ssh.ParsePrivateKey(keyBytes)
	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User: c.user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	log.Println("Connecting to host: ", c.host)
	c.client, err = ssh.Dial("tcp", c.host, config)
	if err != nil {
		log.Println("Connection failed: ", err.Error())
		return err
	}

	log.Println("Connected successfully")
	return nil
}

func (c *SmartOSClient) Close() {
	if c.client != nil {
		c.client.Close()
		c.client = nil
	}
}

func (c *SmartOSClient) CreateMachine(machine *Machine) (*uuid.UUID, error) {
	err := c.Connect()
	if err != nil {
		return nil, err
	}

	session, err := c.client.NewSession()
	if err != nil {
		return nil, err
	}

	defer session.Close()

	json, err := json.Marshal(machine)
	if err != nil {
		log.Fatalln("Failed to create JSON for machine.  Error: ", err.Error())
	}

	log.Println("JSON: ", string(json))

	session.Stdin = bytes.NewReader(json)

	var b bytes.Buffer
	session.Stderr = &b

	err = session.Run("vmadm create")
	if err != nil {
		return nil, err
	}

	output := b.String()
	log.Printf("Returned data: ", output)

	re := regexp.MustCompile("Successfully created VM ([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})")
	matches := re.FindStringSubmatch(output)

	if len(matches) != 2 {
		return nil, fmt.Errorf("Unrecognized response from vmadm: %s", output)
	}

	log.Println("Matched regex: ", matches[1])
	uuid, err := uuid.Parse(matches[1])
	if err != nil {
		return nil, err
	}

	return &uuid, nil
}

func (c *SmartOSClient) GetMachine(id uuid.UUID) (*Machine, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (c *SmartOSClient) DeleteMachine(id uuid.UUID) error {
	err := c.Connect()
	if err != nil {
		return err
	}

	session, err := c.client.NewSession()
	if err != nil {
		return err
	}

	defer session.Close()

	var b bytes.Buffer
	session.Stderr = &b

	err = session.Run("vmadm delete " + id.String())
	if err != nil {
		return err
	}

	output := b.String()
	log.Printf("Returned data: ", output)

	re := regexp.MustCompile("Successfully deleted VM ([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})")
	matches := re.FindStringSubmatch(output)

	if len(matches) != 2 {
		return fmt.Errorf("Unrecognized response from vmadm: %s", output)
	}

	return nil
}
