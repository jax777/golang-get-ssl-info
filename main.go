package main 
import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

func GetDomainsFromCert(ip string,port int)[]string{
	dialer := net.Dialer{Timeout:time.Second*3}
	conn, err := tls.DialWithDialer(&dialer,"tcp", fmt.Sprintf("%s:%d",ip,port), &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Println("failed to connect: " + err.Error())
		return []string{}
	}
	defer conn.Close()
	state := conn.ConnectionState()
	return state.PeerCertificates[0].DNSNames
}
func main(){
	fmt.Println(GetDomainsFromCert("39.156.69.79",443))
}

