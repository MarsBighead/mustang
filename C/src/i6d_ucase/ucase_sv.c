#include <stdio.h>
#include "i6d_ucase.h"

#define BLACKLOG 5

int main(int argc, char *argv[]){
    sockaddr_in6 svaddr, claddr;
    int sfd, j;
    ssize_t numBytes;
    socklen_t len;
    char buf[BUF_SIZE];
    char claddrStr[INET6_ADDRSTRLEN];
    sfd=socket(AF_INET6, SOCK_DGRAM,0);
    if (sfd==-1)
        errExit("socket");
    
    /* Construct server socket address, bind socket to it, 
    and make this a listening socket */

    
    memset(&svaddr,0, sizeof(sockaddr_in6));
    svaddr.sin6_family = AF_INET6;
    svaddr.sin6_addr=in6addr_any;
    svaddr.sin6_port=htons(PORT_NUM);
    if (bind(sfd, (struct sockaddr *)&svaddr, sizeof(sockaddr_in6)) == -1)
        errExit("bind");
 
    for (;;){
        len = sizeof(sockaddr_in6);
        numBytes = recvfrom(sfd, buf, BUF_SIZE,0, 
                            (struct socketaddr *)&claddr, &len);
        if (numBytes == -1)
            errExit("recvfrom");
        if (inet_ntop(AF_INET6, &claddr.sin6_addr, claddrStr, 
                INET6_ADDRSTRLEN)==NULL)
            printf("Couldn't convert client address to string\n");
        else
            printf("Server recieved %ld bytes from (%s, %u)\n", 
                (long)numBytes, claddrStr,ntohs(claddr.sin6_port));
        for (j=0;j<numBytes;j++)
            buf[j]=toupper((unsigned char) buf[j]);
        if (sendto(sfd,buf, numBytes,0, (struct sockaddr *)&claddr, len) != numBytes)
            fatal("sendto");
       
    }
    
}
