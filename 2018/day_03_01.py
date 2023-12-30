class Claim:
    def __init__(self, claim):
        (self.id, self.left_bound, self.top_bound, self.width, self.height) = claim
        self.x_start = self.left_bound + 1
        self.x_end = self.x_start + self.width - 1
        
        self.y_start = self.top_bound + 1
        self.y_end = self.y_start + self.height -1
        
    def __str__(self):
        return "id: {4}, l: {0}, t: {1}, w: {2}, h: {3}".format(self.left_bound, self.top_bound, self.width, self.height, self.id)


class ClaimProcessor:
    def __init__(self):
        self.bins = [[0 for i in range(1, 1002)] for j in range(1, 1002)]

    def record_claim(self, claim):
        for i in range(claim.x_start , claim.x_start + claim.width):
            for j in range(claim.y_start, claim.y_start + claim.height ):
                self.bins[i][j] += 1

    def process(self):
        # file = open("data/day_03_01_test.txt", "r") 
        # file = open("data/day_03_01_sample.txt", "r") 
        file = open("data/day_03_01.txt", "r") 
        import re 
        p = re.compile('#(\d+) @ (\d+),(\d+): (\d+)x(\d+)')

        for line in file: 
            m = p.match(line)
            c = Claim(map(lambda x: int(x), m.groups()))
            self.record_claim(c)


                
c = ClaimProcessor()
c.process()
occupied = 0
print('RESULTS')
for rows in c.bins:
    claimed_more_than_once = list(filter(lambda x: x > 1, rows))
    occupied += len(claimed_more_than_once)
print(occupied)
    
occupied = 0
for rows in c.bins:
    claimed_more_than_once = list(filter(lambda x: x == 1, rows))
    occupied += len(claimed_more_than_once)

print(occupied)

 
