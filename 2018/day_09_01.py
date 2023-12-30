import pprint 
pp = pprint.PrettyPrinter(indent=4)

class MarbleGame:
    def __init__(self, number_players, last_marble_worth):
        self.marbles_in_circle = []
        self.number_players = number_players
        self.last_marble_worth = last_marble_worth
        self.current_marble = 0
        self.current_marble_pos = 0
        self.player_scores = [0 for i in range(0, number_players)]
        self.pending_marbles = [i for i in range(0, last_marble_worth+1)]
        
    def debug(self):
        pp.pprint(
        '{0}, {1}'.format(
            ''.join(list(map(lambda x: '({0})'.format(str(x)) if x == self.current_marble else ' {0} '.format(str(x)), self.marbles_in_circle))),
            self.current_marble_pos
            )
        )
        
    def insert_postions(self):
        circle_length = len(self.marbles_in_circle) 
        return ((self.current_marble_pos + 1) % circle_length, 
            (self.current_marble_pos + 2) % circle_length) if circle_length > 0 else (0, 0)
            
    def seven_anticlockwise_pos(self):
        circle_length = len(self.marbles_in_circle) 
        return (self.current_marble_pos - 7) % circle_length

    def process(self):
        
        player = 0
        
        while self.current_marble < self.last_marble_worth and len(self.pending_marbles) > 0:
            self.current_marble = self.pending_marbles.pop(0)
            if self.current_marble % 1000 == 0:
                print(self.current_marble)
                
            if self.current_marble % 23 == 0 and self.current_marble > 0:
                self.player_scores[player]+= self.current_marble
                removed_node_pos = self.seven_anticlockwise_pos()
                removed_node = self.marbles_in_circle.pop(removed_node_pos)
                self.player_scores[player] += removed_node
                self.current_marble_pos = removed_node_pos
                self.current_marble = self.marbles_in_circle[removed_node_pos]
            else:
                (after, before)  = self.insert_postions()
                self.marbles_in_circle.insert(after + 1, self.current_marble)
                self.current_marble_pos = after + 1
            # self.debug()
            player += 1
            player %= self.number_players
        
# c = MarbleGame(9,25) #32
# c = MarbleGame(10, 1618) # 8317
# c = MarbleGame(13, 7999) # 146373
# c = MarbleGame(17, 1104) # 2764
# c = MarbleGame(21, 6111) # 54718
# c = MarbleGame(411, 5807) # 37305
c = MarbleGame(411, 71058) # 424639
c.process()
print(max(c.player_scores))
