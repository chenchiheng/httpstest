# CA根证书制作   
From: https://www.zhangwenbing.com/blog/linux/JQIo7dGc1

## 1.生成根证书密钥
openssl genrsa -out ca.key 4096 #建议长度为4096,1024长度已经被列为不安全。

## 2.生成自签名根证书
这里在输入CommonName的时候输入名称而不是域名或者ip,就是证书上显示的颁发者，虽然是自签名证书，但是尽量让证书看起来标准一些
openssl req -new -x509 -days 3650 -key ca.key -out ca.crt 

# 使用根证书来签名其它证书
## 1.生成证书密钥
openssl genrsa -out server.key 4096

## 2. Subject Alt Name(SAN)
高版本的Chrome浏览器会要求设置subjectAltName,如果没有设置SAN会报证书错误

参考openssl配置文件,Linux服务器上通常在/etc/pki/tls/openssl.cnf

新建文件san.conf

[req]
default_bits = 4096
distinguished_name = req_distinguished_name
req_extensions = v3_req
 
[req_distinguished_name]
countryName = country
stateOrProvinceName = province
localityName = city
organizationName = company name
commonName = domain name or ip
 
[v3_req]
subjectAltName = @alt_names

[alt_names]
DNS.1=domain #可以使用通配符
IP.1=xxx.xxx.xxx.xxx

## 3. 生成证书签名请求(CSR)
向根证书请求签名一个新的证书，由于用户信任了你的根证书，所以根证书签名的其它证书也会被信任
生成csr 注意要使用sha256算法（推荐是sha256算法，默认算法浏览器会报弱加密算法错误）
openssl req -new -key server.key -out server.csr -config san.conf -sha256

查看csr信息
openssl req -text -in server.csr

csr信息中会有类似的信息
Requested Extensions:
     X509v3 Subject Alternative Name: 
      IP Address:xxxxxx

## 4.使用根证书按照csr给证书签名，生成新证书server.crt
openssl x509 -req -days 365 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt -extfile san.conf -extensions v3_req

查看证书信息
openssl x509 -text -in server.crt

证书信息中会有类似信息
X509v3 extensions:
    X509v3 Subject Alternative Name: 
           IP Address:xxxxxx
