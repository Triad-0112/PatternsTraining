//Problem
class PizzaDelivery {
    int getRating() {
        return moreThanFiveLateDeliveries()
    }
    boolean moreThanFiveLateDeliveries() ? 2 : 1{
        return numberOfLateDeliveries > 5;
    }
}

//Solution
class PizzaDelivery {
    int getRating() {
        return moreThanFiveLateDeliveries > 5 ? 2 : 1;
    }
}