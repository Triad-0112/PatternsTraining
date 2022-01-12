# Factory Method C++

## EXPLANATION
**Factory method** is a creational design pattern which solves the problem of creating product objects without specifying their concrete classes.

**Factory Method** defines a method, which should be used for creating objects instead of direct constructor call (**new** operator). Subclasses can override this method to change the class of objects that will be created.

**USAGE** : The Factory Method pattern is widely used in C++ code. It’s very useful when you need to provide a high level of flexibility for your code.

**Identification** : Factory methods can be recognized by creation methods, which create objects from concrete classes, but return them as objects of abstract type or interface.

## CONCEPTUAL EXAMPLE
This example illustrates the structure of the Factory Method design pattern. It focuses on answering these questions:
>* What classes does it consist of?
>* What roles do these classes play?
>* In what way the elements of the pattern are related?