class Guard:
    def __init__(self, id):
        self.id = id
        self.current_shift = None
        self.shifts = {}
        self.sleep_time = 0
        
    def begin_shift(self, at):
        self.current_shift =  { 'sleep_at': [], 'awake_at': [] }
        self.shifts[at] = self.current_shift
        
    def sleep(self, sleep_at):
        self.current_shift['sleep_at'].append(sleep_at)
        
    def awake(self, awake_at):
        self.current_shift['awake_at'].append(awake_at)
        
    def analyze_sleep_times(self):
        self.sleep_time = 0
        self.minutes_sleeping = [0 for i in range(0, 60)]
        for time, shift in self.shifts.items():
            sleep_time_for_day = 0
            for index, sleep_at in enumerate(shift['sleep_at']):
                sleep_time_for_day += shift['awake_at'][index] - sleep_at
                for sleeping_minute in range(sleep_at, shift['awake_at'][index]):
                    self.minutes_sleeping[sleeping_minute] += 1
                    
            shift['sleep_time_for_day'] = sleep_time_for_day
            self.sleep_time += sleep_time_for_day
            
        self.minutes_slept_most_days = max(enumerate(self.minutes_sleeping), key=lambda v: v[1])
        
    def __str__(self):
        return "Guard: {0}; TS: {1}, DS: {2}".format(self.id, self.sleep_time, ','.join(map(lambda kv: str(kv[1]['sleep_time_for_day']), self.shifts.items())))
        
class SleepProcessor:
    def __init__(self):
        self.guards = {}
        
    def debug(self):
        for k, guard in self.guards.items():
            print(guard)

    def process(self):
        # file = open("data/day_04_01_test.txt", "r") 
        file = open("data/day_04_01.txt", "r") 
        import re 
        # [1518-11-01 00:05] falls asleep
        p = re.compile('\[(\d{4}-\d{2}-\d{2}) (\d{2}):(\d{2})\] (falls|wakes|Guard #(\d*)) .*')
        
        lines = file.readlines()
        lines.sort()
        
        
        current_guard = None
        for line in lines: 
            m = p.match(line)
            (date, hour, min, action, guard_id) = m.groups()

            if action.startswith("Guard"):
                if guard_id in self.guards:
                    current_guard = self.guards[guard_id]
                else:
                    current_guard = Guard(guard_id)
                    self.guards[guard_id] = current_guard

                current_guard.begin_shift('{0} {1}:{2}'.format(date, hour, min))

            if action == "falls":
                current_guard.sleep(int(min))

            if action == "wakes":
                current_guard.awake(int(min))
                
        for id, guard in self.guards.items():
            guard.analyze_sleep_times()
            
        
        max_sleeping_guard = max(self.guards.values(), key=lambda v: v.sleep_time)
        minutes_slept_most_days = max(enumerate(max_sleeping_guard.minutes_sleeping), key=lambda v: v[1])
        # max_sleeping_guard = max(list(map(lambda v: v.sleep_time, self.guards.values())))
        
        #strategy 2
        max_by_most_slept_at_a_minute = max(self.guards.values(), key=lambda v: v.minutes_slept_most_days[1])
        
        print("Max sleeping guard")
        print(max_sleeping_guard)
        print(max_sleeping_guard.minutes_sleeping)
        print(minutes_slept_most_days)
        
        print("strategy: 2")
        print(max_by_most_slept_at_a_minute)
        print(max_by_most_slept_at_a_minute.minutes_sleeping)
        print(max_by_most_slept_at_a_minute.minutes_slept_most_days)

c = SleepProcessor()
c.process()
print("DEBUG")
c.debug()
