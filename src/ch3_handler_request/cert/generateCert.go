package main

import (
	"math/big"
	"crypto/rand"
	"crypto/x509/pkix"
	"crypto/x509"
	"time"
	"net"
	"crypto/rsa"
	"os"
	"encoding/pem"
)

//生成ssl证书以及私钥
func main() {
	//fmt.Printf("ok")
	generateSsl()
}

func generateSsl() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)

	subject := pkix.Name{
		Organization:       []string{"Manning Publication Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}

	template := x509.Certificate{
		SerialNumber:serialNumber,
		Subject:subject,
		NotBefore:time.Now(),
		NotAfter:time.Now().Add(365 * 24 * time.Hour),
		//KeyUsage,ExtKeyUsage:表明这个X.509证书是用于进行服务器身份验证操作的
		KeyUsage:x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:[]x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		//该证书只能在ip地址为127.0.0.1之上运行
		IPAddresses:[]net.IP{net.ParseIP("127.0.0.1")},
	}

	pk,_ := rsa.GenerateKey(rand.Reader,2048)

	derBytes,_ := x509.CreateCertificate(rand.Reader,&template,&template,&pk.PublicKey,pk)
	cerOut,_ := os.Create("cert.pem")
	pem.Encode(cerOut,&pem.Block{Type:"CERTIFICATE",Bytes:derBytes})
	cerOut.Close()

	keyOut,_ := os.Create("key.pem")
	pem.Encode(keyOut,&pem.Block{Type:"RSA PRIVATE KEY",Bytes:x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}
