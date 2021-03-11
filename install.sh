#!/usr/bin/bash
if [ `whoami` != 'root' ]
  then
    echo -e "\e[91mroot privilege required."
    exit
fi
PATH=$PATH:/usr/local/go/bin
go version > /dev/null 2>&1
if [ $? -ne 0 ]; then
   echo -e "\e[31mGo not installed , please install it and run the installer again.\e[0m"
   exit
fi


go build  -o /usr/local/bin/d3pen

echo -e "\e[32mDone!"
