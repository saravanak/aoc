class Amalgamator:
    def __init__(self):
        self.polymer = []
        
    def debug(self):
        pass
        
    def last_unit(self):
        if len(self.polymer) < 1:
            return '.'
            
        return self.polymer[len(self.polymer) - 1]    
        
    def add_unit(self, unit):
        is_self_cancelling_units = abs(ord(self.last_unit()) - ord(unit)) == 32
        if is_self_cancelling_units:
            self.polymer.pop()
        else:
            self.polymer.append(unit)
    
    def reset(self):
        self.polymer = []

    def process(self):
        # file = open("data/day_05_01_test.txt", "r") 
        file = open("data/day_05_01.txt", "r") 
        
        line = file.readline().rstrip()
        
        for index, unit in enumerate(line):
            self.add_unit(unit)
            
        import re
        for removed_unit_id in range(97, 97 + 26):
            removed_unit = chr(removed_unit_id)
            new_line = re.sub(removed_unit, '', line, flags=re.I)
            
            self.reset()
            for index, unit in enumerate(new_line):
                self.add_unit(unit)
        
            print(removed_unit, len(self.polymer))
#11366        
        
c = Amalgamator()
c.process()
# print(''.join(c.polymer))
print(len(c.polymer))
