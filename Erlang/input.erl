-module(input).
-export([start/0]).

start() ->
    io:get_line("").
