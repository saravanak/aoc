ch = 0
File.open("input.txt", "r") do |infile|
    while (line = infile.gets)
        numbers = line.split("\t").map(&:to_i)
        ch += numbers.max - numbers.min
    end
end
puts ch
