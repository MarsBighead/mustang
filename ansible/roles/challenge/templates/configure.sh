#!/bin/bash
config_line="{%for line in config_list %}{{line}} {%endfor%}"
cd {{challenge_dir}}/nginx-{{nginx_version}}
./configure $config_line
