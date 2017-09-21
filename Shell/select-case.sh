#!/bin/bash

echo "What is your favourite OS?"

select var in "Linux" "Windows" "Android"
  do
      case $var in
      Linux|Windows|Android)
        os=$var; break;;
      *)
        continue;;
      esac
   done
echo os is $os
