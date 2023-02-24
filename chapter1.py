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
        
runam = Chapter1()
runam.oneOne('456')
runam.permutation("dadad", "adadf")
runam.URLify("oloba UK")