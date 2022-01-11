# InLine Temp

**PROBLEM** : You have a temporary variable that’s assigned the result of a simple expression and nothing more.
**SOLUTION**: Replace the references to the variable with the expression itself.

## REFACTOR
>* Find all places that use the variable. Instead of the variable, use the expression that had been assigned to it.
>* Delete the declaration of the variable and its assignment line.

## WHY
Inline local variables are almost always used as part of **Replace Temp with Query** or to pave the way for **Extract Method.**

## PROS
This refactoring technique offers almost no benefit in and of itself. However, if the variable is assigned the result of a method, you can marginally improve the readability of the program by getting rid of the unnecessary variable.

## DRAWBACKS
Sometimes seemingly useless temps are used to cache the result of an expensive operation that’s reused several times. So before using this refactoring technique, make sure that simplicity won’t come at the cost of performance.