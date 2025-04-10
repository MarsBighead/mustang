#include <stdio.h>
#include "i6d_ucase.h"

int main(int argc, char *argv[]){
    sockaddr_in6 svaddr;
    int sfd, j;
    size_t msglen;
    ssize_t numBytes;
    char resp[BUF_SIZE];

    if (argc <3 || strcmp(argv[1], "--help")==0)
        usageErr("%s host-adress msg...\n",argv[0]);
    sfd=socket(AF_INET6, SOCK_DGRAM, 0);
    if (sfd==-1)
        errExit("socket");
    
    /* Construct server socket address, 
    and make the connection */
    memset(&svaddr,0, sizeof(sockaddr_in6));
    svaddr.sin6_family = AF_INET6;
    svaddr.sin6_port=htons(PORT_NUM);
   
    //if (bind(sfd, (struct sockaddr *)&claddr, sizeof(sockaddr_un)) == -1)
    if (inet_pton(AF_INET6, argv[1],&svaddr.sin6_addr)<=0)
        fatal("inet_pton failed for address '%s'", argv[1]);
    printf("Conntect succesfully!\n");
    
    for (j=2; j<argc;j++){
        msglen =strlen(argv[j]);
        if (sendto(sfd, argv[j], msglen, 0, (struct sockaddr *)&svaddr,
                sizeof(sockaddr_in6))!=msglen)
            fatal("sendto");
        numBytes = recvfrom(sfd, resp, BUF_SIZE,0, NULL, NULL);
        if (numBytes == -1)
            errExit("recvfrom");
        printf("Response %d: %.*s\n", j, (int)numBytes, resp);
    }

    
    exit(EXIT_SUCCESS); /* close our socker; server sees EOF */
}
