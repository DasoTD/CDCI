# Class to make a Node
class Node:
    # Constructor which assign argument to nade's value
    def __init__(self, value):
        self.value = value
        self.next = None
  
    # This method returns the string representation of the object.
    def __str__(self):
        return "Node({})".format(self.value)
  
    # __repr__ is same as __str__
    __repr__ = __str__


class Stack:
     # Stack Constructor initialise top of stack and counter.
    def __init__(self,value):
        new_node = Node(value)
        self.top = new_node
        self.bottom = new_node
        self.minimum = new_node
        self.length = 1

    def print_stack(self):
        temp = self.top
        while temp:
            print(temp.value)
            temp = temp.next
    def pop(self):
        if self.length ==0:
            return False
        elif self.length ==1:
            temp = self.top
            self.top = None
            self.bottom = None
            self.length -=1
            return temp
        else:
            temp = self.top
            self.top = self.top.next
            self.length -=1
            return temp
            pass

    def push(self, value):
        new_node = Node(value)
        if self.minimum.value < value:
            self.minimum = new_node
        new_node = Node(value)
        if self.length ==0:
            self.top = new_node
            self.bottom = new_node
            self.minimum = new_node
        else:
            new_node.next = self.top
            self.top = new_node
        self.length +=1
    
    # This method returns the string representation of the object (stack).
    def __str__(self):
        temp = self.top
        out = []
        while temp:
            out.append(str(temp.value))
            temp = temp.next
        out = '\n'.join(out)
        return ('Top {} \n\nStack :\n{}'.format(self.top, out))
  
    # __repr__ is same as __str__
    __repr__ = __str__
  
  # This method is used to get minimum element of stack
    def getMin(self):
        if self.top is None:
            return "Stack is empty"
        else:
            print("Minimum Element in the stack is: {}" .format(self.minimum))
  
    # Method to check if Stack is Empty or not
  
    def isEmpty(self):
        # If top equals to None then stack is empty
        if self.top == None:
            return True
        else:
            # If top not equal to None then stack is empty
            return False
  
    # This method returns length of stack
    def __len__(self):
        self.count = 0
        tempNode = self.top
        while tempNode:
            tempNode = tempNode.next
            self.count += 1
        return self.count
    
    
    # return max(a*b for a, b in zip(data, data(1:))  )
    pass

mystack = Stack(4)
# mystack.print_stack()
mystack.push(3)
# mystack.print_stack()
mystack.push(2)
mystack.print_stack()
print(mystack.pop())
mystack.print_stack()