#!/bin/bash

  
  
  
username="hbu"  
password="hbu"  
host="192.168.8.129"  
 

function ssh_cmd(){
    port=22
    expect -c "
	    spawn ssh -p $port $1 $3;  
        expect {
            \"yes/no\" {send \"yes\r\"; exp_continue;}
            \"*password:*\" {send \"$2\r\"; }
        }   
    expect eof;"
    return $?
}

ssh_cmd "$username@$host" $password   ls -lth
