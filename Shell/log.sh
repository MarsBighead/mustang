#!/bin/bash



#==============================================================#
# color log util
#==============================================================#
__CN='\033[0m'    # no color
__CB='\033[0;30m' # black
__CR='\033[0;31m' # red
__CG='\033[0;32m' # green
__CY='\033[0;33m' # yellow
__CB='\033[0;34m' # blue
__CM='\033[0;35m' # magenta
__CC='\033[0;36m' # cyan
__CW='\033[0;37m' # white
function log_info() {  printf "[${__CG} OK ${__CN}] ${__CG}$*${__CN}\n";   }
function log_warn() {  printf "[${__CY}WARN${__CN}] ${__CY}$*${__CN}\n";   }
function log_error() { printf "[${__CR}FAIL${__CN}] ${__CR}$*${__CN}\n";   }
function log_debug() { printf "[${__CB}HINT${__CN}] ${__CB}$*${__CN}\n"; }
function log_input() { printf "[${__CM} IN ${__CN}] ${__CM}$*\n=> ${__CN}"; }
function log_hint()  { printf "${__CB}$*${__CN}"; }

PKG_PATH=$(dirname $0)
log_info "make tarball to ${PKG_PATH}"