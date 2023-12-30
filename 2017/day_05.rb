valid_pass = 0
def process_instruction(array, ip)
  move_by  = array[ip]
  array[ip] += 1
  return {array: array, ip: ip + move_by}  
end

File.open("day_5_input.txt", "r") do |infile|
    array = []
    while (line = infile.gets)
      array << line.to_i
    end
    ip = 0
    steps = 0
    while ip < array.length
      return_val = process_instruction(array, ip)
      ip = return_val[:ip]
      array = return_val[:array]
      steps += 1
    end

    puts steps
end
