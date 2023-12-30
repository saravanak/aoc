import pprint 
pp = pprint.PrettyPrinter(indent=4)

class JobScheduler:
    def __init__(self):
        self.job_dependencies = [['' for i in range(0, 26)] for j in range(0, 26)] 
        self.jobs = []
        
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
        
    def schedule(self):
        self.pending_jobs = list(map(lambda x: x, self.jobs))
        self.completed_jobs = []
        self.queued_jobs = []
        
        while len(self.pending_jobs) > 0 or len(self.queued_jobs) > 0:
            runnables = self.find_runnables()
            print('Runnable', runnables)
            for runnable in runnables:
                self.queued_jobs.append(runnable)
            self.queued_jobs = list(set(self.queued_jobs))
            self.queued_jobs.sort()
            
            job_to_execute = self.queued_jobs.pop(0)
            print('EXEC job', job_to_execute)
            self.completed_jobs.append(job_to_execute)
            
            self.pending_jobs.remove(job_to_execute)

        print(''.join(self.completed_jobs))
        
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
