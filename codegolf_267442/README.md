# codegolf_267442

- Title: ["Implement a simple stack language"](https://codegolf.stackexchange.com/questions/267442)
- Posted at [Code Golf Stack Exchange](https://codegolf.stackexchange.com)
- Posted by user [@RubenVerg](https://codegolf.stackexchange.com/users/118045)
- Posted on December 11, 2023

## Description

In this challenge, you implement an interpreter for a simple stack-based programming
language. Your language must provide the following instructions:

 * push a positive number
 * pop two numbers and push their sum
 * pop two numbers and push their difference (second number - first number)
 * pop a number and push it twice (dup)
 * pop two numbers and push them so that they are in opposite order (swap)
 * pop a number and discard it (drop)

You may assume instructions will never be called with less arguments on the stack than
are needed.

The actual instructions can be chosen for each implementation, please specify how each
instruction is called in the solution. Your program/function must output/return the
stack after all instructions are performed sequentially. Output the stack in whatever
format you prefer. The stack must be empty at the start of your program.
