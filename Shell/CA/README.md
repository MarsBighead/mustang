# Make Certificate for Grafana

## Logs

```Shell
[hbu@hbu CA]$ openssl genrsa -aes128 -out grafana.key
Generating RSA private key, 2048 bit long modulus
.....................................................................................+++
............................................................+++
e is 65537 (0x10001)
Enter pass phrase for grafana.key:
140084624123792:error:28069065:lib(40):UI_set_result:result too small:ui_lib.c:831:You must type in 4 to 1023 characters
Enter pass phrase for grafana.key:
Verifying - Enter pass phrase for grafana.key:
[hbu@hbu CA]$ openssl req -new -key grafana.key -out grafana.csr
Enter pass phrase for grafana.key:
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [XX]:cn
State or Province Name (full name) []:
Locality Name (eg, city) [Default City]:Shanghai
Organization Name (eg, company) [Default Company Ltd]:hbu
Organizational Unit Name (eg, section) []:hbu
Common Name (eg, your name or your server's hostname) []:hbu
Email Address []:root@localhost   

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:Paul
An optional company name []:hbu
[hbu@hbu CA]$ open
open            openjade        openoffice.org  openssl         openvt          
[hbu@hbu CA]$ openssl x509 -req -days 3650 -in grafana.csr -signkey grafana.key -out grafana.crt
Signature ok
subject=/C=cn/L=Shanghai/O=hbu/OU=hbu/CN=hbu/emailAddress=root@localhost
Getting Private key
Enter pass phrase for grafana.key:
[hbu@hbu CA]$ openssl rsa -in  grafana.key -out grafana_priv.key
Enter pass phrase for grafana.key:
writing RSA key

```

## CA and Key

Key:  grafana_priv.key

CA: grafana.crt


## Check CA

```Shell
[hbu@hbu CA]$ openssl x509 -text -in  grafana.crt  -noout
Certificate:
    Data:
        Version: 1 (0x0)
        Serial Number: 12754839385983080749 (0xb1024ad7f0aed92d)
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: C=cn, L=Shanghai, O=hbu, OU=hbu, CN=hbu/emailAddress=root@localhost
        Validity
            Not Before: Jan 30 12:03:14 2019 GMT
            Not After : Jan 27 12:03:14 2029 GMT
        Subject: C=cn, L=Shanghai, O=hbu, OU=hbu, CN=hbu/emailAddress=root@localhost
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:cb:a3:54:cc:08:1b:89:33:d1:70:6a:64:7d:2d:
                    a0:fd:72:1e:b6:a2:5f:7b:45:4d:0d:d2:a5:73:09:
                    a0:e2:37:fa:7c:02:8f:46:81:a0:8e:f2:fa:39:38:
                    fe:8c:8d:d4:1b:f2:f2:24:17:35:a4:ed:46:0e:e7:
                    42:20:05:24:5e:73:1c:54:fc:00:8c:ec:05:5d:dc:
                    5d:17:4a:51:52:47:c4:9c:04:3f:ed:d7:46:eb:d3:
                    4b:cc:d8:1e:5e:2e:84:6c:37:de:9e:16:30:d0:71:
                    43:04:ec:3d:21:7a:ed:33:31:4e:1c:b7:82:f5:ba:
                    97:08:cf:39:07:ef:49:37:16:13:f4:9b:ae:6d:fe:
                    3c:27:26:4b:62:a6:0e:d6:33:89:24:90:6a:3c:4e:
                    15:79:3f:8a:07:83:aa:6e:52:0a:55:23:bf:f4:b3:
                    6e:82:9b:4a:95:50:ad:93:ff:bf:5c:70:36:19:ae:
                    40:8a:56:7e:e4:ad:0b:72:c3:a2:34:ea:c6:a4:c6:
                    2e:27:5f:dd:08:7e:39:2e:9f:ec:ec:c8:c3:4a:a5:
                    91:99:93:7a:9a:e3:6b:3d:93:b8:67:18:8c:f4:b0:
                    8c:d8:68:db:45:10:18:09:f7:9b:b0:3e:37:51:bd:
                    79:15:ed:65:c1:94:4a:a3:dc:6c:7a:28:ed:c5:65:
                    58:9b
                Exponent: 65537 (0x10001)
    Signature Algorithm: sha256WithRSAEncryption
         6e:77:e5:dc:c2:73:af:ff:5e:5f:a0:e3:7d:f2:ea:5c:03:d3:
         e8:7e:2f:37:49:c9:e0:97:62:b1:88:0a:5c:fd:6f:4b:98:e7:
         09:3a:2b:71:68:0c:51:43:c0:79:84:f4:9a:7c:e0:60:3b:7f:
         2c:69:d8:4e:d2:52:e4:a3:f5:75:81:dd:db:4b:ef:44:d5:44:
         1d:0a:69:95:28:b5:b9:91:ab:85:88:15:da:5f:25:e4:e9:54:
         19:4b:3e:da:83:29:94:37:25:1c:8f:04:94:bb:8e:31:76:2e:
         36:96:92:0e:37:2a:c7:05:49:eb:69:6d:c8:d5:75:72:1a:b6:
         2d:a3:c1:50:71:cb:12:d9:d5:2b:d6:c8:3a:a7:d6:ea:3a:b9:
         65:ea:49:3a:6d:a8:86:21:41:5c:70:43:ff:f0:44:79:dd:cc:
         b5:3c:f0:7f:36:ca:34:bd:a0:99:01:7a:aa:f2:22:09:14:5f:
         f7:0f:02:c6:44:d2:cd:7d:70:3d:b2:64:ea:32:e5:f6:0b:d8:
         07:ed:2e:9a:d9:43:05:26:3f:43:f9:1b:30:15:fe:a8:10:14:
         5b:c1:af:65:3c:63:1e:66:d7:13:d1:89:00:77:e5:9d:3a:a0:
         6f:13:52:7e:46:5d:55:f1:23:0e:ff:bb:04:d2:b7:93:59:87:
         df:66:42:6c
```