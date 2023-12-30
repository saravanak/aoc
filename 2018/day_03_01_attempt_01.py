class Claim:
    def __init__(self, claim):
        (self.id, self.left_bound, self.top_bound, self.width, self.height) = claim
        self.x_start = self.left_bound + 1
        self.x_end = self.x_start + self.width - 1
        
        self.y_start = self.top_bound + 1
        self.y_end = self.y_start + self.height -1
        
        self.x_picker = self.x_start // 30 or 1
        self.y_picker = self.y_start // 30 or 1
        
    def within_30_by_30(self):
        if self.width > 30 or self.height > 30: 
            return False
        
        x_inches_needed = 30 - self.left_bound % 30
        if self.width > x_inches_needed:
            return False
            
        y_inches_needed = 30 - self.top_bound % 30
        if self.height > y_inches_needed:
            return False
            
        return True
        
    def sub_divide(self):
        """
        Divide a claim to fall within the defined square
        
        >>> c = Claim((1, 4, 5, 30, 30))
        >>> parts = c.sub_divide()
        >>> str(parts[0])
        'id: 1a, l: 4, t: 5, w: 25, h: 24'
        >>> str(parts[1])
        'id: 1b, l: 30, t: 5, w: 5, h: 24'
        >>> str(parts[2])
        'id: 1c, l: 4, t: 30, w: 25, h: 6'
        >>> str(parts[3])
        'id: 1d, l: 30, t: 30, w: 5, h: 6'
        
        
        Should divide a claim when bounds are greater than 30
        
        >>> c = Claim((1, 134, 135, 30, 30))
        >>> parts = c.sub_divide()
        >>> str(parts[0])
        'id: 1a, l: 134, t: 135, w: 15, h: 14'
        >>> str(parts[1])
        'id: 1b, l: 150, t: 135, w: 15, h: 14'
        >>> str(parts[2])
        'id: 1c, l: 134, t: 150, w: 15, h: 16'
        >>> str(parts[3])
        'id: 1d, l: 150, t: 150, w: 15, h: 16'
        
        When the claim cannot be divided horizontally
        
        >>> c = Claim((1, 4, 5, 25, 30))
        >>> parts = c.sub_divide()
        >>> len(parts)
        2
        >>> str(parts[0])
        'id: 1a, l: 4, t: 5, w: 25, h: 24'
        >>> str(parts[1])
        'id: 1c, l: 4, t: 30, w: 25, h: 6'
        
        When the claim cannot be divided vertically
        
        >>> c = Claim((1, 4, 5, 30, 24))
        >>> parts = c.sub_divide()
        >>> len(parts)
        2
        >>> str(parts[0])
        'id: 1a, l: 4, t: 5, w: 25, h: 24'
        >>> str(parts[1])
        'id: 1b, l: 30, t: 5, w: 5, h: 24'
        
        Handle the recursive case
        
        >>> c = Claim((1, 549, 420, 26, 27))
        >>> parts = c.sub_divide()
        >>> len(parts)
        2
        >>> str(parts[0])
        'id: 1a, l: 549, t: 420, w: 20, h: 27'
        >>> str(parts[1])
        'id: 1b, l: 570, t: 420, w: 6, h: 27'
        """
        parts = []
        
        #x half
        x_inches_needed = 30 - self.left_bound % 30 - 1
        x_is_cut_possible = self.width > x_inches_needed
        x_cut_at = x_inches_needed if x_is_cut_possible else self.width 
        
        #y half
        y_inches_needed = 30 - self.top_bound % 30 - 1
        y_is_cut_possible = self.height > y_inches_needed
        y_cut_at = y_inches_needed if y_is_cut_possible else self.height 
        
        parts.append(Claim((str(self.id) + 'a', self.left_bound, self.top_bound, x_cut_at, y_cut_at)))
        
        x_start = self.x_start + x_cut_at 
        y_start = self.y_start + y_cut_at 
        
        if x_is_cut_possible:
            parts.append(Claim((str(self.id) + 'b', x_start, self.top_bound,  self.width - x_inches_needed , y_cut_at)))
            
        if y_is_cut_possible:
            parts.append(Claim((str(self.id) + 'c', self.left_bound, y_start, x_cut_at, self.height - y_inches_needed)))
            
        if y_is_cut_possible and x_is_cut_possible:
            parts.append(Claim((str(self.id) + 'd', x_start , y_start, self.width - x_inches_needed, self.height - y_inches_needed)))
        
        return parts
        
    def get_board(self):
        this_board = list(map(lambda x: '0', list(range(1, 905))))
        
        # print(str(self), self.y_start, self.height)
        box_y_start = self.y_start % 30
        box_y_start = box_y_start or 30
        
        box_x_start = self.x_start % 30
        box_x_start = box_x_start or 30
        
        # print(box_y_start, box_y_start + self.height)
        for j in range(box_y_start, box_y_start + self.height ):
            for i in range(box_x_start , box_x_start + self.width):
                # print(i, j)
                this_board[((j-1)*30) + i + 1] = '1'
        
        return this_board
        
    def __str__(self):
        return "id: {4}, l: {0}, t: {1}, w: {2}, h: {3}".format(self.left_bound, self.top_bound, self.width, self.height, self.id)


def empty_board():
    return list(map(lambda x: '0', list(range(1,905))))
    
class ClaimProcessor:
    def __init__(self):
        self.bins = [[empty_board() for i in range(41)] for j in range(41)]
        self.crossed_bins = [[[] for i in range(41)] for j in range(41)]

    def sort_claim(self, claim):
        if claim.within_30_by_30():
            self.record_claim(claim)
        else: 
            claims = claim.sub_divide()
            
            # print('Splitting ', str(claim))
            # for this_claim in claims:
            #     print(str(this_claim))
                
            for this_claim in claims:
                self.sort_claim(this_claim)

    def record_claim(self, claim):
        x_picker = claim.x_picker 
        y_picker = claim.y_picker
        
        chess_board = self.bins[x_picker][y_picker]
        this_board = claim.get_board()
        this_board[0] = '1'
        chess_board[0] = '1'
        
        chess_board_numeric = int(''.join(chess_board), 2)        
        this_board_numeric = int(''.join(this_board), 2)
        
        # print('Claim', str(claim))
        # print('binpos', x_picker, y_picker)
        # print('C', bin(chess_board_numeric))
        # print('B', bin(this_board_numeric))
        # print('S', bin(this_board_numeric | chess_board_numeric))
    
        overlapping_claims = list(bin(chess_board_numeric & this_board_numeric))
        
        for i, bit in enumerate(overlapping_claims[2:]):
            if bit == '1' and i > 0:
                # print(i, bit)
                self.crossed_bins[x_picker][y_picker].append(i)

        self.crossed_bins[x_picker][y_picker] = list(set(self.crossed_bins[x_picker][y_picker]))

        # print(y_picker, x_picker, self.crossed_bins[x_picker][y_picker])
        self.bins[x_picker][y_picker] = list(bin(chess_board_numeric | this_board_numeric))[2:]

    def process(self):
        # file = open("data/day_03_01_test.txt", "r") 
        # file = open("data/day_03_01_sample.txt", "r") 
        file = open("data/day_03_01.txt", "r") 
        import re 
        p = re.compile('#(\d+) @ (\d+),(\d+): (\d+)x(\d+)')

        for line in file: 
            m = p.match(line)
            c = Claim(map(lambda x: int(x), m.groups()))
            self.sort_claim(c)


                
c = ClaimProcessor()
c.process()
occupied = 0
print('RESULTS')
for rows in c.crossed_bins:
    for bin in rows:
        claimed_more_than_once = bin #list(filter(lambda x: x > 0, bin))
        if len(claimed_more_than_once) > 0:
            # print(claimed_more_than_once)
            occupied += len(claimed_more_than_once)
print(occupied)

#111023 Too high.
#113586 Too high.
# 110925
#91302 Too low.
#92058 : No clue, guessing 4 times and I've got to wait for 5 mins.

if __name__ == "__main__":
    import doctest
    doctest.testmod()
