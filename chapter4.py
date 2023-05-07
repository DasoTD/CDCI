
class Node:
    def __init__(self, value):
        self.value = value
        self.right = None
        self.left =  None

class BST:
    def __init__(self):
        # new_node = Node(value)
        self.root = None
        

    def insert(self, value):
        new_node = Node(value)
        if self.root is None:
            self.root = new_node
            return True
        temp = self.root

        while (True):
            if new_node.value == temp.value and temp.left is not None:
                return False
            
            # go left
            if new_node.value < temp.value:
                if temp.left is None:
                    temp.left = new_node
                    return True
                temp = temp.left
            else:
                if temp.right is None:
                    temp.right = new_node
                    return True
                temp = temp.right
                
    # Inorder traversal
    # Left -> Root -> Right
    def inorderTraversal(self, root):
        res = []
        if root:
            res = self.inorderTraversal(root.left)
            res.append(root.value)
            res = res + self.inorderTraversal(root.right)
        return res


     # Preorder traversal
    # Root -> Left -> Right
    def preorderTraversal(self, root):
        res = []
        if root:

            res.append(root.value)
            res = res + self.preorderTraversal(root.left)
            res = res + self.preorderTraversal(root.right)
        return res
    
     # Preorder traversal
    # Root -> Left -> Right
    def postorderTraversal(self, root):
        res = []
        if root:

            res = res + self.postorderTraversal(root.left)
            res = res + self.postorderTraversal(root.right)
            res.append(root.value)
        return res