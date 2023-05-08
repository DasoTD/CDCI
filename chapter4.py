
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

    def inorder(self):
        if self.root.value is not None:
            self.inorder(self.root.left)
            print(self.root.value, end= "")
            self.inorder(self.root.right)

    def deleteNode(self, key):
        root = self.root
        # Base Case
        if root.value is None:
            return root

        # If the key to be deleted
        # is smaller than the root's
        # key then it lies in  left subtree
        if key < root.value:
            root.left = self.deleteNode(root.left, key)

        # If the kye to be delete
        # is greater than the root's key
        # then it lies in right subtree
        elif(key > root.value):
            root.right = self.deleteNode(root.right, key)

        # If key is same as root's key, then this is the node
        # to be deleted
        else:

            # Node with only one child or no child
            if root.left is None:
                temp = root.right
                root = None
                return temp

            elif root.right is None:
                temp = root.left
                root = None
                return temp

            # Node with two children:
            # Get the inorder successor
            # (smallest in the right subtree)
            temp = self.get_min(root.right)

            # Copy the inorder successor's
            # content to this node
            root.value = temp.value

            # Delete the inorder successor
            root.right = self.deleteNode(root.right, temp.value)

        return root

    def delete(self, value):
        new_node = Node(value)
        if self.root is None:
            return self.root
        
        temp = self.root
        while temp:
            if new_node.value == temp.value:
                temp.value = self.inOrderSuccessor()
            if new_node.value < temp.value:
                temp = temp.value
                # temp.left = self.delete(value)
            elif new_node.value > temp.value:
                temp = temp.right
                # temp.right = self.delete( value)
        return temp

        # else :
        #     if temp.left == None:
        #         return temp.right
        #     elif temp.right == None:
        #         return temp.left
        #     temp.value = self.inOrderSuccessor(temp.right)
        #     temp.right = self.delete(temp.right, temp.value)
        # return temp

    def inOrderSuccessor(self):
        minimum = self.root.value
        temp = self.root
        while(temp.left):
            minimum = temp.left.value
            temp = temp.left
        return minimum
        

    def search(self, value):
        if self.root is None:
            return self.root
        temp = self.root
        
        while temp:
            if temp.value == value:
                return temp
            if temp.value > value:
                temp = temp.left
            else:
                temp = temp.right
            
        return False

    def get_min(self):
        if self.root == None:
            return None
        root = self.root
        while root.left:
            root = root.left
        return root.value

    def passget_max(self):
        root = self.root
        while root.right:
            root = root.right
        return root.value
    
    
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
    
        

bst = BST()
bst.insert(10)
bst.insert(12)
bst.insert(4)
bst.insert(14)
bst.insert(13)
bst.insert(6)
bst.insert(15)
print(bst.search(6))
bst.deleteNode(14)
print(bst.search(6))
print(bst.search(14))

# def delete(self, value):
#     if self.root == None:
#         return self.root
    
#     temp = self.root
#     if value < temp.value:
#         temp.left = delete(temp.left, value)
#     elif value > temp.value:
#         temp.right = delete(temp.right, value)

#     else :
#         if temp.left == None:
#             return temp.right
#         elif temp.right == None:
#             return temp.left
#         temp.value = inOrderSuccessor(temp.right)
#         temp.right = delete(temp.right, temp.value)
#     return temp

# def inOrderSuccessor(self):
#     minimum = self.root.value
#     temp = self.root
#     while(temp.left):
#         minimum = temp.left.value
#         temp = temp.left
#     return minimum

# public static TreeNode deleteRecursively(TreeNode root, int value) {

#         if (root == null)
#             return root;

#         if (value < (int) root.data) {
#             root.left = deleteRecursively(root.left, value);
#         } else if (value > (int) root.data) {
#             root.right = deleteRecursively(root.right, value);
#         } else {

#             if (root.left == null) {
#                 return root.right;
#             } else if (root.right == null)
#                 return root.left;

#             root.data = inOrderSuccessor(root.right);
#             root.right = deleteRecursively(root.right, (int) root.data);
#         }

#         return root;

#     }

#     public static int inOrderSuccessor(TreeNode root) {
#         int minimum = (int) root.data;
#         while (root.left != null) {
#             minimum = (int) root.left.data;
#             root = root.left;
#         }
#         return minimum;
#     }