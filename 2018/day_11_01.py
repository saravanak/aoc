import pprint 
from turtle import Turtle

import statistics 

pp = pprint.PrettyPrinter(indent=4)

class FuelCellExaminer:
    def __init__(self):
        # self.cells = [['{0},{1}'.format(i, j) for i in range(0, 300)] for j in range(0, 300)]
        self.cells = [[0 for i in range(0, 300)] for j in range(0, 300)]
        self.power_levels = {}

    def debug(self):
        print('\n'.join([' '.join([str(cell) for cell in row]) for row in self.cells]))
        
    def calc_total_power(self, x, y, size):
        power_level = 0
        for rx in range(x, x+size):
            for ry in range(y, y+size):
                power_level += self.cells[rx][ry]
                
        self.power_levels['{0},{1},{2}'.format(x+1,y+1,size+1)]  = power_level        
            
    def find_max(self):
        # for size in range(2, 300):
        size = 3
        for y in range(0, 300-size):
            for x in range(0, 300-size):
                self.calc_total_power(x, y, size)
            
    def process(self, serial_number):
        for i in range(0, 300):
            for j in range(0, 300):
                rack_id = (j+1) + 10
                power_level = rack_id * (i+1)
                power_level += serial_number
                power_level *= rack_id
                power_level = power_level % 1000 // 100
                power_level -= 5
                self.cells[j][i] = power_level
                
        self.find_max()
        print(max(self.power_levels.items(), key=lambda x: x[1]))

c = FuelCellExaminer() 
c.process(3463)
# c.debug()
