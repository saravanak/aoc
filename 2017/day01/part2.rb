File.open("input.txt", "r") do |infile|
  line = infile.gets.strip.chars.map(&:to_i)
  sum = 0
  length = line.length
  skip_len = length/2
  line.each_with_index do |c, i|
    sum += c if line[(i + skip_len) % length] == c
  end
  puts sum
end
