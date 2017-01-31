-module(mylists).
-export([map/2]).
-export([sum/1]).

map(_, [])	->[];
map(F, [H|T])	->[F(H)|map(F, T)].

sum([H|T]) -> H + sum(T);
sum([])	   -> 0.
