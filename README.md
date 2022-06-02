Go implementation of Design Pattern in *Head First Design Pattern*

## strategy

We should notice that Go doesn't have inheritance in struct, so the base type should be an interface. eg, if we need a duck to return from a func newDuck(), the duck should be an interface, and has several methods which a duck should have. And we need a baseDuck(or parentDuck, whatever), to implement the methods which defined by duck interface, and all son ducks should have common fields baseDuck, which contains the basically methods and fields that all sons need.