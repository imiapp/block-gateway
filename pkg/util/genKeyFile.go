package util  
  
import (  
    "crypto/rsa"  
    "crypto/x509"  
    "encoding/pem"  
    "crypto/rand" 
//    "io/ioutil" 
    "bufio"
    "bytes"
    "io"
    "strings"
    "flag"  
    "log"  
    "os"  
)  
  
func main() {  
    var bits int  
    flag.IntVar(&bits, "b", 2048, "密钥长度，默认为1024位")  
    if privatekeyStr,publickeyStr,err := GenRsaKey(bits,"",""); err != nil {  
        log.Fatal("密钥文件生成失败！")  
        log.Println("privatekeyStr:",privatekeyStr)
        log.Println("publickeyStr:",publickeyStr)
    }  
    log.Println("密钥文件生成成功！")  
}  
  
func GenRsaKey(bits int, pathName string,userName string) (privateKeyStr string,publicKeyStr string,_ error) {  
    // 生成私钥文件  
    privateKey, err := rsa.GenerateKey(rand.Reader, bits)  
    if err != nil {  
        return "","",err  
    }  
    derStream := x509.MarshalPKCS1PrivateKey(privateKey)
    block := &pem.Block{  
        Type:  "private key",  
        Bytes: derStream,  
    }  

    //TODO create path of data
    if(!CreateFilePath(pathName)){
        return "","",err
    }
    var privateKeyFileName = pathName + "private_"+userName+".pem"
    log.Println("------privateKeyFileName:",privateKeyFileName)

    file, err := os.Create(privateKeyFileName)  
    if err != nil {  
        return "","",err  
    }
    err = pem.Encode(file, block)  
    if err != nil {  
        return "","",err  
    }  
    
    privateKeyPem := handlerPem(privateKeyFileName)
    log.Println("privateKeyPem=",privateKeyPem)
    
    // 生成公钥文件  
    publicKey := &privateKey.PublicKey  
    derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
    if err != nil {  
        return "","",err  
    }  
    block = &pem.Block{  
        Type:  "public key",  
        Bytes: derPkix,  
    }  

    var publicKeyFileName = pathName + "public_"+userName+".pem"
    log.Println("--------publickKeyFileName:",publicKeyFileName)

    file, err = os.Create(publicKeyFileName)  
    if err != nil {  
        return "","",err  
    }  
    err = pem.Encode(file, block)  
    if err != nil {  
        return "","",err  
    }  
    publicKeyPem := handlerPem(publicKeyFileName)
    log.Println("publicKeyPem=",publicKeyPem)
    return string(privateKeyPem), string(publicKeyPem), nil
}  

func handlerPem(fileName string) (content string){
	fi, err := os.Open(fileName)
    if err != nil {
        log.Printf("Error: %s\n", err)
        return
    }
    defer fi.Close()

    br := bufio.NewReader(fi)
    var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
    for {
        a, _, c := br.ReadLine()
        if c == io.EOF {
            break
        }
        temp := string(a) 
        if(strings.Contains(temp,"---")){
            continue
        }
        buffer.WriteString(temp)
    }
        log.Println("file's content=",buffer.String())
        return buffer.String()
}
