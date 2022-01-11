# Extract Method

**PROBLEM** : You have an expression that’s hard to understand.
**SOLUTION**: Place the result of the expression or its parts in separate variables that are self-explanatory.

**REFACTOR**
>* Insert a new line before the relevant expression and declare a new variable there. Assign part of the complex expression to this variable.
>* Replace that part of the expression with the new variable.
>* Repeat the process for all complex parts of the expression.

**WHY**
The main reason for extracting variables is to make a complex expression more understandable, by dividing it into its intermediate parts. These could be:
>* Condition of the **if()** operator or a part of the **?:** operator in C-based languages
>* A long arithmetic expression without intermediate results
>* Long multipart lines
Extracting a variable may be the first step towards performing **Extract Method** if you see that the extracted expression is used in other places in your code.

**PROS**
More readable code! Try to give the extracted variables good names that announce the variable’s purpose loud and clear. More readability, fewer long-winded comments. Go for names like **customerTaxValue, cityUnemploymentRate, clientSalutationString, etc.**

**DRAWBACKS**
More variables are present in your code. But this is counterbalanced by the ease of reading your code.

When refactoring conditional expressions, remember that the compiler will most likely optimize it to minimize the amount of calculations needed to establish the resulting value. Say you have a following expression **if (a() || b()) ....** The program won’t call the method **b** if the method **a** returns **true** because the resulting value will still be **true**, no matter what value returns **b**.

However, if you extract parts of this expression into variables, both methods will always be called, which might hurt performance of the program, especially if these methods do some heavyweight work.