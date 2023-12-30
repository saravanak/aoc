import pprint 
pp = pprint.PrettyPrinter(indent=4)
NUM_WORKERS = 5
BASE_JOB_COST = 60
class JobScheduler:
    def __init__(self):
        self.job_dependencies = [['' for i in range(0, 26)] for j in range(0, 26)] 
        self.jobs = []
        self.workers = 0
        
    def debug(self):
        print("JOBS:", ','.join(self.jobs))
        print('\n'.join([' '.join([str(cell) for cell in row]) for row in self.job_dependencies]))
        
    def reset(self):
        pass

    def add_job_dependency(self, job_depedency):
        (pred, suc) = job_depedency
        self.job_dependencies[ord(pred) - 65][ord(suc) - 65] = suc
        self.jobs.append(pred)
        self.jobs.append(suc)
        self.jobs = list(set(self.jobs))
        
    def finalze_workers(self):
        for i in range(0, NUM_WORKERS):
            worker = self.executing_jobs[i]
            if worker is not None and worker['job_time'] == worker['spent_time']:
                self.completed_jobs.append(worker['job'])
                self.executing_jobs[i] = None
                
    def process_tick(self):
        for i in range(0, NUM_WORKERS):
            worker = self.executing_jobs[i]
            if worker is not None:
                worker['spent_time'] += 1
                
                
    def next_waiting_worker(self):
        return next((x for x in enumerate(self.executing_jobs) if x[1] is None), False)            
            
    def schedule(self):
        self.pending_jobs = list(map(lambda x: x, self.jobs))
        self.completed_jobs = []
        self.queued_jobs = []
        self.executing_jobs = [None for x in range (0, NUM_WORKERS)]
        
        tick = 0
        while len(self.pending_jobs) > 0 or len(self.queued_jobs) > 0:
            self.finalze_workers()
            runnables = self.find_runnables()
            print('tick', tick, self.executing_jobs, runnables)
            for runnable in runnables:
                self.queued_jobs.append(runnable)
            self.queued_jobs = list(set(self.queued_jobs))
            self.queued_jobs.sort()
            
            next_waiting_worker = self.next_waiting_worker()
                
            while len(self.queued_jobs) > 0 and next_waiting_worker is not None and next_waiting_worker is not False:
                job_to_execute = self.queued_jobs[0]
                if next_waiting_worker[1] is None:
                    self.executing_jobs[next_waiting_worker[0]] = { 'job': job_to_execute, 'spent_time': 0, 'job_time': BASE_JOB_COST + ord(job_to_execute) - 65 + 1 }
                    self.queued_jobs.remove(job_to_execute)
                    self.pending_jobs.remove(job_to_execute)
                    next_waiting_worker = self.next_waiting_worker()
            
            self.process_tick()
            tick+=1

        print(self.executing_jobs, tick)

        while not all( x is None for x  in self.executing_jobs):
            self.finalze_workers()
            self.process_tick()
            tick+=1
            
        print(self.executing_jobs, tick-1)
        
    def predecessors_for_job(self, job):
        job_column = ord(job) - 65
        predecessors = [ chr(65 + row) if len(self.job_dependencies[row][job_column]) == 1 else '' for row in range(0, 26)]
        predecessors = list(filter(lambda x: len(x) == 1 , predecessors))
        
        if len(self.completed_jobs) > 0:
            predecessors = list(filter(lambda x: x not in self.completed_jobs, predecessors))
            
        return predecessors if len(predecessors) > 0 else None
    
    def find_runnables(self):
        return list(filter(lambda job: self.predecessors_for_job(job) is None, self.pending_jobs))
            
    def process(self):
        # file = open("data/day_07_01_test.txt", "r") 
        file = open("data/day_07_01.txt", "r") 
        
        import re 
        p = re.compile('Step (.) must be finished before step (.) can begin.')
        
        for line in file:
            m = p.match(line.rstrip())
            self.add_job_dependency(m.groups())
            
        self.schedule()
        
c = JobScheduler()
c.process()
# c.debug()
