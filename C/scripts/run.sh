#!/bin/bash

CURRENT_DIR=$(cd $(dirname $0);cd ../;pwd)
APPLICATION=${1:-""}
if [ -z $APPLICATION ];then
    echo "Input empty application name, exit..."
    exit
fi

if [ $(uname) == "Darwin" ];then
    export DYLD_LIBRARY_PATH=$DYLD_LIBRARY_PATH:$CURRENT_DIR/src/lib:$CURRENT_DIR/src/common
elif [ $(uname) == "Linux" ];then 
    export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:$CURRENT_DIR/src/lib:$CURRENT_DIR/src/common
fi
pwd
param=${@:2:3}
echo "parameter: $param"
if [[ $APPLICATION == "xfr_sv" ]];then
    echo "src/us_xfr/xfr_sv > $param"
    src/us_xfr/xfr_sv > $param &
elif [[ $APPLICATION == "xfr_cl" ]];then
    src/us_xfr/xfr_cl < src/us_xfr/$param

else
    echo "Unsupport application"
fi
