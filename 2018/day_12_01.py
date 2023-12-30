class Node: 
    def __init__(self, value, pot_number):
        self.value = value
        self.next = None
        self.pot_number = pot_number
        self.previous = None
        self.is_first_node = False
        self.is_last_node = False
    
    def init_prefix(self):
        current_node = self 
        for i in range(5):
            current_node.previous = Node('.', current_node.pot_number - 1)
            current_node.previous.next = current_node 
            current_node = current_node.previous
        return current_node
        
    def init_suffix(self):
        current_node = self 
        for i in range(5):
            current_node.next = Node('.', current_node.pot_number + 1)
            current_node.next.previous = current_node 
            current_node = current_node.next
        
    def add_pot(self, value):
        self.next = Node(value, self.pot_number + 1)
        self.next.previous = self
        return self.next
        
    def prepend_pot(self, value):
        self.previous = Node(value, self.pot_number - 1)
        self.previous.next = self
        return self.previous
        
    def left(self, index):
        if index == 0 or self.previous is None:
            return self.value
        return self.previous.left(index-1)
        
    def right(self, index):
        if index == 0 or self.next is None:
            return self.value
        return self.next.right(index-1)
        
    def finalize_generation(self, first_five, last_five):
        self.value = self.next_gen_value
        if len(first_five) <= 5:
            first_five.append(self)
            
        if len(last_five) <= 5:
            last_five.append(self)
        else:
             last_five.pop(0)
             last_five.append(self)
             
        if self.next is not None:
            self.next.finalize_generation(first_five, last_five)
        
    def process_generation(self, rules):
        rule = self.left(2) + self.left(1) + self.value + self.right(1) + self.right(2) 
        old = self.value
        new = rules[rule] if rule in rules else '.'
        # if rule == '.#...':
        #     print('[{1}] = {2}'.format(old, rule, new))
        self.next_gen_value = new
        if self.next is not None:
            self.next.process_generation(rules)
        
    def print_me(self):
        return '{0}'.format(self.value, self.pot_number) + self.next.print_me() if self.next is not None else self.value
        
    def count_pot_numbers(self):
        return (self.pot_number if self.value == '#' else 0) + self.next.count_pot_numbers() if self.next is not None else 0

class Gardener:
    def __init__(self):
        pass 
        
    def debug(self):
        pass
        
    def live_a_generation(self, rules):
        self.root_node.process_generation(rules)
    
    def process(self):
        # file = open("data/day_12_01_test.txt", "r") 
        file = open("data/day_12_01.txt", "r") 
        
        initial_state = file.readline()
        
        self.initial_state = initial_state.rstrip().replace("initial state: ", '')
        
        current_node = None
        
        for i, c in enumerate(self.initial_state):
            if current_node is None: 
                current_node = Node(c, 0)
                self.start_pot = current_node.init_prefix()
            else:
                current_node = current_node.add_pot(c)
        
        current_node.init_suffix()
        
        file.readline()
        
        rules = {}
        for line in file:
            rule = line.rstrip().split("=>")
            rules[rule[0].strip()] = rule[1].strip()
            
        prev_sum = 0    
        for i in range(121):
            pots = (i, self.start_pot.print_me())
            # print(pots, len(list(filter(lambda x: x== '#', pots[1]))))
            current_sum = self.start_pot.count_pot_numbers()
            
            print('after', i , current_sum, current_sum-prev_sum)
            prev_sum = current_sum
            
            self.start_pot.process_generation(rules)
            first_five = [self.start_pot]
            last_five = []
            self.start_pot.finalize_generation(first_five, last_five)
            
            last_pot_with_plant = max(enumerate(last_five), key=lambda x: x[0] if x[1].value == '#'  else -1)
            first_pot_with_plant = min(enumerate(first_five), key=lambda x: x[0] if x[1].value == '#'  else 999)
            
            if last_pot_with_plant[1].value == '#':
                added_node = last_five[-1]
                for k in range(0, last_pot_with_plant[0]+1):
                    added_node = added_node.add_pot('.')

                    
            if first_pot_with_plant[1].value == '#':
                current_node = first_five[0]
                for k in range(0, 5 - first_pot_with_plant[0]+1):
                    current_node = current_node.prepend_pot('.')
                    
                self.start_pot = current_node
        
            # pots = (i, self.start_pot.print_me())
        
        
c = Gardener() 
c.process()
# 50
# 47
# c.debug()
