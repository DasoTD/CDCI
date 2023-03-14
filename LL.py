class Node:
    def __init__(self, value):
        self.value = value
        self.next = None

class LinkNode:
    def __init__(self, value=0, next = None):
        self.value = value
        self.next = next
        
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

        
def is_palindrome(head):
        # Find the middle of the linked list
        slow = fast = head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next

        # Reverse the second half of the linked list
        prev, curr = None, slow
        while curr:
            next_node = curr.next
            curr.next = prev
            prev = curr
            curr = next_node

        # Compare the first and second halves of the linked list
        left, right = head, prev
        while right:
            if left.value != right.value:
                return False
            left = left.next
            right = right.next

        return True
    

LL = LinkedList(0)
LL.prepend(1)
LL.prepend(2)
LL.prepend(2)
LL.prepend(1)
# print(LL.is_palindrome())

head = LinkNode(1, LinkNode(2, LinkNode(2, LinkNode(1))))
print(is_palindrome(head))

