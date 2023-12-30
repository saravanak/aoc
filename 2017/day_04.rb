valid_pass = 0

File.open("day_4_input.txt", "r") do |infile|
    while (line = infile.gets)
      hash = {}
      words = line.split("\s")
      valid = true
      
      words.each do |word|
        puts word
        if hash.has_key?(word)
          valid = false
        else
          hash[word] = 1
        end
        break unless valid
      end
      puts valid
      valid_pass += 1 if valid 
    end
end
puts valid_pass
