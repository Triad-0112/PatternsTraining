#This can be extract into method. So whenever we need to reuse the code.

class PizzaDelivery:
    def getRating(self):
        return 2 if self.moreThanFiveLateDeliveries() else 1
  
    def moreThanFiveLateDeliveries(self):
        return self.numberOfLateDeliveries > 5
    
#This is how extract method work

class PizzaDelivery:
    def getRating(self):
        return 2 if self.numberOfLateDeliveries > 5 else 1
        