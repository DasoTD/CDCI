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

    def palindrome(self):
        reversed = self.reverved()
        return self.isEqual(reversed, self)
        

    def isEqual(self, A, B):
        A = Node(A)
        B = Node(B)

        while A or B:
            if A != B:
                print(A, B)
                print("e no work")
                # return False
            A = A.next
            B = B.next
        return True
    
    def reverved(self):
        prev = None
        current = self.head

        while current:
            next = current.next
            current.next = prev
            prev = current
            current = next
        self.head = prev
        
    

LL = LinkedList(0)
LL.prepend(1)
LL.prepend(2)
LL.prepend(3)
# LL.print_list()
LL.palindrome()

BB = LinkedList(0)
BB.prepend(1)
BB.prepend(2)
BB.prepend(3)

res = LinkedList(0)

res.isEqual(LL, BB)
