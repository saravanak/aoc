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
        self.bins = [[[] for i in range(1, 1005)] for j in range(1, 1005)]
        self.non_reclaimed_claims = []

    def record_claim(self, claim):
        self.non_reclaimed_claims.append(claim.id)
        
        for i in range(claim.x_start , claim.x_start + claim.width):
            for j in range(claim.y_start, claim.y_start + claim.height):
                if len(self.bins[i][j]) > 0:
                    if claim.id in self.non_reclaimed_claims:
                        self.non_reclaimed_claims.remove(claim.id)
                    for removing_id in self.bins[i][j]:
                        if removing_id in self.non_reclaimed_claims:
                            self.non_reclaimed_claims.remove(removing_id)
                self.bins[i][j].append(claim.id)

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
print(c.non_reclaimed_claims)
