using Fontconfig, Luxor

function buildfontlist()
    fonts = []
    for font in Fontconfig.list()
        families = Fontconfig.format(font, "%{family}")
        for family in split(families, ",")
            push!(fonts, family)
        end
    end
    filter!(f -> !ismatch(r".LastResort|Agenda|Topaz|Bodoni Ornaments|System",
			           f), fonts)
    return sort(unique(fonts))
end

function tabulatepi()
    fonts = buildfontlist()
    ncols = 25
    nrows = convert(Int, ceil(length(fonts))) ÷ ncols
    @svg begin
        background("ivory")
        setopacity(1)
        t = Table(nrows, ncols, 30, 25)
        sethue("black")
        cellnumber = 1
        for n in 1:length(fonts)
            fontface(fonts[n])
            te = textextents("π")
            if te[3] > 0.0
                fontsize(18)
                text("π", t[cellnumber], halign=:center)
                setfont("Lucida-Sans", 3)
                settext(fonts[n], t[cellnumber] + (0, isodd(cellnumber) ? 6 : 10), halign="center")
                cellnumber += 1
            end
        end
    end 800 1200
end

tabulatepi()
