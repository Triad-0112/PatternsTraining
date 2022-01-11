# Inline Method
**PROBLEM** : When a method body is more obvious than the method itself, use this technique.
**SOLUTION**: Replace calls to the method with the method’s content and delete the method itself.

## REFACTOR
>1. Make sure that the method isn’t redefined in subclasses. If the method is redefined, refrain from this technique.
>2. Find all calls to the method. Replace these calls with the content of the method.
>3. Delete the method.

## WHY
> A method simply delegates to another method. In itself, this delegation is no problem. But when there are many such methods, they become a confusing tangle that’s hard to sort through.
> Often methods aren’t too short originally, but become that way as changes are made to the program. So don’t be shy about getting rid of methods that have outlived their use.

## PROS
>* By minimizing the number of unneeded methods, you make the code more straightforward.