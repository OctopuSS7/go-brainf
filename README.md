# go-brainf

[BrainF](https://en.wikipedia.org/wiki/Brainfuck) interpreter / visualiser written in Go.

## Features

- Supports input.
- Visualises memory cells.

## To-Do

- Show position in program execution.
- Allow specification of amount of cells.
- Allow specification of looping, 
  - with cell values, 
  - and pointer values.
- Allow unlimited cells (both directions)
- Add more examples to demonstrate:
  - input capabilities,
  - control characters,
  - some basic maths.

## Limitations

- Escape codes used only work with unix based systems 
  - (or in the new Windows Terminal).

- Option to switch to win32 terminal codes soon.

## Example

Input: `BrainF -file=hello_world`

Output:

```
|-000-|-000-|-072-|-100-|-087-|-032-|-008-|-000-|
              (^)      
              
Output: Hello World
```
