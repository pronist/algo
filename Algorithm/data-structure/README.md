# Data Structure In C

* **Fixed Array**
* **Deque**
* **Doubly Linked List**
* **Queue**
* **Stack**
* **Binary Search Tree**

## Fixed Array

```c
Array arr = array(1, "Hello, world!");
for (void* iter = ar_begin(&arr); 1 != ar_done(iter); iter = ar_next(&arr)) {
  // Hello, world!
  printf("%s", (const char*) iter);
}

ar_clear(&arr);
```

* Array array(size_t length, ...)
* void* ar_at(Array* array, int position)
* void ar_clear(Array* array)

### Array Iterators

* void* ar_begin(Array* array)
* void* ar_next(Array* array)
* void* ar_end(Array* array)
* void* ar_prev(Array* array)
* int ar_done(void* iter)

## Deque

```c
Deque dq = deque(1, "Hello, world!");
dq_push_back(&dq, "Example");

for (void* iter = dq_begin(&dq); 1 != dq_done(iter); iter = dq_next(&dq)) {
  // Hello, world!
  // Example
  printf("%s", (const char*) iter);
}

dq_clear(&dq);
```

* Deque deque(size_t length, ...)
* void dq_push_front(Deque*, deque void* element)
* void dq_pop_front(Deque* deque)
* void dq_push_back(Deque* deque, void* element)
* void dq_pop_back(Deque* deque)
* void dq_insert(Deque* deque, void* element, int position)
* void dq_erase(Deque* deque, int position)
* void* dq_at(Deque* deque, int position)
* void dq_clear(Deque* deque)

### Deque Iterators

* void* dq_begin(Deque* deque)
* void* dq_next(Deque* deque)
* void* dq_end(Deque* deque)
* void* dq_prev(Deque* deque)
* int dq_done(void* iter)

## Linked list

```c
List ls = list(1, "Hello, world!");
ls_push_front(&ls, "Example");

for (Node* iter = ls_begin(&ls); 1 != ls_done(iter); iter = ls_next(iter)) {
  // Example
  // Hello, world!
  printf("%s", (const char*) iter->element);
}

ls_clear(&ls);
```

* List list(size_t length, ...)
* void ls_push_front(List* list, void* element)
* void ls_pop_front(List* list)
* void ls_push_back(List* list, void* element)
* void ls_pop_back(List* list)
* void ls_insert(List* list, Node* node, void* element)
* void ls_erase(List* list, Node* node)
* Node* ls_at(List* list, void* element)
* void ls_clear(List* list)

### Linked List Iterators

* Node* ls_begin(List* list)
* Node* ls_next(Node* iter)
* Node* ls_end(List* list)
* Node* ls_prev(Node* iter)
* int ls_done(Node* iter)

## Queue

```c
Queue qu = queue(1, "Hello, world!");
qu_push(&qu, "Who are you?");
qu_push(&qu, "Goodbye");

// GoodBye
printf("%s", (const char*)qu_rear(&qu)->element);
// Hello, world!
printf("%s", (const char*)qu_front(&qu)->element);

// Remove "Hello, world" Node
qu_pop(&qu);

// Who are you?
printf("%s", (const char*)qu_front(&qu)->element);
```

* Queue queue(size_t length, ...)
* void qu_push(Queue* queue, void* element)
* void qu_pop(Queue* queue)
* void qu_clear(Queue* queue)
* Node* qu_front(Queue* queue)
* Node* qu_rear(Queue* queue)

## Stack

```c
Stack st = stack(1, "Hello, world!");
st_push(&st, "Who are you?");
st_push(&st, "Goodbye");

// GoodBye
printf("%s", (const char*)st_top(&st)->element);

// Remove "GoodBye" Node
st_pop(&st);

// Who are you?
printf("%s", (const char*)st_top(&st)->element);
```

* Stack stack(size_t length, ...)
* void st_push(Stack* stack, void* element)
* void st_pop(Stack* stack)
* void st_clear(Stack* stack)
* Node* st_top(Stack* stack)

## Tree

```c
Tree tr = tree(1, "Hello, world!");
tr_insert(&tr, "Example");

for (Node* iter = tr_begin(&tr); 1 != tr_done(iter); iter = tr_next(&tr, &iter)) {
  // Hello, world
  // Example
  printf("%s", (const char*) iter);
}

tr_clear(&tr);
```

* Tree tree(size_t length, ...)
* void tr_insert(Tree* tree, void* element)
* void tr_erase(Tree* tree, void* element)
* Node* tr_at(Tree* tree, void* element)
* void tr_clear(Tree* tree)

### Tree Iterators

* Node* tr_begin(Tree* tree)
* Node* tr_next(Tree*, Node** iter)
* Node* tr_end(Tree* tree)
* Node* tr_prev(Tree* tree, Node** iter)
* int tr_done(void* iter)

## License

[MIT](https://github.com/pronist/Data-Structure-In-C/blob/master/LICENSE)

Copyright 2019-2020. [SangWoo Jeong](https://github.com/pronist). All rights reserved.
