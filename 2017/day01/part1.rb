File.open("input.txt", "r") do |infile|
  line = infile.gets.strip.chars.map(&:to_i)
  sum = 0
  prev_num = nil
  line.each do |c|
    sum += c if prev_num == c
    prev_num = c
  end

  if line.first == line.last
    sum += line.first.to_i
  end
  puts(sum)
end
