package main

import (
	// "bytes"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
	// "github.com/gorilla/websocket"
)

type tsp_header struct {
	t       byte
	song_id int
}

type tsp_msg struct {
	header tsp_header
	msg    []byte
}

const (
	INFO_FILE = "songs.info"
	INIT      = 0
	LIST      = 1
)

func main() {
	args := os.Args[:]
	if len(args) != 2 {
		fmt.Println("Usage: ", args[0], "<port>")
		os.Exit(1)
	}

	// setup server socket
	ln, err := net.Listen("tcp", ":"+args[1])
	if err != nil {
		panic(err)
	}
	// defer ln.Close() //close connection when function is done

	for {
		conn, err := ln.Accept()
		if err != nil {
			// error accepting this connection
			continue
		}
		go handleConnection(conn)
	}
}

func write_songs_to_info(peer net.Conn, song_bytes []byte) {
	// lock
	song_strs := strings.Split(string(song_bytes[:]), "\n")
	ip := peer.RemoteAddr().String() + ", "

	// info_file, err := os.OpenFile(INFO_FILE, os.O_APPEND, 0666)
	info_file, err := os.OpenFile(INFO_FILE, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic("cant open songs.info file")
	}
	defer info_file.Close()

	for _, s := range song_strs {
		record := "ID, " + ip + s
		info_file.WriteString(record + "\n")
		info_file.Sync()
	}
	// unlock
}

func handleConnection(peer *net.Conn) {
	decoder := gob.NewDecoder(peer)
	in_msg := new(tsp_msg)
	// in_msg := &tsp_msg{}
	decoder.Decode(&in_msg)

	fmt.Println(in_msg)
	// fmt.Printf("Received : %+v", in_msg)
	peer.Close()

	// buff := make([]byte, 4096)
	// _, err := peer.Read(buff)
	// // bytes_read, err := peer.Read(buff)
	// if err != nil {
	//     fmt.Println("error reading song")
	//     os.Exit(1)
	// }
	//
	// tmpbuff := bytes.NewBuffer(buff)
	//
	// tmpstruct := new(tsp_msg)
	//
	// gobobj := gob.NewDecoder(tmpbuff)
	// gobobj.Decode(tmpstruct)
	//
	// fmt.Println(tmpstruct)

	// buff = append(buff[:bytes_read])
	// fmt.Println(string(buff[0]))

	os.Exit(1)

	// switch int(buff[0]) {
	// case 0:
	//     fmt.Println("received type INIT")
	//     break
	// case 1:
	// default:
	//     fmt.Println(buff)
	//     fmt.Println("bad msg header")
	//     return
	// }

	// write_songs_to_info(peer, buff)

	// get songs from client
	// send out info about songs to all hosts

	os.Exit(1)
	// receive songs from client

	// update info doc (locks)

	// send client (all?) updated doc
}
