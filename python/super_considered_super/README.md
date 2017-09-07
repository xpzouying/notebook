# super considered super

## ref urls:

- [super considered super](https://rhettinger.wordpress.com/2011/05/26/super-considered-super/)
- [MOR in python official docs](https://www.python.org/download/releases/2.3/mro/)

## quick note

`super()` is not fater/base class. `super` is the next class in MRO list.

```python
def super(cls, inst):
    mro_list = inst.__class__.mor()
    return mor_list[mor_list(cls) + 1]
```

### example

```python
class A(object): pass
class B(A): pass
class C(A): pass
class D(B, C): pass

help(D)
```

> class D(B, C)  
>  |  Method resolution order:  
>  |      D  
>  |      B  
>  |      C  
>  |      A  
>  |      __builtin__.object  

### example

```python
class A(object):
    def name(self):
        print('A::name()')

class B(A):
    def name(self):
        super(B, self).name()

class C(A):
    def name(self):
        print('C::name()')

class D(B, C):
    pass

D().name()
```

What is the answer and why?

1. The class `D` have not `name()`, so search `name()` function in next class in `MRO`.
2. The next class for `D` in MRO list is `B`.
3. `name()` in `B` is call `super().name()`. Which is **`super()` of B**? A is base class and the next class in MRO list is C, which is right?
4. The result is `C::name()`. Because `super()` is the **next class in MRO list**, not base class.
