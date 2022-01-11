//Problem
int discount(int inputVal, int quantity) {
    if (quantity > 50) {
      inputVal -= 2;
    }
    // ...
  }

//Solution
int discount(int inputVal, int quantity) {
    int result = inputVal;
    if (quantity > 50) {
      result -= 2;
    }
    // ...
  }