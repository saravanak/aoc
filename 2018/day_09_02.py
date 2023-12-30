import pprint 
pp = pprint.PrettyPrinter(indent=4)

class Node: 
    def __init__(self, value):
        self.value = value
        self.next = self
        self.previous = self 
        
    def add_marble(self, value):
        node = self
        
        node = node.next

        newNode = Node(value)
        successor = node.next
        node.next = newNode
        newNode.previous = node 
        newNode.next = successor
        successor.previous = newNode
        return newNode
        
    def remove(self):
        previous = self.previous 
        next = self.next 
        previous.next = next 
        next.previous = previous
        return (self.value, next)
        
    def print_me(self, current_marble, root_node):
        format_string = '({0})'  if self.value == current_marble  else ' {0} '
        return format_string.format(self.value) + (self.next.print_me(current_marble, root_node)  if self.next != root_node else '')
        
class MarbleGame:
    def __init__(self, number_players, last_marble_worth):
        self.first_marble = Node(0)
        self.marbles_in_circle = self.first_marble
        self.number_players = number_players
        self.last_marble_worth = last_marble_worth
        self.player_scores = [0 for i in range(0, number_players)]
        
    def debug(self):
        pp.pprint(
        '{0}'.format(
            self.first_marble.print_me(self.marbles_in_circle.value, self.first_marble)
        ))
            
    def seven_anticlockwise_pos(self):
        for i  in range(0, 7):
            self.marbles_in_circle = self.marbles_in_circle.previous

    def process(self):
        player = 0
        
        next_marble = 1
        while next_marble <= self.last_marble_worth :
            if next_marble % 1000 == 0: 
                print(next_marble)
                
            if next_marble % 23 == 0 and next_marble > 0:
                self.player_scores[player]+= next_marble
                removed_node_pos = self.seven_anticlockwise_pos()
                removed_node =  self.marbles_in_circle.remove()
                self.player_scores[player]+= removed_node[0]
                self.marbles_in_circle = removed_node[1]
            else:
                self.marbles_in_circle = self.marbles_in_circle.add_marble(next_marble)
                
            # self.debug()
            player += 1
            player %= self.number_players
            next_marble += 1
        
# c = MarbleGame(9,25) #32
# c = MarbleGame(10, 1618) # 8317
# c = MarbleGame(13, 7999) # 146373
# c = MarbleGame(17, 1104) # 2764
# c = MarbleGame(21, 6111) # 54718
# c = MarbleGame(30, 5807) # 37305
# c = MarbleGame(411, 71058) # 424639
c = MarbleGame(411, 7105800) # 3516007333 yay!
c.process()
# c.debug()
print(max(c.player_scores))
