import pprint 
from turtle import Turtle

import statistics 

pp = pprint.PrettyPrinter(indent=4)

class FuelCellExaminer:
    def __init__(self):
        self.cells = [[{'power': 0, 'powers': {}} for i in range(0, 300)] for j in range(0, 300)]
        self.power_levels = {}

    def debug(self):
        print('\n'.join([' '.join([str(cell) for cell in row]) for row in self.cells]))
        
    def calc_total_power(self, x, y, size):
        power_level = 0
        
        for rx in range(x, x+size):
            for ry in range(y, y+size):
                power_level += self.cells[rx][ry]['power']
                
        self.power_levels['{0},{1},{2}'.format(x+1,y+1, size)]  = power_level
        self.cells[x][y]['powers'][size] = power_level
            
    def find_max(self):
        for y in range(0, 299): 
            for x in range(0, 299): 
                self.calc_total_power(x, y, 2)
                
        for size in range(2, 299):
            print('finding size', size+1)
            if (size+1) % 2 == 0:
                component_power_level = (size + 1) // 2 
                
                for y in range(0, 300 - size):
                    for x in range(0, 300 - size):
                        fuel_box = [self.cells[x][y], 
                                    self.cells[x+component_power_level][y], 
                                    self.cells[x][y+component_power_level], 
                                    self.cells[x+component_power_level][y+component_power_level]]
                        
                        power_level = 0
                        for cell in fuel_box:
                            power_level += cell['powers'][component_power_level]
                            
                        self.cells[x][y]['powers'][size+1] = power_level    
                        self.power_levels['{0},{1},{2}'.format(x+1,y+1,size+1)]  = power_level
            else:
                # *  *  ry
                # *  *  ry
                # rx rx rx
                 
                # x = 32        
                # y = 44
                for y in range(0, 300 - size):
                    for x in range(0, 300 - size):
                        
                        component_power_level = size 
                        power_level = self.cells[x][y]['powers'][component_power_level]
                        # print(power_level)
                        
                        for rx in range(x, x+size+1):
                            # print('+', self.cells[rx][y+size]['power'], rx, y+size)
                            power_level += self.cells[rx][y+size]['power']
                        
                        for ry in range(y, y+size):
                            # print('+', self.cells[x+size][ry]['power'])
                            power_level += self.cells[x+size][ry]['power']
                        # print(power_level)
            
                        self.cells[x][y]['powers'][size+1] = power_level    
                        self.power_levels['{0},{1},{2}'.format(x+1,y+1,size+1)]  = power_level
            
    def process(self, serial_number):
        for i in range(0, 300):
            for j in range(0, 300):
                rack_id = (j+1) + 10
                power_level = rack_id * (i+1)
                power_level += serial_number
                power_level *= rack_id
                power_level = power_level % 1000 // 100
                power_level -= 5
                self.cells[j][i]['power'] = power_level
                
        self.find_max()
        print(max(self.power_levels.items(), key=lambda x: x[1]))

c = FuelCellExaminer() 
c.process(3463)
# c.debug()
