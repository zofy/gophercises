# Task CLI

Simple task CLI supporting following commands:
- __add__: adds new task to a task list
- __do__: completes task (task is no longer listed among *ToDo tasks*)
- __rm__: removes given task from a task list
- __list__: lists all tasks from the task list
- __completed__: lists all tasks done **today**

## Usage

`task <COMMAND> <ARGS>`

Example:
```
task add "go shopping" "feed the dog"
task list
>>>>
1: go shopping
2: feed the dog
>>>>
```

User can either _remove_ the task from a task list or _complete_ it. Once task is __completed__, it can be listed by: `task completed`

```
task add "cook a dinner"
task list
>>>>
1: cook a dinner
>>>>

task do 1
task list
>>>>
You have no tasks to finish!
>>>>
task completed
>>>>
You've completed following tasks today:
1: cook a dinner
>>>>
```

It is important to note that: __completed__ command lists all the tasks done in a given day.
