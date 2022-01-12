# Introduce Local Extension

**PROBLEM** : A utility class doesn’t contain some methods that you need. But you can’t add these methods to the class.
    ClientClass:
        next(:Date):Date

**SOLUTION**: Create a new class containing the methods and make it either the child or wrapper of the utility class.
    MfDate:
        nextDay(:Date):Date

## REFACTOR
>* Create a new extension class:
    >1. Option A: Make it a child of the utility class.
    >2. Option B: If you have decided to make a wrapper, create a field in it for storing the utility class object to which delegation will be made. When using this option, you will need to also create methods that repeat the public methods of the utility class and contain simple delegation to the methods of the utility object.
>* Create a constructor that uses the parameters of the constructor of the utility class.
>* Also create an alternative “converting” constructor that takes only the object of the original class in its parameters. This will help to substitute the extension for the objects of the original class.
>* Create new extended methods in the class. Move foreign methods from other classes to this class or else delete the foreign methods if their functionality is already present in the extension.
>* Replace use of the utility class with the new extension class in places where its functionality is needed.

## WHY
The class that you’re using doesn’t have the methods that you need. What’s worse, you can’t add these methods (because the classes are in a third-party library, for example). There are two ways out:
>* Create a **subclass** from the relevant class, containing the methods and inheriting everything else from the parent class. This way is easier but is sometimes blocked by the utility class itself (due to final).
>* Create a **wrapper** class that contains all the new methods and elsewhere will delegate to the related object from the utility class. This method is more work since you need not only code to maintain the relationship between the wrapper and utility object, but also a large number of simple delegating methods in order to emulate the public interface of the utility class.

## PROS
By moving additional methods to a separate extension class (wrapper or subclass), you avoid gumming up client classes with code that doesn’t fit. Program components are more coherent and are more reusable.

## DRAWBACKS
The reasons for having the method of a utility class in a client class won’t always be clear to the person maintaining the code after you. If the method can be used in other classes, you could benefit by creating a wrapper for the utility class and placing the method there. This is also beneficial when there are several such utility methods. **Introduce Local Extension** can help with this.