class Node:
    def __init__(self, value):
        self.value = value
        self.next = None
        
class LinkedList:
    def __init__(self, value):
        new_node = Node(value)
        self.value = value
        self.head = new_node
        self.tail = new_node
        self.next = None
        self.length = 1

    
    def print_list(self):
        temp = self.head

        while temp:
            print(temp.value)
            temp = temp.next
        return True
    
    def prepend(self, value):
        new_node = Node(value)
        if self.length == 0:
            self.head = new_node
            self.tail =new_node
        else:
            new_node.next = self.head
            self.head = new_node
        self.length += 1
    

LL = LinkedList(0)
LL.prepend(1)
LL.prepend(2)
LL.print_list()
