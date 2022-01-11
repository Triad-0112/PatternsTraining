# Substitute Algorithm

**PROBLEM** : So you want to replace an existing algorithm with a new one?
**SOLUTION**: Replace the body of the method that implements the algorithm with a new algorithm.

## REFACTOR
>* Make sure that you have simplified the existing algorithm as much as possible. Move unimportant code to other methods using **Extract Method**. The fewer moving parts in your algorithm, the easier it’s to replace.
>* Create your new algorithm in a new method. Replace the old algorithm with the new one and start testing the program.
>* If the results don’t match, return to the old implementation and compare the results. Identify the causes of the discrepancy. While the cause is often an error in the old algorithm, it’s more likely due to something not working in the new one.
>* When all tests are successfully completed, delete the old algorithm for good!

## WHY
* Gradual refactoring isn’t the only method for improving a program. Sometimes a method is so cluttered with issues that it’s easier to tear down the method and start fresh. And perhaps you have found an algorithm that’s much simpler and more efficient. If this is the case, you should simply replace the old algorithm with the new one.
* As time goes on, your algorithm may be incorporated into a well-known library or framework and you want to get rid of your independent implementation, in order to simplify maintenance.
* The requirements for your program may change so heavily that your existing algorithm can’t be salvaged for the task.