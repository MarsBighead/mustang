#include <stdio.h>
#include "us_xfr.h"

#define BLACKLOG 5

int main(int argc, char *argv[]){
    sockaddr_un addr;
    int sfd, cfd;
    ssize_t numRead;
    char buf[BUF_SIZE];
    sfd=socket(AF_UNIX, SOCK_STREAM,0);
    if (sfd==-1)
        errExit("socket");
    
    /* Construct server socket address, bind socket to it, 
    and make this a listening socket */
    if (remove(SV_SOCKET_PATH) == -1 && errno != ENOENT)
        errExit("remove-%s", SV_SOCKET_PATH);
    
    memset(&addr,0, sizeof(sockaddr_un));
    addr.sun_family = AF_UNIX;
    strncpy(addr.sun_path, SV_SOCKET_PATH, sizeof(addr.sun_path)-1);
    if (bind(sfd, (struct sockaddr *)&addr, sizeof(sockaddr_un)) == -1)
        errExit("bind");
    if (listen(sfd, BLACKLOG) == -1)
        errExit("listen");
    for (;;){
        cfd = accept(sfd, NULL, NULL);
        if (cfd == -1)
            errExit("accept");
        while ((numRead = read(cfd, buf, BUF_SIZE))>0)
            if (write(STDOUT_FILENO, buf, numRead) != numRead)
                fatal("partial/failed write");
            printf("Server loop numRead=%ld\n", numRead);
        printf("numRead=%ld\n", numRead);
        if (numRead == -1)
            errExit("read");
        if (close(cfd) == -1)
            errMsg("close");
    }
    
}
