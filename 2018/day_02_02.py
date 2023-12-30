def is_differ_by_one_place(lhs, checksums_list):
    for i, checksum in enumerate(checksums_list):
        number_of_ones =  bin(checksum ^ lhs).count('1')
        if number_of_ones == 2:
            return checksum
    return None

from functools import reduce

def aggregate(accumulator, character):
    if accumulator.get(character):
        accumulator[character] += 1
    else:
        accumulator[character] = 1
    return accumulator

# file = open("data/day_02_02_test.txt", "r") 
# file = open("data/day_02_02_sample.txt", "r") 
file = open("data/day_02_01.txt", "r") 
line_checksum_container = {}
for line in file: 
    frequency_counter = {}
    trimmed_line = line.rstrip()
    
    frequency_counter = {}
    
    for x in list(range(1, 27)):
        frequency_counter[chr(96+x)] = 0
    
    line_checksum = reduce(aggregate, trimmed_line, frequency_counter)
    
    hash_key = int(''.join(map(lambda x: bin(17+x)[2:], line_checksum.values())), 2)
    
    if line_checksum_container.get(hash_key):
        raise "error"
    else:
        line_checksum_container[hash_key] = trimmed_line

checksums = line_checksum_container.keys()
for (k, v)  in line_checksum_container.items():
    result = is_differ_by_one_place(k, checksums)
    if result is not None:
        print(line_checksum_container[result], v)


# Nice problem!
# We are getting a false positive here, but that's ok for now..
# We should encode the order in the `bin(17+x)[2:]`.

# wxlnjevkbodamyiqpufcrhstkg bxlnjevbfwdamyiqpuocrhstkg
# wmlnjevbfodamyiqpuecrhsukg wmlnjevbfodamyiqpuycrhsukg
# wmlnjevbfodamyiqpuycrhsukg wmlnjevbfodamyiqpuecrhsukg
# bxlnjevbfwdamyiqpuocrhstkg wxlnjevkbodamyiqpufcrhstkg
