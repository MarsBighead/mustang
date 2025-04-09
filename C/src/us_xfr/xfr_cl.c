#include <stdio.h>
#include "us_xfr.h"

int main(int argc, char *argv[]){
    sockaddr_un addr;
    int sfd;
    ssize_t numRead;
    char buf[BUF_SIZE];
    sfd=socket(AF_UNIX, SOCK_STREAM, 0);
    if (sfd==-1)
        errExit("socket");
    
    /* Construct server socket address, 
    and make the connection */
    memset(&addr,0, sizeof(sockaddr_un));
    addr.sun_family = AF_UNIX;
    strncpy(addr.sun_path, SV_SOCKET_PATH, sizeof(addr.sun_path)-1);
    if (connect(sfd, (struct sockaddr *)&addr, sizeof(sockaddr_un)) == -1)
        errExit("connect");
    printf("Conntect succesfully!\n");
   
    while ((numRead = read(STDIN_FILENO, buf, BUF_SIZE)) > 0)
        if (write(sfd, buf, numRead) != numRead)
            fatal("partial/failed write");
        printf("loop numRead=%ld\n", numRead);
    printf("numRead=%ld\n", numRead);

    if (numRead == -1)
        errExit("read");
    exit(EXIT_SUCCESS); /* close our socker; server sees EOF */
}
