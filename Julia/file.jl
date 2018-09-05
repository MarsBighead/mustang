function read_and_capitalize(f::IOStream)
    ss = uppercase(read(f, String))
    println(ss)
    return ss
end
f = open("hello.txt","w")
write(f,"Hi, MarsBighead")
close(f)
f = open("hello.txt")
read_and_capitalize(f)
close(f)
