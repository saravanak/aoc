ch = 0
File.open("input.txt", "r") do |infile|
    while (line = infile.gets)
        numbers = line.split(" ").map(&:to_i)
        all_combinations = numbers.permutation(2).to_a
        divisors = all_combinations.find { |x| x[1] % x[0] == 0 }
        q = divisors[1] / divisors[0]
        ch += q
    end
end
puts ch
