#Case

def hasDiscount(order):
    basePrice = order.basePrice()
    return basePrice > 1000
    
#Solution

def hasDiscount(order):
    return order.basePrice() > 1000