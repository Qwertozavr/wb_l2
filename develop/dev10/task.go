package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	var timeout string

	flag.StringVar(&timeout, "timeout", "10s", "time limit to establish connection")
	flag.Parse()

	ok, err := regexp.MatchString(`\ds`, timeout)
	if err != nil {
		log.Fatal(err)
	}

	if !ok {
		log.Fatal("invalid timeout format: " + timeout)
	}

	if len(flag.Args()) < 2 {
		log.Fatal("usage: --timeout=1s host port")
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	timeint, err := strconv.Atoi(timeout[:len(timeout)-1])
	if err != nil {
		log.Fatal(err)
	}

	timer := time.Duration(timeint) * time.Second

	var conn net.Conn

	now := time.Now()

	for time.Since(now) < timer {
		conn, err = net.Dial("tcp", host+":"+port)
		if err == nil {
			break
		}
	}
	if err != nil {
		log.Fatalf("can't establish connection after %v", timer)
	}

	defer conn.Close()

	log.Printf("connected to %s:%s", host, port)

	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
				time.Sleep(5 * time.Millisecond)
				continue
			}

			fmt.Printf("server: %s", message)

			if strings.Contains(message, "Connection: close") {
				break
			}
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in := scanner.Text()
		_, err := fmt.Fprintf(conn, in+"\n")
		if err != nil {
			log.Fatal("connection close")
		}
	}
}
