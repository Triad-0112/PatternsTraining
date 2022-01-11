#This can be extract into method. So whenever we need to reuse the code.

def hasDiscount(order):
    basePrice = order.basePrice()
    return basePrice > 1000
    
#This is how extract method work

def hasDiscount(order):
    return order.basePrice() > 1000