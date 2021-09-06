#!/bin/bash
 pip3 install virtualenv -i {{pypi_mirrors}} --trusted-host {{pypi_trusted_host}} 
 /usr/local/bin/virtualenv {{venv4archery_dir}} --python=python3
 source {{venv4archery_dir}}/bin/activate
 pip3 install -r /opt/requirements.txt -i {{pypi_mirrors}} --trusted-host {{pypi_trusted_host}} 
 #/opt/venv4archery/bin/python /opt/venv4archery/bin/pip3 install -i https://pypi.mirrors.ustc.edu.cn/simple/ --trusted-host pypi.mirrors.ustc.edu.cn -r /opt/requirements.txt
