# Python program to add two numbers 
# represented by linked list
  
# Node class
class Node:
  
    # Constructor to initialize the 
    # node object
    def __init__(self, data):
        self.data = data
        self.next = None
  
  
class LinkedList:
  
    # Function to initialize head
    def __init__(self, value):
        # self.head = None

        new_node = Node(value)
        self.value = value
        self.next = None
        self.head = new_node
        self.tail = new_node
        self.length = 1

    def partition(self, x):
        temp = self.Get(x)
        head = self.head
        res = []
        while head:
            if head.data >= temp:
                res.append(head.data)
                head = head.next
            else:
                res.insert(0, head.data)
                head = head.next
        return res
    
    def print_list(self):
        if self.length ==0:
            return None
        temp = self.head
        while temp:
            print(temp.data)
            temp= temp.next
        pass
  
    # Function to insert a new node at 
    # the beginning
    def prepend(self, new_data):
        new_node = Node(new_data)
        new_node.next = self.head
        self.head = new_node
        self.length +=1
        return True
    
    def append(self, data):
        new_node = Node(data)
        self.tail.next = new_node
        new_node = self.tail
        self.length +=1
        return True
    
    def Pop(self):
        if self.length ==0:
            return None
        if self.length == 1:
            temp = self.head
            self.head = None
            self.tail = None
            self.length -= 1
            return temp.data
        temp = self.head
        pre = self.head
        while(temp.next):
            pre = temp
            temp = temp.next
        self.tail = pre
        self.tail.next = None
        self.length -=1
        return temp.data
        pass
    def first_pop(self):
        if self.length ==0:
            return None
        elif self.length == 1:
            temp = self.head
            self.head = None
            self.tail = None
            self.length -= 1
            return temp.data
        else:
            temp = self.head
            self.head = temp.next
            temp.next = None
            self.length -=1
            return temp.data

        pass

    def loopDetection(self):
        s = set()
        head =  self.head
        while head:
            if (head.data in s):
                return True
            else:
                s.add(head.data)
                head = head.next
        return False
        
    
    def Get(self, index):
        if index < 0 or index > self.length:
            return None
        if index == self.length:
            return self.tail.data
        if index == 0:
            return self.head.data
        temp = self.head
        for _ in range(index):
            temp = temp.next
        return temp.data
    
    def deleteMiddle(self, index):
        if index < 0:
            return None
        if index == self.length:
            return self.Pop()
        if index == 0:
            return self.first_pop()
        temp = self.Get(index-1)
        after = temp.next
        if temp:
            temp.next = after.next
            after.next = None
            self.length -=1
            return after.data
        pass
    # or
    # temp = self.head
        # while temp is not None:
        #     if temp.value == index:
        #         return temp.value
        #     temp = temp.next

    def sum(self, A, B):
        a = ''
        b = ''
        for i in reversed(A):
            a += str(i)
        for j in reversed(B):
            b += str(j)
        a = int(a)
        b = int(b)
        print(a+b)

    def AddTwoLists(self, first, second):
        prev = None
        temp = None
        carry = 0
  
        # While both list exists
        while(first  or second ):
            fdata = 0 if first is None else first.data
            sdata = 0 if second is None else second.data
            Sum = carry + fdata + sdata

             # Update carry for next calculation
            carry = 1 if Sum >= 10 else 0
  
            # Update sum if it is greater than 10
            Sum = Sum if Sum < 10 else Sum % 10
  
            # Create a new node with sum as data
            temp = Node(Sum)
            print(temp)

                        # if this is the first node then set 
            # it as head of resultant list
            if self.head is None:
                self.head = temp
            else:
                prev.next = temp
  
            # Set prev for next insertion
            prev = temp
  
            # Move first and second pointers to 
            # next nodes
            if first is not None:
                first = first.next
            if second is not None:
                second = second.next
  
        if carry > 0:
            temp.next = Node(carry)
  
        
    # Utility function to print the 
    # linked LinkedList
    def printList(self):
        temp = self.head
        while(temp):
            print (temp.data),
            temp = temp.next
            # print (temp.data),
            # temp = temp.next

    def palindrome(self):
        reversed = self.reversed()
        print(reversed)
        return self.isEqual(self, reversed)
        

    def reversed(self):
        prev = None
        current = self.head
        while current :
            next = current.next
            current.next = prev
            prev = current
            current = next
        self.head = prev
        

    def isEqual(self, A, B):
        A = Node(A)
        B = Node(B)

        while(A and B ):
            if(A.data != B.data):
                print(A, B)
                print("e no work")
                return False
            print(A.data, B.data)
            A = A.next
            B = B.next
        return  A == None and B == None

        
  

class Chapter2:
    def __init__(self):
        self = self

    def sum(self, listA, listB):
        A = ''
        B = ''
        for i in  reversed(listA):
            A += str(i)
        for i in reversed(listB):
            B += str(i)
        A = int(A)
        B = int(B)
        print(A+B)


   
c2 = Chapter2()
# c2.sum([1,2,3], [4,5,6])
# c2.palindrome(['1','2','3','2','1'])
first = LinkedList(0)
first.prepend(1)
first.prepend(2)
first.prepend(3)
first.append(10) 
# first.append(10) 
first.prepend(3)
first.prepend(5)
# print(first.Pop())
# print(first.first_pop())
# first.printList() 
# first.reversed()
# first.printList()
# print(first.Get(3))
# print(first.deleteMiddle(1))
# first.palindrome()
print("file")
first.printList()
print(first.partition(3))
print("cook")
# first.printList()
# first.sum([1,2,3],[4,5,6])
# print(first.loopDetection())

