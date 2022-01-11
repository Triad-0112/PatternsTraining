//This can be extract into method. So whenever we need to reuse the code.
class PizzaDelivery 
{
  int GetRating() 
  {
    return MoreThanFiveLateDeliveries() ? 2 : 1;
  }
  bool MoreThanFiveLateDeliveries() 
  {
    return numberOfLateDeliveries > 5;
  }
}

//This is how extract method work
class PizzaDelivery 
{
  int GetRating() 
  {
    return numberOfLateDeliveries > 5 ? 2 : 1;
  }
}