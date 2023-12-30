input = "14	0	15	12	11	11	3	5	1	6	8	4	9	1	8	4"

# input =  "0   2   7  0"

array = input.split(" ").map(&:to_i)


def max_with_ties(array)
  index = array.index(array.max)
end

def redistribute_block(array, index)
  block = array[index]
  size = array.length
  array[index] = 0
  while block > 0
    index += 1
    array[index % size] += 1
    block -= 1
  end
  return array
end

hash = {}



steps = 0
processed = array
hash[processed.join(":")] = 1
puts array
while true 
  processed = redistribute_block(processed, max_with_ties(processed))
  if hash.has_key?(processed.join(":"))
    break
  else 
    hash[processed.join(":")] = 1
  end
  puts processed.join(":")
  steps += 1
end
puts steps+1
