# Python notebook #

- How to create an array in python
- Singleton: from `requests` module
- Lazy Method for Reading Big File in Python?
- function only run once / Efficient way of having a function only execute once in a loop


## How to create an array in python ##
`python3` use `'i'` for int array. [python3 docs](https://docs.python.org/3/library/array.html)

```python
# python3
import array
# i means int
arr1 = array.array('i')

# or need init
arr2 = array.array('i', (0 for i in range(0, 1000)))
```

```python
# python2
arr = range(1000)
```

## Singleton (from `requests` module)

```python
class AuthManager(object):
    """Authentication Manager."""
    
    def __new__(cls):
        singleton = cls.__dict__.get('__singleton__')
        if singleton is not None:
            return singleton

        cls.__singleton__ = singleton = object.__new__(cls)

        return singleton


    def __init__(self):
        self.passwd = {}
        self._auth = {}
```


## [Lazy Method for Reading Big File in Python?](https://stackoverflow.com/questions/519633/lazy-method-for-reading-big-file-in-python) ##

```python
def read_in_chunks(file_object, chunk_size=1024):
    """Lazy function (generator) to read a file piece by piece.
    Default chunk size: 1k."""
    while True:
        data = file_object.read(chunk_size)
        if not data:
            break
        yield data


f = open('really_big_file.dat')
for piece in read_in_chunks(f):
    process_data(piece)
```


## run only once ##
Link: (Efficient way of having a function only execute once in a loop)[https://stackoverflow.com/questions/4103773/efficient-way-of-having-a-function-only-execute-once-in-a-loop]

```python
def run_once(f):
    def wrapper(*args, **kwargs):
        if not wrapper.has_run:
            wrapper.has_run = True
            return f(*args, **kwargs)
    wrapper.has_run = False
    return wrapper


@run_once
def my_function(foo, bar):
    return foo+bar
```
