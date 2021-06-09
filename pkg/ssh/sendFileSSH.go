/**
* @author Jee
* @date 2021/6/9 21:03
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

const (
	user     = "jee"
	ip_port  = "192.168.8.10:22"
	password = "141242"
)

var outCh = make(chan string)

func main() {
	stopCh := make(chan struct{})
	defer close(stopCh)

	conf := ssh.ClientConfig{
		Config:            ssh.Config{},
		User:              user,
		Auth:              []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback:   ssh.InsecureIgnoreHostKey(),
		BannerCallback:    nil,
		ClientVersion:     "",
		HostKeyAlgorithms: nil,
		Timeout:           0,
	}
	Client, err := ssh.Dial("tcp", ip_port, &conf)
	if err != nil {
		log.Fatalf("start client connection failed, err:%v", err)
		return
	}
	defer Client.Close()

	session, err := Client.NewSession()
	if err != nil {
		log.Fatalf("create new session failed, err:%v", err)
		return
	}
	defer session.Close()

	pipStdin, _ := session.StdinPipe()
	pipStdout, _ := session.StdoutPipe()
	pipStderr, _ := session.StderrPipe()
	go stdin(pipStdin, stopCh)
	go stdout(pipStdout, stopCh)
	log.Printf("Run start")
	if err = session.Run(cmd); err != nil {
		if exitErr, ok := err.(*ssh.ExitError); ok {
			if code := exitErr.ExitStatus(); code != 0 {
				err = fmt.Errorf("command run err:%v, exit code is %d", err, code)
			}
		} else {
			err = fmt.Errorf("failed running `%s`:%v", cmd, err)
		}
	}
	log.Printf("Run end")
	bout, _ := ioutil.ReadAll(pipStdout)
	berr, _ := ioutil.ReadAll(pipStderr)
	log.Printf("[bout]:%s  [berr]:%s  [err]%v", string(bout), string(berr), err)
	log.Printf("main over")
}

func stdout(pout io.Reader, stopCh <-chan struct{}) {
	arch, _ := bufio.NewReader(pout).ReadString('\n')
	outCh <- strings.TrimSpace(arch)
	log.Printf("stdout() get %v", strings.TrimSpace(arch))
	return
}

func stdin(pin io.WriteCloser, stopCh <-chan struct{}) {
	// session.Run()会直到pin.Close()之后返回
	// 因为若未close pin，则远端tee一直等待获取输入
	defer pin.Close()
	select {
	case <-stopCh:
		log.Printf("stdin() get stopCh")
	case arch := <-outCh:
		log.Printf("stdin() get %s", arch)
		var file *os.File
		var err error
		if arch == "amd64" {
			file, err = os.Open("x86.txt")
		} else {
			file, err = os.Open("arm.txt")
		}
		if err != nil {
			log.Printf("Open file %s failed, err:%v\n", file, err)
			return
		}
		if _, err := io.Copy(pin, file); err != nil {
			log.Fatalf("copy file failed, err:%v", err)
		}
		file.Close()
		time.Sleep(5 * time.Second)
	}
}

var cmd = `
if [[ "$(arch)"x == "aarch64"x ]];then
ARCH="arm64"
else
ARCH="amd64"
fi
echo $ARCH
tee > /home/jee/test.txt
`
