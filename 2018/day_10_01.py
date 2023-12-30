import pprint 
from turtle import Turtle

import statistics 

pp = pprint.PrettyPrinter(indent=4)

class Point:
    def __init__(self, id, x, y, vx, vy):
        self.x = x
        self.y = y
        self.vx = vx
        self.vy = vy
        self.id = id

    def __str__(self):
        return '{2}@ ({0}, {1})'.format(self.x, self.y, self.id)
        # return 'position=<{: d}, {: d}> velocity=<{: d}, {: d}>'.format(self.x, self.y, self.vx, self.vy)
        
    def move(self):
        self.x += self.vx
        self.y += self.vy
        
    def move_back(self):
        self.x -= self.vx
        self.y -= self.vy
    
class SkyWatcher:
    def __init__(self):
        self.points = []
        self.prev_stdev_x = 0
        self.prev_stdev_y = 0

    def debug(self, i):
        # 15, -6, 11, -4
        
        def transpose(point, min, coordinate):
            x = getattr(point, coordinate)
            if x < 0:
                x_coordinate = abs(min) - abs(x)
            else:
                x_coordinate = x - abs(min)
            return x_coordinate


        if i == 1:
        #     for p in self.points: 
        #         print(str(p))
                
            print('{0}, {1}, {2}, {3}'.format(self.minx, self.maxx, self.miny, self.maxy))
            sky_part = [['.'.format(i, j) for i in range(0, self.maxx - self.minx + 1) ] for j in range(0, self.maxy - self.miny + 1)] 
            # print('\n'.join([' '.join([str(cell) for cell in row]) for row in sky_part]))
            
            
            for p in self.points:
                print(str(p), transpose(p, self.minx, 'x'), transpose(p, self.miny, 'y'))
                sky_part[transpose(p, self.miny, 'y')][transpose(p, self.minx, 'x')] = '#' 
            
            print('\n'.join([' '.join([str(cell) for cell in row]) for row in sky_part]))
        
        self.stdev_x = statistics.stdev(list(map(lambda p: transpose(p, self.minx, 'x'), self.points)))
        self.stdev_y = statistics.stdev(list(map(lambda p: transpose(p, self.miny, 'y'), self.points)))
        # print(self.stdev_x, self.stdev_y)
        
        if self.prev_stdev_x < self.stdev_x and self.prev_stdev_y < self.stdev_y:
            input("BLINK") 
            print(i)
            
        self.prev_stdev_x = self.stdev_x 
        self.prev_stdev_y = self.stdev_y
        
            
            
    def process(self):
        # file = open("data/day_10_01_test.txt", "r") 
        file = open("data/day_10_01_01.txt", "r") 
        # file = open("data/day_10_01.txt", "r") 
        import re
        import time
        p = re.compile("position=<(.\d+), (.\d+)> velocity=<(.\d+), (.\d+)>")
        
        i = 0
        for line in file:
            m = p.match(line)
            (x, y, vx, vy) = list(map(lambda x: int(x), m.groups()))
            self.points.append(Point(i, x, y, vx, vy));
            i+=1
        
        i =0    
        while True:     
            self.maxx = max(self.points, key = lambda p: p.x).x
            self.maxy = max(self.points, key = lambda p: p.y).y
            self.minx = min(self.points, key = lambda p: p.x).x
            self.miny = min(self.points, key = lambda p: p.y).y
            # print('CORRO', self.maxx , self.minx, self.maxy , self.miny)
            
            self.debug(i)
            [point.move() for point in self.points]
            print("XXXXXXXXXXXXXXXX", i)
            # input("Enter your name: ") 
            i+=1
        
c = SkyWatcher() 
c.process()
c.debug()
#PHLGRNFK yay! and 10407 seconds later.
