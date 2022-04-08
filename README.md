<!--suppress HtmlDeprecatedAttribute -->
<img align="center" alt="gix banner" src="/.github/assets/GixBanner.svg">

---

<h4 align="center"><i>Provides utilities, some using 1.18 generics.</i></h4>

---

### Table of Contents

- [Packages](#packages)

    - [cond](#cond)
    - [fmt](#fmt)
    - [math](#math)
    - [slice](#slice)
      <br/><br/>
  > ---
- [Licensing](#licensing)

---

### Packages

#### cond

Package `cond` offers functions to conditionally apply, act, pass, or return values.

| Function  | Description                                                                                                                                                |
|:---------:|:-----------------------------------------------------------------------------------------------------------------------------------------------------------|
| `Ternary` | Returns `a` or `b` depending on `cond`. Great for usage in situations like string concatenation where you want to conditionally choose between two values. |

#### fmt

Package `fmt` offers functions for various formatting purposes.

| Function | Description                                                                                                                                               |
|:--------:|:----------------------------------------------------------------------------------------------------------------------------------------------------------|
| `Bytes`  | Returns the provided `uint64` as a human-readable representation of storage values based on IEC bytes. Useful for displaying memory usage, or file sizes. |

#### math

Package `math` provides alternatives to the standard library's Max and Min functions that accept any
builtin number type that supports arithmetic, as well as utility type constraints.

|   Type    | Description                                                                                |
|:---------:|:-------------------------------------------------------------------------------------------|
|   `Int`   | A type constraint containing all built-in signed integer types.                            |
|  `Uint`   | A type constraint containing all built-in signed integer types.                            |
|  `Float`  | A type constraint containing all built-in float types.                                     |
| `Complex` | A type constraint containing all built-in complex types.                                   |
| `Number`  | A type constraint containing all built-in number types that support arithmetic operations. |

| Function  | Description                                       |
|:---------:|:--------------------------------------------------|
|   `Min`   | Returns the smallest of the two provided numbers. |
| `MinSet`  | Returns the smallest of all the provided numbers. |
|   `Max`   | Returns the largest of the two provided numbers.  |
| `MaxSet`  | Returns the largest of all the provided numbers.  |

#### slice

Package `slice` offers functions to mutate or create new slices from existing ones. None of these
functions mutate the original slice, instead returning a new slice with the changes applied.

| Function  | Description                                                                                                                                     |
|:---------:|:------------------------------------------------------------------------------------------------------------------------------------------------|
|   `Pop`   | Removes and returns the last element of a slice.                                                                                                |
|  `Shift`  | Removes and returns the first element of a slice.                                                                                               |
| `Prepend` | Prepends the passed values to the beginning of a slice, passing only one value acts as a reverse of `Shift`.                                    |
|   `Map`   | Calls a function on each member of a slice, returning a slice with the mutated values.                                                          |
| `Reduce`  | Calls a function on each member of a slice, returning a slice with only values which returned a `true` bool after being passed to the function. |

---

### Licensing

`gix` is licensed under the `BSD-3-Clause` license found [here](/LICENSE).
