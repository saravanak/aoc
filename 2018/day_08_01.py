import pprint 
pp = pprint.PrettyPrinter(indent=4)

class Node:
    def __init__(self, node_id, child_node_count, meta_data_length):
        self.node_id = node_id
        self.child_node_count = child_node_count
        self.meta_data_length = meta_data_length
        self.meta_data = []
        self.child_nodes = []
        
    def pending_children_count(self):
        return self.child_node_count - len(self.child_nodes)
        
    def sum_meta(self):
        child_count = 0
        for child in self.child_nodes:
            child_count += child.sum_meta() 
            
        return sum(self.meta_data) + child_count
    
    def __str__(self):
        return "#{0}, C:{1}, M:{2}".format(self.node_id, self.child_nodes, self.meta_data)
        

class LicenseFileProcessor:
    def __init__(self):
        self.node_list = []
        
    def debug(self):
        pp.pprint(list(map(lambda x: str(x), self.node_list)))
        
    def reset(self):
        pass
        
    def consume_subtree(self):
        parent_node = self.node_list[-1]
        
        pending_children_count = parent_node.pending_children_count()
        
        if pending_children_count > 0:
            self.node_index += 1
            child_node = Node(self.node_index, self.next_in_sequence(), self.next_in_sequence())
            parent_node.child_nodes.append(child_node)
            self.node_list.append(child_node)
            self.consume_subtree()
            
        else:
            [parent_node.meta_data.append(self.next_in_sequence()) for i in range(0, parent_node.meta_data_length)]
            self.node_list.pop()    
            
        print(str(parent_node))

    def next_in_sequence(self):
        next_number = self.numbers[self.instruction_pointer]
        self.instruction_pointer += 1
        return next_number

    def process(self):
        # file = open("data/day_08_01_test.txt", "r") 
        file = open("data/day_08_01.txt", "r") 
        
        line = file.readline().rstrip()

        self.numbers = list(map(lambda  x: int(x), line.split(" ")))
        self.instruction_pointer = 0
        
        self.node_index = 1
        root_node = Node(self.node_index, self.next_in_sequence(), self.next_in_sequence())
        self.node_list.append(root_node)
        
        while self.instruction_pointer < len(self.numbers):
            self.consume_subtree()
            
        print(root_node.sum_meta())
        
c = LicenseFileProcessor()
c.process()
c.debug()
