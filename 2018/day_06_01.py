dim = 500
REGION_CUTOFF = 10000
class Amalgamator:
    def __init__(self):
        self.slate = [[ {'distances': {} } for i in range(0, dim)] for j in range(0, dim)]
        self.y = self.x = 0
        self.coordinates = {}
        
    def debug(self):
        # print(list(map(lambda x: x['nearest_coordinate'], self.slate)))
        print(self.slate)
        
    def reset(self):
        pass
        
    def record_spot(self, id, coordinates):
        (x, y) = list(map(lambda v: int(v), coordinates))
        
        self.x = max(self.x, x)
        self.y = max(self.y, y)
        self.coordinates[id] = (x,y)
        self.slate[x][y]['owner'] = id
        
    def mark_distances(self):
        for owner, owner_coordinate in self.coordinates.items():
            for i in range(0, self.x+1):
                for j in range(0, self.y+1):
                    self.slate[i][j]['distances'][owner] = abs(i - owner_coordinate[0]) + abs(j - owner_coordinate[1])
                        
    def finalize_state(self):
        for i in range(0, dim):
            for j in range(0, dim):
                values = self.slate[i][j]['distances'].values()
                are_all_distances_unique = values is not None and len(values) == len(list(set(values)))
                    
                min_distance  = min(self.slate[i][j]['distances'].items(), key = lambda x: x[1]) if self.slate[i][j]['distances'].items() else ('#', -1)
                total_distance  = sum(self.slate[i][j]['distances'].values())
                
                only_near_to_one_coordinate = len(list(filter(lambda x: x == min_distance[1], self.slate[i][j]['distances'].values()))) == 1
                
                self.slate[i][j]['nearest_coordinate'] = min_distance[0].upper() if only_near_to_one_coordinate else '#'
                self.slate[i][j]['in_region'] = total_distance < REGION_CUTOFF
    
    def count_region(self):
        count = 0
        for i in range(0, self.x+1):
            for j in range(0, self.y+1):
                count += 1 if self.slate[i][j]['in_region'] else 0
        return count
                
    def count_max_area(self):
        coordinate_hash = {}
        for j in range(0, self.x):
            owner = self.slate[0][j]['nearest_coordinate'].lower()
            if owner in self.coordinates:
                self.coordinates.pop(owner)
                
            owner = self.slate[self.x][j]['nearest_coordinate'].lower()
            if owner in self.coordinates:
                self.coordinates.pop(owner)
                
            owner = self.slate[j][self.y]['nearest_coordinate'].lower()
            if owner in self.coordinates:
                self.coordinates.pop(owner)
                
            owner = self.slate[j][0]['nearest_coordinate'].lower()
            if owner in self.coordinates:
                self.coordinates.pop(owner)
                
        for owner, owner_coordinate in self.coordinates.items():
            for i in range(1, self.x):
                for j in range(1, self.y):
                    if self.slate[i][j]['nearest_coordinate'].lower() == owner:
                        if owner in coordinate_hash:
                            coordinate_hash[owner] +=  1
                        else:
                            coordinate_hash[owner] = 1

    def process(self):
        # file = open("data/day_06_01_test.txt", "r") 
        file = open("data/day_06_01.txt", "r") 
        
        import re 
        p = re.compile('(\d+), (\d+)')
        
        index=97
        for line in file:
            m = p.match(line.rstrip())
            self.record_spot('{0}'.format(chr(index)), m.groups())
            index +=1

        self.mark_distances()
        self.finalize_state()
        self.count_max_area()
        count = self.count_region()
        print(count)
                
c = Amalgamator()
c.process()
# c.debug()

#3818  is too high
#2889  is too high
