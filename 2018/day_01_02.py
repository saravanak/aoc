file = open("data/day_01_01.txt", "r") 
change_list = []
for line in file: 
    change_list.append(int(line))

# change_list = [7, 7, -2, -7, -4]
seen_same_frequency = False
frequency_map = {}
running_frequency = 0
iteration = 0
while not seen_same_frequency:
    for change in change_list:
        running_frequency += change
        if frequency_map.get(running_frequency):
            seen_same_frequency = True
            break
        else:
            frequency_map[running_frequency] = True
    iteration+=1    
            
print (running_frequency, iteration)
