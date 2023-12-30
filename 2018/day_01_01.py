file = open("data/day_01_01.txt", "r") 
i = [0];
for line in file: 
    i += int(line)
print (i)
