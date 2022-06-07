Go implementation of Design Pattern in *Head First Design Pattern*. Notice that Go is not an OO language and doesn't have object inheritance, some design patterns might not be included.

## Strategy

We should notice that Go doesn't have inheritance in struct, so the base type should be an interface. eg, if we need a duck to return from a func newDuck(), the duck should be an interface, and has several methods which a duck should have. And we need a baseDuck(or parentDuck, whatever), to implement the methods which defined by duck interface, and all son ducks should have common fields baseDuck, which contains the basically methods and fields that all sons need. See the code for more details.

## Observer

Observer design pattern is quite easy to understand, just see the code.

## Decorator

Decorator is quite interesting, it can help you to add some extra abilities to the original type. Because there is no something like super.Call() in Go, I use a field to help us achieve this.

## Singleton

Singleton is quite different in different languages, while in Go you can use sync.Once to help you with it. You can look up it source code to help you understand how Go to do this job. Go uses atomic int and mutex to achieve singleton's goal.

## Command

Simplify something in book, only 2 commands and no marco command and no interface like Light and fan, just simple light name and fan name which helps to understand.

## Adapter
Adapter helps to convert interface.

## Facade
Facade helps to hide lower interface and gives a simpler interface to users.