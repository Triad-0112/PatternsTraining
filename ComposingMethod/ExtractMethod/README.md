# Extract Method

**PROBLEM** : Have a code fragment that can be grouped together
**SOLUTION**: Move the code to a separate new new method (or function) and replace the old code with a call to the method.

**REFACTOR**:
>1. Create a new method and name it in a way that makes its purpose self-evident.
>2. Copy the relevant code fragment to your new method. Delete the fragment from its old location and put a call for the new method there instead.
>3. Find all variables used in this code fragment. If they’re declared inside the fragment and not used outside of it, simply leave them unchanged—they’ll become local variables for the new method.
>4. If the variables are declared prior to the code that you’re extracting, you will need to pass these variables to the parameters of your new method in order to use the values previously contained in them. Sometimes it’s easier to get rid of these variables by resorting to Replace Temp with Query.
>5. If you see that a local variable changes in your extracted code in some way, this may mean that this changed value will be needed later in your main method. Double-check! And if this is indeed the case, return the value of this variable to the main method to keep everything functioning.