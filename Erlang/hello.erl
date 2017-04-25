-module(hello).
-export([start/0, input/0]).

start() ->
    F = input(),
	io:format("Hello, World.~n~ts",[F]).

input() ->
	io:get_line("").

