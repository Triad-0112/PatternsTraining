//Case
bool HasDiscount(Order order)
{
  double basePrice = order.BasePrice();
  return basePrice > 1000;
}

//Solution
bool HasDiscount(Order order)
{
  return order.BasePrice() > 1000;
}