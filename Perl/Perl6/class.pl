#!/usr/bin/perl6
use v6;


class Human {
    has $.name;
    has $.age;
    has $.sex;
    has $nationality;
    method get_age{
        $nationality="USA";
        return $nationality;
    }
}

my $john = Human.new(name => 'John',
                     age  => 23,
                     sex  => 'M',
                     nationality => 'American');
say $john;
say $john.get_age;
