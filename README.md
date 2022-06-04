Go implementation of Design Pattern in *Head First Design Pattern*

## Strategy

We should notice that Go doesn't have inheritance in struct, so the base type should be an interface. eg, if we need a duck to return from a func newDuck(), the duck should be an interface, and has several methods which a duck should have. And we need a baseDuck(or parentDuck, whatever), to implement the methods which defined by duck interface, and all son ducks should have common fields baseDuck, which contains the basically methods and fields that all sons need. See the code for more details.

## Observer

Observer design pattern is quite easy to understand, just see the code.

## Decorator

Decorator is quite interesting, it can help you to add some extra abilities to the original type. Because there is no something like super.Call() in Go, I use a field to help us achieve this.