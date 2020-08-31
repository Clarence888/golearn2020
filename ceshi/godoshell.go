package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/prerelease", prerelease)
	http.HandleFunc("/online", online)
	log.Fatal(http.ListenAndServe("localhost:9191", nil))
}

func prerelease(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("./testPre.sh")
	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		//fmt.Println(line)
		w.Write([]byte(line))
	}

	cmd.Wait()

}

func online(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("./testOnline.sh")
	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		//fmt.Println(line)
		w.Write([]byte(line))
	}

	cmd.Wait()
}
