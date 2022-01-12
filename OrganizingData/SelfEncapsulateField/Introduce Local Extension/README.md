# Self Encapsulate Field

**PROBLEM** : You use direct access to private fields inside a class.
**SOLUTION**: Create a getter and setter for the field, and use only them for accessing the field.

## WHY
Sometimes directly accessing a private field inside a class just isn’t flexible enough. You want to be able to initiate a field value when the first query is made or perform certain operations on new values of the field when they’re assigned, or maybe do all this in various ways in subclasses.

## PROS
>* Indirect access to fields is when a field is acted on via access methods (getters and setters). This approach is much more flexible than direct access to fields.
    >1. First, you can perform complex operations when data in the field is set or received. Lazy initialization and validation of field values are easily implemented inside field getters and setters.
    >2. Second and more crucially, you can redefine getters and setters in subclasses.
>* You have the option of not implementing a setter for a field at all. The field value will be specified only in the constructor, thus making the field unchangeable throughout the entire object lifespan.

## DRAWBACKS
When direct access to fields is used, code looks simpler and more presentable, although flexibility is diminished.

## REFACTOR
>* Create a getter (and optional setter) for the field. They should be either **protected** or **public**.
>* Find all direct invocations of the field and replace them with getter and setter calls.
