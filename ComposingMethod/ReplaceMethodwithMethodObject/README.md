# Remove Method With Method Object

**PROBLEM** : You have a long method in which the local variables are so intertwined that you can’t apply **Extract Method**.
**SOLUTION**: Transform the method into a separate class so that the local variables become fields of the class. Then you can split the method into several methods within the same class.

## REFACTOR
>* Transform the method into a separate class so that the local variables become fields of the class. Then you can split the method into several methods within the same class.
>* In the new class, create a private field for storing a reference to an instance of the class in which the method was previously located. It could be used to get some required data from the original class if needed.
>* Create a separate private field for each local variable of the method.
>* Create a constructor that accepts as parameters the values of all local variables of the method and also initializes the corresponding private fields.
>* Declare the main method and copy the code of the original method to it, replacing the local variables with private fields.
>* Replace the body of the original method in the original class by creating a method object and calling its main method.

## WHY
A method is too long and you can’t separate it due to tangled masses of local variables that are hard to isolate from each other.

The first step is to isolate the entire method into a separate class and turn its local variables into fields of the class.

Firstly, this allows isolating the problem at the class level. Secondly, it paves the way for splitting a large and unwieldy method into smaller ones that wouldn’t fit with the purpose of the original class anyway.

## PROS
Isolating a long method in its own class allows stopping a method from ballooning in size. This also allows splitting it into submethods within the class, without polluting the original class with utility methods.

## DRAWBACK
Another class is added, increasing the overall complexity of the program.