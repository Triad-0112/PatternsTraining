//Case
public class Order 
{
  // ...
  public double Price() 
  {
    double primaryBasePrice;
    double secondaryBasePrice;
    double tertiaryBasePrice;
    // Perform long computation.
  }
}

//Solution
public class Order 
{
  // ...
  public double Price() 
  {
    return new PriceCalculator(this).Compute();
  }
}

public class PriceCalculator 
{
  private double primaryBasePrice;
  private double secondaryBasePrice;
  private double tertiaryBasePrice;
  
  public PriceCalculator(Order order) 
  {
    // Copy relevant information from the
    // order object.
  }
  
  public double Compute() 
  {
    // Perform long computation.
  }
}