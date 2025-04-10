#include <stdio.h>
#include "ud_ucase.h"

#define BLACKLOG 5

int main(int argc, char *argv[]){
    sockaddr_un svaddr, claddr;
    int sfd, j;
    ssize_t numBytes;
    socklen_t len;
    char buf[BUF_SIZE];
    sfd=socket(AF_UNIX, SOCK_DGRAM,0);
    if (sfd==-1)
        errExit("socket");
    
    /* Construct server socket address, bind socket to it, 
    and make this a listening socket */
    if (remove(SV_SOCKET_PATH) == -1 && errno != ENOENT)
        errExit("remove-%s", SV_SOCKET_PATH);
    
    memset(&svaddr,0, sizeof(sockaddr_un));
    svaddr.sun_family = AF_UNIX;
    strncpy(svaddr.sun_path, SV_SOCKET_PATH, sizeof(svaddr.sun_path)-1);
    if (bind(sfd, (struct sockaddr *)&svaddr, sizeof(sockaddr_un)) == -1)
        errExit("bind");
 
    for (;;){
        len = sizeof(sockaddr_un);
        numBytes = recvfrom(sfd, buf, BUF_SIZE,0, 
                            (struct socketaddr*)&claddr, &len);
        if (numBytes == -1)
            errExit("recvfrom");
        
        printf("Server recieved %ld bytes from %s\n", (long)numBytes,
            claddr.sun_path);
        for (j=0;j<numBytes;j++)
            buf[j]=toupper((unsigned char) buf[j]);
        if (sendto(sfd,buf, numBytes,0, (struct sockaddr *)&claddr, len) != numBytes)
            fatal("sendto");
       
    }
    
}
