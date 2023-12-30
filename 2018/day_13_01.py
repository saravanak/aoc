import functools
from enum import IntFlag

def find(f, seq):
  """Return first item in sequence where f(item) == True."""
  for item in seq:
    if f(item): 
      return item

class Direction(IntFlag):
    T = 8
    B = 4
    R = 2
    L = 1
    
    def __str__(self):
        if self.value == 15:
            return '+'
        if self.value == Direction.T | Direction.B:
            return '|'
        if self.value == Direction.R | Direction.L:
            return '-'
        if self.value == Direction.R | Direction.T:
            return '\\'
        if self.value == Direction.L | Direction.T:
            return '/'
        return self.name
            
    def is_slanting_gate(self):
        return (self.value == Direction.T | Direction.R) or (self.value == Direction.T | Direction.L)
        
    def is_intersecting(self):
        return self.value == 15        
        

class Car: 
    def __init__(self, row_index, column_index, snapshot):
        self.row = row_index
        self.column = column_index
        self.direction_from_snapshot(snapshot)
        self.rotate_state = 0
        self.rotate_handlers = [self.rotate_left, self.go_straight, self.rotate_right]
        
    def __str__(self):
        return 'Car: ({0},{1})'.format(self.row, self.column)

    def on_intersection(self):
        self.rotate_handlers[self.rotate_state % 3]()
        self.rotate_state += 1
        
    def go_straight(self):
        # print("go_straight")
        pass
        
    def move_forward(self):
        snapshot = self.snapshot_from_direction()
        if snapshot == '>':
            return (self.row, self.column + 1)
        if snapshot == '<':
            return (self.row, self.column - 1)
        if snapshot == 'v':
            return (self.row+1, self.column )
        if snapshot == '^':
            return (self.row-1, self.column )
            
    def rotate_left(self):
        print("rotate_left")
        snapshot = self.snapshot_from_direction()
        if snapshot == '>':
            self.direction_from_snapshot('^')
        if snapshot == '<':
            self.direction_from_snapshot('v')
        if snapshot == 'v':
            self.direction_from_snapshot('>')
        if snapshot == '^':
            self.direction_from_snapshot('<')
            
    def rotate_right(self):
        print("rotate_right")
        snapshot = self.snapshot_from_direction()
        if snapshot == '>':
            self.direction_from_snapshot('v')
        if snapshot == '<':
            self.direction_from_snapshot('^')
        if snapshot == 'v':
            self.direction_from_snapshot('<')
        if snapshot == '^':
            self.direction_from_snapshot('>')
        
    def direction_from_snapshot(self, snapshot):
        if snapshot == '>':
            self.path_from_to = [Direction.L, Direction.R]
        if snapshot == '<':
            self.path_from_to = [Direction.R, Direction.L]
        if snapshot == 'v':
            self.path_from_to = [Direction.T, Direction.B]
        if snapshot == '^':
            self.path_from_to = [Direction.B, Direction.T]
            
    def snapshot_from_direction(self):
        start_path = self.path_from_to[0]
        to_path = self.path_from_to[1]
        if start_path == Direction.L and to_path == Direction.R:
            return '>'
        if start_path == Direction.R and to_path == Direction.L:
            return '<'
        if start_path == Direction.T and to_path == Direction.B:
            return 'v'
        if start_path == Direction.B and to_path == Direction.T:
            return '^'
                
    def in_same_postion(self, aCar):
        return self.row == aCar.row and self.column == aCar.column
        
    def realign(self, cell, row, column):
        snapshot = self.snapshot_from_direction()
        entry_from = self.path_from_to[0]

        # print(snapshot, self, row, column, cell, entry_from, entry_from in cell, cell.is_slanting_gate())
        if cell.is_intersecting():
            self.on_intersection()
            
        if not cell.is_slanting_gate(): 
            return entry_from in cell
        
        if cell == Direction.R | Direction.T:
            if snapshot == '>':
                self.rotate_right()
            elif snapshot == 'v': 
                self.rotate_left()
            elif snapshot == '<': 
                self.rotate_right()
            elif snapshot == '^': 
                self.rotate_left()
                
        if cell == Direction.L | Direction.T:
            if snapshot == '<':
                self.rotate_left()
            elif snapshot == '^': 
                self.rotate_right()
            elif snapshot == 'v': 
                self.rotate_right()
            elif snapshot == '>': 
                self.rotate_left()
        
        return True
        
    def move(self, map, cars):
        before_move = (self.row, self.column)
        start_path = self.path_from_to[0]
        
        container_cell = map[self.row][self.column]
        
        (next_row, next_column) = self.move_forward()
        
        if next_row > len(map)-1 or next_column > len(map[next_row])-1 or next_row < 0 or next_column < 0 or map[next_row][next_column] is None:
            return 0
        
        is_realigned = self.realign(map[next_row][next_column], next_row, next_column)
        # print('moving', self, self.path_from_to, is_realigned)
        
        if not is_realigned:
            return 0
        
        (self.row, self.column) = (next_row, next_column)
        
        count_cars = len(list(filter(lambda x: self.in_same_postion(x), cars)))
        if count_cars > 1:
            print("Detected collision")
            return -1
            
        # print('From : {0} to {1}({2})'.format(before_move, str(next_path), map[self.row][self.column]))
        return 1

              
class Circuitry:
    def __init__(self):
        pass 
        
    def car_or_cell_at(self, row, column, cell):
        car = find(lambda x: x.row == row and x.column == column, self.cars)
        return "\033[91m{0}\x1b[0m".format(car.snapshot_from_direction()) if car else str(cell) if cell is not None else ' '

    def debug(self):
        print('\n'.join([''.join([self.car_or_cell_at(row[0], cell[0], cell[1]) for cell in enumerate(row[1])]) for row in enumerate(self.map)]))

    def make_turns(self):
        def car_comparer(a, b):
            a_pos = a.row
            b_pos = b.row
            if a_pos != b_pos:
                return a_pos - b_pos
            else: 
                return a.column - b.column

        self.cars = sorted(self.cars, key=functools.cmp_to_key(car_comparer))

        non_moved_cars = []
        for car in self.cars:
            is_moved = car.move(self.map, self.cars)
            if is_moved == -1: 
                return car
            if is_moved == 0: 
                non_moved_cars.append(car)

        if len(non_moved_cars) == len(self.cars) :
            raise ValueError("No car can move")
            
        if len(self.cars) == 1:
            raise ValueError("Only car")
        
        print("Moved {0} Cars of {1}".format(len(self.cars) - len(non_moved_cars), len(self.cars)))
        # self.cars = [car for car in self.cars if car not in non_moved_cars]
        
        return None

    def process(self):
        # file = open("data/day_13_01_test_01.txt", "r") 
        # file = open("data/day_13_01_test_02.txt", "r") 
        # file = open("data/day_13_01_test_03.txt", "r") 
        # file = open("data/day_13_01_test_04.txt", "r") 
        # file = open("data/day_13_01_test_05.txt", "r") 
        # file = open("data/day_13_01_test_06.txt", "r") 
        # file = open("data/day_13_01_test_07.txt", "r") 
        # file = open("data/day_13_01_test_08.txt", "r") 
        file = open("data/day_13_01_test_09.txt", "r") 
        # file = open("data/day_13_01_test.txt", "r") 
        # file = open("data/day_13_01.txt", "r") 

        import re
        p = re.compile("[<>v^]")

        self.map = []
        self.cars = []
        row_index = 0
        for line in file:
            row = []
            # tbrl 8421
            cell_index = 0
            for cell in line.rstrip():
                current_path = None
                if cell == '/':
                    current_path = Direction(8+1)
                if cell == '\\':
                    current_path = Direction(8+2)
                if cell == '-' or cell == '>' or cell == '<':
                    current_path = Direction(2+1)
                if cell == '|' or cell == '^' or cell == 'v':
                    current_path = Direction(8+4)
                if cell == '+':
                    current_path = Direction(8+4+2+1)
                row.append(current_path)
                if p.match(cell):
                    self.cars.append(Car(row_index, cell_index, cell))
                cell_index += 1

            row_index += 1
            self.map.append(row)

        print("Initial state")
        
        # self.cars = self.cars[6:7]
        self.debug()
        [print(car) for car in self.cars]
        
        crashed = None
        tick = 0
        while crashed is None :
            crashed = self.make_turns()
            print('After Tick {0}'.format(tick))
            # if tick % 5 == 0:
            self.debug()
            #     input("BLINK")
            tick += 1 

        if crashed is not None:
            print(str(crashed), tick)

c = Circuitry() 
c.process()
