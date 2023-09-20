# rm *.pem

# #1. Generate CA's private key and self-signed certificate
# openssl req -x509 -newkey rsa:4096 -days 356 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=MM/ST=Yangon/L=NorthDagon/O=MTM/OU=IT/CN=*.metateammyanmar.org/emailAddress=scm-thukhaaung@gmail.com"

# echo "CA's self-signed certificate"
# openssl x509 -in ca-cert.pem -noout -text


# # cat ca-key.pem   (open file in terminal)
# #C = MM, ST = Yangon, L = NorthDagon, O = MTM, OU = IT, CN = *.metateammyanmar.org, emailAddress = scm-thukhaaung@gmail.com

# #-subj "/C=MM/ST=Yangon/L=NorthDagon/O=MTM/OU=IT/CN=*.metateammyanmar.org/emailAddress=scm-thukhaaung@gmail.com
# #openssl x509 -in ca-cert.pem -noout -text (open file in terminal with no output file and show text format)

# #2. Generate web server's private key and certificate singing request (CSR) #generate private key and csr
# openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=MM/ST=Yangon/L=Botadaung/O=MTM-Company/OU=IT-Server/CN=*.server.metateammyanmar.org/emailAddress=scm.thukhaaung.server@gmail.com"

# # 3. Use CA's private key to sign web server's CSR and get back the signed certificate
# openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf

# echo "Server's signed certificate"
# openssl x509 -in server-cert.pem -noout -text

# # 4. Generate client's private key and certificate signing request (CSR)
# openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem -subj "/C=MM/ST=YangonClient/L=BotadaungClient/O=MTM-Company-Client/OU=IT-Client/CN=*.client.metateammyanmar.org/emailAddress=scm.thukhaaung.client@gmail.com"

# # 5. Use CA's private key to sign client's CSR and get back the signed certificate
# openssl x509 -req -in client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem 
# echo "Client's signed certificate"
# openssl x509 -in client-cert.pem -noout -text


#############################################

# rm *.pem
# #CA means certification authority
# # 1. Generate CA's private key and self-signed certificate
# openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=MM/ST=Occitanie/L=Toulouse/O=Tech School/OU=Education/CN=*.techschool.guru/emailAddress=techschool.guru@gmail.com"

# echo "CA's self-signed certificate"
# openssl x509 -in ca-cert.pem -noout -text

# # 2. Generate web server's private key and certificate signing request (CSR)
# openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=MM/ST=Ile de France/L=Paris/O=PC Book/OU=Computer/CN=*.pcbook.com/emailAddress=pcbook@gmail.com"

# # 3. Use CA's private key to sign web server's CSR and get back the signed certificate
# openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf
# # openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem

# echo "Server's signed certificate"
# openssl x509 -in server-cert.pem -noout -text

# # 4. Generate client's private key and certificate signing request (CSR)
# openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem -subj "/C=MM/ST=Alsace/L=Strasbourg/O=PC Client/OU=Computer/CN=*.pcclient.com/emailAddress=pcclient@gmail.com"

# # 5. Use CA's private key to sign client's CSR and get back the signed certificate
# openssl x509 -req -in client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile client-ext.cnf
# # openssl x509 -req -in client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem

# echo "Client's signed certificate"
# openssl x509 -in client-cert.pem -noout -text

#############################################

rm *.pem
#CA means certification authority
# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=MM/ST=Yangon/L=Botahtaung/O=MTM/OU=ITCompany/CN=*.thukhaaung.com/emailAddress=scm.thukhaaung@gmail.com"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=MM/ST=Yangon/L=Botahtaung/O=MTM/OU=ITCompany/CN=*.thukhaaung.com/emailAddress=scm.thukhaaung@gmail.com"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf
# openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem

echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text

# 4. Generate client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem -subj "/C=MM/ST=Yangon/L=Botahtaung/O=MTM/OU=ITCompany/CN=*.thukhaaung.com/emailAddress=scm.thukhaaung@gmail.com"

# 5. Use CA's private key to sign client's CSR and get back the signed certificate
openssl x509 -req -in client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile client-ext.cnf
# openssl x509 -req -in client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem

echo "Client's signed certificate"
openssl x509 -in client-cert.pem -noout -text