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
        
runam = Chapter1()
runam.oneOne('456')