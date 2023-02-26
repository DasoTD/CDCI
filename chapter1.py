class Chapter1:
    def oneOne(self, val : str) -> str :
        res = ""
        for i in val:
            if i in res:
                print("wahala")
                return False
            else:
                res = f"{res}{i}"
        print(res)

    def permutation(self, str1: str, str2: str) -> bool:
        n1 = len(str1)
        n2= len(str2)
        if n1 != n2:
            print("kole work")
            return False
        a = sorted(str1)
        str1 = " ".join(a)
        b = sorted(str2)
        str2 = " ".join(b)

        for i in range(len(str1)):
            if(str1[i] != str2[i]):
                print("ko work")
                return False
        print("e work seh")
        return True

    def URLify(self, STR: str) -> str:
        if type(STR) != str:
            print("only string is allowed")
        else:
            STR = STR.replace(" ", "%20")    
            print(STR)
    
    def PalidromePermutation(self, string: str, i=0) -> str:
        # palindrome is a word that read same backward and forward e.g madam, lol, daad
        # permutation is the rearrangement of a set

        global output 
        output = False

        # if i == len(string):
        #     print("".join(string), "enter value na")
        #     # return False

        # permutation implementations
        for j in range(i, len(string)):
            words = [c for c in string]
            words[i], words[j] = words[j], words[i]
            # self.PalidromePermutation(words, i + 1)
        
        
        import re
        s = re.sub(r'[^0-9a-zA-Z]', '', string).lower()
        
        left = 0
        right = len(string) -1

        while left < right:
            if s[left] != s[right]:
                output = False
                print("this is not a valid palidrone permutation")
                return output
            left +=1
            right -=1
        print("this is a valid palindrome permutation")
        output = True

    def SinglePermutation(self, string: str, i =0 ):
        if i == len(string):   	 
            print("".join(string))

        for j in range(i, len(string)):

            words = [c for c in string]
    
            # swap
            words[i], words[j] = words[j], words[i]
        
            self.SinglePermutation(words, i + 1)
    # print(SinglePermutation('yup', 5))

    def palindrome(self, string: str) -> bool :
        if len(string) <= 0:
            print("".join(string), "enter value na")
            return False
        import re
        s = re.sub(r'[^0-9a-zA-Z]', '', string).lower()
        
        left = 0
        right = len(string) -1

        while left < right:
            if s[left] != s[right]:
                print("e no work")
                return False
            left +=1
            right -=1
        print("e work")
        return True


runam = Chapter1()
# runam.oneOne('4456')
# runam.permutation("dadad", "adadf")
# runam.URLify("oloba, data UK")
# runam.SinglePermutation("data")
# runam.palindrome("madadm")
runam.PalidromePermutation("")