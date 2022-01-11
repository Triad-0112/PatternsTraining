#This can be extract into method. So whenever we need to reuse the code.

def printOwing(self):
    self.printBanner()
    print("name   :",self.name)
    print("amount :",self.getOutstanding())
    
#This is how extract method work

def printOwing(self):
    self.printBanner()
    self.printDetails(self.getOutstanding())
    
def printDetails(self, outstanding):
    print("name   :",self.name)
    print("amount :",outstanding)
        