package service

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var snmpFile *os.File = nil

func parseTcpFile(file *os.File) (map[string]int, error) {
	var keyLine = ""
	var valueLine = ""
	var tcp = make(map[string]int)

	// find lines that start with "Tcp: "
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Tcp: ") {
			keyLine = strings.TrimPrefix(line, "Tcp: ")
			scanner.Scan()
			line = scanner.Text()

			if strings.HasPrefix(line, "Tcp: ") {
				valueLine = strings.TrimPrefix(line, "Tcp: ")
				break
			}
		}
	}

	if keyLine == "" || valueLine == "" {
		return tcp, fmt.Errorf("can not find Tcp Lines")
	}

	// split lines
	keys := strings.Split(keyLine, " ")
	values := strings.Split(valueLine, " ")

	if len(values) != len(keys) {
		return tcp, fmt.Errorf("length of keys and values are not equal")
	}

	// convert to map
	for i, key := range keys {
		value, err := strconv.Atoi(values[i])

		if err != nil {
			return tcp, err
		}

		tcp[key] = value
	}

	return tcp, nil
}

func GetTcpStatus() (map[string]int, error) {
	if snmpFile == nil {
		var err error

		snmpFile, err = os.Open(os.Getenv("SNMP_PATH"))

		if err != nil {
			return nil, err
		}
	}

	return parseTcpFile(snmpFile)
}

func GetRetransmissionRatio(tcp map[string]int) float64 {
	sent := tcp["OutSegs"]
	retransmited := tcp["RetransSegs"]

	return float64(retransmited) / float64(sent)
}
