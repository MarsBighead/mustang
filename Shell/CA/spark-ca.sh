#!/bin/bash

openssl genrsa -out spark-client-ca.key 2048
openssl req -x509 -new -nodes -key spark-client-ca.key -sha256 -days 3652 -out spark-client-ca.pem

echo "##Execute build client-keystore.p12"
openssl pkcs12 -export -name client-cert -in spark-client-ca.pem -inkey spark-client-ca.key -out clientkeystore.p12

echo "##Import keystore client-keystore.p12"
keytool -importkeystore -destkeystore client.keystore \
        -srckeystore clientkeystore.p12 -srcstoretype pkcs12 \
        -alias client-cert
#
#keytool -import -alias server-cert -file spark-server-ca.pem \
#        -keystore client.truststore

keytool -importkeystore -srckeystore client.keystore -destkeystore client.keystore -deststoretype pkcs12

keytool -import -alias client-cert -file spark-client-ca.pem \
        -keystore client.truststore
