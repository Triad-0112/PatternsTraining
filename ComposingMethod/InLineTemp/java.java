//Problem
boolean hasDiscount(Order order) {
  double basePrice = order.basePrice();
  return basePrice > 1000;
}

//Solution
boolean hasDiscount(Order order) {
  return order.basePrice() > 1000;
}