# Inline Class

**PROBLEM** : A class does almost nothing and isn’t responsible for anything, and no additional responsibilities are planned for it.
**SOLUTION**: Move all features from the class to another one.

## REFACTOR
>* In the recipient class, create the public fields and methods present in the donor class. Methods should refer to the equivalent methods of the donor class.
>* Replace all references to the donor class with references to the fields and methods of the recipient class.
>* Now test the program and make sure that no errors have been added. If tests show that everything is working A-OK, start using Move Method and Move Field to completely transplant all functionality to the recipient class from the original one. Continue doing so until the original class is completely empty.
>* Delete the original class.

## WHY
Often this technique is needed after the features of one class are “transplanted” to other classes, leaving that class with little to do.

## PROS
>* Eliminating needless classes frees up operating memory on the computer—and bandwidth in your head.

## EXAMPLE
Person :

    name
    getTelephoneNumber()

Telephone Number :

    officeAreaCode
    officeNumber
    getTelephoneNumber()

Turn it into 

Person :

    name
    officeAreaCodee
    officeNumber
    getTelephoneNumber()



