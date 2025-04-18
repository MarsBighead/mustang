#include <stdio.h>
#include "ud_ucase.h"

int main(int argc, char *argv[]){
    sockaddr_un svaddr, claddr;
    int sfd, j;
    size_t msglen;
    ssize_t numBytes;
    char resp[BUF_SIZE];

    if (argc <2 || strcmp(argv[1], "--help")==0)
        usageErr("%s msg...\n",argv[0]);
    sfd=socket(AF_UNIX, SOCK_DGRAM, 0);
    if (sfd==-1)
        errExit("socket");
    
    /* Construct server socket address, 
    and make the connection */
    memset(&claddr,0, sizeof(sockaddr_un));
    claddr.sun_family = AF_UNIX;
    snprintf(claddr.sun_path, sizeof(claddr.sun_path),
        "/tmp/ud_ucase_cl.%ld",(long)getpid());
    if (bind(sfd, (struct sockaddr *)&claddr, sizeof(sockaddr_un)) == -1)
        errExit("bind");
    printf("Conntect succesfully!\n");
    memset(&svaddr,0, sizeof(sockaddr_un));
    svaddr.sun_family=AF_UNIX;
    strncpy(svaddr.sun_path, SV_SOCKET_PATH, sizeof(svaddr.sun_path)-1);
   
    for (j=1; j<argc;j++){
        msglen =strlen(argv[j]);
        if (sendto(sfd, argv[j], msglen, 0, (struct sockaddr *)&svaddr,
                sizeof(sockaddr_un))!=msglen)
            fatal("sendto");
        numBytes = recvfrom(sfd, resp, BUF_SIZE,0, NULL, NULL);
        if (numBytes == -1)
            errExit("recvfrom");
        printf("Response %d: %.*s\n", j, (int)numBytes, resp);
    }

    
    remove(claddr.sun_path);
    exit(EXIT_SUCCESS); /* close our socker; server sees EOF */
}
