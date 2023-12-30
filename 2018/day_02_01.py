def is_exactly_n_of_any_letter(frequency_counter, n):
    return 1 if len(list(filter(lambda x: x == n, frequency_counter.values()))) > 0 else 0

# file = open("data/day_02_01_test.txt", "r") 
file = open("data/day_02_01.txt", "r") 
count_twos = 0
count_threes = 0
for line in file: 
    frequency_counter = {}
    
    for c in line.rstrip():
        if frequency_counter.get(c):
            frequency_counter[c] += 1
        else: 
            frequency_counter[c] = 1
            
    count_twos += is_exactly_n_of_any_letter (frequency_counter, 2)
    count_threes += is_exactly_n_of_any_letter (frequency_counter, 3)

print (count_twos * count_threes)
