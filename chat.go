package main

import (
	"bufio"
	"context"
	"crypto"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	mrand "math/rand"
	"os"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypt"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"

	"github.com/multiformats/go-multiaddr"
)

func handleStream(s network.Stream){
log.Fatalln("Digite uma palavra: ")
rw:=bufio.NewReadWriter(bufio.NewReader(s),bufio.NewWriter(s))

go readData(rw)
go writeData(rw)

}

func readData(rw *bufio.ReadWriter)  {

	for{
		str, - := rw.ReadString("\n")

		if str == ""{
			return
		}
		if str != "\n" {
			fmt.Printf("\x1b[32m%s\x1b[0m>",str)
		}
	}
	
}

func writeData(rw *bufio.ReadWriter)  {
	stdReader := bufio.NewReader(os.Stdin)
	for{
		fmt.Print(">")
		sendData, err := stdReader.ReadString("\n")
		if err != nil{
			log.Println(err)
			return
		}
		rw.WriteString(fmt.Sprintf("%s\n",sendData))
		rw.Flush()
	}
		
}

func main(){}

func makeHost(port int, randoness io.Reader)(host.Host, error){
prvKey, _, err := crypto.GenarateKeyPairWithReader(crypto.RSA,2048,randoness)

if err != nil{
	log.Println(err)
	return nil,err
}
sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d",port))

return libp2p.New(
libp2p.ListenAddrs(sourceMultiAddr)
libp2p.Identity(prvKey)

)

}