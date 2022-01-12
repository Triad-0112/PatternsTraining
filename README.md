# Move Field

**PROBLEM** : A field is used more in another class than in its own class.
**SOLUTION**: Create a field in a new class and redirect all users of the old field to it.

## REFACTOR
>* If the field is public, refactoring will be much easier if you make the field private and provide public access methods (for this, you can use **Encapsulate Field**).
>* Create the same field with access methods in the recipient class.
>* Decide how you will refer to the recipient class. You may already have a field or method that returns the appropriate object; if not, you will need to write a new method or field to store the object of the recipient class.unchanged—they’ll become local variables for the new method.
>* Replace all references to the old field with appropriate calls to methods in the recipient class. If the field isn’t private, take care of this in the superclass and subclasses.
>* Delete the field in the original class.

## WHY
Often fields are moved as part of the **Extract Class** technique. Deciding which class to leave the field in can be tough. Here is our rule of thumb:
> **put a field in the same place as the methods that use it** (or else where most of these methods are).
This rule will help in other cases when a field is simply located in the wrong place.

## PROS
>* More readable code! Be sure to give the new method a name that describes the method’s purpose: createOrder(), renderCustomerInfo(), etc.
>* Less code duplication. Often the code that’s found in a method can be reused in other places in your program. So you can replace duplicates with calls to your new method.
>* Isolates independent parts of code, meaning that errors are less likely (such as if the wrong variable is modified).

    Class1: ->   Class1:
    aField       

    Class2: ->   Class2:
                 aField